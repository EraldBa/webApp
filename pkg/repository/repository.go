package repository

import "github.com/EraldBa/webApp/pkg/models"

type DatabaseRepo interface {
	InsertUser(u *models.User) error
	InsertNewStats(s *models.StatsGet) error
	UpdateStats(s *models.StatsGet) error
	GetStats(date string, userID uint) *models.StatsSend
	CheckStats(date string, userID uint) error
	Authenticator(username, testPassword string) (uint, string, error)
	GetUserById(id uint) (*models.User, error)
}
