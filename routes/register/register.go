package register

import (
	"encoding/json"
	"log"
	"net/http"
	"server/db"
	"server/db/models"

	"golang.org/x/crypto/bcrypt"
)

type RegisterRequestBody struct{
	Email *string	`json:"email"`
	Name *string 	`json:"name"`
	Password *string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request)  {
	
	decoder := json.NewDecoder(r.Body)
	var requestBody *RegisterRequestBody

	if err := decoder.Decode(&requestBody); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	formattedPassword := []byte(*requestBody.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(formattedPassword, bcrypt.DefaultCost)
	if err != nil{
		log.Fatal(err)
	}
	rightPassword := string(hashedPassword)

	requestBody.Password = &rightPassword

	 db.Db.Create(&models.User{Name: *requestBody.Name, Email: *requestBody.Email, Password: *requestBody.Password})

	var responseMap map[string]interface{} = map[string]interface{}{"success": true} 

	res, err := json.Marshal(responseMap)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	 w.Write(res)
}