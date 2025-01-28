package server

import (
	"z3ntl3/go-backend-boilerplate/config"
	stripesdk "z3ntl3/go-backend-boilerplate/stripe_sdk"

	"github.com/go-chi/chi"
	"github.com/spf13/viper"
)

type Router struct {
	*chi.Mux
	*stripesdk.StripeSDK
}

type Registry func(*Router)

var RegistryList = []Registry{}

func (r *Router) Bootstrap() *Router {
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
