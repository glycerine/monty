package shadow_os

import "os"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["Args"] = os.Args
    Pkg["Chdir"] = os.Chdir
    Pkg["Chmod"] = os.Chmod
    Pkg["Chown"] = os.Chown
    Pkg["Chtimes"] = os.Chtimes
    Pkg["Clearenv"] = os.Clearenv
    Pkg["Create"] = os.Create
    Pkg["DevNull"] = os.DevNull
    Pkg["Environ"] = os.Environ
    Pkg["ErrClosed"] = os.ErrClosed
    Pkg["ErrExist"] = os.ErrExist
    Pkg["ErrInvalid"] = os.ErrInvalid
    Pkg["ErrNoDeadline"] = os.ErrNoDeadline
    Pkg["ErrNotExist"] = os.ErrNotExist
    Pkg["ErrPermission"] = os.ErrPermission
    Pkg["Executable"] = os.Executable
    Pkg["Exit"] = os.Exit
    Pkg["Expand"] = os.Expand
    Pkg["ExpandEnv"] = os.ExpandEnv
    Ctor["File"] = Shadow_NewStruct_File
    Pkg["FileInfo"] = Shadow_InterfaceConvertTo2_FileInfo
    Pkg["FindProcess"] = os.FindProcess
    Pkg["Getegid"] = os.Getegid
    Pkg["Getenv"] = os.Getenv
    Pkg["Geteuid"] = os.Geteuid
    Pkg["Getgid"] = os.Getgid
    Pkg["Getgroups"] = os.Getgroups
    Pkg["Getpagesize"] = os.Getpagesize
    Pkg["Getpid"] = os.Getpid
    Pkg["Getppid"] = os.Getppid
    Pkg["Getuid"] = os.Getuid
    Pkg["Getwd"] = os.Getwd
    Pkg["Hostname"] = os.Hostname
    Pkg["Interrupt"] = os.Interrupt
    Pkg["IsExist"] = os.IsExist
    Pkg["IsNotExist"] = os.IsNotExist
    Pkg["IsPathSeparator"] = os.IsPathSeparator
    Pkg["IsPermission"] = os.IsPermission
    Pkg["IsTimeout"] = os.IsTimeout
    Pkg["Kill"] = os.Kill
    Pkg["Lchown"] = os.Lchown
    Pkg["Link"] = os.Link
    Ctor["LinkError"] = Shadow_NewStruct_LinkError
    Pkg["LookupEnv"] = os.LookupEnv
    Pkg["Lstat"] = os.Lstat
    Pkg["Mkdir"] = os.Mkdir
    Pkg["MkdirAll"] = os.MkdirAll
    Pkg["NewFile"] = os.NewFile
    Pkg["NewSyscallError"] = os.NewSyscallError
    Pkg["O_APPEND"] = os.O_APPEND
    Pkg["O_CREATE"] = os.O_CREATE
    Pkg["O_EXCL"] = os.O_EXCL
    Pkg["O_RDONLY"] = os.O_RDONLY
    Pkg["O_RDWR"] = os.O_RDWR
    Pkg["O_SYNC"] = os.O_SYNC
    Pkg["O_TRUNC"] = os.O_TRUNC
    Pkg["O_WRONLY"] = os.O_WRONLY
    Pkg["Open"] = os.Open
    Pkg["OpenFile"] = os.OpenFile
    Ctor["PathError"] = Shadow_NewStruct_PathError
    Pkg["PathListSeparator"] = os.PathListSeparator
    Pkg["PathSeparator"] = os.PathSeparator
    Pkg["Pipe"] = os.Pipe
    Ctor["ProcAttr"] = Shadow_NewStruct_ProcAttr
    Ctor["Process"] = Shadow_NewStruct_Process
    Ctor["ProcessState"] = Shadow_NewStruct_ProcessState
    Pkg["Readlink"] = os.Readlink
    Pkg["Remove"] = os.Remove
    Pkg["RemoveAll"] = os.RemoveAll
    Pkg["Rename"] = os.Rename
    Pkg["SEEK_CUR"] = os.SEEK_CUR
    Pkg["SEEK_END"] = os.SEEK_END
    Pkg["SEEK_SET"] = os.SEEK_SET
    Pkg["SameFile"] = os.SameFile
    Pkg["Setenv"] = os.Setenv
    Pkg["Signal"] = Shadow_InterfaceConvertTo2_Signal
    Pkg["StartProcess"] = os.StartProcess
    Pkg["Stat"] = os.Stat
    Pkg["Stderr"] = os.Stderr
    Pkg["Stdin"] = os.Stdin
    Pkg["Stdout"] = os.Stdout
    Pkg["Symlink"] = os.Symlink
    Ctor["SyscallError"] = Shadow_NewStruct_SyscallError
    Pkg["TempDir"] = os.TempDir
    Pkg["Truncate"] = os.Truncate
    Pkg["Unsetenv"] = os.Unsetenv

}
func Shadow_NewStruct_File(src *os.File) *os.File {
    if src == nil {
	   return &os.File{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_FileInfo(x interface{}) (y os.FileInfo, b bool) {
	y, b = x.(os.FileInfo)
	return
}

func Shadow_InterfaceConvertTo1_FileInfo(x interface{}) os.FileInfo {
	return x.(os.FileInfo)
}


func Shadow_NewStruct_LinkError(src *os.LinkError) *os.LinkError {
    if src == nil {
	   return &os.LinkError{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_PathError(src *os.PathError) *os.PathError {
    if src == nil {
	   return &os.PathError{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_ProcAttr(src *os.ProcAttr) *os.ProcAttr {
    if src == nil {
	   return &os.ProcAttr{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Process(src *os.Process) *os.Process {
    if src == nil {
	   return &os.Process{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_ProcessState(src *os.ProcessState) *os.ProcessState {
    if src == nil {
	   return &os.ProcessState{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_Signal(x interface{}) (y os.Signal, b bool) {
	y, b = x.(os.Signal)
	return
}

func Shadow_InterfaceConvertTo1_Signal(x interface{}) os.Signal {
	return x.(os.Signal)
}


func Shadow_NewStruct_SyscallError(src *os.SyscallError) *os.SyscallError {
    if src == nil {
	   return &os.SyscallError{}
    }
    a := *src
    return &a
}

