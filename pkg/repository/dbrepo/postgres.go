package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/EraldBa/webApp/pkg/models"
	"log"
	"time"
)

// InsertUser inserts new user to database
func (m *postgresDBRepo) InsertUser(u *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into users 
    			(username, email, password, access_level, created_at, updated_at)
				values ($1, $2, $3, $4, $5, $6)`
	_, err := m.DB.ExecContext(
		ctx,
		stmt,
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

	stmt := `insert into stats 
				(date, %s, protein, carbs, fats, user_id, created_at, updated_at)
				values($1, $2, $3, $4, $5, $6, $7, $8)`

	stmt = fmt.Sprintf(stmt, s.TimeOfDay)

	_, err := m.DB.ExecContext(
		ctx,
		stmt,
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

// UpdateStats updates the users stats according to date, row has to already exist
func (m *postgresDBRepo) UpdateStats(s *models.StatsGet) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `update stats
				set %s = %s + $1, protein = protein + $2, carbs = carbs + $3, fats = fats + $4, updated_at = $5  
				where user_id = $6 and date = $7`

	stmt = fmt.Sprintf(stmt, s.TimeOfDay, s.TimeOfDay)

	_, err := m.DB.ExecContext(
		ctx,
		stmt,
		s.Calories,
		s.Protein,
		s.Carbs,
		s.Fats,
		time.Now(),
		s.UserID,
		s.Date,
	)

	return err
}

// GetStats returns stats of a user of a particular date
func (m *postgresDBRepo) GetStats(date string, userID int) *models.StatsSend {
	var statsSend models.StatsSend

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `select breakfast, lunch, dinner, snacks, protein, carbs, fats
				from stats
				where user_id = $1 and date = $2`
	row := m.DB.QueryRowContext(ctx, stmt, userID, date)

	if err := row.Err(); err != nil {
		log.Println("Something wrong with getting stats:", err)
		return &statsSend
	}
	err := row.Scan(
		&statsSend.Breakfast,
		&statsSend.Lunch,
		&statsSend.Dinner,
		&statsSend.Snacks,
		&statsSend.Protein,
		&statsSend.Carbs,
		&statsSend.Fats,
	)
	if err != nil {
		log.Println(err)
	}
	return &statsSend
}

// CheckStats checks if stats row exists
func (m *postgresDBRepo) CheckStats(date string, userID int) error {
	var test []byte
	stmt := `select * from stats where user_id = $1 and date = $2`

	err := m.DB.QueryRow(stmt, userID, date).Scan(&test)
	if err != sql.ErrNoRows {
		return nil
	}
	return err
}
