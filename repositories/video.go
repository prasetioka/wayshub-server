package repositories

import (
	"wayshub-server/models"

	"gorm.io/gorm"
)

type VideoRepository interface {
	FindVideosByChannelId(channelID int) (models.Video, error)
	FindMyVideos(channelID int) ([]models.Video, error)
	
	FindVideos() ([]models.Video, error)
	GetVideo(ID int) (models.Video, error)
	CreateVideo(video models.Video) (models.Video, error)
	UpdateVideo(video models.Video) (models.Video, error)
	DeleteVideo(video models.Video) (models.Video, error)

	// ViewsCounter(id int) error

	UpdateViews(video models.Video) (models.Video, error)
}

func RepositoryVideo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindVideos() ([]models.Video, error) {
	var videos []models.Video
	err := r.db.Preload("Channel").Find(&videos).Error

	return videos, err
}

func (r *repository) FindVideosByChannelId(channelID int) (models.Video, error) {
	var video models.Video
	err := r.db.Preload("Channel").Where("channel_id = ?", channelID).Find(&video).Error
	return video, err
}

func (r *repository) FindMyVideos(channelID int) ([]models.Video, error) {
	var videos []models.Video
	err := r.db.Preload("Channel").Where("channel_id = ?", channelID).Find(&videos).Error

	return videos, err
}

func (r *repository) GetVideo(ID int) (models.Video, error) {
	var video models.Video
	err := r.db.Preload("Channel").Preload("Comments.Channel").First(&video, ID).Error

	return video, err
}

func (r *repository) CreateVideo(video models.Video) (models.Video, error) {
	err := r.db.Create(&video).Error

	return video, err
}

func (r *repository) UpdateVideo(video models.Video) (models.Video, error) {
	err := r.db.Preload("Channel").Save(&video).Error

	return video, err
}

func (r *repository) DeleteVideo(video models.Video) (models.Video, error) {
	err := r.db.Delete(&video).Error

	return video, err
}

// func (r *repository) ViewsCounter(id int) error {
// 	var view models.Video
// 	// r.db.First(&view, id)

// 	// view.ViewCount += 1
// 	err := r.db.Save(&view).Error
// 	return err
// }

func (r *repository) UpdateViews(video models.Video) (models.Video, error) {
	err := r.db.Preload("Channel").Save(&video).Error

	return video, err
}
