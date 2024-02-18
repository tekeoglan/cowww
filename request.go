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
	Headers       map[string]string
	ContentLength int
	Host          string
	RemoteAddr    string
	Body          []byte
	Response      *HttpResponse
}
