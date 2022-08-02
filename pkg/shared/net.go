package shared

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	IsStatusSuccess = map[int]bool{
		http.StatusOK:      true,
		http.StatusCreated: true,
	}
)

// TTPRequest
func HTTPRequest(ctx context.Context, method string, url string, body io.Reader, headers ...map[string]string) (res []byte, statusCode int, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	// iterate optional data of headers
	for _, header := range headers {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	timeout, _ := strconv.Atoi(os.Getenv("DEFAULT_TIMEOUT"))
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	r, err := client.Do(req)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	defer func() {
		r.Body.Close()
	}()

	resp := StreamToByte(r.Body)
	if os.Getenv("ENVIRONMENT") == "development" {
		tags := map[string]interface{}{
			"http.headers":    req.Header,
			"http.method":     req.Method,
			"http.url":        req.URL.String(),
			"response.status": r.Status,
			"response.body":   string(resp),
		}
		fmt.Println(tags)
	}

	return resp, r.StatusCode, nil
}

// StreamToString func
func StreamToString(stream io.Reader) string {
	if stream == nil {
		return ""
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}

// StreamToByte ...
func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}
