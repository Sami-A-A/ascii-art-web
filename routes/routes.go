package routes

import (
	"net/http"
	"asciiartweb/controllers"
)

func UseRoutes() {

	// Regular Page Routes
	http.HandleFunc("/", controllers.HomeController)
	http.HandleFunc("/ascii-art", controllers.AsciiArtController)

	// Error Page Routes
	http.HandleFunc("/400", controllers.ErrorBadRequestController)
	http.HandleFunc("/404", controllers.ErrorNotFoundController)
	http.HandleFunc("/500", controllers.ErrorInternalController)
}
