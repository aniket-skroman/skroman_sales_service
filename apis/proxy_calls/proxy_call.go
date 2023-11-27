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
)

type ProxyCalls struct {
	ReqEndpoint    string
	RequestBody    []byte
	RequestMethod  string
	RequestParams  interface{}
	RequestHeaders map[string]string
	IsRequestBody  bool // request need a body or params
	Response       *http.Response
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
