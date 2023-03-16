package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	TodoId          int            `json:"id" gorm:"column:activity_id;primaryKey"`
	Title           string         `json:"title" gorm:"column:title"`
	ActivityGroupId int            `json:"activity_group_id" gorm:"column:activity_group_id"`
	IsActive        bool           `json:"is_active" gorm:"column:is_active"`
	Priority        string         `json:"priority" gorm:"column:priority"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`
}
