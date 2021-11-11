package service

import (
	"context"
	pb "exp-gin/api/user/v1"
	"exp-gin/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

func NewUserService(user *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{
		user: user,
		log:  log.NewHelper(logger),
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	err := s.user.Create(ctx, &biz.User{
		Nickname: req.Nickname,
		Username: req.Username,
		Password: req.Password,
		Gender:   req.Gender,
		Mail:     req.Mail,
		Phone:    req.Phone,
	})

	return &pb.CreateUserReply{}, err
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	err := s.user.Update(ctx, req.Id, &biz.User{
		ID:       req.Id,
		Nickname: req.Nickname,
		Username: req.Username,
		Password: req.Password,
		Gender:   req.Gender,
		Mail:     req.Mail,
		Phone:    req.Phone,
	})

	return &pb.UpdateUserReply{}, err
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	s.log.Infof("input data %v:", req.Id)
	err := s.user.Delete(ctx, req.Id)
	return &pb.DeleteUserReply{}, err
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	user, err := s.user.Get(ctx, req.Id)
	reply := &pb.GetUserReply{User: &pb.User{
		Id:       user.ID,
		Nickname: user.Nickname,
		Username: user.Username,
		Password: user.Password,
		Gender:   user.Gender,
		Mail:     user.Mail,
		Phone:    user.Phone,
	}}

	return reply, err
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	users, err := s.user.List(ctx)
	reply := &pb.ListUserReply{}
	for _, user := range users {
		reply.Users = append(reply.Users, &pb.User{
			Id:       user.ID,
			Nickname: user.Nickname,
			Username: user.Username,
			Password: user.Password,
			Gender:   user.Gender,
			Mail:     user.Mail,
			Phone:    user.Phone,
		})
	}

	return reply, err
}
