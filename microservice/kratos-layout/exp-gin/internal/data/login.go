package data

import (
	"exp-gin/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type loginRepo struct {
	data *Data
	log  *log.Helper
}

func NewLoginRepo(data *Data, logger log.Logger) biz.LoginRepo {
	return &loginRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (l *loginRepo) Login(nickname string) (*biz.User, error) {
	var user biz.User
	err := l.data.db.Get(&user, "select * from tb_user where nickname = ?", nickname)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
