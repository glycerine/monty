package shadow_flate

import "compress/flate"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["BestCompression"] = flate.BestCompression
    Pkg["BestSpeed"] = flate.BestSpeed
    Pkg["DefaultCompression"] = flate.DefaultCompression
    Pkg["HuffmanOnly"] = flate.HuffmanOnly
    Pkg["NewReader"] = flate.NewReader
    Pkg["NewReaderDict"] = flate.NewReaderDict
    Pkg["NewWriter"] = flate.NewWriter
    Pkg["NewWriterDict"] = flate.NewWriterDict
    Pkg["NoCompression"] = flate.NoCompression
    Ctor["ReadError"] = Shadow_NewStruct_ReadError
    Pkg["Reader"] = Shadow_InterfaceConvertTo2_Reader
    Pkg["Resetter"] = Shadow_InterfaceConvertTo2_Resetter
    Ctor["WriteError"] = Shadow_NewStruct_WriteError
    Ctor["Writer"] = Shadow_NewStruct_Writer

}
func Shadow_NewStruct_ReadError(src *flate.ReadError) *flate.ReadError {
    if src == nil {
	   return &flate.ReadError{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_Reader(x interface{}) (y flate.Reader, b bool) {
	y, b = x.(flate.Reader)
	return
}

func Shadow_InterfaceConvertTo1_Reader(x interface{}) flate.Reader {
	return x.(flate.Reader)
}


func Shadow_InterfaceConvertTo2_Resetter(x interface{}) (y flate.Resetter, b bool) {
	y, b = x.(flate.Resetter)
	return
}

func Shadow_InterfaceConvertTo1_Resetter(x interface{}) flate.Resetter {
	return x.(flate.Resetter)
}


func Shadow_NewStruct_WriteError(src *flate.WriteError) *flate.WriteError {
    if src == nil {
	   return &flate.WriteError{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Writer(src *flate.Writer) *flate.Writer {
    if src == nil {
	   return &flate.Writer{}
    }
    a := *src
    return &a
}

