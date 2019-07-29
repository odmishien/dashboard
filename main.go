package main

import (
	"dashboard/db"
	"dashboard/server"
)

func main() {
	db.Init()
	server.Init()
}
