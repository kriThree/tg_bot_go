package main

import (
	"english_learn/internal/app"
	"english_learn/internal/config"
	log "english_learn/internal/lib/logs"
	"fmt"
)

func main() {
	conf := config.MustLoad()

	log := log.LogInitializer(conf.Env)
	fmt.Println(conf)

	application := app.New(log, conf.Token, conf.Storage.Path)

	application.MustRun()

}
