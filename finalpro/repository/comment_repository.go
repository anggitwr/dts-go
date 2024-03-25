package repository

import (
	"finalpro/model"

	"gorm.io/gorm"
)

type CommentRepository interface {
	GetComments() ([]model.Comment, error)
	GetCommentByID(id uint) (model.Comment, error)
	CreateComment(data model.Comment) (model.Comment, error)
	UpdateComment(data model.Comment) (model.Comment, error)
	DeleteComment(ID uint) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r commentRepository) GetComments() ([]model.Comment, error) {
	var comments []model.Comment
	err := r.db.Preload("User").Preload("Photo").Find(&comments).Error
	if err != nil {
		return []model.Comment{}, err
	}
	return comments, nil
}

func (r *commentRepository) GetCommentByID(id uint) (model.Comment, error) {
	var comment model.Comment
	err := r.db.Preload("User").Preload("Photo").Where("id = ?", id).First(&comment).Error
	if err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}

func (r commentRepository) CreateComment(data model.Comment) (model.Comment, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return model.Comment{}, err
	}
	return data, nil
}

func (r *commentRepository) UpdateComment(data model.Comment) (model.Comment, error) {
	err := r.db.Updates(&data).First(&data).Error
	if err != nil {
		return model.Comment{}, err
	}
	return data, nil
}

func (r *commentRepository) DeleteComment(commentID uint) error {
	var comment model.Comment
	comment.ID = commentID
	err := r.db.Where("id = ?", commentID).Delete(&comment).Error
	if err != nil {
		return err
	}
	return nil
}
