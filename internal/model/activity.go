package model

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ActivityId int            `json:"activity_id" gorm:"column:activity_id;primaryKey"`
	Title      string         `json:"title" gorm:"column:title"`
	Email      string         `json:"email" gorm:"column:email" validate:"email"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
