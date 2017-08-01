package main

// ItemSet contains information regarding items of a set and its bonuses.
type ItemSet struct {
	BonusIsSecret bool       `json:"bonusIsSecret"`
	Effects       [][]Effect `json:"effects"`
	ID            int        `json:"id"`
	Items         []int      `json:"items"`
	NameID        I18N       `json:"nameId"`
}
