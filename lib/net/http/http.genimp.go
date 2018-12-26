package shadow_http

import "net/http"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["CanonicalHeaderKey"] = http.CanonicalHeaderKey
    Ctor["Client"] = Shadow_NewStruct_Client
    Pkg["CloseNotifier"] = Shadow_InterfaceConvertTo2_CloseNotifier
    Ctor["Cookie"] = Shadow_NewStruct_Cookie
    Pkg["CookieJar"] = Shadow_InterfaceConvertTo2_CookieJar
    Pkg["DefaultClient"] = http.DefaultClient
    Pkg["DefaultMaxHeaderBytes"] = http.DefaultMaxHeaderBytes
    Pkg["DefaultMaxIdleConnsPerHost"] = http.DefaultMaxIdleConnsPerHost
    Pkg["DefaultServeMux"] = http.DefaultServeMux
    Pkg["DefaultTransport"] = http.DefaultTransport
    Pkg["DetectContentType"] = http.DetectContentType
    Pkg["ErrAbortHandler"] = http.ErrAbortHandler
    Pkg["ErrBodyNotAllowed"] = http.ErrBodyNotAllowed
    Pkg["ErrBodyReadAfterClose"] = http.ErrBodyReadAfterClose
    Pkg["ErrContentLength"] = http.ErrContentLength
    Pkg["ErrHandlerTimeout"] = http.ErrHandlerTimeout
    Pkg["ErrHeaderTooLong"] = http.ErrHeaderTooLong
    Pkg["ErrHijacked"] = http.ErrHijacked
    Pkg["ErrLineTooLong"] = http.ErrLineTooLong
    Pkg["ErrMissingBoundary"] = http.ErrMissingBoundary
    Pkg["ErrMissingContentLength"] = http.ErrMissingContentLength
    Pkg["ErrMissingFile"] = http.ErrMissingFile
    Pkg["ErrNoCookie"] = http.ErrNoCookie
    Pkg["ErrNoLocation"] = http.ErrNoLocation
    Pkg["ErrNotMultipart"] = http.ErrNotMultipart
    Pkg["ErrNotSupported"] = http.ErrNotSupported
    Pkg["ErrServerClosed"] = http.ErrServerClosed
    Pkg["ErrShortBody"] = http.ErrShortBody
    Pkg["ErrSkipAltProtocol"] = http.ErrSkipAltProtocol
    Pkg["ErrUnexpectedTrailer"] = http.ErrUnexpectedTrailer
    Pkg["ErrUseLastResponse"] = http.ErrUseLastResponse
    Pkg["ErrWriteAfterFlush"] = http.ErrWriteAfterFlush
    Pkg["Error"] = http.Error
    Pkg["File"] = Shadow_InterfaceConvertTo2_File
    Pkg["FileServer"] = http.FileServer
    Pkg["FileSystem"] = Shadow_InterfaceConvertTo2_FileSystem
    Pkg["Flusher"] = Shadow_InterfaceConvertTo2_Flusher
    Pkg["Get"] = http.Get
    Pkg["Handle"] = http.Handle
    Pkg["HandleFunc"] = http.HandleFunc
    Pkg["Handler"] = Shadow_InterfaceConvertTo2_Handler
    Pkg["Head"] = http.Head
    Pkg["Hijacker"] = Shadow_InterfaceConvertTo2_Hijacker
    Pkg["ListenAndServe"] = http.ListenAndServe
    Pkg["ListenAndServeTLS"] = http.ListenAndServeTLS
    Pkg["LocalAddrContextKey"] = http.LocalAddrContextKey
    Pkg["MaxBytesReader"] = http.MaxBytesReader
    Pkg["MethodConnect"] = http.MethodConnect
    Pkg["MethodDelete"] = http.MethodDelete
    Pkg["MethodGet"] = http.MethodGet
    Pkg["MethodHead"] = http.MethodHead
    Pkg["MethodOptions"] = http.MethodOptions
    Pkg["MethodPatch"] = http.MethodPatch
    Pkg["MethodPost"] = http.MethodPost
    Pkg["MethodPut"] = http.MethodPut
    Pkg["MethodTrace"] = http.MethodTrace
    Pkg["NewFileTransport"] = http.NewFileTransport
    Pkg["NewRequest"] = http.NewRequest
    Pkg["NewServeMux"] = http.NewServeMux
    Pkg["NoBody"] = http.NoBody
    Pkg["NotFound"] = http.NotFound
    Pkg["NotFoundHandler"] = http.NotFoundHandler
    Pkg["ParseHTTPVersion"] = http.ParseHTTPVersion
    Pkg["ParseTime"] = http.ParseTime
    Pkg["Post"] = http.Post
    Pkg["PostForm"] = http.PostForm
    Ctor["ProtocolError"] = Shadow_NewStruct_ProtocolError
    Pkg["ProxyFromEnvironment"] = http.ProxyFromEnvironment
    Pkg["ProxyURL"] = http.ProxyURL
    Ctor["PushOptions"] = Shadow_NewStruct_PushOptions
    Pkg["Pusher"] = Shadow_InterfaceConvertTo2_Pusher
    Pkg["ReadRequest"] = http.ReadRequest
    Pkg["ReadResponse"] = http.ReadResponse
    Pkg["Redirect"] = http.Redirect
    Pkg["RedirectHandler"] = http.RedirectHandler
    Ctor["Request"] = Shadow_NewStruct_Request
    Ctor["Response"] = Shadow_NewStruct_Response
    Pkg["ResponseWriter"] = Shadow_InterfaceConvertTo2_ResponseWriter
    Pkg["RoundTripper"] = Shadow_InterfaceConvertTo2_RoundTripper
    Pkg["Serve"] = http.Serve
    Pkg["ServeContent"] = http.ServeContent
    Pkg["ServeFile"] = http.ServeFile
    Ctor["ServeMux"] = Shadow_NewStruct_ServeMux
    Pkg["ServeTLS"] = http.ServeTLS
    Ctor["Server"] = Shadow_NewStruct_Server
    Pkg["ServerContextKey"] = http.ServerContextKey
    Pkg["SetCookie"] = http.SetCookie
    Pkg["StatusAccepted"] = http.StatusAccepted
    Pkg["StatusAlreadyReported"] = http.StatusAlreadyReported
    Pkg["StatusBadGateway"] = http.StatusBadGateway
    Pkg["StatusBadRequest"] = http.StatusBadRequest
    Pkg["StatusConflict"] = http.StatusConflict
    Pkg["StatusContinue"] = http.StatusContinue
    Pkg["StatusCreated"] = http.StatusCreated
    Pkg["StatusExpectationFailed"] = http.StatusExpectationFailed
    Pkg["StatusFailedDependency"] = http.StatusFailedDependency
    Pkg["StatusForbidden"] = http.StatusForbidden
    Pkg["StatusFound"] = http.StatusFound
    Pkg["StatusGatewayTimeout"] = http.StatusGatewayTimeout
    Pkg["StatusGone"] = http.StatusGone
    Pkg["StatusHTTPVersionNotSupported"] = http.StatusHTTPVersionNotSupported
    Pkg["StatusIMUsed"] = http.StatusIMUsed
    Pkg["StatusInsufficientStorage"] = http.StatusInsufficientStorage
    Pkg["StatusInternalServerError"] = http.StatusInternalServerError
    Pkg["StatusLengthRequired"] = http.StatusLengthRequired
    Pkg["StatusLocked"] = http.StatusLocked
    Pkg["StatusLoopDetected"] = http.StatusLoopDetected
    Pkg["StatusMethodNotAllowed"] = http.StatusMethodNotAllowed
    Pkg["StatusMovedPermanently"] = http.StatusMovedPermanently
    Pkg["StatusMultiStatus"] = http.StatusMultiStatus
    Pkg["StatusMultipleChoices"] = http.StatusMultipleChoices
    Pkg["StatusNetworkAuthenticationRequired"] = http.StatusNetworkAuthenticationRequired
    Pkg["StatusNoContent"] = http.StatusNoContent
    Pkg["StatusNonAuthoritativeInfo"] = http.StatusNonAuthoritativeInfo
    Pkg["StatusNotAcceptable"] = http.StatusNotAcceptable
    Pkg["StatusNotExtended"] = http.StatusNotExtended
    Pkg["StatusNotFound"] = http.StatusNotFound
    Pkg["StatusNotImplemented"] = http.StatusNotImplemented
    Pkg["StatusNotModified"] = http.StatusNotModified
    Pkg["StatusOK"] = http.StatusOK
    Pkg["StatusPartialContent"] = http.StatusPartialContent
    Pkg["StatusPaymentRequired"] = http.StatusPaymentRequired
    Pkg["StatusPermanentRedirect"] = http.StatusPermanentRedirect
    Pkg["StatusPreconditionFailed"] = http.StatusPreconditionFailed
    Pkg["StatusPreconditionRequired"] = http.StatusPreconditionRequired
    Pkg["StatusProcessing"] = http.StatusProcessing
    Pkg["StatusProxyAuthRequired"] = http.StatusProxyAuthRequired
    Pkg["StatusRequestEntityTooLarge"] = http.StatusRequestEntityTooLarge
    Pkg["StatusRequestHeaderFieldsTooLarge"] = http.StatusRequestHeaderFieldsTooLarge
    Pkg["StatusRequestTimeout"] = http.StatusRequestTimeout
    Pkg["StatusRequestURITooLong"] = http.StatusRequestURITooLong
    Pkg["StatusRequestedRangeNotSatisfiable"] = http.StatusRequestedRangeNotSatisfiable
    Pkg["StatusResetContent"] = http.StatusResetContent
    Pkg["StatusSeeOther"] = http.StatusSeeOther
    Pkg["StatusServiceUnavailable"] = http.StatusServiceUnavailable
    Pkg["StatusSwitchingProtocols"] = http.StatusSwitchingProtocols
    Pkg["StatusTeapot"] = http.StatusTeapot
    Pkg["StatusTemporaryRedirect"] = http.StatusTemporaryRedirect
    Pkg["StatusText"] = http.StatusText
    Pkg["StatusTooManyRequests"] = http.StatusTooManyRequests
    Pkg["StatusUnauthorized"] = http.StatusUnauthorized
    Pkg["StatusUnavailableForLegalReasons"] = http.StatusUnavailableForLegalReasons
    Pkg["StatusUnprocessableEntity"] = http.StatusUnprocessableEntity
    Pkg["StatusUnsupportedMediaType"] = http.StatusUnsupportedMediaType
    Pkg["StatusUpgradeRequired"] = http.StatusUpgradeRequired
    Pkg["StatusUseProxy"] = http.StatusUseProxy
    Pkg["StatusVariantAlsoNegotiates"] = http.StatusVariantAlsoNegotiates
    Pkg["StripPrefix"] = http.StripPrefix
    Pkg["TimeFormat"] = http.TimeFormat
    Pkg["TimeoutHandler"] = http.TimeoutHandler
    Pkg["TrailerPrefix"] = http.TrailerPrefix
    Ctor["Transport"] = Shadow_NewStruct_Transport

}
func Shadow_NewStruct_Client(src *http.Client) *http.Client {
    if src == nil {
	   return &http.Client{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_CloseNotifier(x interface{}) (y http.CloseNotifier, b bool) {
	y, b = x.(http.CloseNotifier)
	return
}

func Shadow_InterfaceConvertTo1_CloseNotifier(x interface{}) http.CloseNotifier {
	return x.(http.CloseNotifier)
}


func Shadow_NewStruct_Cookie(src *http.Cookie) *http.Cookie {
    if src == nil {
	   return &http.Cookie{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_CookieJar(x interface{}) (y http.CookieJar, b bool) {
	y, b = x.(http.CookieJar)
	return
}

func Shadow_InterfaceConvertTo1_CookieJar(x interface{}) http.CookieJar {
	return x.(http.CookieJar)
}


func Shadow_InterfaceConvertTo2_File(x interface{}) (y http.File, b bool) {
	y, b = x.(http.File)
	return
}

func Shadow_InterfaceConvertTo1_File(x interface{}) http.File {
	return x.(http.File)
}


func Shadow_InterfaceConvertTo2_FileSystem(x interface{}) (y http.FileSystem, b bool) {
	y, b = x.(http.FileSystem)
	return
}

func Shadow_InterfaceConvertTo1_FileSystem(x interface{}) http.FileSystem {
	return x.(http.FileSystem)
}


func Shadow_InterfaceConvertTo2_Flusher(x interface{}) (y http.Flusher, b bool) {
	y, b = x.(http.Flusher)
	return
}

func Shadow_InterfaceConvertTo1_Flusher(x interface{}) http.Flusher {
	return x.(http.Flusher)
}


func Shadow_InterfaceConvertTo2_Handler(x interface{}) (y http.Handler, b bool) {
	y, b = x.(http.Handler)
	return
}

func Shadow_InterfaceConvertTo1_Handler(x interface{}) http.Handler {
	return x.(http.Handler)
}


func Shadow_InterfaceConvertTo2_Hijacker(x interface{}) (y http.Hijacker, b bool) {
	y, b = x.(http.Hijacker)
	return
}

func Shadow_InterfaceConvertTo1_Hijacker(x interface{}) http.Hijacker {
	return x.(http.Hijacker)
}


func Shadow_NewStruct_ProtocolError(src *http.ProtocolError) *http.ProtocolError {
    if src == nil {
	   return &http.ProtocolError{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_PushOptions(src *http.PushOptions) *http.PushOptions {
    if src == nil {
	   return &http.PushOptions{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_Pusher(x interface{}) (y http.Pusher, b bool) {
	y, b = x.(http.Pusher)
	return
}

func Shadow_InterfaceConvertTo1_Pusher(x interface{}) http.Pusher {
	return x.(http.Pusher)
}


func Shadow_NewStruct_Request(src *http.Request) *http.Request {
    if src == nil {
	   return &http.Request{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Response(src *http.Response) *http.Response {
    if src == nil {
	   return &http.Response{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_ResponseWriter(x interface{}) (y http.ResponseWriter, b bool) {
	y, b = x.(http.ResponseWriter)
	return
}

func Shadow_InterfaceConvertTo1_ResponseWriter(x interface{}) http.ResponseWriter {
	return x.(http.ResponseWriter)
}


func Shadow_InterfaceConvertTo2_RoundTripper(x interface{}) (y http.RoundTripper, b bool) {
	y, b = x.(http.RoundTripper)
	return
}

func Shadow_InterfaceConvertTo1_RoundTripper(x interface{}) http.RoundTripper {
	return x.(http.RoundTripper)
}


func Shadow_NewStruct_ServeMux(src *http.ServeMux) *http.ServeMux {
    if src == nil {
	   return &http.ServeMux{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Server(src *http.Server) *http.Server {
    if src == nil {
	   return &http.Server{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Transport(src *http.Transport) *http.Transport {
    if src == nil {
	   return &http.Transport{}
    }
    a := *src
    return &a
}

