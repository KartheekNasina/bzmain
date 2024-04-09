package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	dto "github.com/vivekbnwork/bz-backend/bz-main/models/dto"

	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	Redis *redis.Client
}

func NewRedisService(redis *redis.Client) *RedisService {
	return &RedisService{
		Redis: redis,
	}
}

func (s *RedisService) GetTopTrendingBreweriesNew(N int64) ([]dto.BreweryLandingDTO, error) {
	var trendingBreweries []dto.BreweryLandingDTO
	ctx := context.Background()

	luaScript := `
	-- Get the top N trending brewery IDs
	local ids = redis.call('ZREVRANGE', 'trending_breweries', 0, ARGV[1]-1)
	local breweries = {}

	-- For each ID, fetch the corresponding hash
	for _, id in ipairs(ids) do
		local hashName = "brewery:" .. id
		local brewery = redis.call('HGETALL', hashName)
		-- Insert the ID at the beginning of the brewery details
		table.insert(brewery, 1, "ID")
		table.insert(brewery, 2, id)
		table.insert(breweries, brewery)
	end

	return breweries
	`

	script := redis.NewScript(luaScript)
	sha1, err := script.Load(ctx, s.Redis).Result()
	if err != nil {
		return nil, err
	}

	results, err := s.Redis.EvalSha(ctx, sha1, []string{}, N).Result()
	if err != nil {
		// If the error is NOSCRIPT, then reload the script and retry
		if strings.Contains(err.Error(), "NOSCRIPT") {
			sha1, err = script.Load(ctx, s.Redis).Result()
			if err != nil {
				return nil, err
			}
			results, err = s.Redis.EvalSha(ctx, sha1, []string{}, N).Result()
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// Parse the results and populate trendingBreweries slice
	resultSlices, ok := results.([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected result type from Lua script")
	}

	for _, res := range resultSlices {
		fmt.Printf("Raw Data from Redis: %+v\n", res)

		rawData, ok := res.([]interface{})
		if !ok || len(rawData)%2 != 0 {
			continue
		}

		fieldMap := make(map[string]string)
		for i := 0; i < len(rawData); i += 2 {
			key, kOk := rawData[i].(string)
			value, vOk := rawData[i+1].(string)
			if kOk && vOk {
				fieldMap[key] = value
			}
		}
		fmt.Printf("Parsed Map: %+v\n", fieldMap)

		trendingBreweries = append(trendingBreweries, dto.BreweryLandingDTO{
			ID:      fieldMap["ID"],
			Name:    fieldMap["Name"],
			LogoURL: fieldMap["Logo_url"],
		})
	}

	return trendingBreweries, nil
}

func (s *RedisService) FetchBrewClasses() ([]dto.BrewClassDTO, error) {
	ctx := context.Background()
	brewClassIDs, err := s.Redis.Keys(ctx, "brew_classes:*").Result()
	if err != nil {
		return nil, err
	}

	var brewClasses []dto.BrewClassDTO
	for _, brewClassID := range brewClassIDs {
		brewClassDetails, err := s.Redis.HGetAll(ctx, brewClassID).Result()
		if err != nil {
			return nil, err
		}

		// Create a brew class object
		brewClass := dto.BrewClassDTO{
			Id:               brewClassID,
			Title:            brewClassDetails["Title"],
			ShortDescription: brewClassDetails["ShortDescription"],
			Thumbnail:        brewClassDetails["Thumbnail"],
			// Add other brew class-specific fields here
		}

		brewClasses = append(brewClasses, brewClass)
	}

	return brewClasses, nil
}

// FetchEventsByStartDate retrieves events sorted by StartDate from Redis.
func (s *RedisService) FetchEventsByStartDate() ([]dto.EventDTO, error) {
	ctx := context.Background()

	// Fetch event IDs sorted by StartDate from the sorted set
	eventIDs, err := s.Redis.ZRange(ctx, "eventsByStartDate", 0, -1).Result()
	if err != nil {
		return nil, err
	}

	// Fetch event details for each event ID
	var events []dto.EventDTO
	for _, eventID := range eventIDs {
		// Get event details from the Redis hash
		eventDetails, err := s.Redis.HGetAll(ctx, eventID).Result()
		if err != nil {
			return nil, err
		}

		layout := "2006-01-02T15:04:05Z"
		startDate, err := time.Parse(layout, eventDetails["StartDate"])
		if err != nil {
			return nil, err
		}
		endDate, err := time.Parse(layout, eventDetails["EndDate"])
		if err != nil {
			return nil, err
		}

		// Create an event object
		event := dto.EventDTO{
			ID:        eventID,
			Title:     eventDetails["Title"],
			Thumbnail: eventDetails["Thumbnail"],
			BreweryID: eventDetails["BreweryId"],
			StartDate: startDate,
			EndDate:   endDate,
		}

		events = append(events, event)
	}

	return events, nil
}

// FetchTours retrieves tours from Redis.
func (s *RedisService) FetchTours() ([]dto.BrewTourDTO, error) {
	ctx := context.Background()
	tourIDs, err := s.Redis.Keys(ctx, "brew_tours:*").Result()
	if err != nil {
		return nil, err
	}

	var tours []dto.BrewTourDTO
	for _, tourID := range tourIDs {
		tourDetails, err := s.Redis.HGetAll(ctx, tourID).Result()
		if err != nil {
			return nil, err
		}

		// Create a tour object
		tour := dto.BrewTourDTO{
			ID:        tourID,
			Title:     tourDetails["Title"],
			Thumbnail: tourDetails["Thumbnail"],
			// Add other tour-specific fields here
		}

		tours = append(tours, tour)
	}

	return tours, nil
}
