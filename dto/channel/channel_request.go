package channeldto

type UpdateChannelRequest struct {
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	ChannelName string `json:"channelName" form:"channelName"`
	Description string `json:"description" form:"description"`
	Thumbnail   string `json:"thumbnail" form:"thumbnail"`
	Photo       string `json:"photo" form:"photo"`
}