package service

import (
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type MeetupService struct {
	meetupRepo *repository.MeetupRepository
}

func NewMeetupService(r *repository.MeetupRepository) *MeetupService {
	return &MeetupService{meetupRepo: r}
}

func (ms *MeetupService) GetMeetups(limit, offset int) ([]dto.MeetupDTO, error) {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetMeetups",
	}).Debug("Get Meetups - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetMeetups",
	}).Debug("Get Meetups - End")

	meetups, err := ms.meetupRepo.GetMeetups(limit, offset)
	if err != nil {
		return nil, err
	}

	var meetupDTOs []dto.MeetupDTO
	copier.Copy(&meetupDTOs, &meetups)

	return meetupDTOs, nil
}

func (ms *MeetupService) GetMeetup(id string) (*dto.MeetupDTO, error) {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetMeetup",
	}).Debug("Get Meetup - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "GetMeetup",
	}).Debug("Get Meetup - End")

	meetup, err := ms.meetupRepo.GetMeetup(id)
	if err != nil {
		return nil, err
	}

	var meetupDTO dto.MeetupDTO
	copier.Copy(&meetupDTO, meetup)

	return &meetupDTO, nil
}

func (ms *MeetupService) CreateMeetupAttendee(attendeeDTO *dto.MeetupAttendeeDTO) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateMeetupAttendee",
	}).Debug("Create Meetup Attendee - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "CreateMeetupAttendee",
	}).Debug("Create Meetup Attendee - End")

	var attendee models.MeetupAttendee
	if err := copier.Copy(&attendee, attendeeDTO); err != nil {
		return err
	}

	err := ms.meetupRepo.CreateMeetupAttendee(attendee)
	if err != nil {
		return err
	}

	return nil
}

func (ms *MeetupService) UpdateMeetupAttendee(id string, attendeeDTO *dto.MeetupAttendeeDTO) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "UpdateMeetupAttendee",
	}).Debug("Update Meetup Attendee - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "UpdateMeetupAttendee",
	}).Debug("Update Meetup Attendee - End")

	var attendee models.MeetupAttendee
	if err := copier.Copy(&attendee, attendeeDTO); err != nil {
		return err
	}

	err := ms.meetupRepo.UpdateMeetupAttendee(id, attendee)
	if err != nil {
		return err
	}

	return nil
}

func (ms *MeetupService) DeleteMeetupAttendee(meetupID, userID string) error {
	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "DeleteMeetupAttendee",
	}).Debug("Delete Meetup Attendee - Start")

	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "DeleteMeetupAttendee",
	}).Debug("Delete Meetup Attendee - End")

	err := ms.meetupRepo.DeleteMeetupAttendee(meetupID, userID)
	if err != nil {
		return err
	}

	return nil
}
