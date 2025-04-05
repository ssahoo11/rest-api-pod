package route

import (
	"net/http"
	"server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerEvent(cxt *gin.Context) {
	userId := cxt.GetInt64("userId")
	event_id, err := strconv.ParseInt(cxt.Param("id"), 10, 64)

	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error", "error": err.Error()})
		return
	}

	var event *models.Event
	event, err = models.GetEventById(event_id)

	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Event not found", "error": err.Error()})
		return
	}

	err = event.Register(userId)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Register unsuccessful", "error": err.Error()})
		return
	}
	cxt.JSON(http.StatusCreated, gin.H{"message": "Register successful"})
}

func cancelRegistration(cxt *gin.Context) {
	userId := cxt.GetInt64("userId")
	event_id, err := strconv.ParseInt(cxt.Param("id"), 10, 64)

	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error", "error": err.Error()})
		return
	}

	var event models.Event
	event.ID = event_id

	err = event.UnRegister(userId)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Registration cancelation unsuccessful", "error": err.Error()})
		return
	}
	cxt.JSON(http.StatusCreated, gin.H{"message": "Registration canceled!"})
}
