public class Npc {
	public Int id
	public I18N nameId
	public vector<vector<Int>> dialogMessages
	public vector<vector<Int>> dialogReplies
	public vector<UInt> actions
	public Int gender
	public String look
	public vector<AnimFunNpcData> animFunList
	public Bool fastAnimsFun
}

public class AnimFunNpcData {
	public String animName
	public Int animWeight
}

