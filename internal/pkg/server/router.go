package server

import "github.com/gin-gonic/gin"

type Router interface {
	Register(r *gin.Engine)
}

var routers []Router

func RegisterRoute(rs ...Router) {
	routers = append(routers, rs...)
}

func (s *BaseApiServer) initRouter(r *gin.Engine) {
	for _, router := range routers {
		router.Register(r)
	}
}
