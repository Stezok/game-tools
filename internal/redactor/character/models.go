package character

type Character struct {
	ID                        string `character:"//ID" json:"ID"`
	Description               string `character:"Description" json:"Description"`
	ModelCategory             string `character:"Model Category" json:"ModelCategory"`
	ModelType                 string `character:"Model Type" json:"ModelType"`
	LogicType                 string `character:"Logic Type" json:"LogicType"`
	FrameworkNumber           string `character:"Framework Number" json:"FrameworkNumber"`
	SuiteSerial               string `character:"Suite Serial" json:"SuiteSerial"`
	SuiteQuantity             string `character:"Suite Quantity" json:"SuiteQuantity"`
	PartFACE                  string `character:"Part 00 --FACE" json:"PartFACE"`
	PartHeadApparel           string `character:"Part 01-- Head Apparel" json:"PartHeadApparel"`
	PartArmor                 string `character:"Part 02--Armor" json:"PartArmor"`
	PartGloves                string `character:"Part 03-- Gloves" json:"PartGloves"`
	PartShoes                 string `character:"Part 04-- Shoes" json:"PartShoes"`
	PartWeapon                string `character:"Part 05-- Weapon" json:"PartWeapon"`
	Part2ndWeap               string `character:"Part 06-- 2nd Weap (Cruz Only)" json:"Part2ndWeap"`
	PartWings                 string `character:"Part 07-- (Possibly Wings?)" json:"PartWings"`
	FeffID                    string `character:"FeffID" json:"FeffID"`
	EeffID                    string `character:"EeffID" json:"EeffID"`
	SpecialEffectActionSerial string `character:"Special Effect Action Serial" json:"SpecialEffectActionSerial"`
	Shadow                    string `character:"Shadow" json:"Shadow"`
	ActionID                  string `character:"Action ID" json:"ActionID"`
	Transparency              string `character:"Transparency" json:"Transparency"`
	MovingSoundEffect         string `character:"Moving Sound Effect" json:"MovingSoundEffect"`
	BreathingSoundEffect      string `character:"Breathing Sound Effect" json:"BreathingSoundEffect"`
	DeathSoundEffect          string `character:"Death Sound Effect" json:"DeathSoundEffect"`
	CanItBeControlled         string `character:"Can it be controlled" json:"CanItBeControlled"`
	AreaLimitation            string `character:"Area Limitation" json:"AreaLimitation"`
	AltitudeExcursion         string `character:"Altitude Excursion" json:"AltitudeExcursion"`
	ItemsThatCanBeEquipped    string `character:"Items that can be equipped" json:"ItemsThatCanBeEquipped"`
	Length                    string `character:"length" json:"Length"`
	Width                     string `character:"width" json:"Width"`
	Height                    string `character:"height" json:"Height"`
	CollisionRange            string `character:"Collision Range" json:"CollisionRange"`
	Birth                     string `character:"Birth" json:"Birth"`
	Dead                      string `character:"Dead" json:"Dead"`
	BirthEffect               string `character:"Birth Effect" json:"BirthEffect"`
	DeathEffect               string `character:"Death Effect" json:"DeathEffect"`
	HibernateAction           string `character:"Hibernate Action" json:"HibernateAction"`
	DeathInstantAction        string `character:"Death Instant Action" json:"DeathInstantAction"`
	RemainingHPEffectDisplay  string `character:"Remaining HP Effect Display" json:"RemainingHPEffectDisplay"`
	AttackCanBeSwerve         string `character:"Attack can be swerve" json:"AttackCanBeSwerve"`
	ConfirmToUseBlownAway     string `character:"Confirm to use blown away" json:"ConfirmToUseBlownAway"`
	Script                    string `character:"script" json:"Script"`
	WeaponUsed                string `character:"Weapon used" json:"WeaponUsed"`
	SkillID                   string `character:"Skill ID" json:"SkillID"`
	MonsterSkillRate          string `character:"Monster Skill Rate" json:"MonsterSkillRate"`
	ItemIDToDrop              string `character:"Item ID to DROP" json:"ItemIDToDrop"`
	ItemDroprate              string `character:"Item Droprate" json:"ItemDroprate"`
	QuantityLimit             string `character:"Quantity Limit" json:"QuantityLimit"`
	FatalityRate              string `character:"Fatality Rate" json:"FatalityRate"`
	PrefixLevel               string `character:"Prefix Level" json:"PrefixLevel"`
	QuestDropID               string `character:"Quest Drop ID" json:"QuestDropID"`
	Rate                      string `character:"Rate" json:"Rate"`
	AI                        string `character:"ai" json:"AI"`
	Turning                   string `character:"turning" json:"Turning"`
	Vision                    string `character:"vision" json:"Vision"`
	Noise                     string `character:"noise" json:"Noise"`
	GetExp                    string `character:"getexp" json:"GetExp"`
	Light                     string `character:"light" json:"Light"`
	MobExp                    string `character:"mobexp" json:"MobExp"`
	MonsterLVL                string `character:"Monster LVL" json:"MonsterLVL"`
	MaxHP                     string `character:"Max HP" json:"MaxHP"`
	MinHP                     string `character:"Min HP (Hp at which monster dies)" json:"MinHP"`
	MaxSP                     string `character:"Max SP" json:"MaxSP"`
	MinSp                     string `character:"Mn Sp (SP at which monster cant skill)" json:"MinSp"`
	MinAtk                    string `character:"Min atk (Physical)" json:"MinAtk"`
	MaxAtk                    string `character:"Max atk (Physical)" json:"MaxAtk"`
	PhysicalResist            string `character:"Physical Resist (PR)" json:"PhysicalResist"`
	Defense                   string `character:"Defense" json:"Defense"`
	HitRate                   string `character:"Hit Rate (?)" json:"HitRate"`
	Flee                      string `character:"flee (Distance at which monster lose target)" json:"Flee"`
	MonsterCriticalRate       string `character:"Monster Critical Rate  ??" json:"MonsterCriticalRate"`
	MF                        string `character:"MF (DropRate ??)" json:"MF"`
	HREC                      string `character:"H-REC" json:"HREC"`
	SREC                      string `character:"S-REC" json:"SREC"`
	ASPDofMonster             string `character:"ASPD of Monster" json:"ASPDofMonster"`
	AreaOfDetection           string `character:"Area of Detection (adis)" json:"AreaOfDetection"`
	AfterDetectionChaseArea   string `character:"After Detection Chase Area(chase_adis)" json:"AfterDetectionChaseArea"`
	MSPD                      string `character:"M-SPD (monster movement speed)" json:"MSPD"`
	Col                       string `character:"col (?)" json:"Col"`
	Str                       string `character:"str (Stat)" json:"Str"`
	Agi                       string `character:"agi (Stat)" json:"Agi"`
	Dex                       string `character:"dex (Stat ACC)" json:"Dex"`
	Con                       string `character:"con (Stat)" json:"Con"`
	Sta                       string `character:"sta (Stat SPR)" json:"Sta"`
	Luk                       string `character:"luk (Critical Rate chaiten aura gem)" json:"Luk"`
	LeftRad                   string `character:"left_rad" json:"LeftRad"`
	Guild                     string `character:"guild" json:"Guild"`
	Title                     string `character:"title" json:"Title"`
	Job                       string `character:"job" json:"Job"`
	ExpGainedFromKill         string `character:"exp gained from kill" json:"ExpGainedFromKill"`
	Nexp                      string `character:"nexp" json:"Nexp"`
	Fame                      string `character:"fame" json:"Fame"`
	Ap                        string `character:"ap" json:"Ap"`
	Tp                        string `character:"tp" json:"Tp"`
	Gd                        string `character:"gd" json:"Gd"`
	Spri                      string `character:"spri" json:"Spri"`
	Stor                      string `character:"stor" json:"Stor"`
	MxSail                    string `character:"mxsail" json:"MxSail"`
	Sail                      string `character:"sail" json:"Sail"`
	Stasa                     string `character:"stasa" json:"Stasa"`
	Scsm                      string `character:"scsm" json:"Scsm"`
	Tstr                      string `character:"tstr" json:"Tstr"`
	Tagi                      string `character:"tagi" json:"Tagi"`
	Tdex                      string `character:"tdex" json:"Tdex"`
	Tcon                      string `character:"tcon" json:"Tcon"`
	Tsta                      string `character:"tsta" json:"Tsta"`
	Tluk                      string `character:"tluk" json:"Tluk"`
	Tmxhp                     string `character:"tmxhp" json:"Tmxhp"`
	Tmxsp                     string `character:"tmxsp" json:"Tmxsp"`
	Tatk                      string `character:"tatk" json:"Tatk"`
	Tdef                      string `character:"tdef" json:"Tdef"`
	Thit                      string `character:"thit" json:"Thit"`
	Tflee                     string `character:"tflee" json:"Tflee"`
	Tmf                       string `character:"tmf" json:"Tmf"`
	Tcrt                      string `character:"tcrt" json:"Tcrt"`
	Threc                     string `character:"threc" json:"Threc"`
	Tsrec                     string `character:"tsrec" json:"Tsrec"`
	Taspd                     string `character:"taspd" json:"Taspd"`
	Tadis                     string `character:"tadis" json:"Tadis"`
	Tspd                      string `character:"tspd" json:"Tspd"`
	Tspri                     string `character:"tspri" json:"Tspri"`
	Tscsm                     string `character:"tscsm" json:"Tscsm"`
}
