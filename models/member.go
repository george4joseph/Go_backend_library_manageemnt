package models

import "gorm.io/gorm"


type Member struct {
	gorm.Model
	ID          int
	Name        string 
	Email       string
	Books_issued string
	Total_debt  int
}
