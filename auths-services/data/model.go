package data

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	dbTimeout = time.Second * 5
)

type AuthsModel struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Password  string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type authsRepo struct {
	DB *sql.DB
}

// RegisterAuths implements AuthsRepo.
func (a *authsRepo) RegisterAuths(request AuthsModel) (AuthsModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		return AuthsModel{}, err
	}
	var user AuthsModel
	stmt := `insert into users (email, first_name, last_name, password, user_active, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7) returning id, email, first_name, last_name, password, user_active, created_at, updated_at `

	err = a.DB.QueryRowContext(ctx, stmt,
		request.Email,
		request.FirstName,
		request.LastName,
		hashedPassword,
		request.Active,
		time.Now(),
		time.Now(),
	).Scan(&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt)

	if err != nil {
		return AuthsModel{}, err
	}

	return user, nil
}

type AuthsRepo interface {
	FindAuthsByEmail(email string) (AuthsModel, error)
	RegisterAuths(request AuthsModel) (AuthsModel, error)
	PasswordMatches(plainText, hashPassword string) (bool, error)
}

func NewAuthsRepo(DB *sql.DB) AuthsRepo {
	return &authsRepo{DB: DB}
}

func (a authsRepo) FindAuthsByEmail(email string) (AuthsModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `select id, email, first_name, last_name, password, active, created_at, updated_at from auths where email = $1`
	var user AuthsModel
	row := a.DB.QueryRowContext(ctx, query, email)
	log.Println(query)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return AuthsModel{}, err
	}
	return user, nil
}

func (repo *authsRepo) PasswordMatches(plainText, hashPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
