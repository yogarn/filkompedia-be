package main

import "github.com/yogarn/filkompedia-be/pkg/config"

func main() {
	config.LoadEnv()

	app := config.StartFiber()

	config.StartUp(&config.Config{
		App: app,
	})
}
