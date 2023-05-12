package middleware

import (
	"fmt"
	"go_valid/models"
	"log"
	"net/http"

	"encoding/json"

	"go_valid/controllers"

	_ "github.com/lib/pq"
)

type responseValidation struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

func GetAllData(w http.ResponseWriter, r *http.Request) {

	data, err := controllers.GetAllData()
	if err != nil {
		log.Fatalf("data retrival failed! %v", err)
	}
	json.NewEncoder(w).Encode(data)
}

func CreateForm(w http.ResponseWriter, r *http.Request) {
	var form models.Form
	json.NewDecoder(r.Body).Decode(&form)
	if form.Firstname == "" {
		json.NewEncoder(w).Encode(responseValidation{Error: "Firstname is balnk"})
		return
	} else if form.Lastname == "" {
		json.NewEncoder(w).Encode(responseValidation{Error: "Lastname is blank"})
		return
	} else if form.Email == "" {
		json.NewEncoder(w).Encode(responseValidation{Error: "Email is blank"})
		return
	} else if form.Phonenumber == 0 {
		json.NewEncoder(w).Encode(responseValidation{Error: "Phone Number is blank"})
		return
	}

	err := controllers.CreateForm(&form)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("Validation failed : %v ", err))
	}
	json.NewEncoder(w).Encode(form)
}
