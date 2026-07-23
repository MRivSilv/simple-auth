package store

import (
	"database/sql"

	"github.com/lib/pq"

	"fmt"
)

type User struct {
	UserID   string
	Email    string
	Username string
	Hash     string
	AppIDs   []string
}

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{db: db}

}

func (s *UserStore) Create(email, username, hash, appid string) (string, error) {
	var userid string

	appIDs := []string{appid}
	//Logic to check if email already being used before creating the user
	exists, emailerror := s.CheckEmail(email)
	if emailerror != nil {
		return "", emailerror
	}
	if exists {
		return "", fmt.Errorf("Email already in usage")
	}
	//Query in order to create the user
	err := s.db.QueryRow(registerUser, email, username, hash, pq.Array(appIDs)).Scan(&userid)
	if err != nil {
		return "", err
	}
	return userid, nil
}

func (s *UserStore) Get(userID string) (*User, error) {
	var user User
	err := s.db.QueryRow(
		getUser, userID,
	).Scan(&user.UserID, &user.Email, &user.Username, &user.Hash, pq.Array(&user.AppIDs))
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserStore) GetByEmail(email string) (*User, error){
	var user User
	err:= s.db.QueryRow(getUserByEmail, email).Scan(&user.UserID, &user.Email, &user.Username, &user.Hash, pq.Array(&user.AppIDs))
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserStore) Delete(userID string) (bool, error) {
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
