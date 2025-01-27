package server

import (
	"net/http"

	"z3ntl3/boilerplate-gin/config"
	stripesdk "z3ntl3/boilerplate-gin/stripe_sdk"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Router struct {
	*gin.Engine
	*stripesdk.StripeSDK
}

type Route struct {
	method   string
	path     string
	handlers []gin.HandlerFunc
}

var Routes []Route = []Route{}

func init() {
	Routes = append(Routes, []Route{
		{
			method: http.MethodGet,
			path:   "/docs",
			handlers: []gin.HandlerFunc{
				func(ctx *gin.Context) {
					ctx.Redirect(301, "/api/index.html")
				},
			},
		},
	}...)
}

func (app *Router) Bootstrap() *Router {
	for _, route := range Routes {
		app.Handle(route.method, route.path, route.handlers...)
	}

	appMode := config.Development
	if viper.GetBool(config.DebugMode) {
		appMode = config.Production
	}

	var key string = viper.GetStringMapString(appMode)[config.StripeKey]
	app.StripeSDK.Init(key, nil)

	return app
}
