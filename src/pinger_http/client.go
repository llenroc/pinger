package pinger_http

import (
	"flags"
	"github.com/refiito/timeoutclient"
	"time"
)

const connectTimeout = time.Duration(2 * time.Second)

var Client = timeoutclient.NewTimeoutClient(connectTimeout, time.Duration(*flags.PeriodSeconds)*time.Second)
