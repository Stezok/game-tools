package item

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func GenerateItemCode(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'

	row, err := reader.Read()
	if err != nil {
		return "", err
	}

	result := `ID                              string
	Name                            string
	IconName                        string
	ModelGround                     string
	ModelLance                      string
	ModelCarsise                    string
	ModelPhyllis                    string
	ModelAmi                        string
	ShipSymbol                      string
	ShipSize                        string
	ItemType                        string
	ObtainPrefixRate                string
	SetID                           string
	ForgingLevel                    string
	StableLevel                     string
	Repairable                      string
	Tradable                        string
	PickUpable                      string
	Droppable                       string
	Deletable                       string
	MaxStackSize                    string
	Instance                        string
	Price                           string
	CharacterTypes                  string
	CharacterLevel                  string
	CharacterClasses                string
	CharacterNick                   string
	CharacterReputation             string
	EquipableSlots                  string
	ItemSwitchLocations             string
	ItemObtainIntoLocationDetermine string
	STR                             string
	AGI                             string
	ACC                             string
	CON                             string
	SPR                             string
	LUK                             string
	AttackSpeed                     string
	AttackRange                     string
	MinAttack                       string
	MaxAttack                       string
	Defense                         string
	MaxHP                           string
	MaxSP                           string
	DodgeRate                       string
	HitRate                         string
	CriticalRate                    string
	TreasureDropRate                string
	HPRecovery                      string
	SPRecovery                      string
	MovementSpeed                   string
	ResourceGatheringRate           string
	STRMinMax                       string
	AGIMinMax                       string
	ACCMinMax                       string
	CONMinMax                       string
	SPRMinMax                       string
	LUKMinMax                       string
	AttackSpeedAdd                  string
	AttackRangeAdd                  string
	MinAttackAdd                    string
	MaxAttackAdd                    string
	DefenseAdd                      string
	MaxHPAdd                        string
	MaxSPAdd                        string
	DodgeAdd                        string
	HitRateAdd                      string
	CriticalRateAdd                 string
	TreasureDropRateAdd             string
	HPRecoveryAdd                   string
	SPRecoveryAdd                   string
	MovementSpeedAdd                string
	ResourceGatheringRateAdd        string
	PhysicalResistAdd               string
	AftectedByLeftHandEfs           string
	Energy                          string
	Durability                      string
	GemSockets                      string
	ShipDurabilityRecovery          string
	ShipCannonAmount                string
	ShipMemberCount                 string
	ShipLabel                       string
	ShipCargoCapacity               string
	ShipFuelConsumption             string
	ShipAttackSpeed                 string
	ShipMovementSpeed               string
	UsageScript                     string
	DisplayEffect                   string
	BindEffect                      string
	BindEffect2                     string
	FirstInvSlotEffect              string
	DropModelEffect                 string
	ItemUsageEffect                 string
	Description                     string
	Last                            string`

	res := strings.Split(result, "\n")

	insrt := ""
	for i, field := range row {
		fname := ""
		ii := 0
		f := false
		for {
			if (res[i][ii] == ' ' || res[i][ii] == '\t') && f {
				break
			}
			if res[i][ii] != ' ' && res[i][ii] != '\t' && !f {
				f = true
			}

			if f {
				fname += string(res[i][ii])
			}
			ii++
		}
		insrt += fmt.Sprintf("%s `item:\"%s\" json:\"%s\"`\n", res[i], field, fname)
	}

	return insrt, nil
}
