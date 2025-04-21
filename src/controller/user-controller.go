package controller

import (
	"encoding/json"
	"net/http"

	"github.com/vinimostaco/open-finance/src/model"
	"github.com/vinimostaco/open-finance/src/service"
)

func Register(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": false,
            "error":   "Método não permitido",
        })
        return
    }

    defer r.Body.Close()

    var input model.RegisterUserInput
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": false,
            "error":   "Input inválido",
        })
        return
    }

    _, err := service.RegisterUser(input)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": false,
            "error":   err.Error(),
        })
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "success": true,
    })
}

func Login(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": false,
            "error":   "Método não permitido",
        })
        return
    }

    defer r.Body.Close()

    var input model.LoginUserInput
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        w.Header().Set("Content-Type", "application/json")  
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": false,
            "error":   "Input inválido",
        })
        return
    }

    _, err := service.AuthenticateUser(input.Email, input.Password)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": false,
            "error":   err.Error(),
        })
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "success": true,
    })
}