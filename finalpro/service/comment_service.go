package service

import (
	"finalpro/model"
	"finalpro/repository"
	"fmt"
)

type CommentService interface {
	GetAllComments() ([]model.ResponseComment, error)
	GetCommentByID(commentID uint) (model.ResponseComment, error)
	CreateComment(data model.Comment) (model.Comment, error)
	UpdateComment(data model.Comment, commentID uint) (model.Comment, error)
	DeleteComment(commentID uint) error
}

func NewCommentService(commentRepo repository.CommentRepository) CommentService {
	return &commentService{commentRepo: commentRepo}
}

type commentService struct {
	commentRepo repository.CommentRepository
}

func (service *commentService) GetAllComments() ([]model.ResponseComment, error) {
	resComment, err := service.commentRepo.GetComments()
	fmt.Println(resComment)
	if err != nil {
		return []model.ResponseComment{}, err
	}
	var response []model.ResponseComment
	for _, comment := range resComment {
		tempResp := model.ResponseComment{}
		tempResp.ID = comment.ID
		tempResp.PhotoID = comment.PhotoID
		tempResp.Message = comment.Message
		tempResp.User.Username = comment.User.Username
		tempResp.User.Email = comment.User.Email
		tempResp.CreatedAt = comment.CreatedAt
		tempResp.UpdatedAt = comment.UpdatedAt
		tempResp.Photo.Title = comment.Photo.Title
		tempResp.Photo.Caption = comment.Photo.Caption
		tempResp.Photo.PhotoURL = comment.Photo.PhotoURL
		response = append(response, tempResp)
	}
	fmt.Println("p")
	return response, nil
}

func (service *commentService) GetCommentByID(CommentID uint) (model.ResponseComment, error) {
	resComment, err := service.commentRepo.GetCommentByID(CommentID)
	fmt.Println(resComment)
	if err != nil {
		return model.ResponseComment{}, err
	}
	var comment model.ResponseComment
	comment.ID = resComment.ID
	comment.PhotoID = resComment.PhotoID
	comment.Message = resComment.Message
	comment.User.Username = resComment.User.Username
	comment.User.Email = resComment.User.Email
	comment.CreatedAt = resComment.CreatedAt
	comment.UpdatedAt = resComment.UpdatedAt
	comment.Photo.Title = resComment.Photo.Title
	comment.Photo.Caption = resComment.Photo.Caption
	comment.Photo.PhotoURL = resComment.Photo.PhotoURL
	return comment, nil
}

func (service *commentService) CreateComment(data model.Comment) (model.Comment, error) {
	create, err := service.commentRepo.CreateComment(data)
	if err != nil {
		return model.Comment{}, err
	}
	return create, nil
}

func (service *commentService) UpdateComment(data model.Comment, CommentID uint) (model.Comment, error) {
	entityComment := model.Comment{}
	entityComment.ID = uint(CommentID)
	entityComment.Message = data.Message
	update, err := service.commentRepo.UpdateComment(entityComment)
	if err != nil {
		return model.Comment{}, err
	}
	return update, nil
}

func (service *commentService) DeleteComment(commentID uint) error {
	err := service.commentRepo.DeleteComment(commentID)
	if err != nil {
		return err
	}
	return nil
}
