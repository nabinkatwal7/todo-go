package model

import (
	"html"
	"strings"
	"todo-api/db"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)
type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"`
	Todos []ToDo	
}

func (user *User) Save() (*User, error){
	err := db.Database.Create(&user).Error

	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func FindUserByUsername(username string)(*User, error){
	var user User
	err:=db.Database.Where("username = ?", username).First(&user).Error
	if err != nil {
		return &User{}, err
	}
	return &user, err
}

func FindUserById(id uint) (User, error){
	var user User
	err:=db.Database.Preload("Todos").Where("id = ?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}