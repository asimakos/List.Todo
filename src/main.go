
package main

import (
	"fmt"
	"Config"
	"Routes"
	"Models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"	
)

var err error


func main() {

	// Creating a connection to the database
	Config.DB, err = gorm.Open(sqlite.Open("./db/tasks.sqlite3"), &gorm.Config{})

	if err != nil {
		fmt.Println("statuse: ", err)
		return
	}

	//defer Config.DB.Close()

	// run the migrations: todo struct
	Config.DB.AutoMigrate(&Models.Todo{})
   
	//setup routes 
	r := Routes.SetupRouter()

	// running
	r.Run(":3000")
}