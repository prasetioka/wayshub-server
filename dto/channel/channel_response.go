package channeldto

import "wayshub-server/models"

type ChannelResponse struct {
	ID           int                `json:"id"`
	Email        string             `json:"email"`
	ChannelName  string             `json:"channelName"`
	Description  string             `json:"description"`
	Cover        string             `json:"cover"`
	Photo        string             `json:"photo"`
	Video        []models.Video     `json:"video"`
	Subscription []models.Subscribe `json:"subscription"`
	Subscriber   int                `json:"subscriber"`
}

type DeleteResponse struct {
	ID int `json:"id"`
}
