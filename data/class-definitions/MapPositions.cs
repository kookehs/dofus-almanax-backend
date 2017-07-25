public class AmbientSound {
	public Int id
	public Int volume
	public Int criterionId
	public Int silenceMin
	public Int silenceMax
	public Int channel
	public Int type_id
}

public class MapPosition {
	public Int id
	public Int posX
	public Int posY
	public Bool outdoor
	public Int capabilities
	public I18N nameId
	public Bool showNameOnFingerpost
	public vector<AmbientSound> sounds
	public vector<vector<Int>> playlists
	public Int subAreaId
	public Int worldMap
	public Bool hasPriorityOnWorldmap
	public Bool isUnderWater
}

