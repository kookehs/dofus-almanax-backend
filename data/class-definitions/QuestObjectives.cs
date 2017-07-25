public class QuestObjectiveFreeForm {
	public Int stepId
	public Int typeId
	public Int mapId
	public Int id
	public Int dialogId
	public QuestObjectiveParameters parameters
	public Point coords
}

public class QuestObjectiveParameters {
	public UInt numParams
	public Int parameter0
	public Int parameter1
	public Int parameter2
	public Int parameter3
	public Int parameter4
	public Bool dungeonOnly
}

public class QuestObjectiveBringSoulToNpc {
	public Int stepId
	public Int typeId
	public Int mapId
	public Int id
	public Int dialogId
	public QuestObjectiveParameters parameters
	public Point coords
}

public class QuestObjectiveFightMonstersOnMap {
	public Int stepId
	public Int typeId
	public Int mapId
	public Int id
	public Int dialogId
	public QuestObjectiveParameters parameters
	public Point coords
}

public class QuestObjective {
	public Int id
	public Int stepId
	public Int typeId
	public Int dialogId
	public QuestObjectiveParameters parameters
	public Point coords
	public Int mapId
}

public class QuestObjectiveDiscoverMap {
	public Int stepId
	public Int typeId
	public Int mapId
	public Int id
	public Int dialogId
	public QuestObjectiveParameters parameters
	public Point coords
}

public class QuestObjectiveCraftItem {
	public Int stepId
	public Int typeId
	public Int mapId
	public Int id
	public Int dialogId
	public QuestObjectiveParameters parameters
	public Point coords
}

public class Point {
	public Int x
	public Int y
}

public class QuestObjectiveGoToNpc {
	public Int stepId
	public Int typeId
	public Int mapId
	public Int id
	public Int dialogId
	public QuestObjectiveParameters parameters
	public Point coords
}

public class QuestObjectiveMultiFightMonster {
	public Int stepId
	public Int typeId
	public Int mapId
	public Int id
	public Int dialogId
	public QuestObjectiveParameters parameters
	public Point coords
}

public class QuestObjectiveDiscoverSubArea {
	public Int stepId
	public Int typeId
	public Int mapId
	public Int id
	public Int dialogId
	public QuestObjectiveParameters parameters
	public Point coords
}

public class QuestObjectiveFightMonster {
	public Int stepId
	public Int typeId
	public Int mapId
	public Int id
	public Int dialogId
	public QuestObjectiveParameters parameters
	public Point coords
}

public class QuestObjectiveDuelSpecificPlayer {
	public Int stepId
	public Int typeId
	public Int mapId
	public Int id
	public Int dialogId
	public QuestObjectiveParameters parameters
	public Point coords
}

public class QuestObjectiveBringItemToNpc {
	public Int stepId
	public Int typeId
	public Int mapId
	public Int id
	public Int dialogId
	public QuestObjectiveParameters parameters
	public Point coords
}

