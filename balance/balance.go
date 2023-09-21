package balance

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

var backendServers = []string{
	"http://localhost:8081",
	"http://localhost:8082",
	"http://localhost:8083",
}

type LoadBalancer struct {
	backendURLs []*url.URL
	proxy       *httputil.ReverseProxy
	nextServer  int
}

func NewLoadBalancer() *LoadBalancer {
	var backendURLs []*url.URL
	for _, server := range backendServers {
		url, _ := url.Parse(server)
		backendURLs = append(backendURLs, url)
	}

	return &LoadBalancer{
		backendURLs: backendURLs,
		proxy:       &httputil.ReverseProxy{},
		nextServer:  0,
	}
}

func (lb *LoadBalancer) ReverseProxyHandler(w http.ResponseWriter, r *http.Request) {
	lb.nextServer = (lb.nextServer + 1) % len(lb.backendURLs)
	backendURL := lb.backendURLs[lb.nextServer]

	lb.proxy.Director = func(req *http.Request) {
		req.URL = backendURL
		req.Host = backendURL.Host
	}
	lb.proxy.ServeHTTP(w, r)
}
