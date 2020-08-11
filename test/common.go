package test

import (
	"golang.org/x/net/context"
	"net"
	"net/http"
	"net/http/httptest"
)

const (
	BaseTestApiUrl = "http://data.sfgov.org"
	BaseTestGeoApiUrl = "http://maps.googleapis.com"
)

func TestingHTTPClient(handler http.Handler) (*http.Client, func()) {
	s := httptest.NewServer(handler)

	client := &http.Client {
		Transport: &http.Transport {
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
		},
	}

	return client, s.Close
}
