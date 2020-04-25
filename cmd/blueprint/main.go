package main

import (
	"fmt"
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/apis"
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/config"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/docs"
)

// @title Seedtrace Static Site Generator Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email martin7.heinz@gmail.com

// @license.name MIT
// @license.url https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

// @BasePath /api/v1
func main() {
	fmt.Println("Starting to configure Server settings")
	// load application configurations
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}
	fmt.Println("Config successfully set")
	// Creates a router without any middleware by default
	r := gin.New()
	fmt.Println("New Gin Instance created...")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	fmt.Println("Using API Verion: ", v1.BasePath())
	{
		v1.GET("/users/:id", apis.GetUser)
	}

	r.Run(fmt.Sprintf(":%v", config.Config.ServerPort))
}
