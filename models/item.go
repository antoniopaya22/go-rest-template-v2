package models

import "time"

type Item struct {
	Model
	Name   string `gorm:"column:name;not null;" json:"name" form:"name"`
	TaskRef uint
}

func (m *Item) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Item) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
