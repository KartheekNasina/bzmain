package service

import (
	"sync"

	"github.com/sirupsen/logrus"
	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
)

type LandingService struct {
	userRepo     *repository.UserRepository
	redisService *RedisService
}

// Constructor function for UserService
func NewLandingService(rs *RedisService) *LandingService {
	return &LandingService{redisService: rs}
}

func (s *LandingService) FetchLandingData() dto.LandingDTO {

	logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "ListUsers",
	}).Debug("List Users - Start")

	// Using defer to ensure the log is printed regardless of where we exit the function.
	defer logrus.WithFields(logrus.Fields{
		"service": "bz-main",
		"event":   "ListUsers",
	}).Debug("List Users - End")
	var wg sync.WaitGroup
	var err error

	var landingDTO dto.LandingDTO
	var trendingBreweries []dto.BreweryLandingDTO
	var events []dto.EventDTO
	var brewClasses []dto.BrewClassDTO
	var brewTours []dto.BrewTourDTO

	// Fetch trending breweries
	wg.Add(1)
	go func() {
		defer wg.Done()
		trendingBreweries, err = s.redisService.GetTopTrendingBreweriesNew(1000)
		if err != nil {
			// handle error
		}
		landingDTO.TrendingBreweries = trendingBreweries
	}()

	// Fetch events
	wg.Add(1)
	go func() {
		defer wg.Done()
		events, err = s.redisService.FetchEventsByStartDate()
		if err != nil {
			// handle error
		}
		landingDTO.Events = events
	}()

	// Fetch brew classes
	wg.Add(1)
	go func() {
		defer wg.Done()
		brewClasses, err = s.redisService.FetchBrewClasses()
		if err != nil {
			// handle error
		}
		landingDTO.BrewClasses = brewClasses
	}()

	// Fetch brew tours
	wg.Add(1)
	go func() {
		defer wg.Done()
		brewTours, err = s.redisService.FetchTours()
		if err != nil {
			// handle error
		}
		landingDTO.BrewTours = brewTours
	}()

	wg.Wait()

	return landingDTO

}
