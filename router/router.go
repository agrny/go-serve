package router

import "net/http"

type HTTPRouter struct {
	*http.ServeMux
	hostname string
}

func NewHTTPServer(hostname string) *HTTPRouter {
	h := HTTPRouter{}
	h.hostname = hostname
	h.ServeMux = http.NewServeMux()
	return &h
}
