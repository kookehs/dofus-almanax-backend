package main

// ItemSet contains information regarding items of a set and its bonuses.
type ItemSet struct {
	bonusIsSecret bool
	effects       [][]Effect
	id            int
	item          []int
	nameID        I18N
}
