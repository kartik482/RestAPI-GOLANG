package routes

import (
	"net/http"
	"strconv"

	"example.com/api/models"
	"github.com/gin-gonic/gin"
)

func RegisterForEvent(context *gin.Context) {
	userid := context.GetInt64("userid")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	event, err := models.GetEventById(eventId)
	err = event.Register(userid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Registered successfully"})

}

func CancelRegisteration(context *gin.Context) {
	userid := context.GetInt64("userid")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	var event models.Event
	event.ID = int(eventId)

	err = event.CancelRegistration(userid)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Cancelled successfully"})
}
