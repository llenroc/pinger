package pinger_jobs

import (
	"encoding/json"
	"log"
	"pinger_http"
	"strings"
	"time"
)

type Manager struct {
	jobs      []string
	JobServer string
	Period    int
}

func (manager *Manager) Run() error {
	err := manager.loadJobs()
	if err != nil {
		if strings.Contains(err.Error(), "connection refused") {
			log.Printf("Connection to JobServer %s was refused, check the URL and that the JobServer is running", manager.JobServer)
		}
		log.Panicf("Error Loading jobs: %s", err)
	}
	return manager.performJobsLoop()
}

func (manager *Manager) performJobsLoop() error {
	log.Printf("Started loop with %d jobs (run every %d seconds)", len(manager.jobs), manager.Period)
	sleepTime := time.Duration(manager.Period) * time.Second
	for {
		err := manager.performJobs()
		if err != nil {
			return err
		}
		time.Sleep(sleepTime)
	}
}

func (manager *Manager) loadJobs() error {
	data := pinger_http.Get(manager.JobServer)
	data.RequireHTTPOK()

	if data.Error != nil {
		return data.Error
	}

	dec := json.NewDecoder(strings.NewReader(data.Body))
	return dec.Decode(&manager.jobs)
}

func (manager *Manager) performJobs() error {
	results := manager.getAsyncResults()
	b, err := json.Marshal(results)
	if err != nil {
		log.Panicf("Error generating json: %s", err)
	}

	result := pinger_http.PostJSON(manager.JobServer, string(b))

	if result.Error != nil {
		log.Printf("Error posting json to job server: %v\n", result.Error)
	}

	log.Printf("Status = %d", result.Status)

	if result.Status == 410 {
		return result.Error
	}

	return nil
}

func (manager *Manager) getAsyncResults() []*pinger_http.Response {
	results := []*pinger_http.Response{}
	asyncResults := make(chan *pinger_http.Response, len(manager.jobs))

	for _, job := range manager.jobs {
		go func(url string) {
			asyncResults <- pinger_http.Get(url)
		}(job)
	}

	for i := 0; i < len(manager.jobs); i++ {
		results = append(results, <-asyncResults)
	}
	return results
}
