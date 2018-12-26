package shadow_regexp

import "regexp"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["Compile"] = regexp.Compile
    Pkg["CompilePOSIX"] = regexp.CompilePOSIX
    Pkg["Match"] = regexp.Match
    Pkg["MatchReader"] = regexp.MatchReader
    Pkg["MatchString"] = regexp.MatchString
    Pkg["MustCompile"] = regexp.MustCompile
    Pkg["MustCompilePOSIX"] = regexp.MustCompilePOSIX
    Pkg["QuoteMeta"] = regexp.QuoteMeta
    Ctor["Regexp"] = Shadow_NewStruct_Regexp

}
func Shadow_NewStruct_Regexp(src *regexp.Regexp) *regexp.Regexp {
    if src == nil {
	   return &regexp.Regexp{}
    }
    a := *src
    return &a
}

