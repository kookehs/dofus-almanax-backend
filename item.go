package main

// Item contains various statistics.
type Item struct {
	AppearanceID          int         `json:"appearanceId"`
	BonusIsSecret         bool        `json:"bonusIsSecret"`
	ContainerIDs          []int       `json:"containerIds"`
	CraftXpRatio          int         `json:"craftXpRatio"`
	Criteria              string      `json:"criteria"`
	CirteriaTarget        string      `json:"criteriaTarget"`
	Cursed                bool        `json:"cursed"`
	DescriptionID         I18N        `json:"descriptionId"`
	DropMonsterIDs        []int       `json:"dropMonsterIds"`
	Etheral               bool        `json:"etheral"`
	Exchangeable          bool        `json:"exchangeable"`
	HideEffects           bool        `json:"hideEffects"`
	IconID                int         `json:"iconId"`
	ID                    int         `json:"id"`
	IsDestructible        bool        `json:"isDestructible"`
	ItemSetID             int         `json:"itemSetId"`
	FavoriteSubAreas      []int       `json:"favoriteSubAreas"`
	FavoriteSubAreasBonus int         `json:"favoriteSubAreasBonus"`
	Level                 int         `json:"level"`
	NameID                I18N        `json:"nameId"`
	NeedUseConfirm        bool        `json:"needUseConfirm"`
	NonUsableOnAnother    bool        `json:"nonUsableOnAnother"`
	NuggetsBySubArea      [][]float64 `json:"nuggetsBySubarea"`
	PossibleEffects       []Effect    `json:"possibleEffects"`
	Price                 float64     `json:"price"`
	RecipeIDs             []int       `json:"recipeIds"`
	RecipeSlots           int         `json:"recipeSlots"`
	ResourcesBySubArea    [][]int     `json:"resourcesBySubarea"`
	SecretRecipe          bool        `json:"secretRecipe"`
	Targetable            bool        `json:"targetable"`
	TwoHanded             bool        `json:"twoHanded"`
	TypeID                int         `json:"typeId"`
	Usable                bool        `json:"usable"`
	UseAnimationID        int         `json:"useAnimationId"`
}
