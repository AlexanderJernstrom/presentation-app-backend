package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string `gorm:"unique" json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string 
	ID uint `gorm:"primaryKey" json:"id"`
	Presentations []Presentation `json:"presentations"`

}


type Presentation struct{
	ID uint `gorm:"primaryKey;unique" json:"id"`
	Name string `gorm:"unique" json:"name"`
	UserID uint64 `gorm:"not null" json:"user_id"`
	User User `gorm:"foreignKey:UserID" json:"user"`
	Slides []Slide `json:"slides"`
}

type Slide struct{
	ID uint `gorm:"primaryKey" json:"id"`
	PresentationID uint `json:"presentation_id"`
	Elements []Element `json:"elements"`	
}

type Element struct {
	
	
	SlideID uint `json:"slide_id"`
	
	//What type of element it is, it can be image or text
	ID uint `gorm:"primaryKey" json:"id"`
	Type string `json:"type"`
	PositionX float64 `json:"position_x"`
	PositionY float64 `json:"position_y"`
	Width float64 `json:"width"`
	Height float64 `json:"height"`
	//Either the url of the image or the value fo the textfield depending on the type of the element
	Content string `json:"content"`
}
