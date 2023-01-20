package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/EraldBa/webApp/pkg/helpers"
	"github.com/EraldBa/webApp/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

// InsertUser inserts new user to database
func (m *postgresDBRepo) InsertUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return err
	}
	stmt := `insert into users 
    			(username, email, password, access_level, created_at, updated_at)
				values ($1, $2, $3, $4, $5, $6)`

	_, err = m.DB.ExecContext(ctx, stmt,
		user.Username,
		user.Email,
		hashedPassword,
		user.AccessLevel,
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

	_, err := m.DB.ExecContext(ctx, stmt,
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

	_, err := m.DB.ExecContext(ctx, stmt,
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
func (m *postgresDBRepo) GetStats(date string, userID uint) *models.StatsSend {
	var statsSend models.StatsSend

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select breakfast, lunch, dinner, snacks, protein, carbs, fats
				from stats
				where user_id = $1 and date = $2`
	row := m.DB.QueryRowContext(ctx, query, userID, date)

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
	helpers.ErrorCheck(err)

	return &statsSend
}

// CheckStats checks if stats row exists
func (m *postgresDBRepo) CheckStats(date string, userID uint) error {
	var test []byte
	query := `select * from stats where user_id = $1 and date = $2`

	err := m.DB.QueryRow(query, userID, date).Scan(&test)
	if err != sql.ErrNoRows {
		return nil
	}
	return err
}

func (m *postgresDBRepo) Authenticator(username, testPassword string) (uint, string, error) {
	var id uint
	var hashedPassword string

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, "select id, password from users where username = $1", username)

	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return id, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))
	if err != nil {
		return 0, "", err
	}

	return id, hashedPassword, nil
}

func (m *postgresDBRepo) GetUserById(id uint) (*models.User, error) {
	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, username, password, access_level, created_at, updated_at
			from users where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.AccessLevel,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	return &user, err
}
