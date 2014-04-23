package pinger_http

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Response struct {
	Body                    string
	Error                   error
	ErrorStr                string
	ResponseTimeNanoSeconds time.Duration
	ResponseTimeStr         string
	Status                  int
	StatusStr               string
	URL                     string
	Start                   time.Time
}

func newResponseFor(url string) *Response {
	return &Response{URL: url, Start: time.Now()}
}

func (response *Response) AddError(err error) {
	response.Error = err
	response.ErrorStr = fmt.Sprintf("%v", err)
}

func (response *Response) AddErrorMessage(msg string) {
	response.Error = errors.New(msg)
	response.ErrorStr = msg
}

func (response *Response) RequireHTTPOK() {
	if response.Error == nil && response.Status != 200 {
		response.AddErrorMessage(fmt.Sprintf("Got %s (%s)", response.StatusStr, response.Body))
	}
}

func (response *Response) populateFrom(r *http.Response) {
	response.ResponseTimeNanoSeconds = time.Since(response.Start)
	response.ResponseTimeStr = fmt.Sprintf("%v", response.ResponseTimeNanoSeconds)
	response.Status = r.StatusCode
	response.StatusStr = r.Status

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.AddError(err)
	} else {
		response.Body = string(body)
	}
}
