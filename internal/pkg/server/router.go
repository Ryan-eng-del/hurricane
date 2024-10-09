// Copyright 2024 Benjamin Lee <cyan0908@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

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
