package commentdto

type CreateCommentRequest struct {
	Comment string `json:"comment" gorm:"type: varchar (255)"`
}
