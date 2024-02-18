package cowww

type HttpResponse struct {
	Status           string
	StatusCode       int
	Proto            string
	Body             []byte
	ContentLength    int
	TransferEncoding []string
	Request          *HttpRequest
}
