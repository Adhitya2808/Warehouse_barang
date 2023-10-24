package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username	string 
	Password	string 
	Name		string 
	Role		string 
	Email		string 
}

