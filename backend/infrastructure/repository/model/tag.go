package model

import "home/entity"

type Tag struct {
	ID   string
	Name string
}

func (m Tag) ToEntity() *entity.Tag {
	return &entity.Tag{
		ID:   m.ID,
		Name: m.Name,
	}
}
