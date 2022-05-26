package model


type User struct{
	ID uint `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Email string `gorm:"not null unique"`
	Password []byte `json:"password"`
}

