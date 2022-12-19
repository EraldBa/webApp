package dbrepo

import (
	"fmt"
	"github.com/EraldBa/webApp/pkg/models"
	"time"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

func (m *postgresDBRepo) InsertUser(u *models.User) error {
	statement := `insert into users 
    			(username, email, password, access_level, created_at, updated_at)
				values ($1, $2, $3, $4, $5, $6)`
	_, err := m.DB.Exec(statement,
		u.Username,
		u.Email,
		u.Password,
		u.AccessLevel,
		time.Now(),
		time.Now(),
	)
	return err
}

func (m *postgresDBRepo) InsertStats(s *models.StatsForm) error {
	statement := fmt.Sprintf(`insert into stats 
				(date, %s, protein, carbs, fats, user_id, created_at, updated_at)
				values($1, $2, $3, $4, $5, $6, $7, $8)`, s.TimeOfDay)
	_, err := m.DB.Exec(statement,
		s.Date,
		s.Calories,
		s.Protein,
		s.Carbs,
		s.Fats,
		s.UserID,
		time.Now(),
		time.Now(),
	)
	return err
}
