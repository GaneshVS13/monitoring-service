package rest

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// RestService is the rest http request handler interface
type RestService interface {
	Do(ctx context.Context, method, url string, header map[string]string, payload interface{}) ([]byte, int, error)
}

// Service represents the rest service entity
type Service struct {
	Client *http.Client
}

// NewService returns the rest service instance
func NewService() *Service {
	tr := &http.Transport{
		MaxIdleConnsPerHost: 1024,
		TLSHandshakeTimeout: 0 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: 5 * time.Second}

	return &Service{
		Client: client,
	}
}

// Do makes a http request
func (s *Service) Do(ctx context.Context, method, url string,
	header map[string]string, payload interface{}) ([]byte, int, error) {
	reqPayload, err := getIOReader(payload)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqPayload)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to make http request. err: %v ", err)
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := s.Client.Do(req)
	if resp == nil || err != nil {
		return nil, 0, fmt.Errorf("failed to get http response. err: %v ", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to read response body. err: %v ", err)
	}

	return body, resp.StatusCode, nil
}

func getIOReader(body interface{}) (io.Reader, error) {
	var bodyReader io.Reader

	switch body := body.(type) {
	case nil:
	case io.Reader:
		bodyReader = body
	case string:
		bodyReader = bytes.NewBufferString(body)
	case []byte:
		bodyReader = bytes.NewBuffer(body)
	default:
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewBuffer(b)
	}
	return bodyReader, nil
}
