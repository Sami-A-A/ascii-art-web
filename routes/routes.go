package routes

import (
	"net/http"
	"asciiartweb/controllers"
)

func UseRoutes(router *http.ServeMux) {

	// [GET] Regular Page Routes
	router.HandleFunc("/", controllers.HomeController)

	// [GET] Error Page Routes
	router.HandleFunc("/400", controllers.ErrorBadRequestController)
	router.HandleFunc("/404", controllers.ErrorNotFoundController)
	router.HandleFunc("/500", controllers.ErrorInternalController)

	// [POST] Print ascii on home page
	router.HandleFunc("/ascii-art", controllers.PrintAsciiArtController)

}


/*

// ****
// Commented out for study purposes.
// There are multiple ways to achieve the same goal. This was a more straightforward way with less complexity

func UseRoutes() {    // No parameters

	// Regular Page Routes
	http.HandleFunc("/", controllers.HomeController)

	// Error Page Routes
	http.HandleFunc("/400", controllers.ErrorBadRequestController)
	http.HandleFunc("/404", controllers.ErrorNotFoundController)
	http.HandleFunc("/500", controllers.ErrorInternalController)

}

*/