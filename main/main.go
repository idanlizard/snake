package main

import (
	"flag"
	"snake/game"
)

var (
	n int
	m int
)

func init() {
	flag.IntVar(&n, "n", 0, "row size")
	flag.IntVar(&m, "m", 0, "number of rows")
}

func main() {
	flag.Parse()

	g, err := game.New(n, m)
	if err != nil {
		panic(err)
	}

	g.Start()
}
