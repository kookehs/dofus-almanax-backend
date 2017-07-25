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

public class ItemSet {
	public Int id
	public vector<UInt> items
	public I18N nameId
	public Bool bonusIsSecret
	public vector<vector<EffectInstance>> effects
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

