package repository

import "github.com/EraldBa/webApp/pkg/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertUser(u *models.User) error
	InsertStats(s *models.StatsForm) error
}
