package main

import (
	"fmt"
	"z3ntl3/boilerplate-gin/config"
	"z3ntl3/boilerplate-gin/server"
	stripesdk "z3ntl3/boilerplate-gin/stripe_sdk"

	docs "z3ntl3/boilerplate-gin/docs"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stripe/stripe-go/v81/client"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.ExpandEnv()

	app := (&server.
		Router{
		Engine:    gin.Default(),
		StripeSDK: &stripesdk.StripeSDK{API: &client.API{}},
	}).
		Bootstrap()

	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Title = "API"
	docs.SwaggerInfo.Description = fmt.Sprintf("Application Programming Interface for application: %s", viper.GetString(config.AppName))
	docs.SwaggerInfo.Version = "0.1.0"

	app.GET("/api/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	app.Run(fmt.Sprintf(":%d", viper.GetInt("port")))
}
