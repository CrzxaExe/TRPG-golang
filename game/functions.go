package game

import (
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"runtime"
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

func RunCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func CleanTerminal() {
	switch runtime.GOOS {
	case "darwin":
		RunCmd("clear")
	case "linux":
		RunCmd("clear")
	case "windows":
		RunCmd("cmd", "/c", "cls")
	default:
		RunCmd("clear")
	}
}

func Leveling(player *obj.Player, xp float32) {
	playerMaxXP := player.Experience.Lvl*25 + 25

	if xp < float32(playerMaxXP) {
		player.AddXp(player.Experience.Exp + xp)
		return
	}

	player.SetLvl(player.Experience.Lvl + 1)
	player.SetXp(player.Experience.Exp + xp - float32(playerMaxXP))
}
