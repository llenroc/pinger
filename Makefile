export GOPATH=$(shell pwd)

default: new quality build

build:
	go build -o bin/local_pinger local_pinger

clean:
	rm -rf bin

new:
	rm -rf bin/local_pinger

quality: fmt vet lint

fmt:
	go fmt local_pinger pinger_http pinger_jobs

vet:
	go vet local_pinger pinger_http pinger_jobs

lint: bin/golint
	bin/golint src/*pinger*

bin/golint:
	go get github.com/golang/lint/golint
	go build -o bin/golint github.com/golang/lint/golint
	rm -rf src/github.com/golang/lint

dummy-build:
	go build -o bin/dummy_jobs src/dummy/dummy_jobs.go

dummy: dummy-build
	./bin/dummy_jobs