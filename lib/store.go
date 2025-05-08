package lib

import "database/sql"

type Store struct {
	conn *sql.DB
}

func (s *Store) Init() error {
	return nil
}

func (s *Store) GetPasswords() ([]Password, error) {
	return nil, nil
}

func (s *Store) CreatePasswordManual(password CreateNewPasswordManual) error {
	return nil
}
func (s *Store) CreatePasswordAutomatic(password CreateNewPasswordAutomatic) error {
	return nil
}