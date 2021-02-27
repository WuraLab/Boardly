package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robogeek95/ginCrud/models"
)

// GetWelcome shows greeting message
func GetWelcome(c *gin.Context) {
	message := "Hello world :)..."

	c.JSON(http.StatusOK, message)
}

//GetMaterials ... Get all users
func GetMaterials(c *gin.Context) {
	var material []models.Material
	err := models.GetAllMaterials(&material)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, material)
	}
}

//CreateMaterial ... Create User
func CreateMaterial(c *gin.Context) {
	var material models.Material

	if err := c.ShouldBindJSON(&material); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := models.CreateMaterial(&material)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, material)
	}
}

//RegisterAdmin ... Create User
func RegisterAdmin(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := models.RegisterAdmin(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
