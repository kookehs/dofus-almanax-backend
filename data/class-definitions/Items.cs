public class Item {
	public Int id
	public I18N nameId
	public Int typeId
	public I18N descriptionId
	public Int iconId
	public Int level
	public Int realWeight
	public Bool cursed
	public Int useAnimationId
	public Bool usable
	public Bool targetable
	public Bool exchangeable
	public Double price
	public Bool twoHanded
	public Bool etheral
	public Int itemSetId
	public String criteria
	public String criteriaTarget
	public Bool hideEffects
	public Bool enhanceable
	public Bool nonUsableOnAnother
	public Int appearanceId
	public Bool secretRecipe
	public Int recipeSlots
	public vector<UInt> recipeIds
	public vector<UInt> dropMonsterIds
	public Bool bonusIsSecret
	public vector<EffectInstance> possibleEffects
	public vector<UInt> favoriteSubAreas
	public Int favoriteSubAreasBonus
	public Int craftXpRatio
	public Bool needUseConfirm
	public Bool isDestructible
	public vector<vector<Double>> nuggetsBySubarea
	public vector<UInt> containerIds
	public vector<vector<Int>> resourcesBySubarea
}

public class EffectInstanceInteger {
	public String targetMask
	public Bool visibleInBuffUi
	public Bool visibleInFightLog
	public Int targetId
	public Int effectElement
	public Int effectUid
	public String triggers
	public Int duration
	public Int random
	public Int effectId
	public Int delay
	public Bool visibleInTooltip
	public String rawZone
	public Int value
	public Int group
}

public class EffectInstanceDice {
	public String targetMask
	public Int diceNum
	public Bool visibleInBuffUi
	public Bool visibleInFightLog
	public Int targetId
	public Int effectElement
	public Int effectUid
	public String triggers
	public Int duration
	public Int random
	public Int effectId
	public Int delay
	public Int diceSide
	public Bool visibleInTooltip
	public String rawZone
	public Int value
	public Int group
}

public class EffectInstance {
	public Int effectUid
	public Int effectId
	public Int targetId
	public String targetMask
	public Int duration
	public Int random
	public Int group
	public Bool visibleInTooltip
	public Bool visibleInBuffUi
	public Bool visibleInFightLog
	public String rawZone
	public Int delay
	public String triggers
	public Int effectElement
}

public class Weapon {
	public Bool exchangeable
	public Int criticalHitProbability
	public Bool etheral
	public Bool cursed
	public Int appearanceId
	public String criteria
	public Int criticalFailureProbability
	public Bool nonUsableOnAnother
	public Int range
	public Bool castInLine
	public Bool enhanceable
	public vector<EffectInstance> possibleEffects
	public Int criticalHitBonus
	public Int apCost
	public Bool castInDiagonal
	public Bool usable
	public vector<UInt> containerIds
	public I18N descriptionId
	public Bool twoHanded
	public vector<UInt> recipeIds
	public Double price
	public vector<vector<Double>> nuggetsBySubarea
	public String criteriaTarget
	public vector<UInt> favoriteSubAreas
	public Int id
	public Int iconId
	public Bool secretRecipe
	public Bool hideEffects
	public Bool castTestLos
	public Int level
	public Bool needUseConfirm
	public vector<vector<Int>> resourcesBySubarea
	public Int realWeight
	public Int favoriteSubAreasBonus
	public Int craftXpRatio
	public Bool targetable
	public Bool isDestructible
	public Int useAnimationId
	public I18N nameId
	public Int typeId
	public Int itemSetId
	public vector<UInt> dropMonsterIds
	public Int recipeSlots
	public Bool bonusIsSecret
	public Int maxCastPerTurn
	public Int minRange
}

