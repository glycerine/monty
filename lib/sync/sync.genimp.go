package shadow_sync

import "sync"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Ctor["Cond"] = Shadow_NewStruct_Cond
    Pkg["Locker"] = Shadow_InterfaceConvertTo2_Locker
    Ctor["Map"] = Shadow_NewStruct_Map
    Ctor["Mutex"] = Shadow_NewStruct_Mutex
    Pkg["NewCond"] = sync.NewCond
    Ctor["Once"] = Shadow_NewStruct_Once
    Ctor["Pool"] = Shadow_NewStruct_Pool
    Ctor["RWMutex"] = Shadow_NewStruct_RWMutex
    Ctor["WaitGroup"] = Shadow_NewStruct_WaitGroup

}
func Shadow_NewStruct_Cond(src *sync.Cond) *sync.Cond {
    if src == nil {
	   return &sync.Cond{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_Locker(x interface{}) (y sync.Locker, b bool) {
	y, b = x.(sync.Locker)
	return
}

func Shadow_InterfaceConvertTo1_Locker(x interface{}) sync.Locker {
	return x.(sync.Locker)
}


func Shadow_NewStruct_Map(src *sync.Map) *sync.Map {
    if src == nil {
	   return &sync.Map{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Mutex(src *sync.Mutex) *sync.Mutex {
    if src == nil {
	   return &sync.Mutex{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Once(src *sync.Once) *sync.Once {
    if src == nil {
	   return &sync.Once{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Pool(src *sync.Pool) *sync.Pool {
    if src == nil {
	   return &sync.Pool{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_RWMutex(src *sync.RWMutex) *sync.RWMutex {
    if src == nil {
	   return &sync.RWMutex{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_WaitGroup(src *sync.WaitGroup) *sync.WaitGroup {
    if src == nil {
	   return &sync.WaitGroup{}
    }
    a := *src
    return &a
}

