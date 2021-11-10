package data

import (
	"context"
	"exp-sqlx/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ar *articleRepo) ListArticle(ctx context.Context) ([]*biz.Article, error) {
	list := make([]*biz.Article, 0)
	err := ar.data.db.Select(&list, "select id,title,content from articles")
	if err != nil {
		return list, err
	}
	return list, nil
}

func (ar *articleRepo) GetArticle(ctx context.Context, id int64) (*biz.Article, error) {
	var atc *biz.Article
	err := ar.data.db.Get(&atc, "select id,title,content from articles where id = ?", id)
	if err != nil {
		return atc, err
	}
	return atc, nil
}

func (ar *articleRepo) CreateArticle(ctx context.Context, article *biz.Article) error {
	_, err := ar.data.db.Exec("insert into articles(title,content) values(?,?)", article.Title, article.Content)
	if err != nil {
		return err
	}
	return nil
}

func (ar *articleRepo) UpdateArticle(ctx context.Context, id int64, article *biz.Article) error {
	_, err := ar.data.db.Exec("update articles set title = ? where id = ?", article.Title, id)
	if err != nil {
		return err
	}
	return nil
}

func (ar *articleRepo) DeleteArticle(ctx context.Context, id int64) error {
	_, err := ar.data.db.Exec("delete from articles where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
