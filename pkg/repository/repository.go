package repository

import "github.com/EraldBa/webApp/pkg/models"

type DatabaseRepo interface {
	InsertUser(u *models.User) error
	InsertNewStats(s *models.StatsGet) error
	UpdateStats(s *models.StatsGet) error
	GetStats(date, userID string) *models.StatsSend
	CheckStats(date, userID string) error
}
