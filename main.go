package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/maripiyoko/2024-building/features/world"
)

func main() {
	game, err := world.NewGame()
	if err != nil {
		panic(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
