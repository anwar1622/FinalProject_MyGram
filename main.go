package main

import (
	"MyGram/database"
	"MyGram/routers"
)

func main() {
	database.StartDB()
	r := routers.Routers()
	r.Run(":8080")
}
