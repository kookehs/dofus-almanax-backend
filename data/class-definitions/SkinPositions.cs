public class TransformData {
	public Double x
	public Double y
	public Double scaleX
	public Double scaleY
	public Int rotation
	public String originalClip
	public String overrideClip
}

public class SkinPosition {
	public Int id
	public vector<TransformData> transformation
	public vector<String> clip
	public vector<UInt> skin
}

