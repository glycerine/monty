package shadow_runtime

import "runtime"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["BlockProfile"] = runtime.BlockProfile
    Ctor["BlockProfileRecord"] = Shadow_NewStruct_BlockProfileRecord
    Pkg["Breakpoint"] = runtime.Breakpoint
    Pkg["CPUProfile"] = runtime.CPUProfile
    Pkg["Caller"] = runtime.Caller
    Pkg["Callers"] = runtime.Callers
    Pkg["CallersFrames"] = runtime.CallersFrames
    Pkg["Compiler"] = runtime.Compiler
    Pkg["Error"] = Shadow_InterfaceConvertTo2_Error
    Ctor["Frame"] = Shadow_NewStruct_Frame
    Ctor["Frames"] = Shadow_NewStruct_Frames
    Ctor["Func"] = Shadow_NewStruct_Func
    Pkg["FuncForPC"] = runtime.FuncForPC
    Pkg["GC"] = runtime.GC
    Pkg["GOARCH"] = runtime.GOARCH
    Pkg["GOMAXPROCS"] = runtime.GOMAXPROCS
    Pkg["GOOS"] = runtime.GOOS
    Pkg["GOROOT"] = runtime.GOROOT
    Pkg["Goexit"] = runtime.Goexit
    Pkg["GoroutineProfile"] = runtime.GoroutineProfile
    Pkg["Gosched"] = runtime.Gosched
    Pkg["KeepAlive"] = runtime.KeepAlive
    Pkg["LockOSThread"] = runtime.LockOSThread
    Pkg["MemProfile"] = runtime.MemProfile
    Pkg["MemProfileRate"] = runtime.MemProfileRate
    Ctor["MemProfileRecord"] = Shadow_NewStruct_MemProfileRecord
    Ctor["MemStats"] = Shadow_NewStruct_MemStats
    Pkg["MutexProfile"] = runtime.MutexProfile
    Pkg["NumCPU"] = runtime.NumCPU
    Pkg["NumCgoCall"] = runtime.NumCgoCall
    Pkg["NumGoroutine"] = runtime.NumGoroutine
    Pkg["ReadMemStats"] = runtime.ReadMemStats
    Pkg["ReadTrace"] = runtime.ReadTrace
    Pkg["SetBlockProfileRate"] = runtime.SetBlockProfileRate
    Pkg["SetCPUProfileRate"] = runtime.SetCPUProfileRate
    Pkg["SetCgoTraceback"] = runtime.SetCgoTraceback
    Pkg["SetFinalizer"] = runtime.SetFinalizer
    Pkg["SetMutexProfileFraction"] = runtime.SetMutexProfileFraction
    Pkg["Stack"] = runtime.Stack
    Ctor["StackRecord"] = Shadow_NewStruct_StackRecord
    Pkg["StartTrace"] = runtime.StartTrace
    Pkg["StopTrace"] = runtime.StopTrace
    Pkg["ThreadCreateProfile"] = runtime.ThreadCreateProfile
    Ctor["TypeAssertionError"] = Shadow_NewStruct_TypeAssertionError
    Pkg["UnlockOSThread"] = runtime.UnlockOSThread
    Pkg["Version"] = runtime.Version

}
func Shadow_NewStruct_BlockProfileRecord(src *runtime.BlockProfileRecord) *runtime.BlockProfileRecord {
    if src == nil {
	   return &runtime.BlockProfileRecord{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_Error(x interface{}) (y runtime.Error, b bool) {
	y, b = x.(runtime.Error)
	return
}

func Shadow_InterfaceConvertTo1_Error(x interface{}) runtime.Error {
	return x.(runtime.Error)
}


func Shadow_NewStruct_Frame(src *runtime.Frame) *runtime.Frame {
    if src == nil {
	   return &runtime.Frame{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Frames(src *runtime.Frames) *runtime.Frames {
    if src == nil {
	   return &runtime.Frames{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Func(src *runtime.Func) *runtime.Func {
    if src == nil {
	   return &runtime.Func{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_MemProfileRecord(src *runtime.MemProfileRecord) *runtime.MemProfileRecord {
    if src == nil {
	   return &runtime.MemProfileRecord{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_MemStats(src *runtime.MemStats) *runtime.MemStats {
    if src == nil {
	   return &runtime.MemStats{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_StackRecord(src *runtime.StackRecord) *runtime.StackRecord {
    if src == nil {
	   return &runtime.StackRecord{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_TypeAssertionError(src *runtime.TypeAssertionError) *runtime.TypeAssertionError {
    if src == nil {
	   return &runtime.TypeAssertionError{}
    }
    a := *src
    return &a
}

