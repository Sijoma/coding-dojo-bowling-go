package main

import (
	"strconv"
	"testing"
)

func TestGame_BonusPoints(t *testing.T) {
	t.Run("Spare Roll", func(t *testing.T) {
		game := NewGame()
		game.addRoll(5)
		game.addRoll(5)
		game.addRoll(3)
		game.addRoll(1)
		if game.score() != 17 {
			t.Errorf("❌ Score = %d; want 17", game.score())
		}
	})

	t.Run("Strike Roll", func(t *testing.T) {
		game := NewGame()
		game.addRoll(10)
		game.addRoll(2)
		game.addRoll(2)
		if game.score() != 18 {
			t.Errorf("❌ Score = %d; want 18", game.score())
		}
	})

	t.Run("Double Strike", func(t *testing.T) {
		game := NewGame()
		game.addRoll(10)
		game.addRoll(10)
		game.addRoll(5)
		if game.score() != 40 {
			t.Errorf("❌ Score = %d; want 40", game.score())
		}
	})

	t.Run("Strike ➡️ Spare ➡️ Gutter ➡️ 2 Pins", func(t *testing.T) {
		game := NewGame()
		game.addRoll(10)
		game.addRoll(2)
		game.addRoll(8)
		game.addRoll(0)
		game.addRoll(2)
		if game.score() != 32 {
			t.Errorf("❌ Score = %d; want 32", game.score())
		}
	})
}

func TestGame_SingleGameE2E(t *testing.T) {
	rolls := []int{1, 4, 4, 5, 6, 4, 5, 5, 10, 0, 1, 7, 3, 6, 4, 10, 2, 8, 6}
	totalScores := []int{1, 5, 9, 14, 20, 24, 34, 39, 59, 59, 61, 68, 71, 83, 87, 107, 111, 127, 133}
	g := NewGame()
	for index, roll := range rolls {
		t.Run("Game at roll number: "+strconv.Itoa(index), func(t *testing.T) {
			g.addRoll(roll)
			got := g.score() == totalScores[index]
			if got != true {
				t.Errorf("❌ Test at roll %d; score is %d, expected score %d", index, g.score(), totalScores[index])
			}
		})
	}

	t.Run("Game should be over", func(t *testing.T) {
		got := g.gameOver() == true
		if got != true {
			t.Errorf("❌ Game is expected to be over but it is not")
		}
	})

}

// func BenchmarkHello(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		fmt.Sprintf("hello")
// 	}
// }

func Test_IsStrike(t *testing.T) {
	strikeFrame := NewFrame()
	strikeFrame.addRoll(10)
	got := strikeFrame.isStrike()
	if got != true {
		t.Errorf("❌ isStrike = %t; want true", got)
	}
}
func Test_IsSpare(t *testing.T) {
	spareFrame := NewFrame()
	spareFrame.addRoll(5)
	spareFrame.addRoll(5)
	got := spareFrame.isSpare()
	if got != true {
		t.Errorf("❌ isStrike = %t; want true", got)
	}
}
