package repositories

import (
	"github.com/frhnfrnk/blog-platform-microservices/user-service/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *UserRepository) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, "id = ?", userID)
	return &user, result.Error
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) DeleteUser(userID string) error {
	return r.db.Delete(&models.User{}, userID).Error
}
