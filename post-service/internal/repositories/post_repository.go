package repositories

import (
	"github.com/frhnfrnk/blog-platform-microservices/post-service/internal/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) CreatePost(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *PostRepository) GetPostByID(id string) (*models.Post, error) {
	var post models.Post
	if err := r.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) UpdatePost(post *models.Post) error {
	return r.db.Save(post).Error
}

func (r *PostRepository) DeletePost(id string) error {
	return r.db.Delete(&models.Post{}, id).Error
}
