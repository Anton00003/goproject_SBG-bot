package main

import (
	"goproject_SBG-bot/api"
	"goproject_SBG-bot/data"
	"goproject_SBG-bot/repository"
	"goproject_SBG-bot/service"
)

func main() {

	databasereader := data.NewDatabaseReader()
	repo := repository.New(databasereader)
	srv := service.New(repo)
	api := api.New(srv)
	api.Run(srv)
}
