package datatypes

import (
	"encoding/json"
	"os"

	"github.com/charmbracelet/bubbles/table"
)

type Data struct {
	Name     string
	Schema   map[string]string
	DataPack []map[string]string
}

type DataModel interface {
	TableHeader() []table.Column
	TableRow() table.Row
	FileName() string
}

func WriteToFile(d DataModel, path string) error {

	jsn, err := json.Marshal(d)
	if err != nil {
		return err
	}
	filePath := path + d.FileName()
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	f.Write(jsn)
	f.Close()
	return nil
}

func ReadFromFile(schemaPath, dataPath string) ([]Data, error) {

	schemajsn, err := os.ReadFile(schemaPath)
	if err != nil {
		return nil, err
	}
	datajsn, err := os.ReadFile(dataPath)
	if err != nil {
		return nil, err
	}

	return buildDataSlice(schemajsn, datajsn)
}

func buildDataSlice(schemajsn, datajsn []byte) ([]Data, error) {
	schemas := make(map[string]map[string]string, 0)
	data := make(map[string][]map[string]string, 0)

	err := json.Unmarshal(schemajsn, &schemas)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(datajsn, &data)
	if err != nil {
		return nil, err
	}

	editorData := []Data{}
	for key := range schemas {

		pack := make([]map[string]string, 0)
		for _, p := range data[key] {
			pack = append(pack, p)
		}
		d := Data{
			Name:     key,
			Schema:   schemas[key],
			DataPack: pack,
		}

		editorData = append(editorData, d)
	}

	return editorData, nil

}
