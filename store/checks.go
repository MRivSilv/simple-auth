package store

func (s *Storage) CheckEmail(email string) (bool, error) {
	var exists bool
	err := s.db.QueryRow(
		checkEmail,
		email).Scan(&exists)
	return exists, err
}

func (s *Storage) CheckUserId(userid string) (bool, error) {
	var exists bool
	err := s.db.QueryRow(
		checkUserID, userid).Scan(&exists)
	return exists, err
}

func (s *Storage) CheckAppName(name string) (bool, error) {
	var exists bool
	err := s.db.QueryRow(
		checkAppName,
		name).Scan(&exists)
	return exists, err
}

func (s *Storage) CheckAppId(appid string) (bool, error) {
	var exists bool
	err := s.db.QueryRow(
		checkAppID,
		appid).Scan(&exists)
	return exists, err
}
