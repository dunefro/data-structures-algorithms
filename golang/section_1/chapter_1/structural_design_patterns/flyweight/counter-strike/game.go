package main

type game struct {
	terrorists       []*player
	counterTerrorist []*player
}

func newGame() *game {
	return &game{
		terrorists:       make([]*player, 1),
		counterTerrorist: make([]*player, 1),
	}
}

func (c *game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	c.terrorists = append(c.terrorists, player)
}

func (c *game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	c.counterTerrorist = append(c.counterTerrorist, player)
}
