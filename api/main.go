package main

import (
	"team-project/app"
)

func main() {
	a := app.New()
	a.Run()
	a.Shutdown()
}
