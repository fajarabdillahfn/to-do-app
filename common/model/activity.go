package model

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ActivityId int            `json:"id" gorm:"column:activity_id;primaryKey"`
	Title      string         `json:"title" gorm:"column:title"`
	Email      string         `json:"email" gorm:"column:email" validate:"email"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}
