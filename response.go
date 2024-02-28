package cowww

import (
	"errors"
	"net"
)

var (
	InternalServerError = errors.New("Internal Server Error")
	NotImplementedError = errors.New("Not Implemented")
	BadRequestError     = errors.New("Bad Request")
	NotFoundError       = errors.New("Not Found")
	ForbiddenError      = errors.New("Forbidden")
)

const (
	StatusOk                  = 200
	StatusCreated             = 201
	StatusBadRequest          = 400
	StatusNotFound            = 404
	StatusForbidden           = 403
	StatusInternalServerError = 500
	StatusNotImplemented      = 501
)

type ResponseWriter interface {
	Header() Header
	Write([]byte) (int, error)
	WriteHeader(int)
}

type HttpResponse struct {
	Status           string
	StatusCode       int
	Proto            string
	Header           Header
	Body             []byte
	ContentLength    int
	TransferEncoding []string
	Request          *HttpRequest
}

func statusText(status int) string {
	switch status {
	case StatusOk:
		return "OK"
	case StatusCreated:
		return "Created"
	case StatusBadRequest:
		return "Bad Request"
	case StatusNotFound:
		return "Not Found"
	case StatusForbidden:
		return "Forbidden"
	case StatusInternalServerError:
		return "Internal Server Error"
	case StatusNotImplemented:
		return "Not Implemented"
	default:
		return ""
	}
}

type response struct {
	c             net.Conn
	statusCode    int
	status        string
	handlerHeader Header
}

func (r *response) Header() Header {
	return r.handlerHeader
}

func (r *response) Write(b []byte) (int, error) {
	return r.c.Write(b)
}

func (r *response) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.status = statusText(statusCode)
}
