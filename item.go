package main

// Item contains various statistics.
type Item struct {
	appearanceID          int
	bonusIsSecret         bool
	containerIDs          []int
	craftXpRatio          int
	criteria              string
	cirteriaTarget        string
	cursed                bool
	descriptionID         I18N
	dropMonsterIDs        []int
	etheral               bool
	exchangeable          bool
	hideEffects           bool
	iconID                int
	id                    int
	isDestructible        bool
	itemSetID             int
	favoriteSubAreas      []int
	favoriteSubAreasBonus int
	level                 int
	nameID                I18N
	needUseConfirm        bool
	nonUsableOnAnother    bool
	nuggetsBySubArea      [][]float64
	possibleEffects       []Effect
	price                 float64
	realWeight            int
	recipeIDs             []int
	recipeSlots           int
	resourcesBySubArea    [][]int
	secretRecipe          bool
	targetable            bool
	twoHanded             bool
	typeID                int
	usable                bool
	useAnimationID        int
}
