package api

import (
	"github.com/jinzhu/gorm"
)

type Note struct {
	gorm.Model
	UUID        string `json:"uuid"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
