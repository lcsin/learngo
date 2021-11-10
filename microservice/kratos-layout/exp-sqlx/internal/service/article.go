package service

import (
	"context"
	pb "exp-sqlx/api/article/v1"
	"exp-sqlx/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type ArticleService struct {
	pb.UnimplementedArticleServiceServer

	log *log.Helper

	article *biz.ArticleUsecase
}

func NewArticleService(article *biz.ArticleUsecase, logger log.Logger) *ArticleService {
	return &ArticleService{
		article: article,
		log:     log.NewHelper(logger),
	}
}

func (s *ArticleService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleReply, error) {
	return &pb.CreateArticleReply{}, nil
}
func (s *ArticleService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
	return &pb.UpdateArticleReply{}, nil
}
func (s *ArticleService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleReply, error) {
	return &pb.DeleteArticleReply{}, nil
}
func (s *ArticleService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleReply, error) {
	return &pb.GetArticleReply{}, nil
}
func (s *ArticleService) ListArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.ListArticleReply, error) {
	list, err := s.article.List(ctx)
	if err != nil {
		s.log.Error(pb.ErrorUserNotFound("user %s not found", "kratos"))
		return nil, err
	}
	reply := &pb.ListArticleReply{}
	for _, l := range list {
		reply.Results = append(reply.Results, &pb.Article{
			Id:      l.ID,
			Title:   l.Title,
			Content: l.Content,
		})
	}
	return reply, nil
}
