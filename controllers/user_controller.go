package controllers

import (
	"net/http"

	"github.com/sadish665/food-ecommerce/models"
	"github.com/sadish665/food-ecommerce/responses"
)

func HandleGetAllUsers(w http.ResponseWriter, r *http.Request){
	users := models.GetAllUsers()
	responses.SuccessResponds(w, users, "Successfully fetched user list")
}