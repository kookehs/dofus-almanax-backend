package main

// Characteristic contains sources which affect the value.
type Characteristic struct {
	additional           int
	base                 int
	objectsAndMountBonus int
}

// PrimaryCharacteristics contain both basic and advance characteristics.
type PrimaryCharacteristics struct {
	agility      Characteristic
	ap           Characteristic
	chance       Characteristic
	hp           Characteristic
	initiative   Characteristic
	intelligence Characteristic
	maxRange     Characteristic
	mp           Characteristic
	prospecting  Characteristic
	strength     Characteristic
	summons      Characteristic
	vitality     Characteristic
	wisdom       Characteristic
}

// SecondaryCharacteristics contain advance characteristics.
type SecondaryCharacteristics struct {
	apReduction      Characteristic
	apLossResistance Characteristic
	criticalHits     Characteristic
	dodge            Characteristic
	heals            Characteristic
	lock             Characteristic
	mpReduction      Characteristic
	mpLossResistance Characteristic
}

// Damage contains characteristics related to attacks.
type Damage struct {
	damage         Characteristic
	power          Characteristic
	criticalDamage Characteristic
	neutralFixed   Characteristic
	earthFixed     Characteristic
	fireFixed      Characteristic
	waterFixed     Characteristic
	airFixed       Characteristic
	reflect        Characteristic
	weaponSkill    Characteristic
	trapsFixed     Characteristic
	trapsPower     Characteristic
	pushback       Characteristic
	spells         Characteristic
	weapon         Characteristic
	distance       Characteristic
	melee          Characteristic
}

// Resistances contain characteristics related to defense.
type Resistances struct {
	neurtalFixed      Characteristic
	neurtalPercent    Characteristic
	earthFixed        Characteristic
	earthPercent      Characteristic
	fireFixed         Characteristic
	firePercent       Characteristic
	waterFixed        Characteristic
	waterPercent      Characteristic
	airFixed          Characteristic
	airPercent        Characteristic
	criticalHitsFixed Characteristic
	pushbackFixed     Characteristic
	distance          Characteristic
	melee             Characteristic
}
