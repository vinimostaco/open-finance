package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)



func AddValue(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }	
	body,err := io.ReadAll(r.Body)
	if err != nil{
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusInternalServerError)
		return
	}
	fmt.Printf("Corpo da requisição: %s\n", body)

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Erro ao fazer o unmarshal do JSON", http.StatusBadRequest)
		return
	}
	
}

func GetValue() {

}

func RemoveValue() {

}

func UpdateValue() {

}