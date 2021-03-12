package endpoints

import (
	"github.com/io-m/lenses/pkg/controllers"
	"github.com/io-m/lenses/pkg/endpoints/router"
)

var uc = controllers.NewUserController()

// RunApp is an entry point to the app
// It calls endpoints and running server
// It is called from cmd/main.go
func RunApp() {
	router := router.NewMuxRouter()
	router.Get("/users", uc.GetAll)
	router.Get("/users/{id}", uc.GetOne)
	router.Post("/users/new", uc.Save)
	router.Post("/users/login", uc.Login)
	router.Post("/users/logout", uc.Logout)
	router.Put("/users", uc.Update)
	router.Delete("/users", uc.Delete)
	router.ListenAndServe(":9900")
}
