package models

type Subscribe struct {
	ID         int                  `json:"id" gorm:"primary_key:auto_increment"`
	OtherID    int                  `json:"other_id"`
	OtherName  string               `json:"otherName" gorm:"type: varchar(255)"`
	OtherPhoto string               `json:"otherPhoto" gorm:"type: varchar(255)"`
	ChannelID  int                  `json:"channel_id"`
	Channel    ChannelVideoResponse `json:"channels" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OtherVideo []Video              `json:"otherVideo"`
	// OtherVideoID int 				`json:"otherVideo_id"`
}

type SubscribeResponse struct {
	ID        int                  `json:"id"`
	ChannelID int                  `json:"channel_id"`
	Channel   ChannelVideoResponse `json:"-"`
}

type SubscribeChannelResponse struct {
	ID        int                  `json:"id"`
	ChannelID int                  `json:"channel_id"`
	Channel   ChannelVideoResponse `json:"-"`
}

func (SubscribeResponse) TableName() string {
	return "subscribers"
}

func (SubscribeChannelResponse) TableName() string {
	return "subscribers"
}
