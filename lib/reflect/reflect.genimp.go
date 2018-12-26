package shadow_reflect

import "reflect"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["Append"] = reflect.Append
    Pkg["AppendSlice"] = reflect.AppendSlice
    Pkg["ArrayOf"] = reflect.ArrayOf
    Pkg["ChanOf"] = reflect.ChanOf
    Pkg["Copy"] = reflect.Copy
    Pkg["DeepEqual"] = reflect.DeepEqual
    Pkg["FuncOf"] = reflect.FuncOf
    Pkg["Indirect"] = reflect.Indirect
    Pkg["MakeChan"] = reflect.MakeChan
    Pkg["MakeFunc"] = reflect.MakeFunc
    Pkg["MakeMap"] = reflect.MakeMap
    Pkg["MakeMapWithSize"] = reflect.MakeMapWithSize
    Pkg["MakeSlice"] = reflect.MakeSlice
    Pkg["MapOf"] = reflect.MapOf
    Ctor["Method"] = Shadow_NewStruct_Method
    Pkg["New"] = reflect.New
    Pkg["NewAt"] = reflect.NewAt
    Pkg["PtrTo"] = reflect.PtrTo
    Pkg["Select"] = reflect.Select
    Ctor["SelectCase"] = Shadow_NewStruct_SelectCase
    Ctor["SliceHeader"] = Shadow_NewStruct_SliceHeader
    Pkg["SliceOf"] = reflect.SliceOf
    Ctor["StringHeader"] = Shadow_NewStruct_StringHeader
    Ctor["StructField"] = Shadow_NewStruct_StructField
    Pkg["StructOf"] = reflect.StructOf
    Pkg["Swapper"] = reflect.Swapper
    Pkg["Type"] = Shadow_InterfaceConvertTo2_Type
    Pkg["TypeOf"] = reflect.TypeOf
    Ctor["Value"] = Shadow_NewStruct_Value
    Ctor["ValueError"] = Shadow_NewStruct_ValueError
    Pkg["ValueOf"] = reflect.ValueOf
    Pkg["Zero"] = reflect.Zero

}
func Shadow_NewStruct_Method(src *reflect.Method) *reflect.Method {
    if src == nil {
	   return &reflect.Method{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_SelectCase(src *reflect.SelectCase) *reflect.SelectCase {
    if src == nil {
	   return &reflect.SelectCase{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_SliceHeader(src *reflect.SliceHeader) *reflect.SliceHeader {
    if src == nil {
	   return &reflect.SliceHeader{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_StringHeader(src *reflect.StringHeader) *reflect.StringHeader {
    if src == nil {
	   return &reflect.StringHeader{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_StructField(src *reflect.StructField) *reflect.StructField {
    if src == nil {
	   return &reflect.StructField{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_Type(x interface{}) (y reflect.Type, b bool) {
	y, b = x.(reflect.Type)
	return
}

func Shadow_InterfaceConvertTo1_Type(x interface{}) reflect.Type {
	return x.(reflect.Type)
}


func Shadow_NewStruct_Value(src *reflect.Value) *reflect.Value {
    if src == nil {
	   return &reflect.Value{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_ValueError(src *reflect.ValueError) *reflect.ValueError {
    if src == nil {
	   return &reflect.ValueError{}
    }
    a := *src
    return &a
}

