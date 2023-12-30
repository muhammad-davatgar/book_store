package models

import (
	"time"
)

type BooksQuerier interface {
	Create(book Book) (Book, error)
	// Find(bookName string, branchID uint) (Book, error)
	// Update(book Book) (Book, error)
}

type Book struct {
	BookID      *uint `gorm:"primarykey"`
	BookName    string
	BookEdition uint
	PrintYear   time.Time
	BookPrice   uint
	BranchID    uint
	ShelfID     uint
	SpID        uint
}

func (Book) TableName() string {
	return "books"
}
