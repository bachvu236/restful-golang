package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/bachvu236/go-jwt/api/exception"
	"github.com/bachvu236/go-jwt/api/models"
	"github.com/bachvu236/go-jwt/api/response"
	"github.com/bachvu236/go-jwt/api/security"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {

	user := models.User{}
	fmt.Print(email)
	result,err := server.DB.Query("SELECT * FROM employee WHERE email = ?", email )
	if err != nil {
		return "", err
	}
	if !result.Next() {
		return "No data found",err
	}
	result.Scan(&user.ID, &user.Nickname, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
	
}
