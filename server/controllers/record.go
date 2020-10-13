package controllers

import (
	"net/http"
	"time"

	"github.com/gocql/gocql"

	"github.com/BharathKumarRavichandran/k8s-playground/server/db"
	"github.com/BharathKumarRavichandran/k8s-playground/server/models"
	"github.com/BharathKumarRavichandran/k8s-playground/server/utils"
	"github.com/BharathKumarRavichandran/k8s-playground/server/utils/kafka"
	"github.com/gin-gonic/gin"
)

type RecordController struct{}

func (ctrl RecordController) GetRecord(c *gin.Context) {

	id := c.Param("id")
	iter := db.Session.Query("SELECT * FROM records WHERE id = ?", id).Iter()

	var record *models.Record = nil
	m := map[string]interface{}{}

	if iter.MapScan(m) {
		record = &models.Record{
			ID:          m["id"].(gocql.UUID),
			Message:     m["message"].(string),
			CreatedDate: m["created_date"].(time.Time),
		}
		m = map[string]interface{}{}
	}

	responseMessage := "success, ID:" + id
	utils.Logger.Infof("%s", responseMessage)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     responseMessage,
		"data":        record,
	})
	return
}

func (ctrl RecordController) GetAllRecords(c *gin.Context) {

	var records []models.Record
	m := map[string]interface{}{}

	iter := db.Session.Query("SELECT * FROM records").Iter()

	for iter.MapScan(m) {
		records = append(records, models.Record{
			ID:          m["id"].(gocql.UUID),
			Message:     m["message"].(string),
			CreatedDate: m["created_date"].(time.Time),
		})
		m = map[string]interface{}{}
	}

	responseMessage := "success"
	utils.Logger.Infof("%s", responseMessage)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     responseMessage,
		"data":        records,
	})
	return
}

func (ctrl RecordController) PushRecord(c *gin.Context) {

	message := c.PostForm("message")

	// Push message to configured Kafka topic
	kafka.ProduceMessage(message)

	responseMessage := "success"
	utils.Logger.Infof("%s", responseMessage)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     responseMessage,
	})
	return
}
