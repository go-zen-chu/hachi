package handler

import (
	"fmt"
	"net/http"
)

type Handler interface {
	GetHealth(w http.ResponseWriter, r *http.Request)
}

type handler struct {
}

func NewHandler() Handler {
	return &handler{}
}

func (h *handler) GetHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Healthy, wof, wof!")
}
