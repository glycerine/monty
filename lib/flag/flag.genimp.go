package shadow_flag

import "flag"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["Arg"] = flag.Arg
    Pkg["Args"] = flag.Args
    Pkg["Bool"] = flag.Bool
    Pkg["BoolVar"] = flag.BoolVar
    Pkg["CommandLine"] = flag.CommandLine
    Pkg["Duration"] = flag.Duration
    Pkg["DurationVar"] = flag.DurationVar
    Pkg["ErrHelp"] = flag.ErrHelp
    Ctor["Flag"] = Shadow_NewStruct_Flag
    Ctor["FlagSet"] = Shadow_NewStruct_FlagSet
    Pkg["Float64"] = flag.Float64
    Pkg["Float64Var"] = flag.Float64Var
    Pkg["Getter"] = Shadow_InterfaceConvertTo2_Getter
    Pkg["Int"] = flag.Int
    Pkg["Int64"] = flag.Int64
    Pkg["Int64Var"] = flag.Int64Var
    Pkg["IntVar"] = flag.IntVar
    Pkg["Lookup"] = flag.Lookup
    Pkg["NArg"] = flag.NArg
    Pkg["NFlag"] = flag.NFlag
    Pkg["NewFlagSet"] = flag.NewFlagSet
    Pkg["Parse"] = flag.Parse
    Pkg["Parsed"] = flag.Parsed
    Pkg["PrintDefaults"] = flag.PrintDefaults
    Pkg["Set"] = flag.Set
    Pkg["String"] = flag.String
    Pkg["StringVar"] = flag.StringVar
    Pkg["Uint"] = flag.Uint
    Pkg["Uint64"] = flag.Uint64
    Pkg["Uint64Var"] = flag.Uint64Var
    Pkg["UintVar"] = flag.UintVar
    Pkg["UnquoteUsage"] = flag.UnquoteUsage
    Pkg["Usage"] = flag.Usage
    Pkg["Value"] = Shadow_InterfaceConvertTo2_Value
    Pkg["Var"] = flag.Var
    Pkg["Visit"] = flag.Visit
    Pkg["VisitAll"] = flag.VisitAll

}
func Shadow_NewStruct_Flag(src *flag.Flag) *flag.Flag {
    if src == nil {
	   return &flag.Flag{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_FlagSet(src *flag.FlagSet) *flag.FlagSet {
    if src == nil {
	   return &flag.FlagSet{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_Getter(x interface{}) (y flag.Getter, b bool) {
	y, b = x.(flag.Getter)
	return
}

func Shadow_InterfaceConvertTo1_Getter(x interface{}) flag.Getter {
	return x.(flag.Getter)
}


func Shadow_InterfaceConvertTo2_Value(x interface{}) (y flag.Value, b bool) {
	y, b = x.(flag.Value)
	return
}

func Shadow_InterfaceConvertTo1_Value(x interface{}) flag.Value {
	return x.(flag.Value)
}

