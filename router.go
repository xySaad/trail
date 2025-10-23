package trail

import (
	"net/http"
)

type WithContext interface {
	GetBase() *Context
}

type Router[C WithContext] struct {
	base C
	mux  *http.ServeMux
}

func New[C WithContext](base C) *Router[C] {
	return &Router[C]{
		base: base,
		mux:  http.NewServeMux(),
	}
}

func Default() *Router[*Context] {
	return &Router[*Context]{
		base: &Context{},
		mux:  http.NewServeMux(),
	}
}

type Handler[C WithContext] func(c C)
type Middleware[C WithContext] func(c C) bool

func (router *Router[C]) Add(pattern string, handler Handler[C], middlewares ...Middleware[C]) {
	router.mux.HandleFunc(pattern, func(res http.ResponseWriter, req *http.Request) {
		c := router.base
		*c.GetBase() = Context{
			Request:  req,
			Response: res,
		}

		for _, mw := range middlewares {
			if !mw(c) {
				return
			}
		}
		handler(c)
	})
}

func (router *Router[C]) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	router.mux.ServeHTTP(w, req)
}
