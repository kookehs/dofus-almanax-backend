package main

// Weapon contains various statistics.
type Weapon struct {
	APCost                     int         `json:"apCost"`
	AppearanceID               int         `json:"appearanceId"`
	BonusIsSecret              bool        `json:"bonusIsSecret"`
	CastInDiagonal             bool        `json:"castInDiagonal"`
	CastInLine                 bool        `json:"castInLine"`
	CastTestLos                bool        `json:"castTestLos"`
	ContainerIDs               []int       `json:"containerIds"`
	CraftXpRatio               int         `json:"craftXpRatio"`
	Criteria                   string      `json:"criteria"`
	CriteriaTarget             string      `json:"criteriaTarget"`
	CriticalFailureProbability int         `json:"criticalFailureProbability"`
	CriticalHitBonus           int         `json:"criticalHitBonus"`
	CriticalHitProbability     int         `json:"criticalHitProbability"`
	Cursed                     bool        `json:"cursed"`
	DescriptionID              I18N        `json:"descriptionId"`
	DropMonsterIDs             []int       `json:"dropMonstersIds"`
	Enhanceable                bool        `json:"enhanceable"`
	Etheral                    bool        `json:"etheral"`
	Exchangeable               bool        `json:"exchangeable"`
	FavoriteSubAreas           []int       `json:"favoriteSubAreas"`
	FavoriteSubAreasBonus      int         `json:"favoriteSubAreasBonus"`
	HideEffects                bool        `json:"hideEffects"`
	IconID                     int         `json:"iconId"`
	ID                         int         `json:"id"`
	IsDestructible             bool        `json:"isDestructible"`
	ItemSetID                  int         `json:"itemSetId"`
	Level                      int         `json:"level"`
	MaxCastPerTurn             int         `json:"maxCastPerTurn"`
	MaxRange                   int         `json:"range"`
	MinRange                   int         `json:"minRange"`
	NameID                     I18N        `json:"nameId"`
	NeedUseConfirm             bool        `json:"needUseConfirm"`
	NonUsableOnAnother         bool        `json:"nonUsableOnAnother"`
	NuggetsBySubArea           [][]float64 `json:"nuggetsBySubarea"`
	PossibleEffects            []Effect    `json:"possibleEffects"`
	Price                      float64     `json:"price"`
	RealWeight                 int         `json:"realWeight"`
	RecipeIDs                  []int       `json:"recipeIds"`
	RecipeSlots                int         `json:"recipeSlots"`
	ResourcesBySubArea         [][]int     `json:"resourcesBySubarea"`
	SecretRecipe               bool        `json:"secretRecipe"`
	Targetable                 bool        `json:"targetable"`
	TwoHanded                  bool        `json:"twoHanded"`
	TypeID                     int         `json:"typeId"`
	Usable                     bool        `json:"usable"`
	UseAnimationID             int         `json:"useAnimationId"`
}
