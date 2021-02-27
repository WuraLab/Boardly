package controllers

import "github.com/gin-gonic/gin"

// Login controller
func Login(c *gin.Context) {

	var id int
	var name, email, password, createdAt, updatedAt string

	err := row.Scan(&id, &name, &password, &email, &createdAt, &updatedAt)
}
