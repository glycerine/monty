package shadow_lzw

import "compress/lzw"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["NewReader"] = lzw.NewReader
    Pkg["NewWriter"] = lzw.NewWriter

}