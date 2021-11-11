package data

import (
	"context"
	"exp-gin/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *userRepo) ListUser(ctx context.Context) ([]*biz.User, error) {
	users := make([]*biz.User, 0)
	err := u.data.db.Select(&users, "select * from tb_user")
	if err != nil {
		return users, err
	}
	return users, nil
}

func (u *userRepo) GetUser(ctx context.Context, id uint64) (*biz.User, error) {
	var user *biz.User
	err := u.data.db.Get(&user, "select * from tb_user where id = ?", id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	sql := `
insert into tb_user(nickname,username,password,gender,mail,phone)
values (?,?,?,?,?,?)
`
	_, err := u.data.db.Exec(sql, user.Nickname, user.Username, user.Password, user.Gender, user.Mail, user.Phone)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) DeleteUser(ctx context.Context, id uint64) error {
	_, err := u.data.db.Exec("delete from tb_user where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) UpdateUser(ctx context.Context, id uint64, user *biz.User) error {
	_, err := u.data.db.Exec("update user set nickname = ? where id = ?", user.Nickname, id)
	if err != nil {
		return err
	}
	return nil
}
