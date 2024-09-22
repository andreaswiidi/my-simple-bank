package main

import "github.com/andreaswiidi/my-simple-bank/config"

func main() {
	// database connection
	db := config.ConnectDataBase()
	postgresDB, _ := db.DB()
	defer postgresDB.Close()
}
