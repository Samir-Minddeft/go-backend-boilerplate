package controller

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/Samir-Minddeft/go-backend-boilerplate/api/models"
	"github.com/Samir-Minddeft/go-backend-boilerplate/config"
	"github.com/Samir-Minddeft/go-backend-boilerplate/utils/helper"
	"github.com/Samir-Minddeft/go-backend-boilerplate/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgconn"
)

func GetUser(c *gin.Context) {
	user := models.User{}
	config.DB.Select("id, name, email, phone, role, is_active").First(&user, c.Param("id"))

	if user.Id == 0 {
		response.WriteJson(c.Writer, http.StatusNotFound, response.GeneralError(errors.New("user not found")))
		return
	}

	response.WriteJson(c.Writer, http.StatusOK, gin.H{
		"message": "User fetched successfully",
		"user":    user,
	})
}
func GetAllUsers(c *gin.Context) {

	users := []models.User{}
	res := config.DB.Select("id, name, email, phone, role, is_active").Find(&users)

	if res.Error != nil {
		response.WriteJson(c.Writer, http.StatusInternalServerError, response.GeneralError(res.Error))
		return
	}

	response.WriteJson(c.Writer, http.StatusOK, gin.H{
		"message": "Users fetched successfully",
		"users":   users,
	})
}

func CreateUser(c *gin.Context) {

	user := models.User{}
	// Check for EOF first or handle it within ShouldBindJSON error check
	if err := c.ShouldBindJSON(&user); err != nil {
		if errors.Is(err, io.EOF) {
			// Handle EOF as empty body -> empty user -> validation error
			user = models.User{} // Reset to ensure empty
			validate := validator.New()
			err = validate.Struct(user)
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				response.WriteJson(c.Writer, http.StatusBadRequest, response.ValidationError(ve))
				return
			}
		}
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			response.WriteJson(c.Writer, http.StatusBadRequest, response.ValidationError(ve))
			return
		}
		response.WriteJson(c.Writer, http.StatusBadRequest, response.GeneralError(err))
		return
	}

	// Manual validation for non-empty body to respect `validate` tags
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			response.WriteJson(c.Writer, http.StatusBadRequest, response.ValidationError(ve))
			return
		}
		response.WriteJson(c.Writer, http.StatusBadRequest, response.GeneralError(err))
		return
	}

	hashedPassword, salt, err := helper.HashPassword(user.Password, "")
	if err != nil {
		response.WriteJson(c.Writer, http.StatusInternalServerError, response.GeneralError(err))
		return
	}

	user.Password = hashedPassword
	user.Salt = salt

	res := config.DB.Create(&user)
	if res.Error != nil {
		var pgErr *pgconn.PgError
		if errors.As(res.Error, &pgErr) && pgErr.Code == "23505" {
			if strings.Contains(pgErr.ConstraintName, "email") || strings.Contains(pgErr.Message, "email") {
				response.WriteJson(c.Writer, http.StatusConflict, response.GeneralError(errors.New("email already exists")))
				return
			}
			if strings.Contains(pgErr.ConstraintName, "phone") || strings.Contains(pgErr.Message, "phone") {
				response.WriteJson(c.Writer, http.StatusConflict, response.GeneralError(errors.New("phone number already exists")))
				return
			}
		}
		response.WriteJson(c.Writer, http.StatusInternalServerError, response.GeneralError(res.Error))
		return
	}

	response.WriteJson(c.Writer, http.StatusOK, gin.H{
		"message": "User created successfully",
		// "user": gin.H{
		// 	"id":        user.Id,
		// 	"name":      user.Name,
		// 	"email":     user.Email,
		// 	"phone":     user.Phone,
		// 	"role":      user.Role,
		// 	"is_active": user.IsActive,
		// },
	})
}

func UpdateUser(c *gin.Context) {
	user := models.User{}
	config.DB.Select("id, name, email, phone, role, is_active").First(&user, c.Param("id"))

	if user.Id == 0 {
		response.WriteJson(c.Writer, http.StatusNotFound, response.GeneralError(errors.New("user not found")))
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			response.WriteJson(c.Writer, http.StatusBadRequest, response.ValidationError(ve))
			return
		}
		response.WriteJson(c.Writer, http.StatusBadRequest, response.GeneralError(err))
		return
	}

	config.DB.Select("id, name, email, phone, role, is_active").Save(&user)

	response.WriteJson(c.Writer, http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    user,
	})
}

func DeleteUser(c *gin.Context) {
	user := models.User{}
	config.DB.Select("id, name, email, phone, role, is_active").First(&user, c.Param("id"))

	if user.Id == 0 {
		response.WriteJson(c.Writer, http.StatusNotFound, response.GeneralError(errors.New("user not found")))
		return
	}

	config.DB.Delete(&user)

	response.WriteJson(c.Writer, http.StatusOK, gin.H{
		"message": "User deleted successfully",
		"user":    user,
	})
}
