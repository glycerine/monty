package shadow_cmplx

import "math/cmplx"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["Abs"] = cmplx.Abs
    Pkg["Acos"] = cmplx.Acos
    Pkg["Acosh"] = cmplx.Acosh
    Pkg["Asin"] = cmplx.Asin
    Pkg["Asinh"] = cmplx.Asinh
    Pkg["Atan"] = cmplx.Atan
    Pkg["Atanh"] = cmplx.Atanh
    Pkg["Conj"] = cmplx.Conj
    Pkg["Cos"] = cmplx.Cos
    Pkg["Cosh"] = cmplx.Cosh
    Pkg["Cot"] = cmplx.Cot
    Pkg["Exp"] = cmplx.Exp
    Pkg["Inf"] = cmplx.Inf
    Pkg["IsInf"] = cmplx.IsInf
    Pkg["IsNaN"] = cmplx.IsNaN
    Pkg["Log"] = cmplx.Log
    Pkg["Log10"] = cmplx.Log10
    Pkg["NaN"] = cmplx.NaN
    Pkg["Phase"] = cmplx.Phase
    Pkg["Polar"] = cmplx.Polar
    Pkg["Pow"] = cmplx.Pow
    Pkg["Rect"] = cmplx.Rect
    Pkg["Sin"] = cmplx.Sin
    Pkg["Sinh"] = cmplx.Sinh
    Pkg["Sqrt"] = cmplx.Sqrt
    Pkg["Tan"] = cmplx.Tan
    Pkg["Tanh"] = cmplx.Tanh

}