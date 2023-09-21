package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/oluwaferanmiadetunji/simple_bank/api"
	"github.com/oluwaferanmiadetunji/simple_bank/internal/database"
	"github.com/oluwaferanmiadetunji/simple_bank/util"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load environment variables: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := database.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
