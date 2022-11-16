package database

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type Model struct {
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt
}
