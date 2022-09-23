package models

// import "gorm.io/gorm"

type Book struct {
	ID             int `gorm:"primaryKey"`
	Title              string
	Authors            string
	Average_rating     float64
	Isbn               string
	Isbn13             string
	Language_code      string
	Num_pages          int
	Ratings_count      int
	Text_reviews_count int
	Publication_date   string
	Publisher          string
	Total_quantity     int
	Available_quantity int
}
