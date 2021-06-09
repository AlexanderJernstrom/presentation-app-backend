package elementRoutes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/db"
	"server/db/models"
	"strconv"
)

type CreateElementBody struct{
	Type string `json:"type"`
	Content string `json:"content"`
}

type UpdateElementBody struct {
	PositionX float64 `json:"position_x"`
	PositionY float64 `json:"position_y"`
	Width float64 `json:"width"`
	Height float64 `json:"height"`
	Content string `json:"content"`
}

func CreateElement(w http.ResponseWriter, r *http.Request){
	slideID, err := strconv.Atoi(r.URL.Query().Get("id"))
	decoder := json.NewDecoder(r.Body)
	fmt.Println(slideID)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server error"))
	}

	requestBody := new(CreateElementBody)

	if err := decoder.Decode(requestBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server error"))
	}

	newElement := &models.Element{SlideID: uint(slideID), Type: requestBody.Type, PositionX: 0.5, PositionY: 0.5, Width: 0.2, Height: 0.2, Content: ""}

	db.Db.Create(newElement)

	response, err := json.Marshal(newElement)

	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server error"))
	}

	w.Write(response)
}

func UpdateElement(w http.ResponseWriter, r *http.Request){
	// Fields that can be updated: positionX, positionY, width, height, content
	elementID, err := strconv.Atoi(r.URL.Query().Get("id"))
	jsonDecoder := json.NewDecoder(r.Body)
	
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server error"))
	}

	var requestBody UpdateElementBody

	if err := jsonDecoder.Decode(&requestBody); err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server error"))
	}

	db.Db.Model(&models.Element{}).Where("id = ?", elementID).Updates(&models.Element{
		PositionX: requestBody.PositionX, 
		PositionY: requestBody.PositionY,
		Width: requestBody.Width, 
		Height: requestBody.Height,
		Content: requestBody.Content,
	})

	response, err := json.Marshal(requestBody)
	w.Write(response)
}

func DeleteElement(w http.ResponseWriter, r *http.Request) {
	elementID, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server error"))
	}

	db.Db.Delete(&models.Element{}, elementID)

	w.Write([]byte("Element was deleted"))
}