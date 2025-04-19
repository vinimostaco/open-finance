package controller

import (
	"encoding/json"
	"net/http"

	"github.com/vinimostaco/open-finance/src/service"
)


type AddTransactionInput struct {
	Nome  string  `json:"nome"`
	Valor float64 `json:"valor"`
	Tipo   string  `json:"tipo"`
}

func AddValue(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }	

	var input AddTransactionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Input inválido", http.StatusBadRequest)
		return
	}

	if input.Tipo != "income" && input.Tipo != "expense" {
		http.Error(w, "tipo deve ser: 'income' ou 'expense'", http.StatusBadRequest)
		return
	}

	returnedTransaction, err := service.AddTransaction(input.Nome, input.Valor, input.Tipo)

	if err != nil {
		http.Error(w, "Erro ao adicionar transação", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(returnedTransaction)
}

func GetValue() {

}

func RemoveValue() {

}

func UpdateValue() {

}