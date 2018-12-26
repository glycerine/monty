package shadow_debug

import "runtime/debug"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["FreeOSMemory"] = debug.FreeOSMemory
    Ctor["GCStats"] = Shadow_NewStruct_GCStats
    Pkg["PrintStack"] = debug.PrintStack
    Pkg["ReadGCStats"] = debug.ReadGCStats
    Pkg["SetGCPercent"] = debug.SetGCPercent
    Pkg["SetMaxStack"] = debug.SetMaxStack
    Pkg["SetMaxThreads"] = debug.SetMaxThreads
    Pkg["SetPanicOnFault"] = debug.SetPanicOnFault
    Pkg["SetTraceback"] = debug.SetTraceback
    Pkg["Stack"] = debug.Stack
    Pkg["WriteHeapDump"] = debug.WriteHeapDump

}
func Shadow_NewStruct_GCStats(src *debug.GCStats) *debug.GCStats {
    if src == nil {
	   return &debug.GCStats{}
    }
    a := *src
    return &a
}

