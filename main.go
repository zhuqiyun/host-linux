package main

import (
	"linux-host/config"
	"linux-host/router"
	"linux-host/sqlite"
)

func main() {
	config.ParamsInit()
	sqlite.DbInit()
	router.RouterInit()
}