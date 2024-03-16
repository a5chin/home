package repository

import (
	"context"
	"errors"
	"home/entity"
	"home/infrastructure/driver"
	"home/infrastructure/repository/model"
	"net/http"

	"gorm.io/gorm"
)

type ArticleRepository struct {
}

func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{}
}

func (r ArticleRepository) GetArticles(ctx context.Context, limit, offset *int) ([]*entity.Article, error) {
	var records []*model.Article
	articles := []*entity.Article{}

	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	if limit != nil {
		db = db.Limit(*limit)
	}
	if offset != nil {
		db = db.Offset(*offset)
	}
	if err := db.Order("updated_at DESC").Find(&records).Error; err != nil {
		return nil, err
	}

	for _, record := range records {
		articles = append(articles, record.ToEntity())
	}

	return articles, nil
}

func (r ArticleRepository) GetArticleByID(ctx context.Context, articleId string) (*entity.Article, error) {
	var article *model.Article

	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	if err := db.First(&article, "ID = ?", articleId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entity.WrapError(http.StatusNotFound, err)
		}
		return nil, err
	}

	return article.ToEntity(), nil
}

func (r ArticleRepository) IncrementTarget(ctx context.Context, articleId, target string, number uint) error {
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	if err := db.Model(&model.Article{}).Where("ID = ?").Update("viewers", number+1).Error; err != nil {
		return err
	}
	return nil
}
