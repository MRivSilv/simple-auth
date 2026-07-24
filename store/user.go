package store

import (
	"database/sql"
	"fmt"
)

type User struct {
	UserID   string
	Email    string
	Username string
	Hash     string
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{db: db}

}

func (s *Storage) CreateUser(email, username, hash string) (string, error) {
	var userid string
	//Logic to check if email already being used before creating the user
	exists, emailerror := s.CheckEmail(email)
	if emailerror != nil {
		return "", emailerror
	}
	if exists {
		return "", fmt.Errorf("Email already in usage")
	}
	//Query in order to create the user
	err := s.db.QueryRow(registerUser, email, username, hash).Scan(&userid)
	if err != nil {
		return "", err
	}
	return userid, nil
}

func (s *Storage) GetUser(userID string) (*User, error) {
	var user User
	err := s.db.QueryRow(
		getUser, userID,
	).Scan(&user.UserID, &user.Email, &user.Username, &user.Hash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Storage) GetUserByEmail(email string) (*User, error) {
	var user User
	err := s.db.QueryRow(getUserByEmail, email).Scan(&user.UserID, &user.Email, &user.Username, &user.Hash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Storage) DeleteUser(userID string) (bool, error) {
	result, err := s.db.Exec(deleteUser, userID)
	if err != nil {
		return false, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return rows > 0, nil
}
