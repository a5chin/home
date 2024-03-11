package repository

import (
	"context"
	"home/entity"
	"home/infrastructure/driver"
	"home/infrastructure/repository/model"

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
