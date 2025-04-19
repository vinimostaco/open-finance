package controller

import (
	"encoding/json"
	"net/http"

	"github.com/vinimostaco/open-finance/src/model"
	"github.com/vinimostaco/open-finance/src/service"
)


func Add(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }	

	defer r.Body.Close()

	var input model.AddTransactionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Input inválido", http.StatusBadRequest)
		return
	}

	if input.Tipo != "income" && input.Tipo != "expense" {
		http.Error(w, "campo 'tipo' deve ser: 'income' ou 'expense'", http.StatusBadRequest)
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