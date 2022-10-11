package httpclient

import (
	"crypto/tls"
	"net/http"
	"time"
)

type HttpClienter interface {
	Get(url string) (resp *http.Response, err error)
}

func New() *http.Client {
	return &http.Client{
		Timeout: 15 * time.Second, // Default timeout request
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // skip in secure protocol
		},
	}
}
