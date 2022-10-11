SHELL = /bin/bash
covid-serv:
	go run cmd/covid/main.go
test:
	go test -v ./... -cover