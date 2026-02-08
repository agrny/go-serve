package models

import (
	"net/http"
)

type RequesterInfo struct {
	Address   string
	Port      string
	UserAgent string
	Method    string
	Protocol  string
	Headers   map[string][]string
}

func (r *RequesterInfo) ParseRequest(req *http.Request) {
	r.Address = req.RemoteAddr
	r.Port = req.Host
	r.UserAgent = req.UserAgent()
	r.Method = req.Method
	r.Protocol = req.Proto
	if r.Headers == nil {
		r.Headers = make(map[string][]string)
	}
	for name, value := range req.Header {
		r.Headers[name] = append(r.Headers[name], value...)
	}
}
