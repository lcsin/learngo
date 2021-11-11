package service

import (
	pb "exp-gin/api/user/v1"
	"exp-gin/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService, NewLoginService)

type UserService struct {
	pb.UnimplementedUserServiceServer

	user *biz.UserUsecase

	log *log.Helper
}

type LoginService struct {
	login *biz.LoginUsecase

	log *log.Helper
}
