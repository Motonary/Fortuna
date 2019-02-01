package main

import (
	server "github.com/motonary/Fortuna/api/v1"
	"github.com/motonary/Fortuna/database"
)

func main() {
	database.Connect()
	server.Main()
}
