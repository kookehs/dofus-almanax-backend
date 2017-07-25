public class SoundBones {
	public Int id
	public vector<String> keys
	public vector<vector<SoundAnimation>> values
}

public class SoundAnimation {
	public Int id
	public String label
	public String name
	public String filename
	public Int volume
	public Int rolloff
	public Int automationDuration
	public Int automationVolume
	public Int automationFadeIn
	public Int automationFadeOut
	public Bool noCutSilence
	public Int startFrame
}

