package cowww

const (
	MAX_REQUEST_SIZE = 1 << 20
)

const (
	methodGet     = "GET"
	methodPost    = "POST"
	methodPut     = "PUT"
	methodDelete  = "DELETE"
	methodConnect = "CONNECT"
)

type HttpRequest struct {
	Method        string
	Url           string
	Proto         string
	Header        Header
	ContentLength int
	Host          string
	RemoteAddr    string
	Body          []byte
	Response      *HttpResponse
}
