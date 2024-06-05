package services

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/frhnfrnk/blog-platform-microservices/post-service/internal/models"
	"github.com/frhnfrnk/blog-platform-microservices/post-service/internal/repositories"
	"github.com/go-redis/redis/v8"
)

type PostService struct {
	postRepo *repositories.PostRepository
	cache    *redis.Client
}

func NewPostService(postRepo *repositories.PostRepository, cache *redis.Client) *PostService {
	return &PostService{postRepo: postRepo, cache: cache}
}

func (s *PostService) CreatePost(post *models.Post) error {
	err := s.postRepo.CreatePost(post)
	if err == nil {
		s.cache.Del(context.Background(), "all_posts")
	}
	return err
}

func (s *PostService) GetPostByID(id uint) (*models.Post, error) {
	ctx := context.Background()
	cacheKey := "post:" + strconv.Itoa(int(id))
	cachedPost, err := s.cache.Get(ctx, cacheKey).Result()
	if err == nil {
		var post models.Post
		err := json.Unmarshal([]byte(cachedPost), &post)
		if err == nil {
			return &post, nil
		}
	}

	post, err := s.postRepo.GetPostByID(id)
	if err != nil {
		return nil, err
	}

	postJSON, err := json.Marshal(post)
	if err == nil {
		s.cache.Set(ctx, cacheKey, postJSON, time.Hour)
	}

	return post, nil
}

func (s *PostService) GetAllPosts() ([]models.Post, error) {
	ctx := context.Background()
	cacheKey := "all_posts"
	cachedPosts, err := s.cache.Get(ctx, cacheKey).Result()
	if err == nil {
		var posts []models.Post
		err := json.Unmarshal([]byte(cachedPosts), &posts)
		if err == nil {
			return posts, nil
		}
	}

	posts, err := s.postRepo.GetAllPosts()
	if err != nil {
		return nil, err
	}

	postsJSON, err := json.Marshal(posts)
	if err == nil {
		s.cache.Set(ctx, cacheKey, postsJSON, time.Hour)
	}

	return posts, nil
}

func (s *PostService) UpdatePost(post *models.Post) error {
	err := s.postRepo.UpdatePost(post)
	if err == nil {
		s.cache.Del(context.Background(), "post:"+strconv.Itoa(int(post.ID)))
		s.cache.Del(context.Background(), "all_posts")
	}
	return err
}

func (s *PostService) DeletePost(id uint) error {
	err := s.postRepo.DeletePost(id)
	if err == nil {
		s.cache.Del(context.Background(), "post:"+strconv.Itoa(int(id)))
		s.cache.Del(context.Background(), "all_posts")
	}
	return err
}
