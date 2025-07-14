package api

import "fmt"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetHandler() {
	fmt.Println("Get handler")
}
