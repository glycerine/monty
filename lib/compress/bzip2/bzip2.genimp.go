package shadow_bzip2

import "compress/bzip2"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["NewReader"] = bzip2.NewReader

}