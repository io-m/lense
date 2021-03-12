package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type gorillaMux struct{}

// NewMuxRouter is constructor function for creating new instances of Gorilla mux router
func NewMuxRouter() Router {
	return &gorillaMux{}
}

// Instance of Gurilla mux router
var muxRouter = mux.NewRouter().StrictSlash(true)

func (*gorillaMux) Get(path string, handler http.HandlerFunc) {
	muxRouter.HandleFunc(path, handler).Methods(http.MethodGet)
}

func (*gorillaMux) Post(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	muxRouter.HandleFunc(path, handler).Methods(http.MethodPost)
}
func (*gorillaMux) Put(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	muxRouter.HandleFunc(path, handler).Methods(http.MethodPut)
}
func (*gorillaMux) Delete(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	muxRouter.HandleFunc(path, handler).Methods(http.MethodDelete)
}

func (*gorillaMux) ListenAndServe(port string) {
	log.Printf("Server is up and running on PORT : %v", port)
	log.Fatal(http.ListenAndServe(port, muxRouter))
}
