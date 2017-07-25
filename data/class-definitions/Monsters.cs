public class MonsterGrade {
	public Int grade
	public Int monsterId
	public Int level
	public Int lifePoints
	public Int actionPoints
	public Int movementPoints
	public Int vitality
	public Int paDodge
	public Int pmDodge
	public Int wisdom
	public Int earthResistance
	public Int airResistance
	public Int fireResistance
	public Int waterResistance
	public Int neutralResistance
	public Int gradeXp
	public Int damageReflect
	public UInt hiddenLevel
}

public class Monster {
	public Int id
	public I18N nameId
	public Int gfxId
	public Int race
	public vector<MonsterGrade> grades
	public String look
	public Bool useSummonSlot
	public Bool useBombSlot
	public Bool canPlay
	public vector<AnimFunMonsterData> animFunList
	public Bool canTackle
	public Bool isBoss
	public vector<MonsterDrop> drops
	public vector<UInt> subareas
	public vector<UInt> spells
	public Int favoriteSubareaId
	public Bool isMiniBoss
	public Bool isQuestMonster
	public Int correspondingMiniBossId
	public Int speedAdjust
	public Int creatureBoneId
	public Bool canBePushed
	public Bool fastAnimsFun
	public Bool canSwitchPos
	public vector<UInt> incompatibleIdols
	public Bool allIdolsDisabled
	public Bool dareAvailable
	public vector<UInt> incompatibleChallenges
}

public class MonsterDrop {
	public Int dropId
	public Int monsterId
	public Int objectId
	public Double percentDropForGrade1
	public Double percentDropForGrade2
	public Double percentDropForGrade3
	public Double percentDropForGrade4
	public Double percentDropForGrade5
	public Int count
	public Bool hasCriteria
}

public class AnimFunMonsterData {
	public String animName
	public Int animWeight
}

