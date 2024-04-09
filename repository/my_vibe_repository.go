package repository

import (
	"context"

	"github.com/vivekbnwork/bz-backend/bz-main/driver"
	models "github.com/vivekbnwork/bz-backend/bz-main/models/db"
)

type MyVibeRepository struct {
	db *driver.DB
}

func NewMyVibeRepository(database *driver.DB) *MyVibeRepository {
	return &MyVibeRepository{
		db: database,
	}
}

// CreateFoodDrinkRating creates a new food or drink rating record.
func (repo *MyVibeRepository) CreateFoodDrinkRating(foodDrinkRating models.FoodDrinkRating) error {
	query := `
		INSERT INTO food_drink_rating (user_id, item_id, rating, comment_id, rated_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, foodDrinkRating.UserID, foodDrinkRating.ItemID, foodDrinkRating.Rating, foodDrinkRating.CommentID, foodDrinkRating.RatedAt)
	if err != nil {
		return err
	}
	return nil
}

// GetFoodDrinkRatingByID fetches a food or drink rating record by ID.
func (repo *MyVibeRepository) GetFoodDrinkRatingByID(id string) (models.FoodDrinkRating, error) {
	var foodDrinkRating models.FoodDrinkRating
	query := `
		SELECT user_id, item_id, rating, comment_id, rated_at
		FROM food_drink_rating
		WHERE id = $1
	`
	row := repo.db.Pool.QueryRow(context.Background(), query, id)
	err := row.Scan(&foodDrinkRating.UserID, &foodDrinkRating.ItemID, &foodDrinkRating.Rating, &foodDrinkRating.CommentID, &foodDrinkRating.RatedAt)
	if err != nil {
		return models.FoodDrinkRating{}, err
	}
	return foodDrinkRating, nil
}

// UpdateFoodDrinkRating updates an existing food or drink rating record.
func (repo *MyVibeRepository) UpdateFoodDrinkRating(id string, foodDrinkRating models.FoodDrinkRating) error {
	query := `
		UPDATE food_drink_rating
		SET user_id = $1, item_id = $2, rating = $3, comment_id = $4, rated_at = $5
		WHERE id = $6
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, foodDrinkRating.UserID, foodDrinkRating.ItemID, foodDrinkRating.Rating, foodDrinkRating.CommentID, foodDrinkRating.RatedAt, id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteFoodDrinkRating deletes a food or drink rating record by ID.
func (repo *MyVibeRepository) DeleteFoodDrinkRating(id string) error {
	query := `
		DELETE FROM food_drink_rating
		WHERE id = $1
	`
	_, err := repo.db.Pool.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}

// GetFoodDrinkItems fetches food or drink items by BreweryID.
func (repo *MyVibeRepository) GetFoodDrinkItems(breweryID string) ([]models.FoodDrinkItem, error) {
	var foodDrinkItems []models.FoodDrinkItem
	query := `
		SELECT id, brewery_id, name, type, description, image_url
		FROM food_drink_item
		WHERE brewery_id = $1
	`
	rows, err := repo.db.Pool.Query(context.Background(), query, breweryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var foodDrinkItem models.FoodDrinkItem
		err := rows.Scan(&foodDrinkItem.ID, &foodDrinkItem.BreweryID, &foodDrinkItem.Name, &foodDrinkItem.Type, &foodDrinkItem.Description, &foodDrinkItem.ImageURL)
		if err != nil {
			return nil, err
		}
		foodDrinkItems = append(foodDrinkItems, foodDrinkItem)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return foodDrinkItems, nil
}
