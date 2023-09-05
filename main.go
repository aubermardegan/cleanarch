package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/amardegan/cleanarch/infrastructure/repository"
	"github.com/amardegan/cleanarch/usecase/shapes/circle"
	"github.com/amardegan/cleanarch/usecase/shapes/square"
)

func main() {
	dataRepo := repository.NewInMem()

	circleService := circle.NewService(dataRepo)

	circleService.CreateCircle(true, 2)
	circleService.CreateCircle(true, 3)

	squareService := square.NewService(dataRepo)

	squareService.CreateSquare(4, 12)
	squareService.CreateSquare(4, 15)

	circles, err := circleService.ListCircles()
	if err != nil {
		log.Fatal(err)
	}

	for _, circle := range circles {
		jsonCircle, err := json.Marshal(circle)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonCircle))
	}

	squares, err := squareService.ListSquares()
	if err != nil {
		log.Fatal(err)
	}

	for _, square := range squares {
		jsonSquare, err := json.Marshal(square)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonSquare))
	}

	data, err := dataRepo.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range data {
		jsonData, err := json.Marshal(d)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonData))
	}
}
