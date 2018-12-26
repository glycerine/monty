package shadow_path

import "path"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["Base"] = path.Base
    Pkg["Clean"] = path.Clean
    Pkg["Dir"] = path.Dir
    Pkg["ErrBadPattern"] = path.ErrBadPattern
    Pkg["Ext"] = path.Ext
    Pkg["IsAbs"] = path.IsAbs
    Pkg["Join"] = path.Join
    Pkg["Match"] = path.Match
    Pkg["Split"] = path.Split

}