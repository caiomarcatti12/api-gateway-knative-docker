package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func proxyToService(serviceURL *url.URL) http.HandlerFunc {
	proxy := httputil.NewSingleHostReverseProxy(serviceURL)
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}
