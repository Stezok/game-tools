package locale

var EN_INDEX_LOCALE = IndexTemplate{
	Language: "English",

	Text: IndexText{
		ItemDeleted: "The item has been removed.",
		ItemReturn:  "Restore",
	},

	Placeholders: IndexPlaceholdes{
		ItemSearchBar:   "ID or item name...",
		ItemID:          "ID...",
		ItemImage:       "Path...",
		ItemGroundModel: "Model...",
		ItemName:        "Item name...",
		ItemDescription: "Description...",
		Static:          "Static...",
		Percent:         "Percent...",
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
		ItemWarrior:      "Warrior",
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
