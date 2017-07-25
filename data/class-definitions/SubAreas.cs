public class AmbientSound {
	public String id
	public Int volume
	public Int criterionId
	public Int silenceMin
	public Int silenceMax
	public Int channel
	public Int type_id
}

public class Rectangle {
	public Int x
	public Int y
	public Int width
	public Int height
}

public class SubArea {
	public Int id
	public I18N nameId
	public Int areaId
	public vector<AmbientSound> ambientSounds
	public vector<vector<Int>> playlists
	public vector<UInt> mapIds
	public Rectangle bounds
	public vector<Int> shape
	public vector<UInt> customWorldMap
	public UInt packId
	public UInt level
	public Bool isConquestVillage
	public Bool basicAccountAllowed
	public Bool displayOnWorldMap
	public vector<UInt> monsters
	public vector<UInt> entranceMapIds
	public vector<UInt> exitMapIds
	public Bool capturable
	public vector<UInt> achievements
	public vector<vector<Int>> quests
	public vector<vector<Int>> npcs
	public Int exploreAchievementId
	public vector<Int> harvestables
}

