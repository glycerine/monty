package shadow_zlib

import "compress/zlib"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["BestCompression"] = zlib.BestCompression
    Pkg["BestSpeed"] = zlib.BestSpeed
    Pkg["DefaultCompression"] = zlib.DefaultCompression
    Pkg["ErrChecksum"] = zlib.ErrChecksum
    Pkg["ErrDictionary"] = zlib.ErrDictionary
    Pkg["ErrHeader"] = zlib.ErrHeader
    Pkg["HuffmanOnly"] = zlib.HuffmanOnly
    Pkg["NewReader"] = zlib.NewReader
    Pkg["NewReaderDict"] = zlib.NewReaderDict
    Pkg["NewWriter"] = zlib.NewWriter
    Pkg["NewWriterLevel"] = zlib.NewWriterLevel
    Pkg["NewWriterLevelDict"] = zlib.NewWriterLevelDict
    Pkg["NoCompression"] = zlib.NoCompression
    Pkg["Resetter"] = Shadow_InterfaceConvertTo2_Resetter
    Ctor["Writer"] = Shadow_NewStruct_Writer

}
func Shadow_InterfaceConvertTo2_Resetter(x interface{}) (y zlib.Resetter, b bool) {
	y, b = x.(zlib.Resetter)
	return
}

func Shadow_InterfaceConvertTo1_Resetter(x interface{}) zlib.Resetter {
	return x.(zlib.Resetter)
}


func Shadow_NewStruct_Writer(src *zlib.Writer) *zlib.Writer {
    if src == nil {
	   return &zlib.Writer{}
    }
    a := *src
    return &a
}

