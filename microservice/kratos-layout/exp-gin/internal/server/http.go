package server

import (
	v1 "exp-gin/api/login/v1"
	"exp-gin/internal/conf"
	"exp-gin/internal/service"
	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, login *service.LoginService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	// 使用gin提供http服务
	srv := http.NewServer(opts...)
	r := gin.Default()
	r.Use(kgin.Middlewares(recovery.Recovery()))
	srv.HandlePrefix("/", r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "pong!",
		})
	})

	v1.RegisterLoginHTTPServer(r, login)

	return srv
}
