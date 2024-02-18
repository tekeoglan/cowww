package cowww

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseHttpRequest(t *testing.T) {
	// Successful parsing test cases
	testCases := []struct {
		name     string
		input    string
		expected HttpRequest
	}{
		{
			name:  "Valid GET request",
			input: "GET /path/to/resource HTTP/1.1\r\nHost: example.com\r\nContent-Length: 0\r\n\r\n",
			expected: HttpRequest{
				Method: "GET",
				Url:    "/path/to/resource",
				Proto:  "HTTP/1.1",
				Headers: map[string]string{
					"Host":           "example.com",
					"Content-Length": "0",
				},
				Body: []byte{},
			},
		},
		{
			name:  "Valid POST request with body",
			input: "POST /api/data HTTP/1.0\r\nContent-Type: application/json\r\nContent-Length: 18\r\n\r\n{\"key\": \"value\"}",
			expected: HttpRequest{
				Method: "POST",
				Url:    "/api/data",
				Proto:  "HTTP/1.0",
				Headers: map[string]string{
					"Content-Type":   "application/json",
					"Content-Length": "18",
				},
				Body: []byte("{\"key\": \"value\"}"),
			},
		},
		// Add more test cases for different scenarios
	}

	for _, tc := range testCases {
		r := strings.NewReader(tc.input)
		httpRequest, err := parseHttpRequest(r)

		if err != nil {
			t.Errorf("%s: Unexpected error: %v", tc.name, err)
			continue
		}

		if !reflect.DeepEqual(*httpRequest, tc.expected) {
			t.Errorf("%s: Parsed request mismatch:\nGot: %#v\nExpected: %#v", tc.name, *httpRequest, tc.expected)
		}
	}

	// Error handling test cases
	errorTestCases := []struct {
		name          string
		input         string
		expectedError string
	}{
		{
			name:          "Invalid first line",
			input:         "Invalid data\r\n",
			expectedError: "Invalid first line format: Invalid data",
		},
		{
			name:          "Invalid Content-Length",
			input:         "GET / HTTP/1.1\r\nContent-Length: abc\r\n\r\n",
			expectedError: "Invalid content length: abc",
		},
		// Add more test cases for different error scenarios
	}

	for _, tc := range errorTestCases {
		r := strings.NewReader(tc.input)
		_, err := parseHttpRequest(r)

		if err == nil || err.Error() != tc.expectedError {
			t.Errorf("%s: Expected error %q, got %v", tc.name, tc.expectedError, err)
		}
	}
}
