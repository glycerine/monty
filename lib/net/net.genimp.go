package shadow_net

import "net"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["Addr"] = Shadow_InterfaceConvertTo2_Addr
    Ctor["AddrError"] = Shadow_NewStruct_AddrError
    Pkg["CIDRMask"] = net.CIDRMask
    Pkg["Conn"] = Shadow_InterfaceConvertTo2_Conn
    Ctor["DNSConfigError"] = Shadow_NewStruct_DNSConfigError
    Ctor["DNSError"] = Shadow_NewStruct_DNSError
    Pkg["DefaultResolver"] = net.DefaultResolver
    Pkg["Dial"] = net.Dial
    Pkg["DialIP"] = net.DialIP
    Pkg["DialTCP"] = net.DialTCP
    Pkg["DialTimeout"] = net.DialTimeout
    Pkg["DialUDP"] = net.DialUDP
    Pkg["DialUnix"] = net.DialUnix
    Ctor["Dialer"] = Shadow_NewStruct_Dialer
    Pkg["ErrWriteToConnected"] = net.ErrWriteToConnected
    Pkg["Error"] = Shadow_InterfaceConvertTo2_Error
    Pkg["FileConn"] = net.FileConn
    Pkg["FileListener"] = net.FileListener
    Pkg["FilePacketConn"] = net.FilePacketConn
    Ctor["IPAddr"] = Shadow_NewStruct_IPAddr
    Ctor["IPConn"] = Shadow_NewStruct_IPConn
    Ctor["IPNet"] = Shadow_NewStruct_IPNet
    Pkg["IPv4"] = net.IPv4
    Pkg["IPv4Mask"] = net.IPv4Mask
    Pkg["IPv4len"] = net.IPv4len
    Pkg["IPv6len"] = net.IPv6len
    Ctor["Interface"] = Shadow_NewStruct_Interface
    Pkg["InterfaceAddrs"] = net.InterfaceAddrs
    Pkg["InterfaceByIndex"] = net.InterfaceByIndex
    Pkg["InterfaceByName"] = net.InterfaceByName
    Pkg["Interfaces"] = net.Interfaces
    Pkg["JoinHostPort"] = net.JoinHostPort
    Pkg["Listen"] = net.Listen
    Pkg["ListenIP"] = net.ListenIP
    Pkg["ListenMulticastUDP"] = net.ListenMulticastUDP
    Pkg["ListenPacket"] = net.ListenPacket
    Pkg["ListenTCP"] = net.ListenTCP
    Pkg["ListenUDP"] = net.ListenUDP
    Pkg["ListenUnix"] = net.ListenUnix
    Pkg["ListenUnixgram"] = net.ListenUnixgram
    Pkg["Listener"] = Shadow_InterfaceConvertTo2_Listener
    Pkg["LookupAddr"] = net.LookupAddr
    Pkg["LookupCNAME"] = net.LookupCNAME
    Pkg["LookupHost"] = net.LookupHost
    Pkg["LookupIP"] = net.LookupIP
    Pkg["LookupMX"] = net.LookupMX
    Pkg["LookupNS"] = net.LookupNS
    Pkg["LookupPort"] = net.LookupPort
    Pkg["LookupSRV"] = net.LookupSRV
    Pkg["LookupTXT"] = net.LookupTXT
    Ctor["MX"] = Shadow_NewStruct_MX
    Ctor["NS"] = Shadow_NewStruct_NS
    Ctor["OpError"] = Shadow_NewStruct_OpError
    Pkg["PacketConn"] = Shadow_InterfaceConvertTo2_PacketConn
    Pkg["ParseCIDR"] = net.ParseCIDR
    Ctor["ParseError"] = Shadow_NewStruct_ParseError
    Pkg["ParseIP"] = net.ParseIP
    Pkg["ParseMAC"] = net.ParseMAC
    Pkg["Pipe"] = net.Pipe
    Pkg["ResolveIPAddr"] = net.ResolveIPAddr
    Pkg["ResolveTCPAddr"] = net.ResolveTCPAddr
    Pkg["ResolveUDPAddr"] = net.ResolveUDPAddr
    Pkg["ResolveUnixAddr"] = net.ResolveUnixAddr
    Ctor["Resolver"] = Shadow_NewStruct_Resolver
    Ctor["SRV"] = Shadow_NewStruct_SRV
    Pkg["SplitHostPort"] = net.SplitHostPort
    Ctor["TCPAddr"] = Shadow_NewStruct_TCPAddr
    Ctor["TCPConn"] = Shadow_NewStruct_TCPConn
    Ctor["TCPListener"] = Shadow_NewStruct_TCPListener
    Ctor["UDPAddr"] = Shadow_NewStruct_UDPAddr
    Ctor["UDPConn"] = Shadow_NewStruct_UDPConn
    Ctor["UnixAddr"] = Shadow_NewStruct_UnixAddr
    Ctor["UnixConn"] = Shadow_NewStruct_UnixConn
    Ctor["UnixListener"] = Shadow_NewStruct_UnixListener

}
func Shadow_InterfaceConvertTo2_Addr(x interface{}) (y net.Addr, b bool) {
	y, b = x.(net.Addr)
	return
}

func Shadow_InterfaceConvertTo1_Addr(x interface{}) net.Addr {
	return x.(net.Addr)
}


func Shadow_NewStruct_AddrError(src *net.AddrError) *net.AddrError {
    if src == nil {
	   return &net.AddrError{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_Conn(x interface{}) (y net.Conn, b bool) {
	y, b = x.(net.Conn)
	return
}

func Shadow_InterfaceConvertTo1_Conn(x interface{}) net.Conn {
	return x.(net.Conn)
}


func Shadow_NewStruct_DNSConfigError(src *net.DNSConfigError) *net.DNSConfigError {
    if src == nil {
	   return &net.DNSConfigError{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_DNSError(src *net.DNSError) *net.DNSError {
    if src == nil {
	   return &net.DNSError{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Dialer(src *net.Dialer) *net.Dialer {
    if src == nil {
	   return &net.Dialer{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_Error(x interface{}) (y net.Error, b bool) {
	y, b = x.(net.Error)
	return
}

func Shadow_InterfaceConvertTo1_Error(x interface{}) net.Error {
	return x.(net.Error)
}


func Shadow_NewStruct_IPAddr(src *net.IPAddr) *net.IPAddr {
    if src == nil {
	   return &net.IPAddr{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_IPConn(src *net.IPConn) *net.IPConn {
    if src == nil {
	   return &net.IPConn{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_IPNet(src *net.IPNet) *net.IPNet {
    if src == nil {
	   return &net.IPNet{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Interface(src *net.Interface) *net.Interface {
    if src == nil {
	   return &net.Interface{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_Listener(x interface{}) (y net.Listener, b bool) {
	y, b = x.(net.Listener)
	return
}

func Shadow_InterfaceConvertTo1_Listener(x interface{}) net.Listener {
	return x.(net.Listener)
}


func Shadow_NewStruct_MX(src *net.MX) *net.MX {
    if src == nil {
	   return &net.MX{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_NS(src *net.NS) *net.NS {
    if src == nil {
	   return &net.NS{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_OpError(src *net.OpError) *net.OpError {
    if src == nil {
	   return &net.OpError{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_PacketConn(x interface{}) (y net.PacketConn, b bool) {
	y, b = x.(net.PacketConn)
	return
}

func Shadow_InterfaceConvertTo1_PacketConn(x interface{}) net.PacketConn {
	return x.(net.PacketConn)
}


func Shadow_NewStruct_ParseError(src *net.ParseError) *net.ParseError {
    if src == nil {
	   return &net.ParseError{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Resolver(src *net.Resolver) *net.Resolver {
    if src == nil {
	   return &net.Resolver{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_SRV(src *net.SRV) *net.SRV {
    if src == nil {
	   return &net.SRV{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_TCPAddr(src *net.TCPAddr) *net.TCPAddr {
    if src == nil {
	   return &net.TCPAddr{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_TCPConn(src *net.TCPConn) *net.TCPConn {
    if src == nil {
	   return &net.TCPConn{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_TCPListener(src *net.TCPListener) *net.TCPListener {
    if src == nil {
	   return &net.TCPListener{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_UDPAddr(src *net.UDPAddr) *net.UDPAddr {
    if src == nil {
	   return &net.UDPAddr{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_UDPConn(src *net.UDPConn) *net.UDPConn {
    if src == nil {
	   return &net.UDPConn{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_UnixAddr(src *net.UnixAddr) *net.UnixAddr {
    if src == nil {
	   return &net.UnixAddr{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_UnixConn(src *net.UnixConn) *net.UnixConn {
    if src == nil {
	   return &net.UnixConn{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_UnixListener(src *net.UnixListener) *net.UnixListener {
    if src == nil {
	   return &net.UnixListener{}
    }
    a := *src
    return &a
}

