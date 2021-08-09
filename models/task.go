package models

import "time"

type Task struct {
	Model
	Name   string `gorm:"column:name;not null;" json:"name" form:"name"`
	Text   string `gorm:"column:text;not null;" json:"text" form:"text"`
	User   User   `json:"user" gorm:"foreignKey:UserRefer"`
	UserRefer uint
	Items  []Item `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:TaskRef" json:"items"`
}

func (m *Task) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Task) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
