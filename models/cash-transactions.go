package models

import (
	"time"
)

type CashTransactionAttributes struct {
	ID                      string             `json:"id,omitempty"`
	ActualSettlementDate    string             `json:"actual-settlement-date"`
	Amount                  float64            `json:"amount"`
	CommentsLine1           string             `json:"comments-line-1"`
	CommentsLine2           string             `json:"comments-line-2"`
	CommentsLine3           string             `json:"comments-line-3"`
	CommentsLine4           string             `json:"comments-line-4"`
	CreatedAt               time.Time          `json:"actual-settlement-date"`
	Currency                string             `json:"currency"`
	CustomerReference       string             `json:"customer-reference"`
	EffectiveSettlementDate string             `json:"effective-settlement-date"`
	SettlementDate          string             `json:"settlement-date"`
	TradeDate               string             `json:"trade-date"`
	TransactionNumber       int                `json:"transaction-number"`
}

type CashTransactionData struct {
	ID            string            `json:"id,omitempty"`
	Type          string            `json:"type"`
	Attributes    AccountAttributes `json:"attributes"`
	Links         Links             `json:"links"`
	Relationships Relationships     `json:"relationships"`
}

type CashTransaction struct {
	Data CashTransactionData `json:"data"`
}

type CashTransactionsResponse struct {
	CollectionResponse
	Data []CashTransaction `json:"data"`
}
