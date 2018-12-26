package shadow_exec

import "os/exec"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Ctor["Cmd"] = Shadow_NewStruct_Cmd
    Pkg["Command"] = exec.Command
    Pkg["CommandContext"] = exec.CommandContext
    Pkg["ErrNotFound"] = exec.ErrNotFound
    Ctor["Error"] = Shadow_NewStruct_Error
    Ctor["ExitError"] = Shadow_NewStruct_ExitError
    Pkg["LookPath"] = exec.LookPath

}
func Shadow_NewStruct_Cmd(src *exec.Cmd) *exec.Cmd {
    if src == nil {
	   return &exec.Cmd{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Error(src *exec.Error) *exec.Error {
    if src == nil {
	   return &exec.Error{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_ExitError(src *exec.ExitError) *exec.ExitError {
    if src == nil {
	   return &exec.ExitError{}
    }
    a := *src
    return &a
}

