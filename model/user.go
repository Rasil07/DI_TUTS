package model

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


type User struct{
	ID uint `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Email string `gorm:"not null;unique"`
	Password string `json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error){
	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err!=nil{
		fmt.Println("Could not has password")
		return err
	}
	u.Password = string(hashedPassword)

	return nil
}


