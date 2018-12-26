package shadow_csv

import "encoding/csv"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["ErrBareQuote"] = csv.ErrBareQuote
    Pkg["ErrFieldCount"] = csv.ErrFieldCount
    Pkg["ErrQuote"] = csv.ErrQuote
    Pkg["ErrTrailingComma"] = csv.ErrTrailingComma
    Pkg["NewReader"] = csv.NewReader
    Pkg["NewWriter"] = csv.NewWriter
    Ctor["ParseError"] = Shadow_NewStruct_ParseError
    Ctor["Reader"] = Shadow_NewStruct_Reader
    Ctor["Writer"] = Shadow_NewStruct_Writer

}
func Shadow_NewStruct_ParseError(src *csv.ParseError) *csv.ParseError {
    if src == nil {
	   return &csv.ParseError{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Reader(src *csv.Reader) *csv.Reader {
    if src == nil {
	   return &csv.Reader{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Writer(src *csv.Writer) *csv.Writer {
    if src == nil {
	   return &csv.Writer{}
    }
    a := *src
    return &a
}

