package repositories

import (
	"wayshub-server/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	FindComments() ([]models.Comments, error)
	GetComment(ID int) (models.Comments, error)
	CreateComment(comment models.Comments) (models.Comments, error)
	UpdateComment(comment models.Comments) (models.Comments, error)
	DeleteComment(comment models.Comments) (models.Comments, error)
}

func RepositoryComment(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindComments() ([]models.Comments, error) {
	var comments []models.Comments
	err := r.db.Preload("Channel").Preload("Video").Find(&comments).Error

	return comments, err
}

func (r *repository) GetComment(ID int) (models.Comments, error) {
	var comment models.Comments
	err := r.db.Preload("Channel").Preload("Video").First(&comment, ID).Error
  
	return comment, err
}

func (r *repository) CreateComment(comment models.Comments) (models.Comments, error) {
	err := r.db.Create(&comment).Error

	return comment, err
}

// func (r *repository) CreateCommentByVideo(comment models.Comments) (models.Comments, error) {
// 	err := r.db.Preload("Channel").Preload("Video").Create(&comment).Error

// 	return comment, err
// }

func (r *repository) UpdateComment(comment models.Comments) (models.Comments, error) {
	err := r.db.Preload("Channel").Save(&comment).Error

	return comment, err
}

func (r *repository) DeleteComment(comment models.Comments) (models.Comments, error) {
	err := r.db.Delete(&comment).Error

	return comment, err
}