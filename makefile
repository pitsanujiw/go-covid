SHELL = /bin/bash
dev:
	~/go/bin/air --build.cmd "go build -o tmp/covid/main cmd/covid/main.go" --build.bin "./tmp/covid/main"
test:
	go test -v ./... -cover
build-linux:
	env GOOS=linux GOARCH=amd64 go build -o covid ./cmd/covid/main.go
build-darwin:
	env GOOS=darwin GOARCH=amd64 go build -o covid ./cmd/covid/main.go