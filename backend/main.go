package main

import (
	"log"
	"template/backend/config"
	"template/backend/database"
	"template/backend/web"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("Starting application ...")

	// Config
	var CFG = config.GetConfig()

	// Database
	var PSQL = database.GetPSQL(CFG)
	database.InitPSQL(PSQL)
	var VALKEY = database.GetValkey(CFG)
	defer VALKEY.Close()

	// Routines

	// Web
	web.InitWeb(PSQL, VALKEY, CFG)
}
