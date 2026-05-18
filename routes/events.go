package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"db-go.com/api/models"
	"db-go.com/api/tokens"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	claims, ok := context.Get("claims")
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong collecting the data.",
		})
		return
	}
	tokenClaim, ok := claims.(*tokens.TokenClaims)
	userId := tokenClaim.User_ID
	fmt.Println("=================", userId)
	events, err := models.GetAllEvents(userId)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong collecting the data.",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello!",
		"events":  events,
	})
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	events, err := models.GetEvent(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"events": *events,
	})
}
func postEvent(context *gin.Context) {
	var event models.Event
	claims, ok := context.Get("claims")
	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "badrequest",
		})
		return
	}
	tokenClaims, ok := claims.(*tokens.TokenClaims)
	event.UserID = tokenClaims.User_ID
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Binding bad",
		})
		return
	}
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Eevent created",
		"event":   event,
	})

}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	event, err := models.GetEvent(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = context.ShouldBindJSON(event) // from the event we bind or replace the actual event to the new incoming data from the client
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	err = event.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	event, err := models.GetEvent(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Deleted",
	})
}
