package proxy

import (
	"api-gateway/internal/router"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewProxy(target string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	// Optional cloud-native tweak: adjust headers for upstream microservices
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Host = url.Host
	}

	return proxy, nil
}

func GatewayHandler(r *router.Router) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		route := r.Match(req.URL.Path)
		if route == nil {
			http.Error(w, "Route Not Found", http.StatusNotFound)
			return
		}

		proxy, err := NewProxy(route.Upstream)
		if err != nil {
			log.Printf("Failed to create proxy for %s: %v", route.Upstream, err)
			http.Error(w, "Bad Gateway", http.StatusBadGateway)
			return
		}

		proxy.ServeHTTP(w, req)
	})
}
