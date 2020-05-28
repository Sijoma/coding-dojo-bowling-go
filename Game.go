package main

import (
	"fmt"
)

// Game the bowling game
type Game struct {
	frames       []Frame
	currentFrame int
}

func (g *Game) score() int {
	var sum int
	for index := 0; index < len(g.frames); index++ {
		sum += g.frames[index].score()
	}
	return sum
}

func (g *Game) addRoll(rolledPins int) {
	if g.currentFrame >= 1 {
		g.frames[g.currentFrame].calculateBonusPoints(rolledPins, &g.frames[g.currentFrame-1])
	}
	frameFinished := g.frames[g.currentFrame].addRoll(rolledPins)
	if frameFinished && !g.frames[g.currentFrame].lastFrame {
		g.frames = append(g.frames, NewFrame())
		g.currentFrame++
		if len(g.frames) == 10 {
			g.frames[g.currentFrame].lastFrame = true
		}
	}

	if g.gameOver() {
		fmt.Println("✳️ Game OVER!")
		fmt.Println("🏆 your score is", g.score())
	}
}

func (g *Game) gameOver() bool {
	return len(g.frames) >= 10 && g.frames[g.currentFrame].isFinished()
}

// NewGame constructor function
func NewGame() Game {
	newFrame := NewFrame()
	list := []Frame{newFrame}
	game := Game{list, 0}
	return game
}
