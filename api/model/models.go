package api

import (
	"gorm.io/gorm"
	"time"
)

type Plan struct {
	gorm.Model
	Name        string    `json:"name,omitempty"`
	Content     string    `json:"content,omitempty"`
	State       string    `json:"state,omitempty"`
	Start_date  time.Time `json:"start_date,omitempty"`
	Finish_date time.Time `json:"finish_date,omitempty"`
	Username    string    `json:"username,omitempty"`
	UserId      uint      `json:"user_id,omitempty"`
}

type User struct {
	gorm.Model
	UserName string `gorm:"unique;not null"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password" binding:"required,min=8"`
	Email    string `json:"email,omitempty"`
	Plans    []Plan
}
