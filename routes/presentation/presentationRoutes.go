package presentationRoutes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"server/db"
	"server/db/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

type CreatePresentationBody struct {
	Name string `json:"name"`
}



func GetAllPresentations(w http.ResponseWriter, r *http.Request){
	
	jwtMap  := r.Context().Value("id").(jwt.MapClaims)
	
	
	var users = new(models.User)

	
	userID := int(jwtMap["id"].(float64))

	fmt.Println(userID)


	db.Db.Model(&users).Preload("Presentations").First(&users, userID)

	

	response, err := json.Marshal(users.Presentations)
	if err != nil{
		log.Fatal(err)
	}
	
	w.Write(response)
}

func GetPresentation(w http.ResponseWriter, r *http.Request){
	
	/*jwtMap := r.Context().Value("id").(jwt.MapClaims)

	userID := int(jwtMap["id"].(float64))
*/
	var presentation models.Presentation

	fmt.Println(reflect.TypeOf(r.URL.Query().Get("id")))

	presentationID, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("server error"))
	}

	db.Db.Model(&presentation).Preload("Slides.Elements").Preload("User").First(&presentation, presentationID)

	jsonResponse, err := json.Marshal(presentation)

	w.Write(jsonResponse)
}




func CreatePresentation(w http.ResponseWriter, r *http.Request){
	
	jwtMap := r.Context().Value("id").(jwt.MapClaims)
	userID := int(jwtMap["id"].(float64))

	decoder := json.NewDecoder(r.Body)

	newPresentation := new(CreatePresentationBody) 

	if err := decoder.Decode(&newPresentation); err != nil{
		log.Fatal(err)
	}
	var user *models.User

	db.Db.Model(&models.User{}).Find(&user, userID)


	presentation := models.Presentation{Name: newPresentation.Name, UserID: uint64(userID)}


	db.Db.Create(&presentation)
	

	db.Db.Model(user).Association("Presentations").Append(presentation)

	jsonResponse, err := json.Marshal(presentation)

	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server error"))
	}

	w.Write(jsonResponse)

}

func DeletePresentation(w http.ResponseWriter, r *http.Request){
	
	presentationID, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server error"))
	}

	db.Db.Model(models.Slide{}).Where("presentation_id = ?", presentationID).Unscoped().Delete(models.Slide{})
	db.Db.Unscoped().Delete(&models.Presentation{}, presentationID)

	db.Db.Model(&models.User{}).Association("Presentations")

	w.WriteHeader(200)
	w.Write([]byte("Presentation was deleted"))
}

func UpdatePresentation(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	presentationID, err := strconv.Atoi(r.URL.Query().Get("id"))
	var presentation models.Presentation

	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server error"))
	}

	if err := decoder.Decode(&presentation); err != nil{
		log.Fatal(err)
	}

	db.Db.Model(&models.Presentation{}).Where("id = ?", presentationID).Updates(&models.Presentation{
		Name: presentation.Name, 
		Slides: presentation.Slides,
	})
	

	response, err := json.Marshal(&presentation)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(response)
}
