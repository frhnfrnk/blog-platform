package handlers

import (
	"context"
	"strconv"

	"github.com/frhnfrnk/blog-platform-microservices/comment-service/internal/models"
	"github.com/frhnfrnk/blog-platform-microservices/comment-service/internal/services"
	"github.com/frhnfrnk/blog-platform-microservices/comment-service/pb"
)

type CommentHandler struct {
	commentService *services.CommentService
	pb.UnimplementedCommentServiceServer
}

func NewCommentHandler(commentService *services.CommentService) *CommentHandler {
	return &CommentHandler{commentService: commentService}
}

func (h *CommentHandler) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	comment := &models.Comment{
		PostID:  req.PostId,
		UserID:  req.UserId,
		Content: req.Content,
	}

	err := h.commentService.CreateComment(comment)
	if err != nil {
		return nil, err
	}

	return &pb.CreateCommentResponse{Comment: &pb.Comment{Id: string(comment.ID), PostId: req.PostId, UserId: req.UserId, Content: comment.Content}}, nil
}

func (h *CommentHandler) GetComment(ctx context.Context, req *pb.GetCommentRequest) (*pb.GetCommentResponse, error) {
	comment, err := h.commentService.GetCommentByID(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.GetCommentResponse{Comment: &pb.Comment{Id: string(comment.ID), PostId: string(comment.PostID), UserId: comment.UserID, Content: comment.Content}}, nil
}

func (h *CommentHandler) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.UpdateCommentResponse, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	comment, err := h.commentService.GetCommentByID(uint(id))
	if err != nil {
		return nil, err
	}

	comment.PostID = req.PostId
	comment.Content = req.Content
	comment.UserID = req.UserId

	err = h.commentService.UpdateComment(comment)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateCommentResponse{Comment: &pb.Comment{Id: string(comment.ID), PostId: req.PostId, UserId: req.UserId, Content: comment.Content}}, nil
}

func (h *CommentHandler) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	err := h.commentService.DeleteComment(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteCommentResponse{Message: "Comment deleted successfully"}, nil
}

func (h *CommentHandler) GetAllComments(ctx context.Context, req *pb.GetAllCommentsRequest) (*pb.GetAllCommentsResponse, error) {
	comments, err := h.commentService.GetAllComments()
	if err != nil {
		return nil, err
	}

	var pbComments []*pb.Comment
	for _, comment := range comments {
		pbComments = append(pbComments, &pb.Comment{Id: string(comment.ID), PostId: comment.PostID, UserId: comment.UserID, Content: comment.Content})
	}

	return &pb.GetAllCommentsResponse{Comments: pbComments}, nil
}
