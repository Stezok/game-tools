package locale

var EN_INDEX_LOCALE = IndexTemplate{
	Language: "English",

	Text: IndexText{
		ItemDeleted: "The item has been removed.",
		ItemReturn:  "Restore",
		CharDeleted: "The monster has been removed.",
		CharReturn:  "Restore",
	},

	Placeholders: IndexPlaceholdes{
		ItemSearchBar:   "ID or item name...",
		ItemID:          "ID...",
		ItemImage:       "Path...",
		ItemGroundModel: "Model...",
		ItemName:        "Item name...",
		ItemDescription: "Description...",

		CharSearchBar: "ID or monster name...",

		Static:  "Static...",
		Percent: "Percent...",
	},

	Groups: IndexGroups{
		Actions: "Actions",
		Races:   "Races",
		Classes: "Classes",
		Stats:   "Stats",
	},

	Labels: IndexLabels{
		Fixed:            "Fixed",
		ItemType:         "Item type",
		ItemImage:        "Icon",
		ItemGroundModel:  "Model on ground",
		ItemName:         "Item name",
		ItemDescription:  "Description",
		ItemScript:       "Script",
		ItemLevel:        "Level",
		ItemStack:        "Stack",
		ItemPrice:        "Price",
		ItemSlots:        "Slots",
		ItemEffect:       "Display effect",
		ItemUsageEffect:  "Use effect",
		ItemRepair:       "Repairable",
		ItemSell:         "Sellable",
		ItemTake:         "Pickupable",
		ItemThrow:        "Droppable",
		ItemDelete:       "Deletable",
		ItemLance:        "Lance",
		ItemModelLance:   "Model Lance",
		ItemCarsise:      "Carsise",
		ItemModelCarsise: "Model Carsise",
		ItemPhyllis:      "Phyllis",
		ItemModelPhyllis: "Model Phyllis",
		ItemAmi:          "Ami",
		ItemModelAmi:     "Model Ami",
		ItemNewbie:       "Newbie",
		ItemSwordmaster:  "Swordsman",
		ItemHunter:       "Hunter",
		ItemSailor:       "Mariner",
		ItemHealer:       "Herbalist",
		ItemChampion:     "Champion",
		ItemWarrior:      "Crusader",
		ItemArcher:       "Sharpshooter",
		ItemPriest:       "Cleric",
		ItemSealmaster:   "Seal master",
		ItemVoyager:      "Voyager",
		ItemStrength:     "Strength",
		ItemAgility:      "Agility",
		ItemVitality:     "Vitality",
		ItemSpirit:       "Spirit",
		ItemAccuracy:     "Accuracy",
		ItemLuck:         "Luck",
		ItemMinAttack:    "Minimum attack",
		ItemMaxAttack:    "Maximum attack",
		ItemAttackSpeed:  "Attack speed",
		ItemHitChance:    "Hit chance",
		ItemCritChance:   "Critical hit chance",
		ItemMoveSpeed:    "Move speed",
		ItemMaxHP:        "Max HP",
		ItemMaxMP:        "Max MP",
		ItemHPRegen:      "HP Regen",
		ItemMPRegen:      "MP Regen",
		ItemArmour:       "Armour",
		ItemPhysArmour:   "Physical Armour",
		ItemEvasion:      "Evasion",
		ItemDurability:   "Durability",
		ItemEnergy:       "Energy",

		CharID:                        "ID",
		CharDescription:               "Description",
		CharModelCategory:             "Model Category",
		CharModelType:                 "Model Type",
		CharLogicType:                 "Logic Type",
		CharFrameworkNumber:           "Framework Number",
		CharSuiteSerial:               "Suite Serial",
		CharSuiteQuantity:             "Suite Quantity",
		CharPartFACE:                  "Part Face",
		CharPartHeadApparel:           "Part Head Apparel",
		CharPartArmor:                 "Part Armor",
		CharPartGloves:                "Part Gloves",
		CharPartShoes:                 "Part Shoes",
		CharPartWeapon:                "Part Weapon",
		CharPart2ndWeap:               "Part 2en Weapon",
		CharPartWings:                 "Part Wings",
		CharFeffID:                    "Feff ID",
		CharEeffID:                    "Eeff ID",
		CharSpecialEffectActionSerial: "Special Effect",
		CharShadow:                    "Shadow",
		CharActionID:                  "Action ID",
		CharTransparency:              "Transparency",
		CharMovingSoundEffect:         "Moving Sound Effect",
		CharBreathingSoundEffect:      "Breathing Sound Effect",
		CharDeathSoundEffect:          "Death Sound Effect",
		CharCanItBeControlled:         "Can It Be Controlled",
		CharAreaLimitation:            "Area Limitation",
		CharAltitudeExcursion:         "Altitude Excursion",
		CharItemsThatCanBeEquipped:    "Items That Can Be Equipped",
		CharLength:                    "Length",
		CharWidth:                     "Width",
		CharHeight:                    "Height",
		CharCollisionRange:            "Collision Range",
		CharBirth:                     "Birth",
		CharDead:                      "Dead",
		CharBirthEffect:               "Birth Effect",
		CharDeathEffect:               "Death Effect",
		CharHibernateAction:           "Hibernate Action",
		CharDeathInstantAction:        "Death Instant Action",
		CharRemainingHPEffectDisplay:  "Remaining HP Effect Display",
		CharAttackCanBeSwerve:         "Attack Can Be Swerve",
		CharConfirmToUseBlownAway:     "Confirm To Use Blown Away",
		CharScript:                    "Script",
		CharWeaponUsed:                "Weapon Used",
		CharSkillID:                   "Skill ID",
		CharMonsterSkillRate:          "Monster Skill Rate",
		CharItemIDToDrop:              "Item ID To Drop",
		CharItemDroprate:              "Item Droprate",
		CharQuantityLimit:             "Quantity Limit",
		CharFatalityRate:              "Fatality Rate",
		CharPrefixLevel:               "Prefix Level",
		CharQuestDropID:               "Quest Drop ID",
		CharRate:                      "Rate",
		CharAI:                        "AI",
		//
		CharTurning:                 "Turning",
		CharVision:                  "Vision",
		CharNoise:                   "Noise",
		CharGetExp:                  "Get Exp",
		CharLight:                   "Light",
		CharMobExp:                  "Mob Exp",
		CharMonsterLVL:              "Monster LVL",
		CharMaxHP:                   "Max HP",
		CharMinHP:                   "Min HP",
		CharMaxSP:                   "Max SP",
		CharMinSp:                   "Min SP",
		CharMinAtk:                  "Min Atk",
		CharMaxAtk:                  "Max Atk",
		CharPhysicalResist:          "Physical Resist",
		CharDefense:                 "Defense",
		CharHitRate:                 "Hit Rate",
		CharFlee:                    "Flee",
		CharMonsterCriticalRate:     "Monster Critical Rate",
		CharMF:                      "MF",
		CharHREC:                    "HP Recovery",
		CharSREC:                    "SP Recovery",
		CharASPDofMonster:           "ASPD of Monster",
		CharAreaOfDetection:         "Area Of Detection",
		CharAfterDetectionChaseArea: "After Detection Chase Area",
		CharMSPD:                    "MSPD",
		//
		CharCol:               "Col",
		CharStr:               "Str",
		CharAgi:               "Agi",
		CharDex:               "Dex",
		CharCon:               "Con",
		CharSta:               "Spr",
		CharLuk:               "Luck",
		CharLeftRad:           "Left Rad",
		CharGuild:             "Guild",
		CharTitle:             "Title",
		CharJob:               "Job",
		CharExpGainedFromKill: "Exp Gained From Kill",
		CharNexp:              "Nexp",
		CharFame:              "Fame",
		CharAp:                "Ap",
		CharTp:                "Tp",
		CharGd:                "Gd",
		CharSpri:              "Spri",
		CharStor:              "Stor",
		CharMxSail:            "Mx Sail",
		CharSail:              "Sail",
		CharStasa:             "Stasa",
		CharScsm:              "Scsm",
		CharTstr:              "Tstr",
		CharTagi:              "Tagi",
		CharTdex:              "Tdex",
		CharTcon:              "Tcon",
		CharTsta:              "Tsta",
		CharTluk:              "Tluk",
		CharTmxhp:             "Tmxhp",
		CharTmxsp:             "Tmxsp",
		CharTatk:              "Tatk",
		CharTdef:              "Tdef",
		CharThit:              "Thit",
		CharTflee:             "Tflee",
		CharTmf:               "Tmf",
		CharTcrt:              "Tcrt",
		CharThrec:             "Threc",
		CharTsrec:             "Tsrec",
		CharTaspd:             "Taspd",
		CharTadis:             "Tadis",
		CharTspd:              "Tspd",
		CharTspri:             "Tspri",
		CharTscsm:             "Tscsm",
	},

	ItemTypes: []IndexItemType{
		{
			ID:   1,
			Name: "One-handed sword",
		}, {
			ID:   2,
			Name: "Two-handed sword",
		}, {
			ID:   3,
			Name: "Bow",
		}, {
			ID:   4,
			Name: "Musket",
		}, {
			ID:   7,
			Name: "Blade",
		}, {
			ID:   9,
			Name: "Stick",
		}, {
			ID:   11,
			Name: "Shield",
		}, {
			ID:   18,
			Name: "Axe (felling)",
		}, {
			ID:   19,
			Name: "Pickaxe (mining)",
		}, {
			ID:   20,
			Name: "Ami's hat",
		}, {
			ID:   21,
			Name: "Face",
		}, {
			ID:   22,
			Name: "Armour",
		}, {
			ID:   23,
			Name: "Gloves",
		}, {
			ID:   24,
			Name: "Boots",
		}, {
			ID:   25,
			Name: "Necklace",
		}, {
			ID:   26,
			Name: "Ring",
		}, {
			ID:   27,
			Name: "Tattoo",
		}, {
			ID:   28,
			Name: "Hairstyle",
		}, {
			ID:   29,
			Name: "Coral",
		}, {
			ID:   34,
			Name: "Skill book",
		}, {
			ID:   36,
			Name: "Ticket",
		}, {
			ID:   41,
			Name: "Chest",
		}, {
			ID:   56,
			Name: "Fruit for fairies",
		}, {
			ID:   57,
			Name: "Fairy food",
		}, {
			ID:   59,
			Name: "Fairies",
		},
	},
}
