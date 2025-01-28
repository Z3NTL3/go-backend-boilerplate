package docs

import (
	"fmt"
	"net/http"
	"z3ntl3/go-backend-boilerplate/server"

	"github.com/go-chi/chi"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
)

func init() {
	server.RegistryList = append(server.RegistryList, Register)
}

func Register(router *server.Router) {
	router.Group(func(r chi.Router) {
		r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/docs/index.html", http.StatusMovedPermanently)
		})

		r.Get("/docs/*", httpSwagger.Handler(
			httpSwagger.URL(fmt.Sprintf("http://localhost:%d/docs/doc.json", viper.GetInt("port"))), //The url pointing to API definition
		))
	})
}
