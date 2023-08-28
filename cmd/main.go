package main

import (
	"net/http"
	"github.com/nataneb32/go-htmx-template/pkg/home"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	home.Router(r)
	
	http.ListenAndServe("127.0.0.1:3000", r)
}
