package model

import "time"

type Receipt struct {
	ReceiptID int       `json:"receipt_id" gorm:"primary_key"`
	Username  string    `json:"username"`
	GetExp    int       `json:"getexp"`
	CreatedAt time.Time `json:"createdat"`
}

type ReceiptRequest struct {
	Username string `json:"username"`
	Receipt  string `json:"receipt"`
}

type ReceiptResponse struct {
	Result bool `json:"result"`
}

type ReceiptResultResponse struct {
	GetExp int `json:"getexp"`
}

type HistoryList struct {
	CreatedAt time.Time `json:"createdat"`
	GetExp    int       `json:"getexp"`
}

type HistoryListResponse struct {
	HistoryList []HistoryList `json:"historylist"`
}
