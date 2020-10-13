package main

import (
	"github.com/BharathKumarRavichandran/k8s-playground/server/db"
	"github.com/BharathKumarRavichandran/k8s-playground/server/utils"
	"github.com/BharathKumarRavichandran/k8s-playground/server/utils/kafka"
	"github.com/joho/godotenv"

	indexRouter "github.com/BharathKumarRavichandran/k8s-playground/server/routers"

	"github.com/gin-gonic/gin"
)

func RealMain() {
	config := utils.GetConfiguration()
	utils.Init(config)

	db.Init(config)
	defer db.Close()

	kafka.Init(config)

	// Read env file
	err := godotenv.Load()
	if err != nil {
		utils.Logger.Fatal("Error loading .env file")
	}

	// Configure router
	router := gin.Default()

	// Configure routes; Redirect all routes to indexRouter
	indexRouter.Routes(router)

	// Start router and serve application
	utils.Logger.Infof("Listening and serving HTTP on %s", string(config.ServerPort))
	utils.Logger.Fatal(router.Run(config.ServerPort))
}

func main() {
	for {
		RealMain()
	}
}
