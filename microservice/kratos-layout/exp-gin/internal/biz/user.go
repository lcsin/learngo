package biz

import "context"

type User struct {
	ID       uint64 `db:"id"`
	Nickname string `db:"nickname"`
	Username string `db:"username"`
	Password string `db:"password"`
	Gender   int64  `db:"gender"`
	Mail     string `db:"mail"`
	Phone    string `db:"phone"`
}

type UserRepo interface {
	ListUser(ctx context.Context) ([]*User, error)
	GetUser(ctx context.Context, id uint64) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id uint64) error
	UpdateUser(ctx context.Context, id uint64, user *User) error
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) List(ctx context.Context) (users []*User, err error) {
	users, err = uc.repo.ListUser(ctx)
	if err != nil {
		return
	}
	return
}

func (uc *UserUsecase) Get(ctx context.Context, id uint64) (user *User, err error) {
	user, err = uc.repo.GetUser(ctx, id)
	if err != nil {
		return
	}
	return
}

func (uc *UserUsecase) Create(ctx context.Context, user *User) error {
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUsecase) Delete(ctx context.Context, id uint64) error {
	return uc.repo.DeleteUser(ctx, id)
}

func (uc *UserUsecase) Update(ctx context.Context, id uint64, user *User) error {
	return uc.repo.UpdateUser(ctx, id, user)
}
