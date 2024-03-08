package controllers

import (
	"github.com/Today-budget-app/backend/initializers"
	"github.com/Today-budget-app/backend/models"
	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	//get date from request body
	//create user
	//return it
	//GORM create guide

	var content struct {
		Username string
		Email    string
	}

	c.Bind(&content)

	user := models.User{Username: content.Username, Email: content.Email}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UserFetch(c *gin.Context) {
	//get user
	var content struct {
		Username string
		Email    string
	}

	c.Bind(&content)
	user := models.User{Username: content.Username, Email: content.Email}
	result := initializers.DB.Find(&user)

	if result.Error != nil {
		c.Status(404)
		return
	}

	//respond
	c.JSON(200, gin.H{
		"user": user,
	})
}
