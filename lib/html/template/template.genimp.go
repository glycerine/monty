package shadow_template

import "html/template"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Ctor["Error"] = Shadow_NewStruct_Error
    Pkg["HTMLEscape"] = template.HTMLEscape
    Pkg["HTMLEscapeString"] = template.HTMLEscapeString
    Pkg["HTMLEscaper"] = template.HTMLEscaper
    Pkg["IsTrue"] = template.IsTrue
    Pkg["JSEscape"] = template.JSEscape
    Pkg["JSEscapeString"] = template.JSEscapeString
    Pkg["JSEscaper"] = template.JSEscaper
    Pkg["Must"] = template.Must
    Pkg["New"] = template.New
    Pkg["ParseFiles"] = template.ParseFiles
    Pkg["ParseGlob"] = template.ParseGlob
    Ctor["Template"] = Shadow_NewStruct_Template
    Pkg["URLQueryEscaper"] = template.URLQueryEscaper

}
func Shadow_NewStruct_Error(src *template.Error) *template.Error {
    if src == nil {
	   return &template.Error{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Template(src *template.Template) *template.Template {
    if src == nil {
	   return &template.Template{}
    }
    a := *src
    return &a
}

