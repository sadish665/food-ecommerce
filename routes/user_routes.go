package routes

import (
	"github.com/gorilla/mux"
	"github.com/sadish665/food-ecommerce/controllers"
)

func InitUserRoutes(router *mux.Router){
	router.HandleFunc("/users", controllers.HandleGetAllUsers).Methods("GET")
}