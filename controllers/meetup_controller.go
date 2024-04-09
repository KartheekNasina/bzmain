package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type MeetupController struct {
	meetupService *service.MeetupService
}

// Constructor function for MeetupController
func NewMeetupController(s *service.MeetupService) *MeetupController {
	return &MeetupController{meetupService: s}
}

func (mc *MeetupController) GetMeetups(c *gin.Context) {
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
	meetups, err := mc.meetupService.GetMeetups(limit, offset)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, meetups)
}

func (mc *MeetupController) GetMeetup(c *gin.Context) {
	id := c.Param("id")
	meetup, err := mc.meetupService.GetMeetup(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, meetup)
}

func (mc *MeetupController) CreateMeetupAttendee(c *gin.Context) {
	var attendeeDTO dto.MeetupAttendeeDTO
	if err := c.BindJSON(&attendeeDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := mc.meetupService.CreateMeetupAttendee(&attendeeDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "Attendee added successfully"})
}

func (mc *MeetupController) UpdateMeetupAttendee(c *gin.Context) {
	id := c.Param("id")
	var attendeeDTO dto.MeetupAttendeeDTO
	if err := c.BindJSON(&attendeeDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := mc.meetupService.UpdateMeetupAttendee(id, &attendeeDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Attendee updated successfully"})
}

func (mc *MeetupController) DeleteMeetupAttendee(c *gin.Context) {
	meetupID := c.Param("meetupID")
	userID := c.Param("userID")
	err := mc.meetupService.DeleteMeetupAttendee(meetupID, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Attendee removed successfully"})
}
