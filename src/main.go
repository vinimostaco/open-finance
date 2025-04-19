package main

import (
	"fmt"
	"net/http"

	"github.com/vinimostaco/open-finance/src/config"
	"github.com/vinimostaco/open-finance/src/model"
	"github.com/vinimostaco/open-finance/src/routes"
)

func main() {
	config.Connect()
	config.DB.AutoMigrate(&model.Transaction{})

	routes.SetupRoutes()

	fmt.Println("🚀 Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil{
		fmt.Println("❌ Error starting server:", err)
		return
	}
}