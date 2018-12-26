package shadow_tar

import "archive/tar"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["ErrFieldTooLong"] = tar.ErrFieldTooLong
    Pkg["ErrHeader"] = tar.ErrHeader
    Pkg["ErrWriteAfterClose"] = tar.ErrWriteAfterClose
    Pkg["ErrWriteTooLong"] = tar.ErrWriteTooLong
    Pkg["FileInfoHeader"] = tar.FileInfoHeader
    Ctor["Header"] = Shadow_NewStruct_Header
    Pkg["NewReader"] = tar.NewReader
    Pkg["NewWriter"] = tar.NewWriter
    Ctor["Reader"] = Shadow_NewStruct_Reader
    Pkg["TypeBlock"] = tar.TypeBlock
    Pkg["TypeChar"] = tar.TypeChar
    Pkg["TypeCont"] = tar.TypeCont
    Pkg["TypeDir"] = tar.TypeDir
    Pkg["TypeFifo"] = tar.TypeFifo
    Pkg["TypeGNULongLink"] = tar.TypeGNULongLink
    Pkg["TypeGNULongName"] = tar.TypeGNULongName
    Pkg["TypeGNUSparse"] = tar.TypeGNUSparse
    Pkg["TypeLink"] = tar.TypeLink
    Pkg["TypeReg"] = tar.TypeReg
    Pkg["TypeRegA"] = tar.TypeRegA
    Pkg["TypeSymlink"] = tar.TypeSymlink
    Pkg["TypeXGlobalHeader"] = tar.TypeXGlobalHeader
    Pkg["TypeXHeader"] = tar.TypeXHeader
    Ctor["Writer"] = Shadow_NewStruct_Writer

}
func Shadow_NewStruct_Header(src *tar.Header) *tar.Header {
    if src == nil {
	   return &tar.Header{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Reader(src *tar.Reader) *tar.Reader {
    if src == nil {
	   return &tar.Reader{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Writer(src *tar.Writer) *tar.Writer {
    if src == nil {
	   return &tar.Writer{}
    }
    a := *src
    return &a
}

