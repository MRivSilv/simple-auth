package store

import (
	"database/sql"
	"fmt"
	"time"
)

type App struct {
	AppID       string
	AppName     string
	Description string
	CreatedAt   time.Time
}

type Storage struct {
	db *sql.DB
}

func (s *Storage) CreateApp(appName, description string) (string, error) {
	exists, namecheckerror := s.CheckAppName(appName)
	if namecheckerror != nil {
		return "", namecheckerror
	}
	if exists {
		return "", fmt.Errorf("App name already registered")
	}
	var appid string
	err := s.db.QueryRow(
		registerApp, appName, description).Scan(&appid)
	if err != nil {
		return "", err
	}
	return appid, nil
}

func (s *Storage) DeleteApp(appid string) (bool, error) {
	result, err := s.db.Exec(deleteApp, appid)
	if err != nil {
		return false, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return rows > 0, nil
}
