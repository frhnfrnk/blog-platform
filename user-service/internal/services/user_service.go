package services

import (
	"context"
	"encoding/json"
	"github.com/frhnfrnk/blog-platform-microservices/user-service/internal/models"
	"github.com/frhnfrnk/blog-platform-microservices/user-service/internal/repositories"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

type UserService struct {
	userRepo *repositories.UserRepository
	cache    *redis.Client
}

func NewUserService(userRepo *repositories.UserRepository, cache *redis.Client) *UserService {
	return &UserService{
		userRepo: userRepo,
		cache:    cache,
	}
}

func (us *UserService) CreateUser(user *models.User) error {
	err := us.userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	// Invalidate cache after creating user
	err = us.cache.Del(context.Background(), "all_users").Err()
	if err != nil {
		return err
	}

	err = us.cache.Del(context.Background(), "user:"+strconv.Itoa(int(user.ID))).Err()
	if err != nil {
		// Handle error (optional)
		log.Printf("Failed to invalidate cache for user:%d: %v", user.ID, err)
	}

	return nil
}

func (us *UserService) UpdateUser(user *models.User) error {
	err := us.userRepo.UpdateUser(user)
	if err != nil {
		return err
	}

	// Invalidate cache after updating user
	err = us.cache.Del(context.Background(), "all_users").Err()
	if err != nil {
		return err
	}

	err = us.cache.Del(context.Background(), "user:"+strconv.Itoa(int(user.ID))).Err()
	if err != nil {
		// Handle error (optional)
		log.Printf("Failed to invalidate cache for user:%d: %v", user.ID, err)
	}

	return nil
}

func (us *UserService) DeleteUser(userID string) error {
	err := us.userRepo.DeleteUser(userID)
	if err != nil {
		return err
	}

	// Invalidate cache after deleting user
	err = us.cache.Del(context.Background(), "all_users").Err()
	if err != nil {
		return err
	}

	err = us.cache.Del(context.Background(), "user:"+userID).Err()
	if err != nil {
		// Handle error (optional)
		log.Printf("Failed to invalidate cache for user:%s: %v", userID, err)
	}

	return nil
}

func (us *UserService) GetAllUsers() ([]*models.User, error) {
	cachedUsers, err := us.cache.Get(context.Background(), "all_users").Result()
	if err == nil {
		var users []*models.User
		err := json.Unmarshal([]byte(cachedUsers), &users)
		if err != nil {
			return nil, err
		}
		return users, nil
	}

	users, err := us.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	usersJSON, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}
	err = us.cache.Set(context.Background(), "all_users", usersJSON, 0).Err()
	if err != nil {
		return nil, err
	}

	var usersPtr []*models.User
	for _, user := range users {
		usersPtr = append(usersPtr, &user)
	}

	return usersPtr, nil
}

func (us *UserService) GetUserByID(userID string) (*models.User, error) {
	cachedUser, err := us.cache.Get(context.Background(), "user:"+userID).Result()
	if err == nil {
		var user models.User
		err := json.Unmarshal([]byte(cachedUser), &user)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}

	user, err := us.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	err = us.cache.Set(context.Background(), "user:"+userID, userJSON, 0).Err()
	if err != nil {
		return nil, err
	}

	return user, nil
}
