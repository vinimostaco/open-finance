package main

import (
	"fmt"
	"net/http"

	"github.com/vinimostaco/open-finance/src/routes"
)

func main() {
	routes.SetupRoutes()

	fmt.Println("🚀 Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}