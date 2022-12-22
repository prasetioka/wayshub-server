package videodto

type CreateVideoRequest struct {
	Title       string `json:"title" gorm:"type: varchar(255)"`
	Thumbnail   string `json:"thumbnail" gorm:"type: varchar(255)"`
	Description string `json:"description" gorm:"type: varchar(255)"`
	Video       string `json:"video" gorm:"type: varchar(255)"`
}

type UpdateVideoRequest struct {
	Title       string `json:"title" gorm:"type: varchar(255)"`
	Thumbnail   string `json:"thumbnail" gorm:"type: varchar(255)"`
	Description string `json:"description" gorm:"type: varchar(255)"`
	Video       string `json:"video" gorm:"type: varchar(255)"`
}

type UpdateViewsRequest struct {
	ViewCount int `json:"viewCount" gorm:"type: int"`
}
