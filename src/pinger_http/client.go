package pinger_http

import (
	"net"
	"net/http"
	"time"
)

const timeout = time.Duration(2 * time.Second)

var Client http.Client

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

func init() {
	transport := http.Transport{
		Dial: dialTimeout,
	}

	Client = http.Client{
		Transport: &transport,
	}
}
