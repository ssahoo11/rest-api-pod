package route

import (
	"database/sql"
	"fmt"
	"net/http"
	"server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(cxt *gin.Context) {

	events, err := models.GetAllEvents()
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		return
	}
	cxt.JSON(http.StatusOK, events)
}

func getEvent(cxt *gin.Context) {
	eventId, err := strconv.ParseInt(cxt.Param("id"), 10, 64)
	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event Id"})
		return
	}

	fmt.Printf("Received request for event ID: %d\n", eventId) // Log the eventId

	event, err := models.GetEventById(eventId)
	if err != nil {
		// Check if it's a no rows error (this means the event doesn't exist in DB)
		if err == sql.ErrNoRows {
			cxt.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		} else {
			cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error", "error": err.Error()})
		}
		return
	}

	cxt.JSON(http.StatusOK, event)
}

func createEvent(cxt *gin.Context) {

	userId := cxt.GetInt64("userId")
	var event models.Event
	err := cxt.ShouldBindJSON(&event)
	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse input data"})
		return
	}

	event.UserID = userId
	err = event.Save()
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save events"})
		return
	}
	cxt.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(cxt *gin.Context) {
	eventId, err := strconv.ParseInt(cxt.Param("id"), 10, 64)
	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event Id"})
		return
	}
	userId := cxt.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if event.UserID != userId {
		cxt.JSON(http.StatusUnauthorized, gin.H{"message": "User is not authorized to update the event"})
	}

	if err != nil {
		// Check if it's a no rows error (this means the event doesn't exist in DB)
		if err == sql.ErrNoRows {
			cxt.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		} else {
			cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error", "error": err.Error()})
		}
		return
	}

	fmt.Printf("Received update request for event ID: %d\n", eventId) // Log the eventId

	var updatedEvent models.Event
	err = cxt.ShouldBindJSON(&updatedEvent)
	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse input data"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.UpdateEventById()
	if err != nil {
		// Check if it's a no rows error (this means the event doesn't exist in DB)
		if err == sql.ErrNoRows {
			cxt.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		} else {
			cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't update the event", "error": err.Error()})
		}
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"message": "Event updated!", "event": updatedEvent})
}

func deleteEvent(cxt *gin.Context) {
	eventId, err := strconv.ParseInt(cxt.Param("id"), 10, 64)
	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event Id"})
		return
	}
	userId := cxt.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if event.UserID != userId {
		cxt.JSON(http.StatusUnauthorized, gin.H{"message": "User is not authorized to delete the event"})
	}

	if err != nil {
		// Check if it's a no rows error (this means the event doesn't exist in DB)
		if err == sql.ErrNoRows {
			cxt.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		} else {
			cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error", "error": err.Error()})
		}
		return
	}

	fmt.Printf("Received delete request for event ID: %d\n", eventId) // Log the eventId

	err = models.DeleteEventById(eventId)
	if err != nil {
		// Check if it's a no rows error (this means the event doesn't exist in DB)
		if err == sql.ErrNoRows {
			cxt.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		} else {
			cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't delete the event", "error": err.Error()})
		}
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"message": "Event deleted!"})
}
