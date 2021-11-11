package biz

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
)

type LoginRepo interface {
	Login(nickname string) (*User, error)
}

type LoginUsecase struct {
	repo LoginRepo
	log  *log.Helper
}

func NewLoginUsecase(repo LoginRepo, logger log.Logger) *LoginUsecase {
	return &LoginUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (lu *LoginUsecase) Login() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var user User
		err := ctx.BindJSON(&user)
		if err != nil {
			lu.log.Errorf("invalid param: %v", err)
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "数据格式错误",
				"data": nil,
			})
			return
		}

		dbUser, err := lu.repo.Login(user.Nickname)
		switch {
		case err == sql.ErrNoRows:
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "用户名或密码错误",
				"data": nil,
			})
			return
		case err != nil:
			lu.log.Errorf("get user failed: %v", err)
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "系统异常",
				"data": nil,
			})
			return
		}

		if dbUser.Password != user.Password {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "用户名或密码错误",
				"data": nil,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "登录成功",
			"data": nil,
		})
	}
}
