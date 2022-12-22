package authdto

type LoginResponse struct {
	ID    int    `gorm:"type: int" json:"id"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	Photo string `gorm:"type: varchar(255)" json:"photo"`
	Token string `gorm:"type: varchar(255)" json:"token"`
}

type RegisterResponse struct {
	Email string `gorm:"type: varchar(255)" json:"email"`
}

type CheckAuthResponse struct {
	ID    int    `gorm:"type: int" json:"id"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	Photo string `gorm:"type: varchar(255)" json:"photo"`
	Token string `gorm:"type: varchar(255)" json:"token"`
}
