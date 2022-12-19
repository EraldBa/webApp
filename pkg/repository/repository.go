package repository

import "github.com/EraldBa/webApp/pkg/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertUser(u *models.User) error
	InsertNewStats(s *models.StatsGet) error
	UpdateStats(s *models.StatsGet) error
	GetStats(s *models.StatsSend) error
	CheckStats(date string) error
}
