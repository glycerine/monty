package shadow_context

import "context"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["Background"] = context.Background
    Pkg["Canceled"] = context.Canceled
    Pkg["Context"] = Shadow_InterfaceConvertTo2_Context
    Pkg["DeadlineExceeded"] = context.DeadlineExceeded
    Pkg["TODO"] = context.TODO
    Pkg["WithCancel"] = context.WithCancel
    Pkg["WithDeadline"] = context.WithDeadline
    Pkg["WithTimeout"] = context.WithTimeout
    Pkg["WithValue"] = context.WithValue

}
func Shadow_InterfaceConvertTo2_Context(x interface{}) (y context.Context, b bool) {
	y, b = x.(context.Context)
	return
}

func Shadow_InterfaceConvertTo1_Context(x interface{}) context.Context {
	return x.(context.Context)
}

