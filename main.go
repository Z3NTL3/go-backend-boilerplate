package main

import (
	"fmt"
	"log"
	"net/http"
	"z3ntl3/go-backend-boilerplate/config"
	"z3ntl3/go-backend-boilerplate/server"
	stripesdk "z3ntl3/go-backend-boilerplate/stripe_sdk"

	docs "z3ntl3/go-backend-boilerplate/docs"

	"github.com/go-chi/chi"
	"github.com/spf13/viper"
	"github.com/stripe/stripe-go/v81/client"

	_ "z3ntl3/go-backend-boilerplate/server/routes/docs"
)

func main() {
	config.ExpandEnv()

	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Title = "API"
	docs.SwaggerInfo.Description = fmt.Sprintf("Application Programming Interface for application: %s", viper.GetString(config.AppName))
	docs.SwaggerInfo.Version = "0.1.0"

	app := (&server.
		Router{
		Mux:       chi.NewRouter(),
		StripeSDK: &stripesdk.StripeSDK{API: &client.API{}},
	}).Bootstrap()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("port")), app))
}
