package subscribedto

type SubscribeRequest struct {
	ChannelID int `json:"channel_id"`
	OtherID   int `json:"other_id"`
}
