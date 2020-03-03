package main

import (
	"database/sql"
	"errors"
)

func (m *member) getUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (m *member) updateUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (m *member) deleteUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (m *member) createUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getUsers(db *sql.DB, start, count int) ([]user, error) {
	return nil, errors.New("Not implemented")
}
