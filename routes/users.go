package routes

import (
	"fmt"
	"net/http"

	"db-go.com/api/models"
	"db-go.com/api/tokens"
	"db-go.com/api/utils"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User Created",
	})

}
func logout(context *gin.Context) {
	refreshToken, err := context.Cookie("refresh_token")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = models.Delete(refreshToken)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	context.SetCookie("refresh_token", "", -1, "/", "", true, true)
	context.JSON(http.StatusAccepted, gin.H{
		"message": "LOGOUT",
	})
}
func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	fmt.Println(err)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}
	incomingPasssword := user.Password

	err = user.GetUser(user.Email)
	fmt.Println(incomingPasssword, user.Password)
	isPasswordValid := utils.VerifyPassword(incomingPasssword, user.Password)
	fmt.Println(isPasswordValid)
	if isPasswordValid == true {
		ss, err := tokens.SaveRefreshTokenDB(&user)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})

			return
		}
		token, err := tokens.GenerateAccesToken(&user)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})

			return
		}
		fmt.Println(err)
		context.SetCookie("refresh_token", ss, 24*60*60, "/", "", true, true)
		context.JSON(http.StatusAccepted, gin.H{
			"token": token,
		})
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Bas password",
		})

		return
	}
}
