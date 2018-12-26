package shadow_fmt

import "fmt"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["Errorf"] = fmt.Errorf
    Pkg["Formatter"] = Shadow_InterfaceConvertTo2_Formatter
    Pkg["Fprint"] = fmt.Fprint
    Pkg["Fprintf"] = fmt.Fprintf
    Pkg["Fprintln"] = fmt.Fprintln
    Pkg["Fscan"] = fmt.Fscan
    Pkg["Fscanf"] = fmt.Fscanf
    Pkg["Fscanln"] = fmt.Fscanln
    Pkg["GoStringer"] = Shadow_InterfaceConvertTo2_GoStringer
    Pkg["Print"] = fmt.Print
    Pkg["Printf"] = fmt.Printf
    Pkg["Println"] = fmt.Println
    Pkg["Scan"] = fmt.Scan
    Pkg["ScanState"] = Shadow_InterfaceConvertTo2_ScanState
    Pkg["Scanf"] = fmt.Scanf
    Pkg["Scanln"] = fmt.Scanln
    Pkg["Scanner"] = Shadow_InterfaceConvertTo2_Scanner
    Pkg["Sprint"] = fmt.Sprint
    Pkg["Sprintf"] = fmt.Sprintf
    Pkg["Sprintln"] = fmt.Sprintln
    Pkg["Sscan"] = fmt.Sscan
    Pkg["Sscanf"] = fmt.Sscanf
    Pkg["Sscanln"] = fmt.Sscanln
    Pkg["State"] = Shadow_InterfaceConvertTo2_State
    Pkg["Stringer"] = Shadow_InterfaceConvertTo2_Stringer

}
func Shadow_InterfaceConvertTo2_Formatter(x interface{}) (y fmt.Formatter, b bool) {
	y, b = x.(fmt.Formatter)
	return
}

func Shadow_InterfaceConvertTo1_Formatter(x interface{}) fmt.Formatter {
	return x.(fmt.Formatter)
}


func Shadow_InterfaceConvertTo2_GoStringer(x interface{}) (y fmt.GoStringer, b bool) {
	y, b = x.(fmt.GoStringer)
	return
}

func Shadow_InterfaceConvertTo1_GoStringer(x interface{}) fmt.GoStringer {
	return x.(fmt.GoStringer)
}


func Shadow_InterfaceConvertTo2_ScanState(x interface{}) (y fmt.ScanState, b bool) {
	y, b = x.(fmt.ScanState)
	return
}

func Shadow_InterfaceConvertTo1_ScanState(x interface{}) fmt.ScanState {
	return x.(fmt.ScanState)
}


func Shadow_InterfaceConvertTo2_Scanner(x interface{}) (y fmt.Scanner, b bool) {
	y, b = x.(fmt.Scanner)
	return
}

func Shadow_InterfaceConvertTo1_Scanner(x interface{}) fmt.Scanner {
	return x.(fmt.Scanner)
}


func Shadow_InterfaceConvertTo2_State(x interface{}) (y fmt.State, b bool) {
	y, b = x.(fmt.State)
	return
}

func Shadow_InterfaceConvertTo1_State(x interface{}) fmt.State {
	return x.(fmt.State)
}


func Shadow_InterfaceConvertTo2_Stringer(x interface{}) (y fmt.Stringer, b bool) {
	y, b = x.(fmt.Stringer)
	return
}

func Shadow_InterfaceConvertTo1_Stringer(x interface{}) fmt.Stringer {
	return x.(fmt.Stringer)
}

