package main

import (
	"log"
)

func main(){
		log.Println("Testing dependency injection wi go fx")

	
		evnt := InitializeEvent()

		evnt.Start()



}