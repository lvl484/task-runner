package main

import (
	"context"
	"github.com/lvl484/task-runner/config"
	"github.com/lvl484/task-runner/database"
	"github.com/lvl484/task-runner/runner"
	"github.com/lvl484/task-runner/scheduler"
	"github.com/lvl484/task-runner/server"
	"github.com/lvl484/task-runner/service"
	"log"
)

func main() {
	config := config.Init()
	db := database.NewMemory()
	run := runner.NewBash()
	sch, err := scheduler.NewScheduler(context.Background(), run, db)
	if err != nil {
		log.Fatalln(err)
	}
	srv := service.NewService(db, sch)
	s := server.NewHTTP(srv, config.Address())
	err = s.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
