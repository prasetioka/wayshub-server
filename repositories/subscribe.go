package repositories

import (
	"wayshub-server/models"

	"gorm.io/gorm"
)

type SubscribeRepository interface {
	FindSubscribes(ChannelID int) ([]models.Subscribe, error)
	GetSubscribe(ChannelID int) (models.Subscribe, error)
	
	GetSubscribeByOther(OtherID int) (models.Subscribe, error)

	GetOther(ID int) (models.Channel, error)

	CreateSubscribe(subscribe models.Subscribe) (models.Subscribe, error)

	DeleteSubscribe(subscribe models.Subscribe) (models.Subscribe, error)

	GetSubscription(ChannelID int) ([]models.Subscribe, error)
}

func RepositorySubscribe(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetOther(ID int) (models.Channel, error) {
	var channel models.Channel
	err := r.db.Preload("Subscription.Channel").Preload("Video.Channel").Preload("Video").First(&channel, ID).Error

	return channel, err
}

func (r *repository) FindSubscribes(ChannelID int) ([]models.Subscribe, error) {
	var subscribes []models.Subscribe
	err := r.db.Preload("Channel").Where("channel_id = ?", ChannelID).Find(&subscribes).Error

	return subscribes, err
}

func (r *repository) GetSubscribe(ChannelID int) (models.Subscribe, error) {
	var subscribe models.Subscribe
	err := r.db.Preload("Channel").Where("channel_id = ?", ChannelID).First(&subscribe).Error

	return subscribe, err
}

func (r *repository) GetSubscribeByOther(OtherID int) (models.Subscribe, error) {
	var subscribe models.Subscribe
	err := r.db.Preload("Channel").Where("other_id = ?", OtherID).First(&subscribe).Error

	return subscribe, err
}

func (r *repository) CreateSubscribe(subscribe models.Subscribe) (models.Subscribe, error) {
	err := r.db.Create(&subscribe).Error

	return subscribe, err
}

// func (r *repository) DeleteSubscribe(OtherID int) (models.Subscribe, error) {
// 	var subscribe models.Subscribe
// 	err := r.db.Where("other_id = ?", OtherID).Delete(&subscribe).Error

// 	return subscribe, err
// }

func (r *repository) DeleteSubscribe(subscribe models.Subscribe) (models.Subscribe, error) {
	err := r.db.Delete(&subscribe).Error

	return subscribe, err
}

func (r *repository) GetSubscription(ChannelID int) ([]models.Subscribe, error) {
	var subscription []models.Subscribe
	err := r.db.Preload("Channel").Preload("OtherVideo.Channel").Find(&subscription).Error

	return subscription, err
}

func (r *repository) GetOtherVid(ID int) (models.Video, error) {
	var channel models.Video
	err := r.db.First(&channel, ID).Error

	return channel, err
}