package slidesRoutes

import (
	"encoding/json"
	"net/http"
	"server/db"
	"server/db/models"
	"strconv"
)


func CreateSlide(w http.ResponseWriter, r *http.Request){
	
	presentationID, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server error"))
	}

	newSlide := &models.Slide{PresentationID: uint(presentationID)}

	db.Db.Create(newSlide)

	response, err := json.Marshal(newSlide)

	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server error"))
	}

	w.Write(response)

}

func DeleteSlide(w http.ResponseWriter, r *http.Request){
	
	slideID, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server error"))
	}

	db.Db.Delete(&models.Slide{}, slideID)

	w.Write([]byte("Slide  was deleted"))
}