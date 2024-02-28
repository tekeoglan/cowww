package cowww

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var carriageReturn string = "\r\n"
var columnSign string = ": "

// parseHttpRequest parses the HTTP request from the provided io.Reader and returns the HttpRequest object.
func parseHttpRequest(r io.Reader) (*HttpRequest, error) {
	if r == nil {
		return nil, errors.New("Reader is nil")
	}

	scanner := bufio.NewReaderSize(r, MAX_REQUEST_SIZE)

	firstLineBytes, _, err := scanner.ReadLine()
	if err != nil {
		return nil,
			errors.New("Can't parse first line")
	}

	firstLine := string(firstLineBytes)
	parts := strings.Split(firstLine, " ")
	if len(parts) != 3 {
		return nil,
			errors.New(fmt.Sprintf("Invalid first line format: %s", firstLine))
	}

	httpRequest := new(HttpRequest)

	method := parts[0]
	httpRequest.Method = method

	url := parts[1]
	httpRequest.Url = url

	proto := parts[2]
	httpRequest.Proto = proto

	headers := map[string]string{}
	var rawHeader []byte
	var header string
	var idx int
	for true {
		rawHeader, _, err = scanner.ReadLine()
		if err != nil {
			return nil,
				errors.New("Can't parse headers")
		}

		if len(rawHeader) == 0 {
			break
		}

		header = string(rawHeader)
		idx = strings.Index(header, columnSign)
		if idx == -1 {
			break
		}

		fieldKey := header[:idx]
		fieldValue := header[idx+2:]
		headers[fieldKey] = fieldValue
	}
	httpRequest.Header = headers
	body := []byte{}

	contentLength := getHeaderByKey(headers, "Content-Length")
	if contentLength != "" {
		_, err = strconv.Atoi(contentLength)
		if err != nil {
			return nil,
				errors.New(fmt.Sprintf("Invalid content length: %s", contentLength))
		}

		var b byte
		for true {
			b, err = scanner.ReadByte()
			if err != nil {
				if err.Error() == "EOF" {
					break
				}

				return nil, errors.New(fmt.Sprintf("Can't parse body: %v", err))
			}
			body = append(body, b)
		}
	}

	httpRequest.Body = body

	return httpRequest, nil
}

// parseResponseToBytes parses the HTTP response to bytes
// it returns nil if the response is nil
func parseResponseToBytes(r *HttpResponse) []byte {
	if r == nil {
		return nil
	}
	buffer := bytes.NewBuffer([]byte{})

	// write status line
	buffer.WriteString(r.Proto)
	buffer.WriteString(fmt.Sprintf(" %d %s%s", r.StatusCode, r.Status, carriageReturn))

	if r.Header == nil {
		return buffer.Bytes()
	}

	for key, value := range r.Header {
		buffer.WriteString(key)
		buffer.WriteString(columnSign)
		buffer.WriteString(value)
		buffer.WriteString(carriageReturn)
	}

	if r.Body == nil || len(r.Body) == 0 {
		return buffer.Bytes()
	}

	buffer.WriteString(carriageReturn)
	buffer.Write(r.Body)

	return buffer.Bytes()
}
