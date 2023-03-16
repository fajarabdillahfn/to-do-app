package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	TodoId          int            `json:"todo_id" gorm:"column:activity_id;primaryKey"`
	ActivityGroupId int            `json:"activity_group_id" gorm:"column:activity_group_id"`
	Title           string         `json:"title" gorm:"column:title"`
	Priority        string         `json:"priority" gorm:"column:priority"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
