package shadow_gzip

import "compress/gzip"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["BestCompression"] = gzip.BestCompression
    Pkg["BestSpeed"] = gzip.BestSpeed
    Pkg["DefaultCompression"] = gzip.DefaultCompression
    Pkg["ErrChecksum"] = gzip.ErrChecksum
    Pkg["ErrHeader"] = gzip.ErrHeader
    Ctor["Header"] = Shadow_NewStruct_Header
    Pkg["HuffmanOnly"] = gzip.HuffmanOnly
    Pkg["NewReader"] = gzip.NewReader
    Pkg["NewWriter"] = gzip.NewWriter
    Pkg["NewWriterLevel"] = gzip.NewWriterLevel
    Pkg["NoCompression"] = gzip.NoCompression
    Ctor["Reader"] = Shadow_NewStruct_Reader
    Ctor["Writer"] = Shadow_NewStruct_Writer

}
func Shadow_NewStruct_Header(src *gzip.Header) *gzip.Header {
    if src == nil {
	   return &gzip.Header{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Reader(src *gzip.Reader) *gzip.Reader {
    if src == nil {
	   return &gzip.Reader{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Writer(src *gzip.Writer) *gzip.Writer {
    if src == nil {
	   return &gzip.Writer{}
    }
    a := *src
    return &a
}

