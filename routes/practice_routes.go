package routes

import (
	"github.com/gorilla/mux"
	"github.com/sadish665/food-ecommerce/controllers"
)

func InitAllPracticeRoutes(router *mux.Router) {
	router.HandleFunc("/success", controllers.HandleSuccessRequest)
	router.HandleFunc("/error",controllers.HandleErrorRequest)
}