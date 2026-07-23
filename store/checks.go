package store

func (s *UserStore) CheckEmail(email string) (bool, error) {
	var exists bool
	err:=s.db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)",
		email).Scan(&exists)
		return exists, err
}