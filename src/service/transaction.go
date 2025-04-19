package service

import (
	"github.com/google/uuid"
	"github.com/vinimostaco/open-finance/src/config"
	"github.com/vinimostaco/open-finance/src/model"
)

func AddTransaction(title string, amount float64, txType string) (*model.Transaction, error){
	transaction := &model.Transaction{
		ID: uuid.New(),
		Title: title,	
		Amount: amount,
		Type: txType,
	}

	if err := config.DB.Create(transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func GetTransactions() ([]model.Transaction, error) {
	var transactions []model.Transaction
	if err := config.DB.Find(&transactions).Error; err != nil {
		return nil, err
	}	
	return transactions, nil
}

func GetTransactionsByName(name string) ([]model.Transaction, error) {
	var transactions []model.Transaction
	if err := config.DB.Where("title = ?", name).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func GetTransactionsByType(txType string) ([]model.Transaction, error) {
	var transactions []model.Transaction
	if err := config.DB.Where("type = ?", txType).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}