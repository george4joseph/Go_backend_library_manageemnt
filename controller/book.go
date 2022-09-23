package controller

import (

	"hello/config"
	"hello/models"


	"github.com/gin-gonic/gin"
)

func Books_Get(c *gin.Context) {
	books := []models.Book{}
	config.DB.Find(&books)
	c.JSON(200, books)
}

func Book_Get(c *gin.Context) {
	var bok models.Book
	config.DB.Where("id=?", c.Param("id")).First(&bok)
	c.JSON(200, &bok)

}

func Books_Post(c *gin.Context) {

	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		return
	}
	config.DB.Create(&book)
	c.JSON(200, &book)
}

func Books_Update(c *gin.Context) {
	var book_db models.Book
	var book models.Book
	config.DB.Where("id=?", c.Param("id")).First(&book_db)
	if err := c.BindJSON(&book); err != nil {
		return
	}
	config.DB.Model(&book_db).Updates(book)
	c.JSON(200, &book)
}

