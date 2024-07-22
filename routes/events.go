package routes

import (
	"net/http"
	"strconv"

	"example.com/example/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"There was an error getting the events.": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"There was an error getting the event.": err.Error()})
		return
	}
	event, err := models.GetEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"There was an error getting the event.": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"There was an error saving the event.": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, strConvErr := strconv.ParseInt(context.Param("id"), 10, 64)
	if strConvErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"There was an error updating the event.": strConvErr.Error()})
		return
	}

	_, getEventErr := models.GetEvent(eventId)
	if getEventErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Could not fetch the event": getEventErr.Error()})
		return
	}

	var updatedEvent models.Event
	bindJSONErr := context.BindJSON(&updatedEvent)

	if bindJSONErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Could not parse the event data": bindJSONErr.Error()})
		return

	}

	updatedEvent.ID = eventId
	updateEventErr := updatedEvent.Update()

	if updateEventErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"There was an error updating the event.": updateEventErr.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	eventId, strConvErr := strconv.ParseInt(context.Param("id"), 10, 64)
	if strConvErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"There was an error updating the event.": strConvErr.Error()})
		return
	}

	event, getEventErr := models.GetEvent(eventId)
	if getEventErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Could not fetch the event": getEventErr.Error()})
		return
	}

	deleteEventError := event.Delete()
	if deleteEventError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"There was an error deleting the event.": deleteEventError.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
