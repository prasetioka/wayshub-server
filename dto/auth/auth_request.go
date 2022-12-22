package authdto

type RegisterRequest struct {
	Email       string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password    string `gorm:"type: varchar(255)" json:"password" validate:"required"`
	ChannelName string `gorm:"type: varchar(255)" json:"channelName" validate:"required"`
	Description string `gorm:"type: varchar(255)" json:"description" validate:"required"`
}

type LoginRequest struct {
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}
