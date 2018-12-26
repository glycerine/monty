package shadow_zip

import "archive/zip"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["Deflate"] = zip.Deflate
    Pkg["ErrAlgorithm"] = zip.ErrAlgorithm
    Pkg["ErrChecksum"] = zip.ErrChecksum
    Pkg["ErrFormat"] = zip.ErrFormat
    Ctor["File"] = Shadow_NewStruct_File
    Ctor["FileHeader"] = Shadow_NewStruct_FileHeader
    Pkg["FileInfoHeader"] = zip.FileInfoHeader
    Pkg["NewReader"] = zip.NewReader
    Pkg["NewWriter"] = zip.NewWriter
    Pkg["OpenReader"] = zip.OpenReader
    Ctor["ReadCloser"] = Shadow_NewStruct_ReadCloser
    Ctor["Reader"] = Shadow_NewStruct_Reader
    Pkg["RegisterCompressor"] = zip.RegisterCompressor
    Pkg["RegisterDecompressor"] = zip.RegisterDecompressor
    Pkg["Store"] = zip.Store
    Ctor["Writer"] = Shadow_NewStruct_Writer

}
func Shadow_NewStruct_File(src *zip.File) *zip.File {
    if src == nil {
	   return &zip.File{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_FileHeader(src *zip.FileHeader) *zip.FileHeader {
    if src == nil {
	   return &zip.FileHeader{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_ReadCloser(src *zip.ReadCloser) *zip.ReadCloser {
    if src == nil {
	   return &zip.ReadCloser{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Reader(src *zip.Reader) *zip.Reader {
    if src == nil {
	   return &zip.Reader{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Writer(src *zip.Writer) *zip.Writer {
    if src == nil {
	   return &zip.Writer{}
    }
    a := *src
    return &a
}

