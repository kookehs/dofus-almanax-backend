public class BreedRoleByBreed {
	public Int breedId
	public Int roleId
	public I18N descriptionId
	public Int value
	public Int order
}

public class Breed {
	public Int id
	public I18N shortNameId
	public I18N longNameId
	public I18N descriptionId
	public I18N gameplayDescriptionId
	public String maleLook
	public String femaleLook
	public Int creatureBonesId
	public Int maleArtwork
	public Int femaleArtwork
	public vector<vector<UInt>> statsPointsForStrength
	public vector<vector<UInt>> statsPointsForIntelligence
	public vector<vector<UInt>> statsPointsForChance
	public vector<vector<UInt>> statsPointsForAgility
	public vector<vector<UInt>> statsPointsForVitality
	public vector<vector<UInt>> statsPointsForWisdom
	public vector<UInt> breedSpellsId
	public vector<BreedRoleByBreed> breedRoles
	public vector<UInt> maleColors
	public vector<UInt> femaleColors
	public Int spawnMap
	public Int complexity
	public Int sortIndex
}

