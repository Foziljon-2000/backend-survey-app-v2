package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouterCompl() http.Handler {
	r := mux.NewRouter()

	return r
}
