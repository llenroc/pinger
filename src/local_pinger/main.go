package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"pinger_jobs"
)

func main() {
	log.Printf("Preparing to run")

	serverURL := flag.String("jobs_url", "http://127.0.0.1:3000/http_checks", "URL for jobs server")
	periodSeconds := flag.Int("period", 30, "Delay between checks")
	flag.Parse()

	host, err := os.Hostname()
	if err != nil {
		log.Panicf("Unable to get hostname: %s", err)
	}

	manager := pinger_jobs.Manager{Period: *periodSeconds,
		JobServer: fmt.Sprintf("%s?host=%s", *serverURL, host)}
	manager.Run()
}
