package login

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"server/db"
	"server/db/models"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)


type LoginBodyStruct struct {
	Email *string `json:"email"`
	Password *string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request)  {
	
	decoder := json.NewDecoder(r.Body)
	
	requestBody := new(LoginBodyStruct)

	notLoggedInUser := new(models.User)

	if err := decoder.Decode(&requestBody); err != nil {
		log.Fatal(err)
	}
	db.Db.Where("email = ?", requestBody.Email).First(notLoggedInUser)

	if err := bcrypt.CompareHashAndPassword([]byte(notLoggedInUser.Password), []byte(*requestBody.Password)); err != nil{
		log.Fatal(err)
	}


	claims := jwt.MapClaims{}
	claims["id"] = notLoggedInUser.ID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encodedToken, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil{
		log.Fatal(err)
	}

	 w.Write([]byte(encodedToken))
}