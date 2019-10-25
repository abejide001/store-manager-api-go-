package controllers

import (
	"github.com/abejide001/Store_Manager_GO/api/middlewares"
)

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/api/v1", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/api/v1/auth/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Register route
	s.Router.HandleFunc("/api/v1/auth/signup", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")

	// Users route
	s.Router.HandleFunc("/api/v1/auth/user", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/api/v1/auth/user/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/api/v1/auth/user/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/api/v1/auth/user/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Products routes
	s.Router.HandleFunc("/api/v1/products", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.CreateProduct))).Methods("POST")
	s.Router.HandleFunc("/api/v1/products", middlewares.SetMiddlewareJSON((s.GetProducts))).Methods("GET")
	s.Router.HandleFunc("/api/v1/products/{id}", middlewares.SetMiddlewareJSON(s.GetProduct)).Methods("GET")
	s.Router.HandleFunc("/api/v1/products/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateProduct))).Methods("PUT")
	s.Router.HandleFunc("/api/v1/products/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteProduct)).Methods("DELETE")
}
