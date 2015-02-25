package main

import (
	"fmt"
	"log"
	"os"
	"pinger_flags"
	"pinger_jobs"
)

func main() {
	log.Printf("Preparing to run")

	host, err := os.Hostname()
	if err != nil {
		log.Panicf("Unable to get hostname: %s", err)
	}

	manager := pinger_jobs.Manager{Period: *pinger_flags.PeriodSeconds,
		JobServer: fmt.Sprintf("%s?host=%s", *pinger_flags.ServerURL, host)}
	err = manager.Run()
	if err != nil {
		os.Exit(1)
	}
}
