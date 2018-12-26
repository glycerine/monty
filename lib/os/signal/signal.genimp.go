package shadow_signal

import "os/signal"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["Ignore"] = signal.Ignore
    Pkg["Notify"] = signal.Notify
    Pkg["Reset"] = signal.Reset
    Pkg["Stop"] = signal.Stop

}