package server

import "github.com/gin-gonic/gin"

type Middleware interface {
	Fetch() []gin.HandlerFunc
}

var injectMiddlewares = []gin.HandlerFunc{}

func AddMiddlewares(middlewares ...Middleware) error {
	for _, middleware := range middlewares {
		injectMiddlewares = append(injectMiddlewares, middleware.Fetch()...)
	}
	return nil
}
