package responses

import (
	"encoding/json"
	"net/http"
)

func customResponse(w http.ResponseWriter, statusCode int, data interface{}, msg string){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	dataToSend := map[string]interface{}{
		"message": msg,
		"data": data,
	}
	dataByte, _ := json.Marshal(dataToSend)
	w.Write(dataByte)
}

func SuccessResponds(w http.ResponseWriter, data interface{}, msg string){
	customResponse(w, http.StatusOK, data, msg)
}

func ErrosResponds(w http.ResponseWriter,statusCode int, msg string){
	customResponse(w, statusCode, nil, msg)
}