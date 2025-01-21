package game

import (
	"encoding/json"
	"io"
	"os"
	"trpg/obj"
)

func SaveToFile(filename string, data obj.Player) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(data)
}

func LoadFromFile(filename string) (obj.Player, error) {
	var data obj.Player
	file, err := os.Open(filename)

	if err != nil {
		return data, err
	}

	defer file.Close()
	bytes, err := io.ReadAll(file)

	if err != nil {
		return data, err
	}

	err = json.Unmarshal(bytes, &data)
	return data, err
}
