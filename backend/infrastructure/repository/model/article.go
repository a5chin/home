package model

import (
	"home/entity"

	"gorm.io/gorm"
)

type Article struct {
	ID        string
	Title     string
	Content   string
	Viewers   uint
	Favorites uint
	gorm.Model
}

func (m Article) ToEntity() *entity.Article {
	return &entity.Article{
		ID:        m.ID,
		Title:     m.Title,
		Content:   m.Content,
		Viewers:   m.Viewers,
		Favorites: m.Favorites,
	}
}
