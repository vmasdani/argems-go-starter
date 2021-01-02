package main

import (
	"time"
)

type GormModel struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type Todo struct {
	GormModel
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
