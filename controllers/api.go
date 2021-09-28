package controllers

import (
	"api/di"
	"api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiController struct {
}

func (t *ApiController) Index(c *gin.Context) {
	db := di.Gorm()
	user := model.User{}
	db.Where("username = ?", "admin").Find(&user)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "api",
		"data":    user,
	})
}
