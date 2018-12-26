package shadow_base64

import "encoding/base64"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Ctor["Encoding"] = Shadow_NewStruct_Encoding
    Pkg["NewDecoder"] = base64.NewDecoder
    Pkg["NewEncoder"] = base64.NewEncoder
    Pkg["NewEncoding"] = base64.NewEncoding
    Pkg["NoPadding"] = base64.NoPadding
    Pkg["RawStdEncoding"] = base64.RawStdEncoding
    Pkg["RawURLEncoding"] = base64.RawURLEncoding
    Pkg["StdEncoding"] = base64.StdEncoding
    Pkg["StdPadding"] = base64.StdPadding
    Pkg["URLEncoding"] = base64.URLEncoding

}
func Shadow_NewStruct_Encoding(src *base64.Encoding) *base64.Encoding {
    if src == nil {
	   return &base64.Encoding{}
    }
    a := *src
    return &a
}

