package main

import (
	"backend/internal/pb"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"log"
)

func main() {
	app := pb.CreateApp()

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: true,
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
