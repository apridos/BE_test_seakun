package controller

import (
	"crypto/md5"
	"net/http"

	"encoding/hex"

	"github.com/gin-gonic/gin"

	Model "seakun/Model"
	"seakun/Utils"
)

func AdminLogin(c *gin.Context) {
	var adminRequest Model.Admin
	c.BindJSON(&adminRequest)

	hash := md5.Sum([]byte(adminRequest.Password))

	encodedHash := hex.EncodeToString(hash[:])

	var admin Model.Admin
	err := Model.FindAdminByUsername(&admin, adminRequest.Username)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if encodedHash != admin.Password {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := Utils.GenerateJWTToken(admin)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.SetCookie("token", token, 60*15, "/", "seakun.aprido.my.id", true, true)

	refreshToken, err := Utils.GenerateJWTRefreshToken(admin)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.SetCookie("refresh_token", refreshToken, 60*15*24, "/", "seakun.aprido.my.id", true, true)

	c.JSON(http.StatusOK, admin)
}

func Authorize(c *gin.Context) bool {
	// Validating the token

	// retrieve token from cookie and abort the request if token not present
	token, err := c.Cookie("token")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return false
	}

	// validate and extract the username from token
	var username string

	code := Utils.ValidateJWTToken(token, &username)

	// if the validation failed, utilize the refresh token instead
	if code != 200 {
		refreshToken, err := c.Cookie("refreshToken")

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return false
		}

		code := Utils.ValidateJWTToken(refreshToken, &username)

		// if the refresh token doesn't work, abort request
		if code != 200 {
			c.AbortWithStatus(code)
			return false
		}

		// create a new token and refresh token
		var admin Model.Admin

		err = Model.FindAdminByUsername(&admin, username)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return false
		}

		token, err := Utils.GenerateJWTToken(admin)

		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return false
		}

		c.SetCookie("token", token, 60*15, "/", "seakun.aprido.my.id", true, true)

		refreshToken, err = Utils.GenerateJWTRefreshToken(admin)

		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return false
		}

		c.SetCookie("refresh_token", refreshToken, 60*15*24, "/", "seakun.aprido.my.id", true, true)
	}

	return true
}
