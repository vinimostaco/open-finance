package routes

import (
	"net/http"

	"github.com/vinimostaco/open-finance/src/controller"
)

func SetupRoutes() {
	http.HandleFunc("/addTransactions", controller.Add)
	http.HandleFunc("/getTransactions", controller.Get)
	http.HandleFunc("/getTransactionsByName", controller.GetByName)
	http.HandleFunc("/getTransactionsByType", controller.GetByType)
}