package repositories

import (
  "wayshub-server/models"

  "gorm.io/gorm"
)

type AuthRepository interface {
  Register(channel models.Channel) (models.Channel, error)
  Login(email string) (models.Channel, error)
  Getchannel(ID int) (models.Channel, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
  return &repository{db}
}

func (r *repository) Register(channel models.Channel) (models.Channel, error) {
  err := r.db.Create(&channel).Error

  return channel, err
}

func (r *repository) Login(email string) (models.Channel, error) {
  var channel models.Channel
  err := r.db.First(&channel, "email=?", email).Error

  return channel, err
}

func (r *repository) Getchannel(ID int) (models.Channel, error) {
	var channel models.Channel
	err := r.db.First(&channel, ID).Error

	return channel, err
}