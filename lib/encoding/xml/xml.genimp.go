package shadow_xml

import "encoding/xml"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Ctor["Attr"] = Shadow_NewStruct_Attr
    Pkg["CopyToken"] = xml.CopyToken
    Ctor["Decoder"] = Shadow_NewStruct_Decoder
    Ctor["Encoder"] = Shadow_NewStruct_Encoder
    Ctor["EndElement"] = Shadow_NewStruct_EndElement
    Pkg["Escape"] = xml.Escape
    Pkg["EscapeText"] = xml.EscapeText
    Pkg["HTMLAutoClose"] = xml.HTMLAutoClose
    Pkg["HTMLEntity"] = xml.HTMLEntity
    Pkg["Header"] = xml.Header
    Pkg["Marshal"] = xml.Marshal
    Pkg["MarshalIndent"] = xml.MarshalIndent
    Pkg["Marshaler"] = Shadow_InterfaceConvertTo2_Marshaler
    Pkg["MarshalerAttr"] = Shadow_InterfaceConvertTo2_MarshalerAttr
    Ctor["Name"] = Shadow_NewStruct_Name
    Pkg["NewDecoder"] = xml.NewDecoder
    Pkg["NewEncoder"] = xml.NewEncoder
    Pkg["NewTokenDecoder"] = xml.NewTokenDecoder
    Ctor["ProcInst"] = Shadow_NewStruct_ProcInst
    Ctor["StartElement"] = Shadow_NewStruct_StartElement
    Ctor["SyntaxError"] = Shadow_NewStruct_SyntaxError
    Ctor["TagPathError"] = Shadow_NewStruct_TagPathError
    Pkg["Token"] = Shadow_InterfaceConvertTo2_Token
    Pkg["TokenReader"] = Shadow_InterfaceConvertTo2_TokenReader
    Pkg["Unmarshal"] = xml.Unmarshal
    Pkg["Unmarshaler"] = Shadow_InterfaceConvertTo2_Unmarshaler
    Pkg["UnmarshalerAttr"] = Shadow_InterfaceConvertTo2_UnmarshalerAttr
    Ctor["UnsupportedTypeError"] = Shadow_NewStruct_UnsupportedTypeError

}
func Shadow_NewStruct_Attr(src *xml.Attr) *xml.Attr {
    if src == nil {
	   return &xml.Attr{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Decoder(src *xml.Decoder) *xml.Decoder {
    if src == nil {
	   return &xml.Decoder{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Encoder(src *xml.Encoder) *xml.Encoder {
    if src == nil {
	   return &xml.Encoder{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_EndElement(src *xml.EndElement) *xml.EndElement {
    if src == nil {
	   return &xml.EndElement{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_Marshaler(x interface{}) (y xml.Marshaler, b bool) {
	y, b = x.(xml.Marshaler)
	return
}

func Shadow_InterfaceConvertTo1_Marshaler(x interface{}) xml.Marshaler {
	return x.(xml.Marshaler)
}


func Shadow_InterfaceConvertTo2_MarshalerAttr(x interface{}) (y xml.MarshalerAttr, b bool) {
	y, b = x.(xml.MarshalerAttr)
	return
}

func Shadow_InterfaceConvertTo1_MarshalerAttr(x interface{}) xml.MarshalerAttr {
	return x.(xml.MarshalerAttr)
}


func Shadow_NewStruct_Name(src *xml.Name) *xml.Name {
    if src == nil {
	   return &xml.Name{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_ProcInst(src *xml.ProcInst) *xml.ProcInst {
    if src == nil {
	   return &xml.ProcInst{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_StartElement(src *xml.StartElement) *xml.StartElement {
    if src == nil {
	   return &xml.StartElement{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_SyntaxError(src *xml.SyntaxError) *xml.SyntaxError {
    if src == nil {
	   return &xml.SyntaxError{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_TagPathError(src *xml.TagPathError) *xml.TagPathError {
    if src == nil {
	   return &xml.TagPathError{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_Token(x interface{}) (y xml.Token, b bool) {
	y, b = x.(xml.Token)
	return
}

func Shadow_InterfaceConvertTo1_Token(x interface{}) xml.Token {
	return x.(xml.Token)
}


func Shadow_InterfaceConvertTo2_TokenReader(x interface{}) (y xml.TokenReader, b bool) {
	y, b = x.(xml.TokenReader)
	return
}

func Shadow_InterfaceConvertTo1_TokenReader(x interface{}) xml.TokenReader {
	return x.(xml.TokenReader)
}


func Shadow_InterfaceConvertTo2_Unmarshaler(x interface{}) (y xml.Unmarshaler, b bool) {
	y, b = x.(xml.Unmarshaler)
	return
}

func Shadow_InterfaceConvertTo1_Unmarshaler(x interface{}) xml.Unmarshaler {
	return x.(xml.Unmarshaler)
}


func Shadow_InterfaceConvertTo2_UnmarshalerAttr(x interface{}) (y xml.UnmarshalerAttr, b bool) {
	y, b = x.(xml.UnmarshalerAttr)
	return
}

func Shadow_InterfaceConvertTo1_UnmarshalerAttr(x interface{}) xml.UnmarshalerAttr {
	return x.(xml.UnmarshalerAttr)
}


func Shadow_NewStruct_UnsupportedTypeError(src *xml.UnsupportedTypeError) *xml.UnsupportedTypeError {
    if src == nil {
	   return &xml.UnsupportedTypeError{}
    }
    a := *src
    return &a
}

