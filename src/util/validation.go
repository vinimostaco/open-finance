package util

import (
	"errors"

	"github.com/vinimostaco/open-finance/src/model"
)


func ValidateAddTransactionInput(input model.AddTransactionInput) error {
    if input.Nome == "" {
        return errors.New("campo 'nome' é obrigatório")
    }
    if input.Valor == 0 {
        return errors.New("campo 'valor' deve ser diferente de zero")
    }
    if input.Tipo != "income" && input.Tipo != "expense" {
        return errors.New("campo 'tipo' deve ser 'income' ou 'expense'")
    }
    return nil
}

func ValidateGetTransactionByNameInput(input model.GetTransactionByNameInput) error {
    if input.Nome == "" {
        return errors.New("campo 'nome' é obrigatório")
    }
    return nil
}

func ValidateGetTransactionByTypeInput(input model.GetTransactionByTypeInput) error {
    if input.Tipo == "" {
        return errors.New("campo 'tipo' é obrigatório")
    }
    if input.Tipo != "income" && input.Tipo != "expense" {
        return errors.New("campo 'tipo' deve ser 'income' ou 'expense'")
    }
    return nil
}