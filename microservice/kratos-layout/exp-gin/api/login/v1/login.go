package v1

import (
	"exp-gin/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterLoginHTTPServer(r *gin.Engine, service *service.LoginService) {
	LoginServer(r, service)
}

func LoginServer(r *gin.Engine, service *service.LoginService) {
	r.GET("/login", service.Login())
}
