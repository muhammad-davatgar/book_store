package repositories

import (
	"book_store/models"

	"gorm.io/gorm"
)

type StoreBranchRepo struct {
	DB *gorm.DB
}

func (d *StoreBranchRepo) Create(storeBranch models.StoreBranch) (models.StoreBranch, error) {
	err := d.DB.Create(&storeBranch).Error

	return storeBranch, err
}

func (d *StoreBranchRepo) List() ([]models.StoreBranch, error) {
	res := []models.StoreBranch{}

	err := d.DB.Find(&res).Error

	return res, err
}
