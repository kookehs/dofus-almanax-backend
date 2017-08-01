package main

// Characteristic contains sources which affect the value.
type Characteristic struct {
	Additional           int
	Base                 int
	ObjectsAndMountBonus int
}

// PrimaryCharacteristics contain both basic and advance characteristics.
type PrimaryCharacteristics struct {
	Agility      Characteristic
	AP           Characteristic
	Chance       Characteristic
	HP           Characteristic
	Initiative   Characteristic
	Intelligence Characteristic
	MaxRange     Characteristic
	MP           Characteristic
	Prospecting  Characteristic
	Strength     Characteristic
	Summons      Characteristic
	Vitality     Characteristic
	Wisdom       Characteristic
}

// SecondaryCharacteristics contain advance characteristics.
type SecondaryCharacteristics struct {
	APReduction      Characteristic
	APLossResistance Characteristic
	CriticalHits     Characteristic
	Dodge            Characteristic
	Heals            Characteristic
	Lock             Characteristic
	MPReduction      Characteristic
	MPLossResistance Characteristic
}

// Damage contains characteristics related to attacks.
type Damage struct {
	AirFixed       Characteristic
	CriticalDamage Characteristic
	Damage         Characteristic
	Distance       Characteristic
	EarthFixed     Characteristic
	FireFixed      Characteristic
	Melee          Characteristic
	NeutralFixed   Characteristic
	Power          Characteristic
	Pushback       Characteristic
	Reflect        Characteristic
	Spells         Characteristic
	TrapsFixed     Characteristic
	TrapsPower     Characteristic
	WaterFixed     Characteristic
	Weapon         Characteristic
	WeaponSkill    Characteristic
}

// Resistances contain characteristics related to defense.
type Resistances struct {
	AirFixed          Characteristic
	AirPercent        Characteristic
	CriticalHitsFixed Characteristic
	Distance          Characteristic
	EarthFixed        Characteristic
	EarthPercent      Characteristic
	FireFixed         Characteristic
	FirePercent       Characteristic
	Melee             Characteristic
	NeurtalFixed      Characteristic
	NeurtalPercent    Characteristic
	PushbackFixed     Characteristic
	WaterFixed        Characteristic
	WaterPercent      Characteristic
}
