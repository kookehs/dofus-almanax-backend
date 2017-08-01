package main

// Effect contains information about buffs and debuffs.
type Effect struct {
	Delay             int    `json:"delay"`
	DiceNum           int    `json:"diceNum"`
	DiceSide          int    `json:"diceSide"`
	Duration          int    `json:"duration"`
	EffectElement     int    `json:"effectElement"`
	EffectID          int    `json:"effectId"`
	EffectUID         int    `json:"effectUid"`
	Group             int    `json:"group"`
	Random            int    `json:"random"`
	RawZone           string `json:"rawZone"`
	TargetID          int    `json:"targetId"`
	TargetMask        string `json:"targetMask"`
	Triggers          string `json:"triggers"`
	Value             int    `json:"value"`
	VisibleInBuffUI   bool   `json:"visibleInBuffUi"`
	VisibleInFightLog bool   `json:"visibleInFightLog"`
	VisibleInTooltip  bool   `json:"visibleInTooltip"`
}
