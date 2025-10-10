package trail

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
type Middleware[T any] = func(c *Context[T]) bool

func (router *Router[T]) Add(pattern string, handler Handler[T], middlewares ...Middleware[T]) {
	router.ServeMux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		ctx := &Context[T]{
			Dep:      router.Dependency,
			Request:  r,
			Response: w,
		}
		for _, mw := range middlewares {
			if !mw(ctx) {
				return
			}
		}
		handler(ctx)
	})
}

func (router Router[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.ServeMux.ServeHTTP(w, r)
}
