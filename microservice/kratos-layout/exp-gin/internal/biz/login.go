package biz

import (
	"database/sql"
	"exp-gin/internal/tool"
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
			tool.Error(ctx, http.StatusInternalServerError, "参数有误")
			return
		}

		dbUser, err := lu.repo.Login(user.Nickname)
		switch {
		case err == sql.ErrNoRows:
			tool.Error(ctx, http.StatusNotFound, "用户名或密码错误")
			return
		case err != nil:
			lu.log.Errorf("get user failed: %v", err)
			tool.Error(ctx, http.StatusInternalServerError, "系统异常")
			return
		}

		if dbUser.Password != user.Password {
			tool.Error(ctx, http.StatusNotFound, "用户名或密码错误")
			return
		}

		claims := tool.MyClaims{
			ID:       dbUser.ID,
			Username: dbUser.Username,
			Nickname: dbUser.Nickname,
		}
		token, err := tool.GenToken(claims)
		if err != nil {
			lu.log.Errorf("cannot generate jwt, err:", err)
			tool.Error(ctx, http.StatusInternalServerError, "系统异常")
		}

		tool.OK(ctx, token, "登录成功")
	}
}
