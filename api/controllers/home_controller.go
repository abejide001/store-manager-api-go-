package controllers

import (
	"net/http"

	"github.com/abejide001/Store_Manager_GO/api/responses"
)

// Home Method
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
