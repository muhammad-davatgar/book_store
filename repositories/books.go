package repositories

import (
	"book_store/models"

	"gorm.io/gorm"
)

type BookRepo struct {
	DB *gorm.DB
}

func (d *BookRepo) Create(book models.Book) (models.Book, error) {
	err := d.DB.Create(&book).Error

	return book, err
}
