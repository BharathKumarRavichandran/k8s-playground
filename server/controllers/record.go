package controllers

import (
	"net/http"

	//"github.com/BharathKumarRavichandran/k8s-playground/server/db"
	//"github.com/BharathKumarRavichandran/k8s-playground/server/models"

	"github.com/BharathKumarRavichandran/k8s-playground/server/utils"
	"github.com/gin-gonic/gin"
)

type RecordController struct{}

func (ctrl RecordController) GetRecord(c *gin.Context) {

	id := c.Param("id")

	responseMessage := "What are ya doing here?" + id
	utils.Logger.Infof("%s", responseMessage)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     responseMessage,
	})
	return
}

func (ctrl RecordController) GetAllRecords(c *gin.Context) {

	responseMessage := "What are ya doing here?"
	utils.Logger.Infof("%s", responseMessage)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     responseMessage,
	})
	return
}

func (ctrl RecordController) PushRecord(c *gin.Context) {

	responseMessage := "What are ya doing here?"
	utils.Logger.Infof("%s", responseMessage)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     responseMessage,
	})
	return
}
