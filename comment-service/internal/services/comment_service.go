package services

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/frhnfrnk/blog-platform-microservices/comment-service/internal/models"
	"github.com/frhnfrnk/blog-platform-microservices/comment-service/internal/repositories"
	postPb "github.com/frhnfrnk/blog-platform-microservices/post-service/pb"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

type CommentService struct {
	commentRepo *repositories.CommentRepository
	cache       *redis.Client
	postClient  postPb.PostServiceClient
}

func NewCommentService(commentRepo *repositories.CommentRepository, cache *redis.Client, postClient postPb.PostServiceClient) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
		cache:       cache,
		postClient:  postClient,
	}
}

func (s *CommentService) CreateComment(comment *models.Comment) error {
	_, err := s.postClient.GetPost(context.Background(), &postPb.GetPostRequest{Id: comment.PostID})
	if err != nil {
		return errors.New("post not found")
	}

	err = s.commentRepo.CreateComment(comment)
	if err == nil {
		s.cache.Del(context.Background(), "all_comments")
	}
	return err
}

func (s *CommentService) GetCommentByID(id uint) (*models.Comment, error) {
	ctx := context.Background()
	cacheKey := "comment:" + strconv.Itoa(int(id))
	cachedComment, err := s.cache.Get(ctx, cacheKey).Result()
	if err == nil {
		var comment models.Comment
		err := json.Unmarshal([]byte(cachedComment), &comment)
		if err == nil {
			return &comment, nil
		}
	}
	comment, err := s.commentRepo.GetCommentByID(strconv.Itoa(int(id)))
	if err != nil {
		return nil, err
	}

	commentJSON, err := json.Marshal(comment)
	if err == nil {
		s.cache.Set(ctx, cacheKey, commentJSON, time.Hour)
	}

	return comment, nil
}

func (s *CommentService) UpdateComment(comment *models.Comment) error {
	_, err := s.postClient.GetPost(context.Background(), &postPb.GetPostRequest{Id: comment.PostID})
	if err != nil {
		return errors.New("post not found")
	}

	err = s.commentRepo.UpdateComment(comment)
	if err == nil {
		s.cache.Del(context.Background(), "comment:"+strconv.Itoa(int(comment.ID)))
		s.cache.Del(context.Background(), "all_comments")
	}
	return err
}

func (s *CommentService) DeleteComment(id string) error {
	err := s.commentRepo.DeleteComment(id)
	if err == nil {
		s.cache.Del(context.Background(), "comment:"+id)
		s.cache.Del(context.Background(), "all_comments")
	}
	return err
}

func (s *CommentService) GetAllComments() ([]models.Comment, error) {
	ctx := context.Background()
	cacheKey := "all_comments"
	cachedComment, err := s.cache.Get(ctx, cacheKey).Result()
	if err == nil {
		var comments []models.Comment
		err := json.Unmarshal([]byte(cachedComment), &comments)
		if err == nil {
			return comments, nil
		}
	}

	comments, err := s.commentRepo.GetAllComments()
	if err != nil {
		return nil, err
	}

	commentsJSON, err := json.Marshal(comments)
	if err == nil {
		s.cache.Set(ctx, cacheKey, commentsJSON, time.Hour)
	}

	return comments, nil
}
