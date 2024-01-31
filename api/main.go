package main

import (
	"fmt"
	"log"
	api "vgapi/handler"

	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatal("Error getting swagger")
	}
	swagger.Servers = nil

	ssi := api.NewGameApi()
	handler := api.NewStrictHandler(ssi, nil)

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hello, World!"})
	})
	r.Use(middleware.OapiRequestValidator(swagger))

	api.RegisterHandlers(r, handler)

	fmt.Println("listening on localhost:8080...")
	r.Run()
}
