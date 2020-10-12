package router

import (
	controllers "github.com/BharathKumarRavichandran/k8s-playground/server/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	rootController := new(controllers.RootController)
	recordController := new(controllers.RecordController)

	router.GET("/", rootController.PongController)

	record := router.Group("/record")
	record.Use()
	{
		record.GET("/id/:id", recordController.GetRecord)
		record.GET("/all", recordController.GetAllRecords)

		record.POST("/push", recordController.PushRecord)
	}
}
