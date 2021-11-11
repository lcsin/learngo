package service

import (
	"exp-gin/internal/biz"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

func NewLoginService(login *biz.LoginUsecase, logger log.Logger) *LoginService {
	return &LoginService{
		login: login,
		log:   log.NewHelper(logger),
	}
}

func (ls *LoginService) Login() func(ctx *gin.Context) {
	return ls.login.Login()
}
