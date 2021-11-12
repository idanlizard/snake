package main

import (
	"flag"
	"snake/game"
)

var (
	n   int
	m   int
	fps int
)

func init() {
	flag.IntVar(&n, "n", 20, "row size")
	flag.IntVar(&m, "m", 20, "number of rows")
	flag.IntVar(&fps, "fps", 8, "the number of frames per second (each frame advances the snake)")
}

func main() {
	flag.Parse()
	if n < 5 || m < 5 {
		panic("board size should be at leas 5 by 5")
	}

	if fps < 1 {
		panic("fps must be a positive number")
	}

	g, err := game.New(n, m, fps)
	if err != nil {
		panic(err)
	}

	g.Start()
}
