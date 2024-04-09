package routes

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/vivekbnwork/bz-backend/bz-main/controllers"
	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	"github.com/vivekbnwork/bz-backend/bz-main/repository"
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

func InitializeRoutes(router *gin.Engine, db *driver.DB, rdb *redis.Client, s3Client *s3.S3) {

	var redisService *service.RedisService = service.NewRedisService(rdb)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	initializeUserRoutes(router, userController)

	cityRepo := repository.NewCityRepository(db)
	cityService := service.NewCityService(cityRepo)
	cityController := controllers.NewCityController(cityService)
	initializeCityRoutes(router, cityController)

	BreweryRepo := repository.NewBreweryRepository(db)
	BreweryService := service.NewBreweryService(BreweryRepo)
	BreweryController := controllers.NewBreweryController(BreweryService)

	//  Brewery routes
	initializeBreweryRoutes(router, BreweryController)

	beerRepo := repository.NewBeerRepository(db)
	beerService := service.NewBeerService(beerRepo)
	beerController := controllers.NewBeerController(beerService)
	initializeBeerRoutes(router, beerController)

	brewClassRepo := repository.NewBrewClassRepository(db)
	brewClassService := service.NewBrewClassService(brewClassRepo)
	brewClassController := controllers.NewBrewClassController(brewClassService)
	initializeBrewClassesRoutes(router, brewClassController)

	brewTourRepo := repository.NewBrewTourRepository(db)
	brewTourService := service.NewBrewTourService(brewTourRepo)
	brewTourController := controllers.NewBrewToursController(brewTourService)
	initializeBrewToursRoutes(router, brewTourController)

	breweryReviewRepo := repository.NewBreweryReviewRepository(db)
	breweryReviewService := service.NewBreweryReviewService(breweryReviewRepo)
	breweryReviewController := controllers.NewBreweryReviewController(breweryReviewService)
	initializeBreweryReviewRoutes(router, breweryReviewController)

	communityRepo := repository.NewCommunityRepository(db)
	communityService := service.NewCommunityService(communityRepo)
	communityController := controllers.NewCommunityController(communityService)
	initializeCommunityRoutes(router, communityController)

	drinkPurchaseRequestRepo := repository.NewDrinkPurchaseRequestRepository(db)
	drinkPurchaseRequestService := service.NewDrinkPurchaseRequestService(drinkPurchaseRequestRepo)
	drinkPurchaseRequestController := controllers.NewDrinkPurchaseRequestController(drinkPurchaseRequestService)
	initializeDrinkPurchaseRequestRoutes(router, drinkPurchaseRequestController)

	eventRepo := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepo)
	eventController := controllers.NewEventController(eventService)
	initializeEventRoutes(router, eventController)

	meetupRepo := repository.NewMeetupRepository(db)
	meetupService := service.NewMeetupService(meetupRepo)
	meetupController := controllers.NewMeetupController(meetupService)
	initializeMeetupRoutes(router, meetupController)

	myVibeRepo := repository.NewMyVibeRepository(db)
	myVibeService := service.NewMyVibeService(myVibeRepo)
	myVibeController := controllers.NewMyVibeController(myVibeService)
	initializeMyVibeRoutes(router, myVibeController)

	paymentRepo := repository.NewPaymentsRepository(db)
	paymentService := service.NewPaymentService(paymentRepo)
	paymentController := controllers.NewPaymentController(paymentService)
	initializePaymentRoutes(router, paymentController)

	userBlockedListRepo := repository.NewUserBlockedListRepository(db)
	userBlockedListService := service.NewUserBlockedListService(userBlockedListRepo)
	userBlockedListController := controllers.NewUserBlockedListController(userBlockedListService)
	initializeUserBlockedListRoutes(router, userBlockedListController)

	userBreweryCheckinRepo := repository.NewUserBreweryCheckinRepository(db)
	userBreweryCheckinsService := service.NewUserBreweryCheckinsService(userBreweryCheckinRepo)
	userBreweryCheckinsController := controllers.NewUserBreweryCheckinsController(userBreweryCheckinsService)
	initializeUserBreweryCheckinsRoutes(router, userBreweryCheckinsController)

	userConnectionsRepo := repository.NewUserConnectionRepository(db)
	userConnectionsService := service.NewUserConnectionsService(userConnectionsRepo)
	userConnectionsController := controllers.NewUserConnectionsController(userConnectionsService)
	initializeUserConnectionsRoutes(router, userConnectionsController)

	userFavoriteBreweryRepo := repository.NewUserFavoriteBreweryRepository(db)
	userFavoriteBreweryService := service.NewUserFavoriteBreweryService(userFavoriteBreweryRepo)
	userFavoriteBreweryController := controllers.NewUserFavoriteBreweryController(userFavoriteBreweryService)
	initializeUserFavoriteBreweryRoutes(router, userFavoriteBreweryController)

	userFeedbackRepo := repository.NewUserFeedbackRepository(db)
	userFeedbackService := service.NewUserFeedbackService(userFeedbackRepo)
	userFeedbackController := controllers.NewUserFeedbackController(userFeedbackService)
	initializeUserFeedbackRoutes(router, userFeedbackController)

	userHistoryRepo := repository.NewUserHistoryRepository(db)
	userHistoryService := service.NewUserHistoryService(userHistoryRepo)
	userHistoryController := controllers.NewUserHistoryController(userHistoryService)
	initializeUserHistoryRoutes(router, userHistoryController)

	userNotificationRepo := repository.NewUserNotificationRepository(db)
	userNotificationService := service.NewUserNotificationService(userNotificationRepo)
	userNotificationController := controllers.NewUserNotificationController(userNotificationService)
	initializeUserNotificationRoutes(router, userNotificationController)

	userReferralRepo := repository.NewUserReferralRepository(db)
	userReferralService := service.NewUserReferralService(userReferralRepo)
	userReferralController := controllers.NewUserReferralController(userReferralService)
	initializeUserReferralRoutes(router, userReferralController)

	userSubscriptionRepo := repository.NewUserSubscriptionRepository(db)
	userSubscriptionService := service.NewUserSubscriptionService(userSubscriptionRepo)
	userSubscriptionController := controllers.NewUserSubscriptionController(userSubscriptionService)
	initializeUserSubscriptionRoutes(router, userSubscriptionController)

	var awsService *service.AWSService = service.NewAWSService(s3Client)
	awsController := controllers.NewAWSController(awsService)
	initializeAwsRoutes(router, awsController)

	landingService := service.NewLandingService(redisService)
	landingController := controllers.NewLandingController(landingService)

	// User routes
	initializeLandingRoutes(router, landingController)
}
