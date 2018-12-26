package shadow_bufio

import "bufio"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["ErrAdvanceTooFar"] = bufio.ErrAdvanceTooFar
    Pkg["ErrBufferFull"] = bufio.ErrBufferFull
    Pkg["ErrFinalToken"] = bufio.ErrFinalToken
    Pkg["ErrInvalidUnreadByte"] = bufio.ErrInvalidUnreadByte
    Pkg["ErrInvalidUnreadRune"] = bufio.ErrInvalidUnreadRune
    Pkg["ErrNegativeAdvance"] = bufio.ErrNegativeAdvance
    Pkg["ErrNegativeCount"] = bufio.ErrNegativeCount
    Pkg["ErrTooLong"] = bufio.ErrTooLong
    Pkg["MaxScanTokenSize"] = bufio.MaxScanTokenSize
    Pkg["NewReadWriter"] = bufio.NewReadWriter
    Pkg["NewReader"] = bufio.NewReader
    Pkg["NewReaderSize"] = bufio.NewReaderSize
    Pkg["NewScanner"] = bufio.NewScanner
    Pkg["NewWriter"] = bufio.NewWriter
    Pkg["NewWriterSize"] = bufio.NewWriterSize
    Ctor["ReadWriter"] = Shadow_NewStruct_ReadWriter
    Ctor["Reader"] = Shadow_NewStruct_Reader
    Pkg["ScanBytes"] = bufio.ScanBytes
    Pkg["ScanLines"] = bufio.ScanLines
    Pkg["ScanRunes"] = bufio.ScanRunes
    Pkg["ScanWords"] = bufio.ScanWords
    Ctor["Scanner"] = Shadow_NewStruct_Scanner
    Ctor["Writer"] = Shadow_NewStruct_Writer

}
func Shadow_NewStruct_ReadWriter(src *bufio.ReadWriter) *bufio.ReadWriter {
    if src == nil {
	   return &bufio.ReadWriter{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Reader(src *bufio.Reader) *bufio.Reader {
    if src == nil {
	   return &bufio.Reader{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Scanner(src *bufio.Scanner) *bufio.Scanner {
    if src == nil {
	   return &bufio.Scanner{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Writer(src *bufio.Writer) *bufio.Writer {
    if src == nil {
	   return &bufio.Writer{}
    }
    a := *src
    return &a
}

