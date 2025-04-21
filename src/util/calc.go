package util

import "github.com/vinimostaco/open-finance/src/model"


func CalculateTotalAmount(transactions []model.Transaction) (float64, error) {
	var total float64
	for _, transaction := range transactions {
		total += transaction.Amount
	}
	return total, nil
}