package infrastructure

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



type Database struct{
	*gorm.DB
}
type DatabaseConf struct{
	HOST string
	PORT string
	DB_NAME string
	DB_USER string
	DB_PASSWORD string
}

func NewDatabase() Database{

		dbConf := DatabaseConf{
			HOST: os.Getenv("DB_HOST"),
			PORT: os.Getenv("DB_PORT"),
			DB_NAME: os.Getenv("DB_NAME"),
			DB_USER: os.Getenv("DB_USER"),
			DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		}
	
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", dbConf.DB_USER, dbConf.DB_PASSWORD, dbConf.HOST, dbConf.PORT)

	db,err:= gorm.Open(mysql.Open(dbUrl), &gorm.Config{})

	if err!=nil{
		log.Fatalf("Couldnot open database %v",err)
		log.Panic(err)
	}

	fmt.Println("Creating database if not exists")

		
	if err= db.Exec("CREATE DATABASE IF NOT EXISTS "+ dbConf.DB_NAME).Error;err!=nil{
		fmt.Println("Couldnot create db")
		log.Panicln(err)
	}

	fmt.Println("useing given database")

	if err= db.Exec(fmt.Sprintf("USE %s",dbConf.DB_NAME)).Error; err!=nil{
		fmt.Println("Couldnot use given db")
		log.Panicln(err)
	}

	fmt.Println("Database connected successfully")

	return Database{DB:db}

}