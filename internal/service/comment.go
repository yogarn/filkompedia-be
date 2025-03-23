package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/yogarn/filkompedia-be/entity"
	"github.com/yogarn/filkompedia-be/internal/repository"
	"github.com/yogarn/filkompedia-be/model"
)

type ICommentService interface {
	GetComment(id uuid.UUID) (*entity.Comment, error)
	GetCommentByBook(bookId uuid.UUID) (*[]entity.Comment, error)
	CreateComment(commentReq *model.CreateComment, userId uuid.UUID) error
	UpdateComment(commentReq *model.UpdateComment, userId uuid.UUID, bookId uuid.UUID, commentId uuid.UUID) error
	DeleteComment(id uuid.UUID, userId uuid.UUID) error
}

type CommentService struct {
	commentRepository repository.ICommentRepository
}

func NewCommentService(commentRepository repository.ICommentRepository) ICommentService {
	return &CommentService{
		commentRepository: commentRepository,
	}
}

func (s *CommentService) GetComment(id uuid.UUID) (*entity.Comment, error) {
	return s.commentRepository.GetComment(id)
}

func (s *CommentService) GetCommentByBook(bookId uuid.UUID) (*[]entity.Comment, error) {
	comments, err := s.commentRepository.GetCommentByBook(bookId)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *CommentService) CreateComment(commentReq *model.CreateComment, userId uuid.UUID) error {
	comment := &entity.Comment{
		Id:        uuid.New(),
		UserId:    userId,
		BookId:    commentReq.BookId,
		Comment:   commentReq.Comment,
		Rating:    commentReq.Rating,
		CreatedAt: time.Now(),
	}

	return s.commentRepository.CreateComment(comment)
}

func (s *CommentService) UpdateComment(commentReq *model.UpdateComment, userId uuid.UUID, bookId uuid.UUID, commentId uuid.UUID) error {
	comment := &entity.Comment{
		Id:        commentId,
		UserId:    userId,
		BookId:    bookId,
		Comment:   commentReq.Comment,
		Rating:    commentReq.Rating,
		CreatedAt: time.Now(),
	}

	return s.commentRepository.UpdateComment(comment)
}

func (s *CommentService) DeleteComment(id uuid.UUID, userId uuid.UUID) error {
	return s.commentRepository.DeleteComment(id, userId)
}
