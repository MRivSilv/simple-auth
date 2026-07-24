package store

import "fmt"

func (s *Storage) AddUserToApp(userid, appid string) error {
	app_exists, err_app_check := s.CheckAppId(appid)
	user_exists, err_user_check := s.CheckUserId(userid)
	if err_app_check != nil || err_user_check != nil {
		return fmt.Errorf("Error validating app/user")
	}
	if !app_exists || !user_exists {
		return fmt.Errorf("app or user not found")
	}

	_, err := s.db.Exec(
		registerUserInApp, userid, appid)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) RemoveUserFromApp(userid, appid string) error {
	app_exists, err_app_check := s.CheckAppId(appid)
	user_exists, err_user_check := s.CheckUserId(userid)
	if err_app_check != nil || err_user_check != nil {
		return fmt.Errorf("Error validating app/user")
	}
	if !app_exists || !user_exists {
		return fmt.Errorf("app or user not found")
	}

	_, err := s.db.Exec(
		deleteUserFromApp, userid, appid)
	if err != nil {
		return err
	}
	return nil
}
