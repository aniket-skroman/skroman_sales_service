package proxycalls

import (
	"bytes"
	"context"
	"net/http"
	"time"
)

var (
	BASE_URL     = "http://3.7.18.55:3000/skroman/"
	USER_SERVICE = "http://15.207.19.172:8080/api/"
	// USER_SERVICE = "http://15.207.19.172:808/api/"
)

type APIRequest interface {
	MakeRequest() (*http.Response, error)
}

type ProxyCalls struct {
	ReqEndpoint    string
	RequestMethod  string
	RequestBody    []byte
	RequestParams  interface{}
	RequestHeaders map[string]string
	IsRequestBody  bool // request need a body or params
	Response       *http.Response
}

func NewAPIRequest(endpoint, method string, isRequestBody bool, requestBody []byte, params interface{}, headers map[string]string) APIRequest {
	return &ProxyCalls{
		ReqEndpoint:    endpoint,
		RequestMethod:  method,
		IsRequestBody:  isRequestBody,
		RequestBody:    requestBody,
		RequestParams:  params,
		RequestHeaders: headers,
	}
}

func (apiRequest *ProxyCalls) MakeRequest() (*http.Response, error) {
	var request *http.Request
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if apiRequest.IsRequestBody {
		request, err = http.NewRequestWithContext(ctx, apiRequest.RequestMethod, USER_SERVICE+apiRequest.ReqEndpoint, bytes.NewReader(apiRequest.RequestBody))
	} else {
		request, err = http.NewRequestWithContext(ctx, apiRequest.RequestMethod, USER_SERVICE+apiRequest.ReqEndpoint, nil)
	}

	if err != nil {
		return nil, err
	}

	for key, val := range apiRequest.RequestHeaders {
		request.Header.Set(key, val)
	}

	request.Close = true
	request.Header.Set("Content-Type", "application/json")

	return http.DefaultClient.Do(request)
}

func (procall *ProxyCalls) MakeRequestWithBody() (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var request *http.Request
	var err error

	if procall.IsRequestBody {
		request, err = http.NewRequestWithContext(ctx, procall.RequestMethod, USER_SERVICE+procall.ReqEndpoint, bytes.NewReader(procall.RequestBody))
	} else {
		request, err = http.NewRequestWithContext(ctx, procall.RequestMethod, USER_SERVICE+procall.ReqEndpoint, nil)
	}

	if err != nil {
		return nil, err
	}
	request.Close = true

	for key, val := range procall.RequestHeaders {
		request.Header.Set(key, val)
	}

	request.Header.Set("Content-Type", "application/json")

	return http.DefaultClient.Do(request)
}
