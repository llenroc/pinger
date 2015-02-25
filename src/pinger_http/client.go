package pinger_http

import (
	"github.com/refiito/timeoutclient"
	"pingerFlags"
	"time"
)

const connectTimeout = time.Duration(2 * time.Second)

var Client = timeoutclient.NewTimeoutClient(connectTimeout, time.Duration(*pingerFlags.PeriodSeconds))
