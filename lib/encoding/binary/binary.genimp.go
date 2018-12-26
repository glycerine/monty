package shadow_binary

import "encoding/binary"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["BigEndian"] = binary.BigEndian
    Pkg["ByteOrder"] = Shadow_InterfaceConvertTo2_ByteOrder
    Pkg["LittleEndian"] = binary.LittleEndian
    Pkg["MaxVarintLen16"] = binary.MaxVarintLen16
    Pkg["MaxVarintLen32"] = binary.MaxVarintLen32
    Pkg["MaxVarintLen64"] = binary.MaxVarintLen64
    Pkg["PutUvarint"] = binary.PutUvarint
    Pkg["PutVarint"] = binary.PutVarint
    Pkg["Read"] = binary.Read
    Pkg["ReadUvarint"] = binary.ReadUvarint
    Pkg["ReadVarint"] = binary.ReadVarint
    Pkg["Size"] = binary.Size
    Pkg["Uvarint"] = binary.Uvarint
    Pkg["Varint"] = binary.Varint
    Pkg["Write"] = binary.Write

}
func Shadow_InterfaceConvertTo2_ByteOrder(x interface{}) (y binary.ByteOrder, b bool) {
	y, b = x.(binary.ByteOrder)
	return
}

func Shadow_InterfaceConvertTo1_ByteOrder(x interface{}) binary.ByteOrder {
	return x.(binary.ByteOrder)
}

