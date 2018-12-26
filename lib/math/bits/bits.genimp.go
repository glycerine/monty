package shadow_bits

import "math/bits"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["LeadingZeros"] = bits.LeadingZeros
    Pkg["LeadingZeros16"] = bits.LeadingZeros16
    Pkg["LeadingZeros32"] = bits.LeadingZeros32
    Pkg["LeadingZeros64"] = bits.LeadingZeros64
    Pkg["LeadingZeros8"] = bits.LeadingZeros8
    Pkg["Len"] = bits.Len
    Pkg["Len16"] = bits.Len16
    Pkg["Len32"] = bits.Len32
    Pkg["Len64"] = bits.Len64
    Pkg["Len8"] = bits.Len8
    Pkg["OnesCount"] = bits.OnesCount
    Pkg["OnesCount16"] = bits.OnesCount16
    Pkg["OnesCount32"] = bits.OnesCount32
    Pkg["OnesCount64"] = bits.OnesCount64
    Pkg["OnesCount8"] = bits.OnesCount8
    Pkg["Reverse"] = bits.Reverse
    Pkg["Reverse16"] = bits.Reverse16
    Pkg["Reverse32"] = bits.Reverse32
    Pkg["Reverse64"] = bits.Reverse64
    Pkg["Reverse8"] = bits.Reverse8
    Pkg["ReverseBytes"] = bits.ReverseBytes
    Pkg["ReverseBytes16"] = bits.ReverseBytes16
    Pkg["ReverseBytes32"] = bits.ReverseBytes32
    Pkg["ReverseBytes64"] = bits.ReverseBytes64
    Pkg["RotateLeft"] = bits.RotateLeft
    Pkg["RotateLeft16"] = bits.RotateLeft16
    Pkg["RotateLeft32"] = bits.RotateLeft32
    Pkg["RotateLeft64"] = bits.RotateLeft64
    Pkg["RotateLeft8"] = bits.RotateLeft8
    Pkg["TrailingZeros"] = bits.TrailingZeros
    Pkg["TrailingZeros16"] = bits.TrailingZeros16
    Pkg["TrailingZeros32"] = bits.TrailingZeros32
    Pkg["TrailingZeros64"] = bits.TrailingZeros64
    Pkg["TrailingZeros8"] = bits.TrailingZeros8
    Pkg["UintSize"] = bits.UintSize

}