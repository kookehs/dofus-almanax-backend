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

public class SpellLevel {
	public Int id
	public Int spellId
	public Int grade
	public Int spellBreed
	public Int apCost
	public Int minRange
	public Int range
	public Bool castInLine
	public Bool castInDiagonal
	public Bool castTestLos
	public Int criticalHitProbability
	public Bool needFreeCell
	public Bool needTakenCell
	public Bool needFreeTrapCell
	public Bool rangeCanBeBoosted
	public Int maxStack
	public Int maxCastPerTurn
	public Int maxCastPerTarget
	public Int minCastInterval
	public Int initialCooldown
	public Int globalCooldown
	public Int minPlayerLevel
	public Bool hideEffects
	public Bool hidden
	public Bool playAnimation
	public vector<Int> statesRequired
	public vector<Int> statesForbidden
	public vector<EffectInstanceDice> effects
	public vector<EffectInstanceDice> criticalEffect
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

