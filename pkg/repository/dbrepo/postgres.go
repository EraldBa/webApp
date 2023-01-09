package dbrepo

import (
	"context"
	"github.com/EraldBa/webApp/pkg/models"
	"log"
	"time"
)

// InsertUser inserts new user to database
func (m *postgresDBRepo) InsertUser(u *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `insert into users 
    			(username, email, password, access_level, created_at, updated_at)
				values ($1, $2, $3, $4, $5, $6)`
	_, err := m.DB.ExecContext(
		ctx,
		statement,
		u.Username,
		u.Email,
		u.Password,
		u.AccessLevel,
		time.Now(),
		time.Now(),
	)
	return err
}

// InsertNewStats inserts new stats row to database for a user
func (m *postgresDBRepo) InsertNewStats(s *models.StatsGet) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `insert into stats 
				(date, $1, protein, carbs, fats, user_id, updated_at)
				values($2, $3, $4, $5, $6, $7, $8)`

	_, err := m.DB.ExecContext(
		ctx,
		statement,
		s.TimeOfDay,
		s.Date,
		s.Calories,
		s.Protein,
		s.Carbs,
		s.Fats,
		s.UserID,
		time.Now(),
	)
	log.Println("It's insertStats")
	return err
}

// UpdateStats updates the users stats according to date, row has to already exist
func (m *postgresDBRepo) UpdateStats(s *models.StatsGet) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `update stats
				set $1 = $1 + $2, protein = protein + $3, carbs = carbs + $4, fats = fats + $5, updated_at = $6  
				where user_id = $7`
	_, err := m.DB.ExecContext(
		ctx,
		statement,
		s.TimeOfDay,
		s.Calories,
		s.Protein,
		s.Carbs,
		s.Fats,
		time.Now(),
		s.UserID,
	)
	log.Println("It's updateStats")
	return err
}

// GetStats returns stats of a user of a particular date
func (m *postgresDBRepo) GetStats(date string, userID int) *models.StatsSend {
	var statsSend models.StatsSend

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `select (breakfast, lunch, dinner, snacks, protein, carbs, fats) 
				from stats
				where user_id = $1 and date = $2`
	row := m.DB.QueryRowContext(ctx, statement, userID, date)

	if err := row.Err(); err != nil {
		log.Println("Something wrong with getting stats:", err)
		return &statsSend
	}
	_ = row.Scan(
		&statsSend.Breakfast,
		&statsSend.Lunch,
		&statsSend.Dinner,
		&statsSend.Snacks,
		&statsSend.Protein,
		&statsSend.Carbs,
		&statsSend.Fats,
	)
	return &statsSend
}

// CheckStats checks if stats row exists
func (m *postgresDBRepo) CheckStats(date string, userID int) error {
	statement := `select * from stats where user_id = $1 and date = $2`

	row := m.DB.QueryRow(statement, userID, date)

	return row.Err()
}
