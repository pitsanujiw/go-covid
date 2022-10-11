FROM golang:1.18.5-alpine3.16 as build

WORKDIR /build

ENV GO111MODULE=on 
ENV CGO_ENABLED=0 
ENV GOOS=linux

COPY go.mod go.sum /build/
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o covid ./cmd/covid/main.go
RUN go clean -cache

RUN adduser -S -D -H -h /build appuser

FROM scratch as production

WORKDIR /app

COPY --from=build /build /app/

EXPOSE 8080

CMD ["./covid"]