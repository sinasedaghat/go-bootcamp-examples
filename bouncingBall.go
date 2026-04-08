package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

// VERSION 0
var (
	matrixV0 [][]bool
	widthV0  int
	hightV0  int

	xAxisBorderV0 string
)

const (
	empty = "  "
	ball  = "🎾"
)

func _initV0(w, h int) {
	widthV0 = w + 2
	hightV0 = h + 2

	matrixV0 = make([][]bool, widthV0)
	for i := range matrixV0 {
		matrixV0[i] = make([]bool, hightV0)
	}

	for w := range widthV0 {
		switch w {
		case 0:
			xAxisBorderV0 = "+-"
		case widthV0 - 1:
			xAxisBorderV0 = fmt.Sprintf("%s-+\n", xAxisBorderV0)
		default:
			xAxisBorderV0 = fmt.Sprintf("%s--", xAxisBorderV0)
		}
	}
}

func _courtV0() {
	screen.Clear()
	screen.MoveTopLeft()
	for h := range hightV0 {
		if h == 0 || h == hightV0-1 {
			fmt.Print(xAxisBorderV0)
		} else {
			for w := range widthV0 {
				switch w {
				case 0:
					fmt.Printf("| ")
				case widthV0 - 1:
					fmt.Printf(" |\n")
				default:
					if matrixV0[w][h] {
						fmt.Print(ball)
					} else {
						fmt.Print(empty)
					}
				}
			}
		}
	}
}

func bouncingBallV0() {
	ww := 60 // give from terminal
	hh := 24 // give from terminal
	_initV0(ww, hh)

	var (
		w int = 1
		h int = 4

		increaseW bool = true
		increaseH bool = true
	)

	matrixV0[w][h] = true
	_courtV0()
	time.Sleep(time.Second / 3)

	for {
		matrixV0[w][h] = false

		if increaseH {
			h++
			if h == hh {
				increaseH = false
			}
		} else {
			h--
			if h == 1 {
				increaseH = true
			}
		}

		if increaseW {
			w++
			if w == ww {
				increaseW = false
			}
		} else {
			w--
			if w == 1 {
				increaseW = true
			}
		}

		matrixV0[w][h] = true
		_courtV0()
		time.Sleep(time.Second / 10)
	}
}

// VERSION 1
const (
	widthV1 int = 50
	hightV1 int = 10

	cellEmpty rune = ' '
	cellBall  rune = '🏀'

	maxFrame int = 1000
	speed        = time.Second / 20
)

func bouncingBallV1() {
	board := make([][]bool, widthV1)
	for w := range board {
		board[w] = make([]bool, hightV1)
	}

	buff := make([]rune, widthV1*hightV1)

	var (
		px, py int
		vx, vy = 1, 1
	)

	screen.Clear()
	board[px][py] = true

	for range maxFrame {
		buff = buff[:0]

		for y := range board[0] {
			for x := range board {
				if board[x][y] {
					buff = append(buff, cellBall, ' ')
					board[x][y] = false
					break
				}
				buff = append(buff, cellEmpty, ' ')
			}
			buff = append(buff, '\n')
		}

		px += vx
		if px <= 0 || px >= widthV1-1 {
			vx *= -1
		}

		py += vy
		if py <= 0 || py >= hightV1-1 {
			vy *= -1
		}

		board[px][py] = true

		screen.MoveTopLeft()
		fmt.Print(string(buff))
		time.Sleep(speed)
	}
}
