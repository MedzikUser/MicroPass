package database

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	Id        string    `gorm:"size:40,primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func defaultModel() Model {
	id := uuid.New()

	return Model{Id: id.String()}
}
