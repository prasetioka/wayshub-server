package models

type Channel struct {
	ID           int         `json:"id" gorm:"primary_key:auto_increment"`
	Email        string      `json:"email" gorm:"type: varchar(255)"`
	Password     string      `json:"password" gorm:"type: varchar(255)"`
	ChannelName  string      `json:"channelName" gorm:"type: varchar(255)"`
	Description  string      `json:"description" gorm:"type: varchar(255)"`
	Cover        string      `json:"cover" gorm:"type: varchar(255)"`
	Photo        string      `json:"photo" gorm:"type: varchar(255)"`
	Video        []Video     `json:"video"`
	Subscription []Subscribe `json:"subscription"`
	Subscriber   int         `json:"subscriber" gorm:"type: int"`
}

type ChannelVideoResponse struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	ChannelName string `json:"channelName"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
	Photo       string `json:"photo"`
}

type OtherChannel struct {
	ID          int    `json:"id" gorm:"primary_key:auto_increment"`
	Email       string `json:"email"`
	ChannelName string `json:"channelName"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
	Photo       string `json:"photo"`
	// Video       []Video `json:"video"`
}

func (ChannelVideoResponse) TableName() string {
	return "channels"
}

func (OtherChannel) TableName() string {
	return "channels"
}
