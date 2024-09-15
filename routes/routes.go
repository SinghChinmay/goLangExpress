package routes

import (
	"goLangExpress/handlers"
	"goLangExpress/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	// Public routes
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/logout", handlers.Logout).Methods("POST")

	// Protected routes
	r.HandleFunc("/me", middleware.JWTMiddleware(handlers.Me)).Methods("GET")
	r.HandleFunc("/protected1", middleware.JWTMiddleware(handlers.Protected1)).Methods("GET")
	r.HandleFunc("/protected2", middleware.JWTMiddleware(handlers.Protected2)).Methods("GET")
	r.HandleFunc("/protected3", middleware.JWTMiddleware(handlers.Protected3)).Methods("GET")
	r.HandleFunc("/protected4", middleware.JWTMiddleware(handlers.Protected4)).Methods("GET")
	r.HandleFunc("/protected5", middleware.JWTMiddleware(handlers.Protected5)).Methods("GET")
}