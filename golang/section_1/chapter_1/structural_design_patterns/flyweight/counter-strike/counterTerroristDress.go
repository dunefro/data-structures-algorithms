package main

type counterTerroristDress struct {
	color string
}

func (ct *counterTerroristDress) getColor() string {
	return ct.color
}

func newCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{
		color: "green",
	}
}
