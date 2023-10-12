package usecase

import (
	"backend/model"
	"backend/repository"
)

type IReceiptUsecase interface {
	CreateReceipt(receipt model.Receipt) error
	GetCreateReceiptResult(username string) (model.ReceiptResultResponse, error)
	GetHistoryList(username string) (model.HistoryListResponse, error)
}

type receiptUsecase struct {
	rr repository.IReceiptRepository
}

func NewReceiptUsecase(rr repository.IReceiptRepository) IReceiptUsecase {
	return &receiptUsecase{rr}
}

func (ru *receiptUsecase) CreateReceipt(receipt model.Receipt) error {
	if err := ru.rr.CreateReceipt(&receipt); err != nil {
		return err
	}
	return nil
}

func (ru *receiptUsecase) GetCreateReceiptResult(username string) (model.ReceiptResultResponse, error) {
	receipts := []model.Receipt{}
	if err := ru.rr.GetReceiptByUsername(&receipts, username); err != nil {
		return model.ReceiptResultResponse{}, err
	}
	rrr := model.ReceiptResultResponse{}
	rrr.GetExp = receipts[0].GetExp
	return rrr, nil
}

func (ru *receiptUsecase) GetHistoryList(username string) (model.HistoryListResponse, error) {
	receipts := []model.Receipt{}
	if err := ru.rr.GetReceiptByUsername(&receipts, username); err != nil {
		return model.HistoryListResponse{}, err
	}
	hlr := model.HistoryListResponse{}
	for _, receipt := range receipts {
		hl := model.HistoryList{}
		hl.CreatedAt = receipt.CreatedAt
		hl.GetExp = receipt.GetExp
		hlr.HistoryList = append(hlr.HistoryList, hl)
	}
	return hlr, nil
}
