package handlers

import (
	"context"
	"github.com/frhnfrnk/blog-platform-microservices/post-service/internal/models"
	"strconv"

	"github.com/frhnfrnk/blog-platform-microservices/post-service/internal/services"
	"github.com/frhnfrnk/blog-platform-microservices/post-service/pb"
)

type PostHandler struct {
	postService *services.PostService
	pb.UnimplementedPostServiceServer
}

func NewPostHandler(postService *services.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	post := &models.Post{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: uint(req.AuthorId),
	}

	err := h.postService.CreatePost(post)
	if err != nil {
		return nil, err
	}

	return &pb.CreatePostResponse{
		Post: &pb.Post{
			Id:       strconv.Itoa(int(post.ID)),
			Title:    post.Title,
			Content:  post.Content,
			AuthorId: int32(post.AuthorID),
		},
	}, nil
}

func (h *PostHandler) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	post, err := h.postService.GetPostByID(uint(id))
	if err != nil {
		return nil, err
	}

	return &pb.GetPostResponse{
		Post: &pb.Post{
			Id:       strconv.Itoa(int(post.ID)),
			Title:    post.Title,
			Content:  post.Content,
			AuthorId: int32(post.AuthorID),
		},
	}, nil
}

func (h *PostHandler) GetAllPosts(ctx context.Context, req *pb.GetAllPostsRequest) (*pb.GetAllPostsResponse, error) {
	posts, err := h.postService.GetAllPosts()
	if err != nil {
		return nil, err
	}

	var pbPosts []*pb.Post
	for _, post := range posts {
		pbPost := &pb.Post{
			Id:       strconv.Itoa(int(post.ID)),
			Title:    post.Title,
			Content:  post.Content,
			AuthorId: int32(post.AuthorID),
		}
		pbPosts = append(pbPosts, pbPost)
	}

	return &pb.GetAllPostsResponse{Posts: pbPosts}, nil
}

func (h *PostHandler) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	post, err := h.postService.GetPostByID(uint(id))
	if err != nil {
		return nil, err
	}

	post.Title = req.Title
	post.Content = req.Content
	post.AuthorID = uint(req.AuthorId)

	err = h.postService.UpdatePost(post)
	if err != nil {
		return nil, err
	}

	return &pb.UpdatePostResponse{
		Post: &pb.Post{
			Id:       strconv.Itoa(int(post.ID)),
			Title:    post.Title,
			Content:  post.Content,
			AuthorId: int32(post.AuthorID),
		},
	}, nil
}

func (h *PostHandler) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	err = h.postService.DeletePost(uint(id))
	if err != nil {
		return nil, err
	}

	return &pb.DeletePostResponse{Message: "Post deleted successfully"}, nil
}
