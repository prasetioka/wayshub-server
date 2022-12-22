package models

import "time"

type Video struct {
	ID          int                  `json:"id" gorm:"primary_key:auto_increment"`
	Title       string               `json:"title" gorm:"type: varchar(255)"`
	Thumbnail   string               `json:"thumbnail" gorm:"type: varchar(255)"`
	Description string               `json:"description" gorm:"type: varchar(255)"`
	Video       string               `json:"video" gorm:"type: varchar(255)"`
	CreatedAt   time.Time            `json:"createdAt"`
	FormatTime  string               `json:"formatTime" gorm:"type: varchar(255)"`
	ViewCount   int                  `json:"viewCount" gorm:"type: int"`
	ChannelID   int                  `json:"channel_id"`
	Channel     ChannelVideoResponse `json:"channel" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Comments    []Comments           `json:"comments"`
	SubscribeID int                  `json:"subscribe_id"`
	// Subscribe   Subscribe            `json:"subscribe" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type VideoComments struct {
	Title       string `json:"title" gorm:"type: varchar(255)"`
	Thumbnail   string `json:"thumbnail" gorm:"type: varchar(255)"`
	Description string `json:"description" gorm:"type: varchar(255)"`
	Video       string `json:"video" gorm:"type: varchar(255)"`
}

type CreateVideoRequest struct {
	Title       string `json:"title" gorm:"type: varchar(255)"`
	Thumbnail   string `json:"thumbnail" gorm:"type: varchar(255)"`
	Description string `json:"description" gorm:"type: varchar(255)"`
	Video       string `json:"video" gorm:"type: varchar(255)"`
}

func (CreateVideoRequest) TableName() string {
	return "subscribers"
}
