package shadow_json

import "encoding/json"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["Compact"] = json.Compact
    Ctor["Decoder"] = Shadow_NewStruct_Decoder
    Ctor["Encoder"] = Shadow_NewStruct_Encoder
    Pkg["HTMLEscape"] = json.HTMLEscape
    Pkg["Indent"] = json.Indent
    Ctor["InvalidUTF8Error"] = Shadow_NewStruct_InvalidUTF8Error
    Ctor["InvalidUnmarshalError"] = Shadow_NewStruct_InvalidUnmarshalError
    Pkg["Marshal"] = json.Marshal
    Pkg["MarshalIndent"] = json.MarshalIndent
    Pkg["Marshaler"] = Shadow_InterfaceConvertTo2_Marshaler
    Ctor["MarshalerError"] = Shadow_NewStruct_MarshalerError
    Pkg["NewDecoder"] = json.NewDecoder
    Pkg["NewEncoder"] = json.NewEncoder
    Ctor["SyntaxError"] = Shadow_NewStruct_SyntaxError
    Pkg["Token"] = Shadow_InterfaceConvertTo2_Token
    Pkg["Unmarshal"] = json.Unmarshal
    Ctor["UnmarshalFieldError"] = Shadow_NewStruct_UnmarshalFieldError
    Ctor["UnmarshalTypeError"] = Shadow_NewStruct_UnmarshalTypeError
    Pkg["Unmarshaler"] = Shadow_InterfaceConvertTo2_Unmarshaler
    Ctor["UnsupportedTypeError"] = Shadow_NewStruct_UnsupportedTypeError
    Ctor["UnsupportedValueError"] = Shadow_NewStruct_UnsupportedValueError
    Pkg["Valid"] = json.Valid

}
func Shadow_NewStruct_Decoder(src *json.Decoder) *json.Decoder {
    if src == nil {
	   return &json.Decoder{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Encoder(src *json.Encoder) *json.Encoder {
    if src == nil {
	   return &json.Encoder{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_InvalidUTF8Error(src *json.InvalidUTF8Error) *json.InvalidUTF8Error {
    if src == nil {
	   return &json.InvalidUTF8Error{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_InvalidUnmarshalError(src *json.InvalidUnmarshalError) *json.InvalidUnmarshalError {
    if src == nil {
	   return &json.InvalidUnmarshalError{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_Marshaler(x interface{}) (y json.Marshaler, b bool) {
	y, b = x.(json.Marshaler)
	return
}

func Shadow_InterfaceConvertTo1_Marshaler(x interface{}) json.Marshaler {
	return x.(json.Marshaler)
}


func Shadow_NewStruct_MarshalerError(src *json.MarshalerError) *json.MarshalerError {
    if src == nil {
	   return &json.MarshalerError{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_SyntaxError(src *json.SyntaxError) *json.SyntaxError {
    if src == nil {
	   return &json.SyntaxError{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_Token(x interface{}) (y json.Token, b bool) {
	y, b = x.(json.Token)
	return
}

func Shadow_InterfaceConvertTo1_Token(x interface{}) json.Token {
	return x.(json.Token)
}


func Shadow_NewStruct_UnmarshalFieldError(src *json.UnmarshalFieldError) *json.UnmarshalFieldError {
    if src == nil {
	   return &json.UnmarshalFieldError{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_UnmarshalTypeError(src *json.UnmarshalTypeError) *json.UnmarshalTypeError {
    if src == nil {
	   return &json.UnmarshalTypeError{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_Unmarshaler(x interface{}) (y json.Unmarshaler, b bool) {
	y, b = x.(json.Unmarshaler)
	return
}

func Shadow_InterfaceConvertTo1_Unmarshaler(x interface{}) json.Unmarshaler {
	return x.(json.Unmarshaler)
}


func Shadow_NewStruct_UnsupportedTypeError(src *json.UnsupportedTypeError) *json.UnsupportedTypeError {
    if src == nil {
	   return &json.UnsupportedTypeError{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_UnsupportedValueError(src *json.UnsupportedValueError) *json.UnsupportedValueError {
    if src == nil {
	   return &json.UnsupportedValueError{}
    }
    a := *src
    return &a
}

