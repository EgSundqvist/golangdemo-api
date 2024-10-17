package main

import (
	"github.com/EgSundqvist/config"
	"github.com/EgSundqvist/data"
	"github.com/EgSundqvist/routes"
)

var cfg config.Config

func main() {
	config.ReadConfig(&cfg)
	data.Init(cfg.Database.File,
		cfg.Database.Server,
		cfg.Database.Database,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Port)

	r := routes.SetupRouter()
	r.Run()
}
