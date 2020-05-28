package main

import (
	"fmt"
)

func main() {
	game := NewGame()
	fmt.Println("the game is", game.frames, "score", game.currentFrame)
	game.addRoll(10)
	game.addRoll(10)
	game.addRoll(3)
	game.addRoll(3)
	game.addRoll(5)
	fmt.Println("the current frame of the game is", game.frames[game.currentFrame].standingPins)
	fmt.Println("the current score is ", game.score())
}
