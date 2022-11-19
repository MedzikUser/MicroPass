package database

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Model struct {
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt
}

func UpdatedOrDeletedAfter(t time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("updated_at >= ? OR deleted_at >= ?", t, t.Unix())
	}
}
