package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type httpHeaders map[string]string

// An error response from the Statuspage API
type httpError struct {
	Message string `json:"message"` // Message describing the error
}

func (httpErr httpError) Error() string {
	return httpErr.Message
}

// Client used to perform HTTP requests
var client = &http.Client{}

func route(path string) string {
	const prefix = "https://api.statuspage.io/v1"
	return prefix + path
}

// Perform an HTTP request with authentication. Adapted from DoRequest in CSplan-API's tests
func doRequest(method string, url string, body any, headers httpHeaders, expectedStatus int) (r *http.Response, e error) {
	// Initialize headers
	if headers == nil {
		headers = make(httpHeaders)
	}

	// Encode body into a buffer
	var buffer bytes.Buffer
	if body != nil {
		// Set content-type to json by default
		if len(headers["Content-Type"]) == 0 {
			headers["Content-Type"] = "application/json"
		}

		// Handle different content types
		switch headers["Content-Type"] {
		case "application/json":
			marshalled, err := json.Marshal(body)
			if err != nil {
				return nil, err
			} else if strings.Contains(string(marshalled), "null") {
				fmt.Printf("%s:%s - null value in marshal found:\n%s", method, url, string(marshalled))
			}
			buffer.Write(marshalled)

		default:
			fmt.Println("unknown content type:", headers["Content-Type"])
		}
	}

	// Create request
	req, err := http.NewRequest(method, url, &buffer)
	if err != nil {
		return nil, err
	}

	// Add auth key
	req.Header.Set("Authorization", fmt.Sprintf("OAuth %s", key))

	// Attach headers to request
	for header, value := range headers {
		req.Header.Set(header, value)
	}

	// Perform the request
	r, e = client.Do(req)
	if e == nil && r.StatusCode != expectedStatus {
		var httpErr httpError
		json.NewDecoder(r.Body).Decode(&httpErr)
		// Format an error based on status if no response message is given
		if len(httpErr.Message) == 0 {
			httpErr.Message = fmt.Sprintf("Unknown error - expected status %d, received status %d", expectedStatus, r.StatusCode)
		}
		e = httpErr
	}
	return r, e
}
