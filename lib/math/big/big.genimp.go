package shadow_big

import "math/big"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Ctor["ErrNaN"] = Shadow_NewStruct_ErrNaN
    Ctor["Float"] = Shadow_NewStruct_Float
    Ctor["Int"] = Shadow_NewStruct_Int
    Pkg["Jacobi"] = big.Jacobi
    Pkg["MaxBase"] = big.MaxBase
    Pkg["MaxExp"] = big.MaxExp
    Pkg["MaxPrec"] = big.MaxPrec
    Pkg["MinExp"] = big.MinExp
    Pkg["NewFloat"] = big.NewFloat
    Pkg["NewInt"] = big.NewInt
    Pkg["NewRat"] = big.NewRat
    Pkg["ParseFloat"] = big.ParseFloat
    Ctor["Rat"] = Shadow_NewStruct_Rat

}
func Shadow_NewStruct_ErrNaN(src *big.ErrNaN) *big.ErrNaN {
    if src == nil {
	   return &big.ErrNaN{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Float(src *big.Float) *big.Float {
    if src == nil {
	   return &big.Float{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Int(src *big.Int) *big.Int {
    if src == nil {
	   return &big.Int{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Rat(src *big.Rat) *big.Rat {
    if src == nil {
	   return &big.Rat{}
    }
    a := *src
    return &a
}

