package locale

var RU_INDEX_LOCALE = IndexTemplate{
	Language: "Русский",

	Placeholders: IndexPlaceholdes{
		ItemSearchBar:   "ID или название предмета...",
		ItemID:          "ID...",
		ItemImage:       "Путь...",
		ItemGroundModel: "Модель...",
		ItemName:        "Название предмета...",
		ItemDescription: "Описание...",
		Static:          "Статик...",
		Percent:         "Процент...",
	},

	Groups: IndexGroups{
		Actions: "Действия с предметом",
		Races:   "Расы",
		Classes: "Классы",
		Stats:   "Характеристики",
	},

	Labels: IndexLabels{
		Fixed:            "Зафиксировать",
		ItemType:         "Тип предмета",
		ItemImage:        "Иконка",
		ItemGroundModel:  "Модель на земле",
		ItemName:         "Название предмета",
		ItemDescription:  "Описание",
		ItemScript:       "Скрипт",
		ItemLevel:        "Уровень",
		ItemStack:        "Стэк",
		ItemPrice:        "Цена",
		ItemSlots:        "Слоты",
		ItemEffect:       "Эф. отображения",
		ItemUsageEffect:  "Эф. использования",
		ItemRepair:       "Ремонтируется",
		ItemSell:         "Продается",
		ItemTake:         "Подбирается",
		ItemThrow:        "Выбрасывается",
		ItemDelete:       "Удаляется",
		ItemLance:        "Ланс",
		ItemModelLance:   "Модель Ланс",
		ItemCarsise:      "Карциз",
		ItemModelCarsise: "Модель Карциз",
		ItemPhyllis:      "Филлис",
		ItemModelPhyllis: "Модель Филли",
		ItemAmi:          "Ами",
		ItemModelAmi:     "Модель Ами",
		ItemNewbie:       "Новичок",
		ItemSwordmaster:  "Мечник",
		ItemHunter:       "Охотник",
		ItemSailor:       "Моряк",
		ItemHealer:       "Лекарь",
		ItemChampion:     "Чемпион",
		ItemWarrior:      "Воитель",
		ItemArcher:       "Стрелок",
		ItemPriest:       "Клерик",
		ItemSealmaster:   "Seal master",
		ItemVoyager:      "Вой",
		ItemStrength:     "Сила",
		ItemAgility:      "Ловкость",
		ItemVitality:     "Телосложение",
		ItemSpirit:       "Дух",
		ItemAccuracy:     "Точность",
		ItemLuck:         "Удача",
		ItemMinAttack:    "Мин. атака",
		ItemMaxAttack:    "Макс.атака",
		ItemAttackSpeed:  "Скор. атаки",
		ItemHitChance:    "Шанс попад.",
		ItemCritChance:   "Шанс крит.",
		ItemMoveSpeed:    "Скор. бега",
		ItemMaxHP:        "Макс. ЖЗ",
		ItemMaxMP:        "Макс. МН",
		ItemHPRegen:      "Вост. ЖЗ",
		ItemMPRegen:      "Вост. МН",
		ItemArmour:       "Защита",
		ItemPhysArmour:   "Физическая",
		ItemEvasion:      "Уклонение",
		ItemDurability:   "Прочность",
		ItemEnergy:       "Энергия",
	},

	ItemTypes: []IndexItemType{
		{
			ID:   1,
			Name: "Одноручный меч",
		}, {
			ID:   2,
			Name: "Двуручный меч",
		}, {
			ID:   3,
			Name: "Лук",
		}, {
			ID:   4,
			Name: "Мушкет",
		}, {
			ID:   7,
			Name: "Клинок",
		}, {
			ID:   9,
			Name: "Посох",
		}, {
			ID:   11,
			Name: "Щит",
		}, {
			ID:   18,
			Name: "Топор (рубка)",
		}, {
			ID:   19,
			Name: "Кирка (рудодобыча)",
		}, {
			ID:   20,
			Name: "Шапка Ами",
		}, {
			ID:   21,
			Name: "Лицо",
		}, {
			ID:   22,
			Name: "Доспехи",
		}, {
			ID:   23,
			Name: "Перчатки",
		}, {
			ID:   24,
			Name: "Ботинки",
		}, {
			ID:   25,
			Name: "Ожерелье",
		}, {
			ID:   26,
			Name: "Кольцо",
		}, {
			ID:   27,
			Name: "Тату",
		}, {
			ID:   28,
			Name: "Прическа",
		}, {
			ID:   29,
			Name: "Коралл",
		}, {
			ID:   34,
			Name: "Книга скиллов",
		}, {
			ID:   36,
			Name: "Билет",
		}, {
			ID:   41,
			Name: "Сундук",
		}, {
			ID:   56,
			Name: "Фрукты для фей",
		}, {
			ID:   57,
			Name: "Еда для фей",
		}, {
			ID:   59,
			Name: "Феи",
		},
	},
}
