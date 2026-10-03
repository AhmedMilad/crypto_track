package controllers

import (
	"crypto_track/database"
	"crypto_track/models"
	"crypto_track/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignUp(c *gin.Context) {

	var req utils.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Error: "Body binding failed " + err.Error(),
		})

		return
	}

	var user models.User

	result := database.Conn.Where("email = ?", req.Email).First(&user)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {

		c.JSON(http.StatusBadRequest, utils.Response{
			Error: "Could not complete the process successfully, please try again later.",
		})

		return
	}

	if user != (models.User{}) {

		c.JSON(http.StatusBadRequest, utils.Response{
			Error: "The email is already taken.",
		})

		return
	}

	hash, err := utils.HashPasswordSecure(req.Password)

	if err != nil {

		c.JSON(http.StatusBadRequest, utils.Response{
			Error: "Could not complete the process successfully, please try again later.",
		})

		return
	}

	token, err := utils.CreateToken(req.Email)

	if err != nil {

		c.JSON(http.StatusBadRequest, utils.Response{
			Error: "Could not complete the process successfully, please try again later.",
		})

		return

	}

	user.Email = req.Email
	user.Password = hash
	user.FullName = req.FullName

	if err := database.Conn.Create(&user).Error; err != nil {

		c.JSON(http.StatusBadRequest, utils.Response{
			Error: "Could not complete the process successfully, please try again later.",
		})

		return
	}

	var resp struct {
		Token string
	}

	resp.Token = token

	c.JSON(http.StatusOK, utils.Response{
		Message: "You have successfully signed up.",
		Data:    resp,
	})

}

func Login(c *gin.Context) {

	var req utils.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Error: "Body binding failed " + err.Error(),
		})

		return
	}

	var user models.User

	result := database.Conn.Where("email = ?", req.Email).First(&user)

	if result.Error != nil {

		message := "Something went wrong, please try again later."

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {

			message = "Invalid credentials."
		}

		c.JSON(http.StatusBadRequest, utils.Response{
			Error: message,
		})

		return
	}

	isVerified, err := utils.VerifyPasswordSecure(user.Password, req.Password)

	if err != nil {

		c.JSON(http.StatusBadRequest, utils.Response{
			Error: "Could not complete the process successfully, please try again later.",
		})

		return

	}

	if !isVerified {

		c.JSON(http.StatusBadRequest, utils.Response{
			Error: "Invalid credentials",
		})

		return
	}

	token, err := utils.CreateToken(req.Email)

	if err != nil {

		c.JSON(http.StatusBadRequest, utils.Response{
			Error: "Could not complete the process successfully, please try again later.",
		})

		return

	}

	var resp struct {
		Token string
	}

	resp.Token = token

	c.JSON(http.StatusOK, utils.Response{
		Message: "You have successfully logged in.",
		Data:    resp,
	})
}
