package pinger_http

func Get(url string) (resp *Response) {
	resp = newResponseFor(url)

	r, err := Client.Get(url)

	if err != nil {
		resp.AddError(err)
		return
	}

	defer r.Body.Close()

	resp.populateFrom(r)

	return
}
