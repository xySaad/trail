package main

import "net/http"

type Router[T any] struct {
	Dependency T
	ServeMux   *http.ServeMux
}

func Default() Router[any] {
	return Router[any]{}
}

func New[T any](dependency T) Router[T] {
	return Router[T]{
		Dependency: dependency,
		ServeMux:   http.DefaultServeMux,
	}
}

type Handler[T any] = func(c *Context[T])

func (router *Router[T]) Add(pattern string, handler Handler[T]) {
	router.ServeMux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		handler(&Context[T]{
			Dep:      router.Dependency,
			Request:  r,
			Response: w,
		})
	})
}

func (router Router[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.ServeMux.ServeHTTP(w, r)
}
