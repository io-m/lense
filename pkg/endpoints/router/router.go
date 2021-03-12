package router

import "net/http"

// Router interface defines http methods, endpoints and handlers
type Router interface {
	Get(path string, handler http.HandlerFunc)
	Post(path string, handler func(w http.ResponseWriter, r *http.Request))
	Put(path string, handler func(w http.ResponseWriter, r *http.Request))
	Delete(path string, handler func(w http.ResponseWriter, r *http.Request))
	ListenAndServe(port string)
}
