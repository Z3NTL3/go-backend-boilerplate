package server

import (
	"html/template"
	"z3ntl3/go-backend-boilerplate/config"
	requestsize "z3ntl3/go-backend-boilerplate/server/middlewares/request_size"
	stripesdk "z3ntl3/go-backend-boilerplate/stripe_sdk"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/viper"
)

type Router struct {
	*chi.Mux
	*stripesdk.StripeSDK
	Templates *template.Template
}

type Registry func(*Router)

var RegistryList = []Registry{}

func (r *Router) Bootstrap() *Router {
	// global middlewares
	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.StripSlashes,
		requestsize.RequestSize(viper.GetInt64(config.MaxReqBodySize)),
	)

	for _, registry := range RegistryList {
		registry(r)
	}

	key := viper.GetStringMapString(config.Production)[config.StripeKey]
	if viper.GetBool(config.DebugMode) {
		key = viper.GetStringMapString(config.Development)[config.StripeKey]
	}

	r.StripeSDK.Init(key, nil)
	return r
}
