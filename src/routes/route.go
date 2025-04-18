package routes

import (
	"net/http"

	"github.com/vinimostaco/open-finance/src/controller"
)

func SetupRoutes() {
	http.HandleFunc("/addValue", controller.AddValue)
}