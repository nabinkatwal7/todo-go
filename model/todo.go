package model

import (
	"todo-api/db"

	"gorm.io/gorm"
)

type ToDo struct {
	gorm.Model
	Content string `gorm:"type:text;not null;" json:"content"`
	Done    bool   `gorm:"not null;" json:"done"`
	UserID  uint 
}

func (todo *ToDo) Save() (*ToDo, error){
	err := db.Database.Create(&todo).Error
	if err != nil {
		return &ToDo{}, err
	}
	return todo, nil
}

func (todo *ToDo) Update() (*ToDo, error){
	err := db.Database.Save(&todo).Error
	if err != nil {
		return &ToDo{}, err
	}
	return todo, nil
}

func (todo *ToDo) Delete() error{
	return db.Database.Delete(&todo).Error
}