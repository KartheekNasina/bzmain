package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type EventController struct {
	eventService *service.EventService
}

// Constructor function for EventController
func NewEventController(s *service.EventService) *EventController {
	return &EventController{eventService: s}
}

// EventController.go
func (ec *EventController) GetEvents(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid offset parameter"})
		return
	}

	events, err := ec.eventService.GetEvents(limit, offset)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, events)
}

func (ec *EventController) GetEventAndSchedule(c *gin.Context) {
	eventID := c.Param("eventID")
	eventInfo, eventSchedule, err := ec.eventService.GetEventAndSchedule(eventID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"eventInfo": eventInfo, "eventSchedule": eventSchedule})
}

func (ec *EventController) CreateEventJoinRequest(c *gin.Context) {
	var requestDTO dto.EventJoinRequestDTO
	if err := c.BindJSON(&requestDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := ec.eventService.CreateEventJoinRequest(&requestDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Event Join Request created successfully"})
}

// ... Similarly for UpdateEventJoinRequest and DeleteEventJoinRequest
