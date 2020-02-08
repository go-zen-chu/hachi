package hachi

import (
	"fmt"
	"net/http"
	"github.com/go-zen-chu/hachi/pkg/interface/handler"
)

type HttpServer interface {
	ConfigureRoute(h handler.Handler)
	Run(port int) error
}

type httpServer struct {
	mux *http.ServeMux
}

// NewHttpServer returns actual http server that handles http requests
func NewHttpServer() HttpServer {
	return &httpServer{
		mux: http.NewServeMux(),
	}
}

// ConfigureRoute configures http server routes
func (hs *httpServer) ConfigureRoute(h handler.Handler) {
	hs.mux.HandleFunc("/health", h.GetHealth)
}

// Run runs server
func (hs *httpServer) Run(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), hs.mux)
}
