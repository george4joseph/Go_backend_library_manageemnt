package controller

import (
	"hello/config"
	"hello/models"

	"github.com/gin-gonic/gin"
)

func Members_Get(c *gin.Context) {
	members := []models.Member{}
	config.DB.Find(&members)
	c.JSON(200, &members)
}

func Members_Post(c *gin.Context) {
	// mem := models.Member{Name: c.PostForm("name"), Email: c.PostForm("email")}
	// config.DB.Create(&mem)
	// c.JSON(200, &mem)
	
	var mem models.Member
	if err := c.BindJSON(&mem); err != nil {
		return
	}
	config.DB.Create(&mem)
	c.JSON(200, &mem)
	
}

func Member_Del(c *gin.Context) {
	var mem models.Member
	config.DB.Where("id=?", c.Param("id")).Delete(&mem)
	c.JSON(200, &mem)
}

func Member_Patch(c *gin.Context) {
	// var mem models.Member
	// config.DB.Where("id=?", c.Param("id")).First(&mem)
	// mem.Name = c.Query("name")
	// mem.Email = c.Query("email")
	// config.DB.Save(&mem)
	// c.JSON(200, &mem)
	var mem_db models.Member
	var mem models.Member
	config.DB.Where("id=?", c.Param("id")).First(&mem_db)
	if err := c.BindJSON(&mem); err != nil {
		return
	}
	config.DB.Model(&mem_db).Updates(mem)
	c.JSON(200, &mem)
}

func Member_Get(c *gin.Context) {
	var mem models.Member
	config.DB.Where("id=?", c.Param("id")).First(&mem)
	c.JSON(200, &mem)
}