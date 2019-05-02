package main

import (
	"github.com/bb3104/salmon/db"
	"github.com/bb3104/salmon/server"
)

func main() {
	db.Init()
	server.Init()
}
