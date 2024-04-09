package service

import (
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type EventService struct {
	eventRepo *repository.EventRepository
}

func NewEventService(r *repository.EventRepository) *EventService {
	return &EventService{eventRepo: r}
}

// EventService
func (es *EventService) GetEvents(limit, offset int) ([]models.EventBasicInfo, error) {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetEvents",
	}).Debug("Get Events - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetEvents",
	}).Debug("Get Events - End")

	return es.eventRepo.GetEvents(limit, offset)
}

func (es *EventService) GetEventAndSchedule(eventID string) (*models.EventBasicInfo, *models.EventSchedule, error) {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetEventAndSchedule",
	}).Debug("Get Event and Schedule - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetEventAndSchedule",
	}).Debug("Get Event and Schedule - End")

	return es.eventRepo.GetEventAndSchedule(eventID)
}

func (es *EventService) CreateEventJoinRequest(requestDTO *dto.EventJoinRequestDTO) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateEventJoinRequest",
	}).Debug("Create Event Join Request - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateEventJoinRequest",
	}).Debug("Create Event Join Request - End")

	var request models.EventJoinRequest
	copier.Copy(&request, requestDTO)

	return es.eventRepo.CreateEventJoinRequest(request)
}

func (es *EventService) UpdateEventJoinRequest(id string, requestDTO *dto.EventJoinRequestDTO) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "UpdateEventJoinRequest",
	}).Debug("Update Event Join Request - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "UpdateEventJoinRequest",
	}).Debug("Update Event Join Request - End")

	var request models.EventJoinRequest
	copier.Copy(&request, requestDTO)

	return es.eventRepo.UpdateEventJoinRequest(id, request)
}

func (es *EventService) DeleteEventJoinRequest(requestID string) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "DeleteEventJoinRequest",
	}).Debug("Delete Event Join Request - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "DeleteEventJoinRequest",
	}).Debug("Delete Event Join Request - End")

	return es.eventRepo.DeleteEventJoinRequest(requestID)
}
