package model

import (
	"home/entity"
	"log"

	"gorm.io/gorm"
)

type Article struct {
	ID        string
	Title     string
	Content   string
	Tags      []*Tag `gorm:"many2many:article_tags"`
	Viewers   uint
	Favorites uint
	gorm.Model
}

func (m Article) ToEntity() *entity.Article {
	tags := []*entity.Tag{}
	if m.Tags != nil {
		for _, tag := range m.Tags {
			tags = append(tags, tag.ToEntity())
		}
	}
	log.Print(m.Tags)

	return &entity.Article{
		ID:        m.ID,
		Title:     m.Title,
		Content:   m.Content,
		Tags:      tags,
		Viewers:   m.Viewers,
		Favorites: m.Favorites,
	}
}
