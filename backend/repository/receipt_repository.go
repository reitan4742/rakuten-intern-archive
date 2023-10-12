package repository

import (
	"backend/model"

	"gorm.io/gorm"
)

type IReceiptRepository interface {
	GetReceiptByUsername(receipts *[]model.Receipt, username string) error
	CreateReceipt(receipt *model.Receipt) error
}

type receiptRepository struct {
	db *gorm.DB
}

func NewReceiptRepository(db *gorm.DB) IReceiptRepository {
	return &receiptRepository{db}
}

func (rr *receiptRepository) GetReceiptByUsername(receipts *[]model.Receipt, username string) error {
	if err := rr.db.Where("username = ?", username).Order("created_at desc").Limit(10).Find(receipts).Error; err != nil {
		return err
	}
	return nil
}

func (rr *receiptRepository) CreateReceipt(receipt *model.Receipt) error {
	if err := rr.db.Create(receipt).Error; err != nil {
		return err
	}
	return nil
}
