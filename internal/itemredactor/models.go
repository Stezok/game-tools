package itemredactor

type Item struct {
	ID                              string `item:"//ID" json:"ID"`
	Name                            string `item:"Name" json:"Name"`
	IconName                        string `item:"Icon Name" json:"IconName"`
	ModelGround                     string `item:"Model (Ground)" json:"ModelGround"`
	ModelLance                      string `item:"Model (Lance)" json:"ModelLance"`
	ModelCarsise                    string `item:"Model (Carsise)" json:"ModelCarsise"`
	ModelPhyllis                    string `item:"Model (Phyllis)" json:"ModelPhyllis"`
	ModelAmi                        string `item:"Model (Ami)" json:"ModelAmi"`
	ShipSymbol                      string `item:"Ship Symbol" json:"ShipSymbol"`
	ShipSize                        string `item:"Ship Size" json:"ShipSize"`
	ItemType                        string `item:"Item Type" json:"ItemType"`
	ObtainPrefixRate                string `item:"Obtain Prefix Rate" json:"ObtainPrefixRate"`
	SetID                           string `item:"Set ID" json:"SetID"`
	ForgingLevel                    string `item:"Forging Level" json:"ForgingLevel"`
	StableLevel                     string `item:"Stable Level" json:"StableLevel"`
	Repairable                      string `item:"Repairable" json:"Repairable"`
	Tradable                        string `item:"Tradable" json:"Tradable"`
	PickUpable                      string `item:"Pick-upable" json:"PickUpable"`
	Droppable                       string `item:"Droppable" json:"Droppable"`
	Deletable                       string `item:"Deletable" json:"Deletable"`
	MaxStackSize                    string `item:"Max Stack Size" json:"MaxStackSize"`
	Instance                        string `item:"Instance" json:"Instance"`
	Price                           string `item:"Price" json:"Price"`
	CharacterTypes                  string `item:"Character Types" json:"CharacterTypes"`
	CharacterLevel                  string `item:"Character Level" json:"CharacterLevel"`
	CharacterClasses                string `item:"Character Classes" json:"CharacterClasses"`
	CharacterNick                   string `item:"Character Nick" json:"CharacterNick"`
	CharacterReputation             string `item:"Character Reputation" json:"CharacterReputation"`
	EquipableSlots                  string `item:"Equipable Slots" json:"EquipableSlots"`
	ItemSwitchLocations             string `item:"Item Switch Locations" json:"ItemSwitchLocations"`
	ItemObtainIntoLocationDetermine string `item:"Item Obtain Into Location Determine" json:"ItemObtainIntoLocationDetermine"`
	STR                             string `item:"+STR %" json:"STR"`
	AGI                             string `item:"+AGI %" json:"AGI"`
	ACC                             string `item:"+ACC %" json:"ACC"`
	CON                             string `item:"+CON %" json:"CON"`
	SPR                             string `item:"+SPR %" json:"SPR"`
	LUK                             string `item:"+LUK %" json:"LUK"`
	AttackSpeed                     string `item:"+Attack Speed %" json:"AttackSpeed"`
	AttackRange                     string `item:"+Attack Range %" json:"AttackRange"`
	MinAttack                       string `item:"+Min Attack %" json:"MinAttack"`
	MaxAttack                       string `item:"+Max Attack %" json:"MaxAttack"`
	Defense                         string `item:"+Defense %" json:"Defense"`
	MaxHP                           string `item:"+Max HP %" json:"MaxHP"`
	MaxSP                           string `item:"+Max SP %" json:"MaxSP"`
	DodgeRate                       string `item:"+Dodge Rate %" json:"DodgeRate"`
	HitRate                         string `item:"+Hit Rate %" json:"HitRate"`
	CriticalRate                    string `item:"+Critical Rate %" json:"CriticalRate"`
	TreasureDropRate                string `item:"+Treasure Drop Rate %" json:"TreasureDropRate"`
	HPRecovery                      string `item:"+HP Recovery %" json:"HPRecovery"`
	SPRecovery                      string `item:"+SP Recovery %" json:"SPRecovery"`
	MovementSpeed                   string `item:"+Movement Speed %" json:"MovementSpeed"`
	ResourceGatheringRate           string `item:"+Resource Gathering Rate %" json:"ResourceGatheringRate"`
	STRMinMax                       string `item:"+STR (Min,Max)" json:"STRMinMax"`
	AGIMinMax                       string `item:"+AGI (Min,Max)" json:"AGIMinMax"`
	ACCMinMax                       string `item:"+ACC (Min,Max)" json:"ACCMinMax"`
	CONMinMax                       string `item:"+CON (Min,Max)" json:"CONMinMax"`
	SPRMinMax                       string `item:"+SPR (Min,Max)" json:"SPRMinMax"`
	LUKMinMax                       string `item:"+LUK (Min,Max)" json:"LUKMinMax"`
	AttackSpeedAdd                  string `item:"+Attack Speed" json:"AttackSpeedAdd"`
	AttackRangeAdd                  string `item:"+Attack Range" json:"AttackRangeAdd"`
	MinAttackAdd                    string `item:"+Min Attack" json:"MinAttackAdd"`
	MaxAttackAdd                    string `item:"+Max Attack" json:"MaxAttackAdd"`
	DefenseAdd                      string `item:"+Defense" json:"DefenseAdd"`
	MaxHPAdd                        string `item:"+Max HP" json:"MaxHPAdd"`
	MaxSPAdd                        string `item:"+Max SP" json:"MaxSPAdd"`
	DodgeAdd                        string `item:"+Dodge" json:"DodgeAdd"`
	HitRateAdd                      string `item:"+Hit Rate" json:"HitRateAdd"`
	CriticalRateAdd                 string `item:"+Critical Rate" json:"CriticalRateAdd"`
	TreasureDropRateAdd             string `item:"+Treasure Drop Rate" json:"TreasureDropRateAdd"`
	HPRecoveryAdd                   string `item:"+HP Recovery" json:"HPRecoveryAdd"`
	SPRecoveryAdd                   string `item:"+SP Recovery" json:"SPRecoveryAdd"`
	MovementSpeedAdd                string `item:"+Movement Speed" json:"MovementSpeedAdd"`
	ResourceGatheringRateAdd        string `item:"+Resource Gathering Rate" json:"ResourceGatheringRateAdd"`
	PhysicalResistAdd               string `item:"+Physical Resist" json:"PhysicalResistAdd"`
	AftectedByLeftHandEfs           string `item:"Aftected By Left Hand Efs" json:"AftectedByLeftHandEfs"`
	Energy                          string `item:"+Energy" json:"Energy"`
	Durability                      string `item:"+Durability" json:"Durability"`
	GemSockets                      string `item:"Gem Sockets" json:"GemSockets"`
	ShipDurabilityRecovery          string `item:"Ship Durability Recovery" json:"ShipDurabilityRecovery"`
	ShipCannonAmount                string `item:"Ship Cannon Amount" json:"ShipCannonAmount"`
	ShipMemberCount                 string `item:"Ship Member Count" json:"ShipMemberCount"`
	ShipLabel                       string `item:"Ship Label" json:"ShipLabel"`
	ShipCargoCapacity               string `item:"Ship Cargo Capacity" json:"ShipCargoCapacity"`
	ShipFuelConsumption             string `item:"Ship Fuel Consumption" json:"ShipFuelConsumption"`
	ShipAttackSpeed                 string `item:"Ship Attack Speed" json:"ShipAttackSpeed"`
	ShipMovementSpeed               string `item:"Ship Movement Speed" json:"ShipMovementSpeed"`
	UsageScript                     string `item:"Usage Script" json:"UsageScript"`
	DisplayEffect                   string `item:"Display Effect" json:"DisplayEffect"`
	BindEffect                      string `item:"Bind Effect" json:"BindEffect"`
	BindEffect2                     string `item:"Bind Effect 2" json:"BindEffect2"`
	FirstInvSlotEffect              string `item:"1st Inv Slot Effect" json:"FirstInvSlotEffect"`
	DropModelEffect                 string `item:"Drop Model Effect" json:"DropModelEffect"`
	ItemUsageEffect                 string `item:"Item Usage Effect" json:"ItemUsageEffect"`
	Description                     string `item:"Description" json:"Description"`
	Last                            string `item:"0" json:"Last"`
}
