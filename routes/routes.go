package routes

import (
	"net/http"
	"asciiartweb/controllers"
)

func UseRoutes() {

	// Regular Page Routes
	http.HandleFunc("/home", controllers.HomeController)

	// Error Page Routes
	http.HandleFunc("/400", controllers.ErrorBadRequestController)
	http.HandleFunc("/404", controllers.ErrorNotFoundController)
	http.HandleFunc("/500", controllers.ErrorInternalController)

	// All other routes to be redirected to 404
	http.HandleFunc("/", controllers.ErrorNotFoundController)

}
