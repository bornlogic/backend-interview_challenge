package api

import (
	"database/sql"
	"errors"
)

func (m *entities.Member) GetMember(db *sql.DB) error {
	return errors.New("Not implemented")
}

// func (m *entities.Member) UpdateMember(db *sql.DB) error {
// 	return errors.New("Not implemented")
// }

// func (m *entities.Member) DeleteMember(db *sql.DB) error {
// 	return errors.New("Not implemented")
// }

// func (m *entities.Member) CreateMember(db *sql.DB) error {
// 	return errors.New("Not implemented")
// }

// func (m *entities.Member) Authenticate(db *sql.DB) error {
// 	return errors.New("Not implemented")
// }

// func GetMembers(db *sql.DB, start, count int) ([]entities.Member, error) {
// 	return nil, errors.New("Not implemented")
// }
