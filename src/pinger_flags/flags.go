package pinger_flags

import (
	"flag"
)

var (
	ServerURL     = flag.String("jobs_url", "http://127.0.0.1:3000/http_checks", "URL for jobs server")
	PeriodSeconds = flag.Int("period", 30, "Delay between checks")
)

func init() {
	flag.Parse()
}
