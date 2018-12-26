package shadow_encoding

import "encoding"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["BinaryMarshaler"] = Shadow_InterfaceConvertTo2_BinaryMarshaler
    Pkg["BinaryUnmarshaler"] = Shadow_InterfaceConvertTo2_BinaryUnmarshaler
    Pkg["TextMarshaler"] = Shadow_InterfaceConvertTo2_TextMarshaler
    Pkg["TextUnmarshaler"] = Shadow_InterfaceConvertTo2_TextUnmarshaler

}
func Shadow_InterfaceConvertTo2_BinaryMarshaler(x interface{}) (y encoding.BinaryMarshaler, b bool) {
	y, b = x.(encoding.BinaryMarshaler)
	return
}

func Shadow_InterfaceConvertTo1_BinaryMarshaler(x interface{}) encoding.BinaryMarshaler {
	return x.(encoding.BinaryMarshaler)
}


func Shadow_InterfaceConvertTo2_BinaryUnmarshaler(x interface{}) (y encoding.BinaryUnmarshaler, b bool) {
	y, b = x.(encoding.BinaryUnmarshaler)
	return
}

func Shadow_InterfaceConvertTo1_BinaryUnmarshaler(x interface{}) encoding.BinaryUnmarshaler {
	return x.(encoding.BinaryUnmarshaler)
}


func Shadow_InterfaceConvertTo2_TextMarshaler(x interface{}) (y encoding.TextMarshaler, b bool) {
	y, b = x.(encoding.TextMarshaler)
	return
}

func Shadow_InterfaceConvertTo1_TextMarshaler(x interface{}) encoding.TextMarshaler {
	return x.(encoding.TextMarshaler)
}


func Shadow_InterfaceConvertTo2_TextUnmarshaler(x interface{}) (y encoding.TextUnmarshaler, b bool) {
	y, b = x.(encoding.TextUnmarshaler)
	return
}

func Shadow_InterfaceConvertTo1_TextUnmarshaler(x interface{}) encoding.TextUnmarshaler {
	return x.(encoding.TextUnmarshaler)
}

