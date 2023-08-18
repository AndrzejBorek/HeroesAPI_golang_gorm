package main

import (
	"github.com/AndrzejBorek/HeroesAPI_golang_gorm/internal/app"
)

func main() {
	r := app.SetupRouter()
	err := r.Run()
	if err != nil {
		return
	}
}
