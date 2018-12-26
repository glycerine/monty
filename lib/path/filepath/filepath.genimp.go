package shadow_filepath

import "path/filepath"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["Abs"] = filepath.Abs
    Pkg["Base"] = filepath.Base
    Pkg["Clean"] = filepath.Clean
    Pkg["Dir"] = filepath.Dir
    Pkg["ErrBadPattern"] = filepath.ErrBadPattern
    Pkg["EvalSymlinks"] = filepath.EvalSymlinks
    Pkg["Ext"] = filepath.Ext
    Pkg["FromSlash"] = filepath.FromSlash
    Pkg["Glob"] = filepath.Glob
    Pkg["HasPrefix"] = filepath.HasPrefix
    Pkg["IsAbs"] = filepath.IsAbs
    Pkg["Join"] = filepath.Join
    Pkg["ListSeparator"] = filepath.ListSeparator
    Pkg["Match"] = filepath.Match
    Pkg["Rel"] = filepath.Rel
    Pkg["Separator"] = filepath.Separator
    Pkg["SkipDir"] = filepath.SkipDir
    Pkg["Split"] = filepath.Split
    Pkg["SplitList"] = filepath.SplitList
    Pkg["ToSlash"] = filepath.ToSlash
    Pkg["VolumeName"] = filepath.VolumeName
    Pkg["Walk"] = filepath.Walk

}