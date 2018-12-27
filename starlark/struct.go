// Copyright 2017 The Bazel Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package starlark

// Package starlarkstruct defines the Starlark 'struct' type,
// and was merged with package starlark.

// It is tempting to introduce a variant of Struct that is a wrapper
// around a Go struct value, for stronger typing guarantees and more
// efficient and convenient field lookup. However:
// 1) all fields of Starlark structs are optional, so we cannot represent
//    them using more specific types such as String, Int, *Depset, and
//    *File, as such types give no way to represent missing fields.
// 2) the efficiency gain of direct struct field access is rather
//    marginal: finding the index of a field by binary searching on the
//    sorted list of field names is quite fast compared to the other
//    overheads.
// 3) the gains in compactness and spatial locality are also rather
//    marginal: the array behind the []entry slice is (due to field name
//    strings) only a factor of 2 larger than the corresponding Go struct
//    would be, and, like the Go struct, requires only a single allocation.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/glycerine/monty/syntax"
)

// Make is the implementation of a built-in function that instantiates
// an immutable struct from the specified keyword arguments.
//
// An application can add 'struct' to the Starlark environment like so:
//
// 	globals := starlark.StringDict{
// 		"struct":  starlark.NewBuiltin("struct", starlarkstruct.Make),
// 	}
//
func StructMake(_ *Thread, _ *Builtin, args Tuple, kwargs []Tuple) (Value, error) {
	if len(args) > 0 {
		return nil, fmt.Errorf("struct: unexpected positional arguments")
	}
	return StructFromKeywords(Default, kwargs), nil
}

// FromKeywords returns a new struct instance whose fields are specified by the
// key/value pairs in kwargs.  (Each kwargs[i][0] must be a starlark.String.)
func StructFromKeywords(ctor string, kwargs []Tuple) *Struct {
	if ctor == "" {
		ctor = Default
	}
	s := NewStruct(String(ctor))
	for _, kwarg := range kwargs {
		k := string(kwarg[0].(String))
		v := kwarg[1]
		s.Fields = append(s.Fields, field{k, v})
		s.Map[k] = len(s.Fields) - 1
	}
	return s
}

func NewStruct(constructor Value) *Struct {
	s, ok := constructor.(String)
	if !ok {
		panic(fmt.Sprintf("error: NewStruct must be called with String ctor, got %T", constructor))
	}
	return &Struct{
		Ctor: string(s),
		Map:  make(map[string]int),
	}
}

// FromStringDict returns a whose elements are those of d.
// The constructor parameter specifies the constructor; use Default for an ordinary struct.
func FromStringDict(constructor Value, d *StringDict) *Struct {
	if constructor == nil {
		panic("nil constructor")
	}
	s := NewStruct(constructor)
	for k, v := range d.Map {
		s.Fields = append(s.Fields, field{k, v})
		s.Map[k] = len(s.Fields) - 1
	}
	return s
}

// Struct is an immutable Starlark type that maps field names to values.
// It is not iterable and does not support len.
//
// A struct has a constructor, a distinct value that identifies a class
// of structs, and which appears in the struct's string representation.
//
// Operations such as x+y fail if the constructors of the two operands
// are not equal.
//
// The default constructor, Default, is the string "struct", but
// clients may wish to 'brand' structs for their own purposes.
// The constructor value appears in the printed form of the value,
// and is accessible using the Constructor method.
//
// Use Attr to access its fields and AttrNames to enumerate them.
type Struct struct {
	Ctor   string
	Fields []field        // insertion order
	Map    map[string]int // INVAR: fields[Map[field]] == field
}

// Default is the default constructor for structs.
// It is merely the string "struct".
const Default = "struct"

type fields []field

func (a fields) Len() int           { return len(a) }
func (a fields) Less(i, j int) bool { return a[i].name < a[j].name }
func (a fields) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type field struct {
	name  string // not to_{proto,json}
	value Value
}

var (
	_ HasAttrs  = (*Struct)(nil)
	_ HasBinary = (*Struct)(nil)
)

// ToStringDict adds a name/value field to d for each field of the struct.
func (s *Struct) ToStringDict(d StringDict) {
	for _, e := range s.Fields {
		d.Map[e.name] = e.value
	}
}

func (s *Struct) String() string {
	var buf bytes.Buffer
	if s.Ctor == Default {
		// NB: The Java implementation always prints struct
		// even for Bazel provider instances.
		buf.WriteString("struct") // avoid String()'s quotation
	} else {
		buf.WriteRune('$')
		buf.WriteString(s.Ctor)
	}
	buf.WriteByte('{')

	// write in sorted order
	sorted := s.AttrNames()
	for i, fld := range sorted {
		e := s.Fields[s.Map[fld]]
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(e.name)
		buf.WriteString(": ")
		buf.WriteString(e.value.String())
	}
	buf.WriteByte('}')
	return buf.String()
}

// Constructor returns the constructor used to create this struct.
func (s *Struct) Constructor() Value { return String(s.Ctor) }

func (s *Struct) Type() string {
	if s.Ctor != Default {
		return "struct_literal<$" + s.Ctor + ">"
	}
	return "struct"
}
func (s *Struct) Truth() Bool { return true } // even when empty
func (s *Struct) Hash() (uint32, error) {
	// Same algorithm as Tuple.hash, but with different primes.
	var x, m uint32 = 8731, 9839
	for _, e := range s.Fields {
		namehash, _ := String(e.name).Hash()
		x = x ^ 3*namehash
		y, err := e.value.Hash()
		if err != nil {
			return 0, err
		}
		x = x ^ y*m
		m += 7349
	}
	return x, nil
}
func (s *Struct) Freeze() {
	for _, e := range s.Fields {
		e.value.Freeze()
	}
}

func (x *Struct) Binary(op syntax.Token, y Value, side Side) (Value, error) {
	if y, ok := y.(*Struct); ok && op == syntax.PLUS {
		if side == Right {
			x, y = y, x
		}

		if x.Ctor != y.Ctor {
			return nil, fmt.Errorf("cannot add structs of different constructors: %s + %s",
				x.Ctor, y.Ctor)
		}

		z := NewStringDict(x.len() + y.len())
		for _, e := range x.Fields {
			z.Map[e.name] = e.value
		}
		for _, e := range y.Fields {
			z.Map[e.name] = e.value
		}

		return FromStringDict(String(x.Ctor), z), nil
	}
	return nil, nil // unhandled
}

// update if already present, insert if not.
func (s *Struct) Upsert(key Value, val Value) error {
	k, ok := key.(String)
	if !ok {
		return fmt.Errorf("'%s' fields must be string labels; we got %T", s.Ctor, key)
	}

	ks := string(k)
	where, already := s.Map[ks]
	if already {
		s.Fields[where] = field{ks, val}
	} else {
		s.Fields = append(s.Fields, field{ks, val})
		s.Map[ks] = len(s.Fields) - 1
	}
	return nil
}

// SetField is required for
// the HasSetField interface. This means we have fields
// that may be written by a dot expression such as "x.f = y".
func (s *Struct) SetField(name string, val Value) error {

	where, already := s.Map[name]
	if already {
		s.Fields[where] = field{name, val}
		return nil
	}
	return fmt.Errorf("struct $%s has no .%s attribute", s.Ctor, name)
}

// Attr returns the value of the specified field,
// or deprecated method if the name is "to_json" or "to_proto"
// and the struct has no field of that name.
func (s *Struct) Attr(name string) (Value, error) {
	k, ok := s.Map[name]
	if ok {
		return s.Fields[k].value, nil
	}

	// TODO(adonovan): there may be a nice feature for core
	// Value here, especially for JSON, but the current
	// features are incomplete and underspecified.
	//
	// to_{json,proto} are deprecated, appropriately; see Google issue b/36412967.
	switch name {
	case "to_json", "to_proto":
		return NewBuiltin(name, func(thread *Thread, fn *Builtin, args Tuple, kwargs []Tuple) (Value, error) {
			var buf bytes.Buffer
			var err error
			if name == "to_json" {
				err = writeJSON(&buf, s)
			} else {
				err = writeProtoStruct(&buf, 0, s)
			}
			if err != nil {
				// TODO(adonovan): improve error error messages
				// to show the path through the object graph.
				return nil, err
			}
			return String(buf.String()), nil
		}), nil
	}

	var ctor string
	if s.Ctor != Default {
		ctor = s.Ctor + " "
	}
	return nil, fmt.Errorf("%sstruct has no .%s attribute", ctor, name)
}

func writeProtoStruct(out *bytes.Buffer, depth int, s *Struct) error {
	for _, e := range s.Fields {
		if err := writeProtoField(out, depth, e.name, e.value); err != nil {
			return err
		}
	}
	return nil
}

func writeProtoField(out *bytes.Buffer, depth int, field string, v Value) error {
	if depth > 16 {
		return fmt.Errorf("to_proto: depth limit exceeded")
	}

	switch v := v.(type) {
	case *Struct:
		fmt.Fprintf(out, "%*s%s {\n", 2*depth, "", field)
		if err := writeProtoStruct(out, depth+1, v); err != nil {
			return err
		}
		fmt.Fprintf(out, "%*s}\n", 2*depth, "")
		return nil

	case *List, Tuple:
		iter := Iterate(v)
		defer iter.Done()
		var elem Value
		for iter.Next(&elem) {
			if err := writeProtoField(out, depth, field, elem); err != nil {
				return err
			}
		}
		return nil
	}

	// scalars
	fmt.Fprintf(out, "%*s%s: ", 2*depth, "", field)
	switch v := v.(type) {
	case Bool:
		fmt.Fprintf(out, "%t", v)

	case Int:
		// TODO(adonovan): limits?
		out.WriteString(v.String())

	case Float:
		fmt.Fprintf(out, "%g", v)

	case String:
		fmt.Fprintf(out, "%q", string(v))

	default:
		return fmt.Errorf("cannot convert %s to proto", v.Type())
	}
	out.WriteByte('\n')
	return nil
}

// writeJSON writes the JSON representation of a Starlark value to out.
// TODO(adonovan): there may be a nice feature for core Value here,
// but the current feature is incomplete and underspecified.
func writeJSON(out *bytes.Buffer, v Value) error {
	switch v := v.(type) {
	case NoneType:
		out.WriteString("null")
	case Bool:
		fmt.Fprintf(out, "%t", v)
	case Int:
		// TODO(adonovan): test large numbers.
		out.WriteString(v.String())
	case Float:
		// TODO(adonovan): test.
		fmt.Fprintf(out, "%g", v)
	case String:
		s := string(v)
		if goQuoteIsSafe(s) {
			fmt.Fprintf(out, "%q", s)
		} else {
			// vanishingly rare for text strings
			data, _ := json.Marshal(s)
			out.Write(data)
		}
	case Indexable: // Tuple, List
		out.WriteByte('[')
		for i, n := 0, Len(v); i < n; i++ {
			if i > 0 {
				out.WriteString(", ")
			}
			if err := writeJSON(out, v.Index(i)); err != nil {
				return err
			}
		}
		out.WriteByte(']')
	case *Struct:
		out.WriteByte('{')
		for i, e := range v.Fields {
			if i > 0 {
				out.WriteString(", ")
			}
			if err := writeJSON(out, String(e.name)); err != nil {
				return err
			}
			out.WriteString(": ")
			if err := writeJSON(out, e.value); err != nil {
				return err
			}
		}
		out.WriteByte('}')
	default:
		return fmt.Errorf("cannot convert %s to JSON", v.Type())
	}
	return nil
}

func goQuoteIsSafe(s string) bool {
	for _, r := range s {
		// JSON doesn't like Go's \xHH escapes for ASCII control codes,
		// nor its \UHHHHHHHH escapes for runes >16 bits.
		if r < 0x20 || r >= 0x10000 {
			return false
		}
	}
	return true
}

func (s *Struct) len() int { return len(s.Fields) }

// AttrNames returns a new sorted list of the struct fields.
//
// Unlike in the Java implementation,
// the deprecated struct methods "to_json" and "to_proto" do not
// appear in AttrNames, and hence dir(struct), since that would force
// the majority to have to ignore them, but they may nonetheless be
// called if the struct does not have fields of these names. Ideally
// these will go away soon. See Google issue b/36412967.
func (s *Struct) AttrNames() []string {
	names := make([]string, len(s.Fields))
	for i, e := range s.Fields {
		names[i] = e.name
	}
	sort.Strings(names)
	return names
}

func (x *Struct) CompareSameType(op syntax.Token, y_ Value, depth int) (bool, error) {
	y := y_.(*Struct)
	switch op {
	case syntax.EQL:
		return structsEqual(x, y, depth)
	case syntax.NEQ:
		eq, err := structsEqual(x, y, depth)
		return !eq, err
	default:
		return false, fmt.Errorf("%s %s %s not implemented", x.Type(), op, y.Type())
	}
}

func structsEqual(x, y *Struct, depth int) (bool, error) {
	if x.len() != y.len() {
		return false, nil
	}

	if x.Ctor != y.Ctor {
		return false, nil
	}

	for i, n := 0, x.len(); i < n; i++ {
		if x.Fields[i].name != y.Fields[i].name {
			return false, nil
		} else if eq, err := EqualDepth(x.Fields[i].value, y.Fields[i].value, depth-1); err != nil {
			return false, err
		} else if !eq {
			return false, nil
		}
	}
	return true, nil
}
