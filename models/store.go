package models

type StoreBranchQuerier interface {
	Create(book StoreBranch) (StoreBranch, error)
	List() ([]StoreBranch, error)
	// Update(book Book) (Book, error)
}

type StoreBranch struct {
	BranchID      *uint  `gorm:"primarykey"`
	BranchAddress string `form:"BranchAddress"`
}

func (StoreBranch) TableName() string {
	return "store_branch"
}
