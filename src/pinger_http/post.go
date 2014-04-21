package pinger_http

import (
	"strings"
)

const jsonType = "application/json"

func PostJson(url, json string) (resp *Response) {
	resp = newResponseFor(url)

	b := strings.NewReader(json)

	r, err := Client.Post(url, jsonType, b)

	if err != nil {
		resp.AddError(err)
		return
	}

	defer r.Body.Close()

	resp.populateFrom(r)
	resp.RequireHttpOk()

	return
}
