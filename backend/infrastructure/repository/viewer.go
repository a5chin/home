package repository

import (
	"context"
	"home/infrastructure/driver"

	"gorm.io/gorm"
)

type ViewerRepository struct {
}

func NewViewerRepository() *ViewerRepository {
	return &ViewerRepository{}
}

func (r ViewerRepository) GetTotalViewers(ctx context.Context) (uint, error) {
	var record uint

	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)

	if err := db.Table("lp").Select("viewers").Find(&record).Error; err != nil {
		return 0, err
	}
	return record, nil
}
