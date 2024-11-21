package pushnotifications

import (
	"net/http"
	"time"
)

type Option func(*pushNotifications)

func WithRequestTimeout(timeout time.Duration) Option {
	return func(pn *pushNotifications) {
		pn.httpClient.Timeout = timeout
	}
}

func WithCustomBaseURL(url string) Option {
	return func(pn *pushNotifications) {
		pn.baseEndpoint = url
	}
}

func WithHTTPClient(httpClient *http.Client) Option {
	return func(pn *pushNotifications) {
		pn.httpClient = httpClient
	}
}
