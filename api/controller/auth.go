package controller

import (
	"errors"
	"net/http"

	"github.com/Samir-Minddeft/go-backend-boilerplate/api/models"
	"github.com/Samir-Minddeft/go-backend-boilerplate/config"
	"github.com/Samir-Minddeft/go-backend-boilerplate/utils/helper"
	"github.com/Samir-Minddeft/go-backend-boilerplate/utils/response"
	"github.com/Samir-Minddeft/go-backend-boilerplate/utils/types"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	login := types.Login{}
	if err := c.ShouldBindJSON(&login); err != nil {
		response.WriteJson(c.Writer, http.StatusBadRequest, response.GeneralError(err))
		return
	}

	if login.Email == "" || login.Password == "" {
		response.WriteJson(c.Writer, http.StatusBadRequest, response.GeneralError(errors.New("Email and Password are required")))
		return
	}

	user := models.User{}
	config.DB.Select("id, email, role, password, salt").First(&user, "email = ?", login.Email)

	if user.Id == 0 {
		response.WriteJson(c.Writer, http.StatusNotFound, response.GeneralError(errors.New("user not found")))
		return
	}

	verified, err := helper.VerifyPassword(login.Password, user.Salt, user.Password)
	if err != nil {
		response.WriteJson(c.Writer, http.StatusInternalServerError, response.GeneralError(err))
		return
	}

	if !verified {
		response.WriteJson(c.Writer, http.StatusUnauthorized, response.GeneralError(errors.New("invalid email or password")))
		return
	}

	token, err := helper.CreateJwtToken(uint(user.Id), user.Email, user.Role)
	if err != nil {
		response.WriteJson(c.Writer, http.StatusInternalServerError, response.GeneralError(err))
		return
	}

	response.WriteJson(c.Writer, http.StatusOK, gin.H{
		"message": "Logged in successfully",
		"token":   token,
		"user": gin.H{
			"id":    user.Id,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}
