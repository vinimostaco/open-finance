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

	response := map[string]interface{}{
		"success": true,
		"data": returnedTransaction,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
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

	tsx, err := util.CalculateTotalAmount(transactions)

	if err != nil{
		http.Error(w, "Erro ao calcular o total", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"success": true,
		"data": transactions,
		"totalAmount": tsx,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GetByName(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	
	defer r.Body.Close()

	var input model.GetTransactionByNameInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Input inválido", http.StatusBadRequest)
		return
	}

	if err := util.ValidateGetTransactionByNameInput(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	returnedTransactions, err := service.GetTransactionsByName(input.Nome)
	if err != nil {
		http.Error(w, "Erro ao buscar transação", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"success": true,
		"data": returnedTransactions,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GetByType(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	
	defer r.Body.Close()

	var input model.GetTransactionByTypeInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Input inválido", http.StatusBadRequest)
		return
	}

	if err := util.ValidateGetTransactionByTypeInput(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	returnedTransactions, err := service.GetTransactionsByType(input.Tipo)
	if err != nil {
		http.Error(w, "Erro ao buscar transação", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
			"success": true,
			"data": returnedTransactions,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}