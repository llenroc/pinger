package pinger_http

import (
	"net/http"
)

func Get(url string) (resp *Response) {
	resp = newResponseFor(url)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		resp.AddError(err)
		return
	}

	req.Header.Set("User-Agent", "local_pinger")

	r, err := Client.Do(req)

	if err != nil {
		resp.AddError(err)
		return
	}

	defer r.Body.Close()

	resp.populateFrom(r)

	return
}
