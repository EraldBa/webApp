package dbrepo

import (
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

func (m *postgresDBRepo) InsertNewStats(s *models.StatsGet) error {
	statement := `insert into stats 
				(date, $1, protein, carbs, fats, user_id, updated_at)
				values($2, $3, $4, $5, $6, $7, $8)`
	_, err := m.DB.Exec(statement,
		s.TimeOfDay,
		s.Date,
		s.Calories,
		s.Protein,
		s.Carbs,
		s.Fats,
		s.UserID,
		time.Now(),
	)

	return err
}

func (m *postgresDBRepo) UpdateStats(s *models.StatsGet) error {
	statement := `update stats
				set $1 = $1 + $2, protein = protein + $3, carbs = carbs + $4, fats = fats + $5, updated_at = $6  
				where user_id = $7`
	_, err := m.DB.Exec(statement,
		s.TimeOfDay,
		s.Calories,
		s.Protein,
		s.Carbs,
		s.Fats,
		time.Now(),
		s.UserID,
	)

	return err
}

func (m *postgresDBRepo) GetStats(s *models.StatsSend) error {
	return nil
}

func (m *postgresDBRepo) CheckStats(date string) error {
	statement := `select * from stats where date = $1`

	row := m.DB.QueryRow(statement, date)

	return row.Err()
}
