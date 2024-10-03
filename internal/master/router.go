package master

import (
	"hurricane/internal/pkg/server"

	"github.com/gin-gonic/gin"
)

func init() {
	server.RegisterRoute(&MasterRoute{})
}

type MasterRoute struct {
}

func (mr *MasterRoute) Register(r *gin.Engine) {

}
