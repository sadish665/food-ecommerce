package routes

import "github.com/gorilla/mux"

func InitAllRoutes() *mux.Router{
	router := mux.NewRouter()
	InitAllPracticeRoutes(router)
	InitUserRoutes(router)
	return router
}