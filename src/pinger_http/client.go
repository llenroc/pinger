package pinger_http

import (
	"github.com/refiito/timeoutclient"
	"pinger_flags"
	"time"
)

const connectTimeout = time.Duration(2 * time.Second)

var Client = timeoutclient.NewTimeoutClient(connectTimeout, time.Duration(*pinger_flags.PeriodSeconds))
