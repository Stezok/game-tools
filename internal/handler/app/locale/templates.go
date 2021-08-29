package locale

type IndexPlaceholdes struct {
	ItemSearchBar   string
	ItemID          string
	ItemImage       string
	ItemGroundModel string
	ItemName        string
	ItemDescription string
	Static          string
	Percent         string
}

type IndexLabels struct {
	Fixed            string
	ItemType         string
	ItemImage        string
	ItemGroundModel  string
	ItemName         string
	ItemDescription  string
	ItemScript       string
	ItemLevel        string
	ItemStack        string
	ItemPrice        string
	ItemSlots        string
	ItemEffect       string
	ItemUsageEffect  string
	ItemRepair       string
	ItemSell         string
	ItemTake         string
	ItemThrow        string
	ItemDelete       string
	ItemLance        string
	ItemModelLance   string
	ItemCarsise      string
	ItemModelCarsise string
	ItemPhyllis      string
	ItemModelPhyllis string
	ItemAmi          string
	ItemModelAmi     string
	ItemNewbie       string
	ItemSwordmaster  string
	ItemHunter       string
	ItemSailor       string
	ItemHealer       string
	ItemChampion     string
	ItemWarrior      string
	ItemArcher       string
	ItemPriest       string
	ItemSealmaster   string
	ItemVoyager      string
	ItemStrength     string
	ItemAgility      string
	ItemVitality     string
	ItemSpirit       string
	ItemAccuracy     string
	ItemLuck         string
	ItemMinAttack    string
	ItemMaxAttack    string
	ItemAttackSpeed  string
	ItemHitChance    string
	ItemCritChance   string
	ItemMoveSpeed    string
	ItemMaxHP        string
	ItemMaxMP        string
	ItemHPRegen      string
	ItemMPRegen      string
	ItemArmour       string
	ItemPhysArmour   string
	ItemEvasion      string
	ItemDurability   string
	ItemEnergy       string
}

type IndexItemType struct {
	Name string
	ID   int
}

type IndexGroups struct {
	Actions string
	Races   string
	Classes string
	Stats   string
}

type IndexTemplate struct {
	Language string

	Placeholders IndexPlaceholdes
	Labels       IndexLabels
	Groups       IndexGroups

	ItemTypes []IndexItemType
}
