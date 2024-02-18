package cowww

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// parseHttpRequest parses the HTTP request from the provided io.Reader and returns the HttpRequest object.
func parseHttpRequest(r io.Reader) (*HttpRequest, error) {
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
		idx = strings.Index(header, ": ")
		if idx == -1 {
			break
		}

		fieldKey := header[:idx]
		fieldValue := header[idx+2:]
		headers[fieldKey] = fieldValue
	}
	httpRequest.Headers = headers

	var cl int
	contentLength := getHeaderByKey(headers, "Content-Length")
	cl, err = strconv.Atoi(contentLength)
	if err != nil {
		return nil,
			errors.New(fmt.Sprintf("Invalid content length: %s", contentLength))
	}

	body := []byte{}
	if cl > 0 {
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
