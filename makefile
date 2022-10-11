SHELL = /bin/bash
dev:
	~/go/bin/air --build.cmd "go build -o tmp/covid/main cmd/covid/main.go" --build.bin "./tmp/covid/main"
test:
	go test -v ./... -cover
build-linux:
	GOOS=linux GOARCH=amd64 go build -o covid ./cmd/covid/main.go
build-darwin:
	eGOOS=darwin GOARCH=amd64 go build -o covid ./cmd/covid/main.go