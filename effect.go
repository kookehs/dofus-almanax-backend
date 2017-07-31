package main

// Effect contains information about buffs and debuffs.
type Effect struct {
	delay             int
	duration          int
	effectElement     int
	effectID          int
	effectUID         int
	group             int
	random            int
	rawZone           string
	targetID          int
	targetMask        string
	triggers          string
	visibleInBuffUI   bool
	visibleInFightLog bool
	visibleInTooltip  bool
}
