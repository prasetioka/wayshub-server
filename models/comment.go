package models

type Comments struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	Comment   string               `json:"comment" gorm:"type: varchar (255)"`
	ChannelID int                  `json:"channel_id"`
	Channel   ChannelVideoResponse `json:"channel" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	VideoID   int                  `json:"video_id"`
}

type CommentsVideoResponse struct {
	ID      int                  `json:"id"`
	Comment string               `json:"comment"`
	Channel ChannelVideoResponse `json:"channel"`
}

func (CommentsVideoResponse) TableName() string {
	return "comments"
}
