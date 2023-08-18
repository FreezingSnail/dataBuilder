package datatypes

import (
	"reflect"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
)

type CreatureModel struct {
	creatures []Creature
}

type Creature struct {
	ID         int `json:"ID"`
	Type1      int `json:"Type1"`
	Type2      int `json:"Type2"`
	EvoLevel   int `json:"EvoLevel"`
	AtkSeed    int `json:"AtkSeed"`
	DefSeed    int `json:"DefSeed"`
	SpcAtkSeed int `json:"SpcAtkSeed"`
	SpcDefSeed int `json:"SpcDefSeed"`
	HpSeed     int `json:"HpSeed"`
	SpdSeed    int `json:"SpdSeed"`
	Move1      int `json:"Move1"`
	Move2      int `json:"Move2"`
	Move3      int `json:"Move3"`
	Move4      int `json:"Move4"`
}

func (c Creature) FileName() string {
	return "creatureData.json"
}

func (c Creature) TableHeader() []table.Column {
	val := reflect.ValueOf(c)
	headers := []table.Column{}
	for i := 0; i < val.Type().NumField(); i++ {
		header := val.Type().Field(i).Name
		column := table.Column{
			Title: header,
			Width: len(header),
		}
		headers = append(headers, column)
	}
	return headers
}

func (c Creature) TableRow() table.Row {
	return table.Row{
		strconv.Itoa(c.ID),
		strconv.Itoa(c.Type1),
		strconv.Itoa(c.Type2),
		strconv.Itoa(c.EvoLevel),
		strconv.Itoa(c.AtkSeed),
		strconv.Itoa(c.DefSeed),
		strconv.Itoa(c.SpcAtkSeed),
		strconv.Itoa(c.SpcDefSeed),
		strconv.Itoa(c.HpSeed),
		strconv.Itoa(c.SpdSeed),
		strconv.Itoa(c.Move1),
		strconv.Itoa(c.Move2),
		strconv.Itoa(c.Move3),
		strconv.Itoa(c.Move4),
	}
}
