package models

import "time"

type GormModel struct {
	ID        uint       `gorm:"not null;primaryKey" json:"id" form:"id"`
	CreatedAt *time.Time `gorm:"not null" json:"created_at,omitempty" form:"created_at"`
	UpdatedAt *time.Time `gorm:"not null" json:"updated_at,omitempty" form:"updated_at"`
}
