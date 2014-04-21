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
	Url                     string
	Start                   time.Time
}

func newResponseFor(url string) *Response {
	return &Response{Url: url, Start: time.Now()}
}

func (this *Response) AddError(err error) {
	this.Error = err
	this.ErrorStr = fmt.Sprintf("%v", err)
}

func (this *Response) AddErrorMessage(msg string) {
	this.Error = errors.New(msg)
	this.ErrorStr = msg
}

func (this *Response) RequireHttpOk() {
	if this.Error == nil && this.Status != 200 {
		this.AddErrorMessage(fmt.Sprintf("Got %s (%s)", this.StatusStr, this.Body))
	}
}

func (this *Response) populateFrom(r *http.Response) {
	this.ResponseTimeNanoSeconds = time.Since(this.Start)
	this.ResponseTimeStr = fmt.Sprintf("%v", this.ResponseTimeNanoSeconds)
	this.Status = r.StatusCode
	this.StatusStr = r.Status

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		this.AddError(err)
	} else {
		this.Body = string(body)
	}
}
