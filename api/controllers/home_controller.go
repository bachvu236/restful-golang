package controllers

import (
	"net/http"
	responses "github.com/bachvu236/go-jwt/api/response"
)
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
