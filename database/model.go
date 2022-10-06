package database

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	Uuid      string    `gorm:"size:40,primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func defaultModel() Model {
	uuid := uuid.New()

	return Model{Uuid: uuid.String()}
}
