package controller

import (
	"encoding/json"
	"net/http"

	"github.com/vinimostaco/open-finance/src/model"
	"github.com/vinimostaco/open-finance/src/service"
	"github.com/vinimostaco/open-finance/src/util"
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

	if err := util.ValidateAddTransactionInput(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}	
								
	returnedTransaction, err := service.AddTransaction(input.Nome, input.Valor, input.Tipo)

	if err != nil {
		http.Error(w, "Erro ao adicionar transação", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(returnedTransaction)
}

func Get(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	transactions, err := service.GetTransactions()
	if err != nil {
		http.Error(w, "Erro ao buscar transações", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}