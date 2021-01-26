package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type request struct {
	*http.Request
	err error
}

// NewRequest
func NewRequest(method, url string, body io.Reader) *request {
	req := new(request)
	request, err := http.NewRequest(method, url, body)

	if err != nil {
		req.err = err
	}

	req.Request = request
	return req
}

// Get
func Get(url string) (*response, error) {
	return NewRequest(http.MethodGet, url, nil).Send()
}

// Post
func Post(url string, body io.Reader) (*response, error) {
	return NewRequest(http.MethodPost, url, body).Send()
}

// Send send http request
func (req *request) Send() (*response, error) {
	if req.err != nil {
		return nil, req.err
	}
	client := http.Client{}
	resp, err := client.Do(req.Request)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, fmt.Errorf("no response")
	}

	// close response
	defer func() {
		_ = resp.Body.Close()
	}()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}
	return &response{body: body}, nil
}

func (req *request) Err() error {
	return req.err
}

// response http response wrapper
type response struct {
	body []byte
}

// String return response as string
func (wrap *response) String() string {
	return string(wrap.body)
}

// Byte return response as byte
func (wrap *response) Byte() []byte {
	return wrap.body
}

// BindJson bind json object with pointer
func (wrap *response) BindJson(object interface{}) error {
	return JsonToStruct(string(wrap.body), object)
}

// Reader
func (wrap *response) Reader() io.Reader {
	return bytes.NewReader(wrap.body)
}

func JsonToStruct(jsonStr string, obj interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
