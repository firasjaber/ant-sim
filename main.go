package main

import (
	"log"
	"os"

	"github.com/firasjaber/ant-sim/examples"
	"github.com/firasjaber/ant-sim/scene"
)

func main() {
	sceneArg := os.Args[1:]
	if (sceneArg == nil) || (len(sceneArg) == 0) {
		log.Panicln("Please choose a scene to simulate")
	}
	chosenScene := sceneArg[0]

	switch chosenScene {
	case "home":
		scene := examples.NewHomeScene()
		scene.Run()
	case "wandering":
		scene := examples.NewWanderingScene()
		scene.Run()
	case "food":
		scene := examples.NewFoodScene()
		scene.Run()
	case "pheromone":
		scene := examples.NewPhScene()
		scene.Run()
	case "one_nest":
		scene := scene.NewScene(1)
		scene.Run()
	case "two_nests":
		scene := scene.NewScene(2)
		scene.Run()
	case "four_nests":
		scene := scene.NewScene(3)
		scene.Run()
	case "surrounding":
		scene := scene.NewScene(4)
		scene.Run()
	default:
		log.Panicln("Invalid scene name")
	}
}
