package api

import "gorm.io/gorm"

type Plan struct {
	gorm.Model
	Name        string `json:"name"`
	Content     string `json:"content"`
	Start_date  string `json:"start_date"`
	Finish_date string `json:"finish_date"`
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
