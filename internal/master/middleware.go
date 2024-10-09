package master

import (
	"github.com/Ryan-eng-del/hurricane/internal/pkg/server"

	"github.com/gin-gonic/gin"
)

var masterMiddlewares = make([]gin.HandlerFunc, 0)

type MasterMiddleware struct{}

func init() {
	server.AddMiddlewares(&MasterMiddleware{})
}

func (m *MasterMiddleware) Fetch() []gin.HandlerFunc {
	return masterMiddlewares
}
