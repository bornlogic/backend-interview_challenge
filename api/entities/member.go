package main

import (
	"github.com/jinzhu/gorm"
	"github.com/marciusvinicius/Interview-Backend-Code-Challenge/pkg/matrix"
)

type Member struct {
	gorm.Model
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
	Notes     []Note `json:"notes"`
}

type Authentication struct {
	gorm.Model
	Member Member        `json:"member"`
	Matrix matrix.Matrix `json:"matrix"`
}
