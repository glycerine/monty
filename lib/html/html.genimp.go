package shadow_html

import "html"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["EscapeString"] = html.EscapeString
    Pkg["UnescapeString"] = html.UnescapeString

}