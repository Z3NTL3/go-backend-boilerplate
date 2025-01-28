package api

import (
	"bytes"
	"net/http"
	"z3ntl3/go-backend-boilerplate/server"

	"github.com/go-chi/chi"
	"github.com/gorilla/csrf"
	"github.com/spf13/viper"
)

func init() {
	server.RegistryList = append(server.RegistryList, Register)
}

func Register(router *server.Router) {
	router.Group(func(r1 chi.Router) {
		r1.Use(csrf.Protect([]byte(viper.GetString("csrf_token"))))
		r1.Route("/api", func(r2 chi.Router) {
			r2.Get("/echo", func(w http.ResponseWriter, r *http.Request) {
				tempOut := bytes.NewBufferString("")
				err := router.Templates.ExecuteTemplate(tempOut, "login.html", map[string]string{"text": "hello world"})
				if err != nil {
					w.Write([]byte("some error"))
					return
				}

				w.Write(tempOut.Bytes())
			})

			r2.Post("/echo", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("ok"))
			})
		})
	})
}
