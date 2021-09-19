package locale

var RU_INDEX_LOCALE = IndexTemplate{
	Language: "Русский",

	Text: IndexText{
		ItemDeleted: "Предмет был удален.",
		ItemReturn:  "Восстановить",
		CharDeleted: "Монстр был удален.",
		CharReturn:  "Восстановить",
	},

	Placeholders: IndexPlaceholdes{
		ItemSearchBar:   "ID или название предмета...",
		ItemID:          "ID...",
		ItemImage:       "Путь...",
		ItemGroundModel: "Модель...",
		ItemName:        "Название предмета...",
		ItemDescription: "Описание...",

		CharSearchBar: "ID или имя монстра...",

		Static:  "Статик...",
		Percent: "Процент...",
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

		CharID:                        "ID",
		CharDescription:               "Название",
		CharModelCategory:             "Категория модели",
		CharModelType:                 "Тип модели",
		CharLogicType:                 "Тип логической модели",
		CharFrameworkNumber:           "Номер структуры",
		CharSuiteSerial:               "Suite Serial",
		CharSuiteQuantity:             "Suite Quantity",
		CharPartFACE:                  "Вид лица",
		CharPartHeadApparel:           "Вид головы",
		CharPartArmor:                 "Вид тела",
		CharPartGloves:                "Вид рук",
		CharPartShoes:                 "Вид ног",
		CharPartWeapon:                "Вид оружия в левой руке",
		CharPart2ndWeap:               "Вид оружия в правой руке",
		CharPartWings:                 "Крылья",
		CharFeffID:                    "ID эффекта",
		CharEeffID:                    "EeffID",
		CharSpecialEffectActionSerial: "Номера действ. эффектов",
		CharShadow:                    "Тень",
		CharActionID:                  "ID действия",
		CharTransparency:              "Прозрачность",
		CharMovingSoundEffect:         "Звук ходьбы",
		CharBreathingSoundEffect:      "Звук дыхания",
		CharDeathSoundEffect:          "Звук смерти",
		CharCanItBeControlled:         "Возможность контроля",
		CharAreaLimitation:            "Ограничение в области",
		CharAltitudeExcursion:         "Уровень над землей",
		CharItemsThatCanBeEquipped:    "Одетые вещи",
		CharLength:                    "Длина",
		CharWidth:                     "Ширина",
		CharHeight:                    "Высота",
		CharCollisionRange:            "Дистанция столкновения",
		CharBirth:                     "Возраждение",
		CharDead:                      "Смерть",
		CharBirthEffect:               "Эффект возраждения",
		CharDeathEffect:               "Эффект смерти",
		CharHibernateAction:           "Эффект бездействия",
		CharDeathInstantAction:        "Сразу после смерти",
		CharRemainingHPEffectDisplay:  "Эффект отображения хп",
		CharAttackCanBeSwerve:         "Атк. может отклониться",
		CharConfirmToUseBlownAway:     "Использование ветра",
		CharScript:                    "Сценарий",
		CharWeaponUsed:                "Использует оружие",
		CharSkillID:                   "ID Навыков",
		CharMonsterSkillRate:          "Шанс использования нав.",
		CharItemIDToDrop:              "Дроп лист",
		CharItemDroprate:              "Шанс дропа",
		CharQuantityLimit:             "Предельное количество",
		CharFatalityRate:              "Шанс неудачи",
		CharPrefixLevel:               "Приставка уровня",
		CharQuestDropID:               "Дроп лист (Квест)",
		CharRate:                      "Шанс выпадения",
		CharAI:                        "ИИ",
		CharTurning:                   "Поворот",
		CharVision:                    "Радиус обзора",
		CharNoise:                     "Шум",
		CharGetExp:                    "Получение опыта",
		CharLight:                     "Свет",
		CharMobExp:                    "Mob exp",
		CharMonsterLVL:                "Уровень",
		CharMaxHP:                     "Макс. ХП",
		CharMinHP:                     "Мин. ХП",
		CharMaxSP:                     "Макс. МН",
		CharMinSp:                     "Мин. МН",
		CharMinAtk:                    "Мин. атк",
		CharMaxAtk:                    "Макс. атк",
		CharPhysicalResist:            "Физ Сопротивление",
		CharDefense:                   "Защита",
		CharHitRate:                   "Шанс попадания",
		CharFlee:                      "Шанс уклонения",
		CharMonsterCriticalRate:       "Шанс крит удара",
		CharMF:                        "Шанс выпадания",
		CharHREC:                      "Реген ХП",
		CharSREC:                      "Регех МН",
		CharASPDofMonster:             "Скорость атк",
		CharAreaOfDetection:           "Зона видимости",
		CharAfterDetectionChaseArea:   "Зона преследования",
		CharMSPD:                      "Скорость бега",
		CharCol:                       "Шанс выпадения ресурсов",
		CharStr:                       "Сила",
		CharAgi:                       "Ловкость",
		CharDex:                       "Точность",
		CharCon:                       "Телосложение",
		CharSta:                       "Дух",
		CharLuk:                       "Удача",
		CharLeftRad:                   "left_rad",
		CharGuild:                     "Гильдия",
		CharTitle:                     "Название",
		CharJob:                       "Профессия",
		CharExpGainedFromKill:         "Опыт за убийство",
		CharNexp:                      "Опыт до след уроовня",
		CharFame:                      "Слава",
		CharAp:                        "Очки статов",
		CharTp:                        "Очки навыков",
		CharGd:                        "Деньги",
		CharSpri:                      "Вид камеры",
		CharStor:                      "Stor",
		CharMxSail:                    "Mx Sail",
		CharSail:                      "Sail",
		CharStasa:                     "Stasa",
		CharScsm:                      "Scsm",
		CharTstr:                      "Tstr",
		CharTagi:                      "Tagi",
		CharTdex:                      "Tdex",
		CharTcon:                      "Tcon",
		CharTsta:                      "Tsta",
		CharTluk:                      "Tluk",
		CharTmxhp:                     "Tmxhp",
		CharTmxsp:                     "Tmxsp",
		CharTatk:                      "Tatk",
		CharTdef:                      "Tdef",
		CharThit:                      "Thit",
		CharTflee:                     "Tflee",
		CharTmf:                       "Tmf",
		CharTcrt:                      "Tcrt",
		CharThrec:                     "Threc",
		CharTsrec:                     "Tsrec",
		CharTaspd:                     "Taspd",
		CharTadis:                     "Tadis",
		CharTspd:                      "Tspd",
		CharTspri:                     "Tspri",
		CharTscsm:                     "Tscsm",
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
