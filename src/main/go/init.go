package main

import (
	"AdityaHegde/apricate-ai/src/main/go/apricate"
	"AdityaHegde/apricate-ai/src/main/go/database"
	"fmt"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	apricateService, err := apricate.NewApricateService(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = apricateService.LoadInitialData()
	if err != nil {
		fmt.Println(err)
		return
	}
}
