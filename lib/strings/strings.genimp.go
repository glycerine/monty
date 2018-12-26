package shadow_strings

import "strings"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Ctor["Builder"] = Shadow_NewStruct_Builder
    Pkg["Compare"] = strings.Compare
    Pkg["Contains"] = strings.Contains
    Pkg["ContainsAny"] = strings.ContainsAny
    Pkg["ContainsRune"] = strings.ContainsRune
    Pkg["Count"] = strings.Count
    Pkg["EqualFold"] = strings.EqualFold
    Pkg["Fields"] = strings.Fields
    Pkg["FieldsFunc"] = strings.FieldsFunc
    Pkg["HasPrefix"] = strings.HasPrefix
    Pkg["HasSuffix"] = strings.HasSuffix
    Pkg["Index"] = strings.Index
    Pkg["IndexAny"] = strings.IndexAny
    Pkg["IndexByte"] = strings.IndexByte
    Pkg["IndexFunc"] = strings.IndexFunc
    Pkg["IndexRune"] = strings.IndexRune
    Pkg["Join"] = strings.Join
    Pkg["LastIndex"] = strings.LastIndex
    Pkg["LastIndexAny"] = strings.LastIndexAny
    Pkg["LastIndexByte"] = strings.LastIndexByte
    Pkg["LastIndexFunc"] = strings.LastIndexFunc
    Pkg["Map"] = strings.Map
    Pkg["NewReader"] = strings.NewReader
    Pkg["NewReplacer"] = strings.NewReplacer
    Ctor["Reader"] = Shadow_NewStruct_Reader
    Pkg["Repeat"] = strings.Repeat
    Pkg["Replace"] = strings.Replace
    Ctor["Replacer"] = Shadow_NewStruct_Replacer
    Pkg["Split"] = strings.Split
    Pkg["SplitAfter"] = strings.SplitAfter
    Pkg["SplitAfterN"] = strings.SplitAfterN
    Pkg["SplitN"] = strings.SplitN
    Pkg["Title"] = strings.Title
    Pkg["ToLower"] = strings.ToLower
    Pkg["ToLowerSpecial"] = strings.ToLowerSpecial
    Pkg["ToTitle"] = strings.ToTitle
    Pkg["ToTitleSpecial"] = strings.ToTitleSpecial
    Pkg["ToUpper"] = strings.ToUpper
    Pkg["ToUpperSpecial"] = strings.ToUpperSpecial
    Pkg["Trim"] = strings.Trim
    Pkg["TrimFunc"] = strings.TrimFunc
    Pkg["TrimLeft"] = strings.TrimLeft
    Pkg["TrimLeftFunc"] = strings.TrimLeftFunc
    Pkg["TrimPrefix"] = strings.TrimPrefix
    Pkg["TrimRight"] = strings.TrimRight
    Pkg["TrimRightFunc"] = strings.TrimRightFunc
    Pkg["TrimSpace"] = strings.TrimSpace
    Pkg["TrimSuffix"] = strings.TrimSuffix

}
func Shadow_NewStruct_Builder(src *strings.Builder) *strings.Builder {
    if src == nil {
	   return &strings.Builder{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Reader(src *strings.Reader) *strings.Reader {
    if src == nil {
	   return &strings.Reader{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Replacer(src *strings.Replacer) *strings.Replacer {
    if src == nil {
	   return &strings.Replacer{}
    }
    a := *src
    return &a
}

