package controllers

import (
	"net/http"

	//"github.com/BharathKumarRavichandran/k8s-playground/server/db"
	//"github.com/BharathKumarRavichandran/k8s-playground/server/models"

	"github.com/BharathKumarRavichandran/k8s-playground/server/utils"
	"github.com/gin-gonic/gin"
)

type RootController struct{}

func (ctrl RootController) PongController(c *gin.Context) {

	responseMessage := "What are you doing here?"
	utils.Logger.Infof("%s", responseMessage)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     responseMessage,
	})
	return
}
