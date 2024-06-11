package repositories

import (
	"github.com/frhnfrnk/blog-platform-microservices/comment-service/internal/models"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) CreateComment(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *CommentRepository) GetCommentByID(id string) (*models.Comment, error) {
	var comment models.Comment
	result := r.db.First(&comment, id)
	return &comment, result.Error
}

func (r *CommentRepository) UpdateComment(comment *models.Comment) error {
	return r.db.Save(comment).Error
}

func (r *CommentRepository) DeleteComment(id string) error {
	return r.db.Delete(&models.Comment{}, id).Error
}

func (r *CommentRepository) GetAllComments() ([]models.Comment, error) {
	var comments []models.Comment
	result := r.db.Find(&comments)
	return comments, result.Error
}
