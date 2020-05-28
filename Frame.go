package main

// Frame the frame of a bowling game
type Frame struct {
	pinsRolled   []int
	lastFrame    bool
	standingPins int
	bonusPoints  int
	rolls        int
}

// NewFrame constructor function
func NewFrame() Frame {
	frame := Frame{[]int{0, 0, 0}, false, 10, 0, 0}
	return frame
}

func (f *Frame) addRoll(rolledPins int) bool {
	f.pinsRolled[f.rolls] = rolledPins
	f.rolls++
	f.standingPins = f.standingPins - rolledPins
	if f.lastFrame && f.standingPins == 0 {
		f.addBonusRoll()
	}
	return f.isFinished()
}

func (f *Frame) isFinished() bool {
	maxRolls := 2
	if f.lastFrame {
		maxRolls++
	}
	return f.rolls == maxRolls || f.standingPins == 0
}

func (f *Frame) addBonusRoll() {
	f.standingPins = 10
}

func (f *Frame) isStrike() bool {
	return f.pinsRolled[0] == 10
}

func (f *Frame) isSpare() bool {
	return f.pinsRolled[0]+f.pinsRolled[1] == 10 && !(f.isStrike())
}

func (f *Frame) score() int {
	var sum int
	for index := 0; index < len(f.pinsRolled); index++ {
		sum += f.pinsRolled[index]
	}
	sum += f.bonusPoints
	return sum
}

func (f *Frame) calculateBonusPoints(rolledPins int, previousFrame *Frame) {
	if previousFrame.isStrike() && f.rolls != 2 {
		previousFrame.bonusPoints += rolledPins
	} else if previousFrame.isSpare() && f.rolls == 0 {
		previousFrame.bonusPoints += rolledPins
	}
}
