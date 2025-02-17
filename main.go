package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"trpg/game"
	"trpg/obj"
)

// Var
var isRunning bool = true

func main() {
	game.CleanTerminal()

	user, err := game.LoadFromFile("data.json")

	if err != nil {
		user = obj.Player{
			Money:  0,
			Health: obj.Health{Current: 10, Max: 10},
			Items:  []obj.Item{}}
	}
	fmt.Print("Welcome to TRPG (Terminal RPG)\n\n")

	if user.Name == "" {
		fmt.Print("Input Username: ")
		fmt.Scan(&user.Name)
	}

	// Reading Input
	fmt.Println("Start command: ")
	for isRunning {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("$> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		args := strings.Split(input, " ")

		inputUser(args, &user)
	}
}

func inputUser(args []string, user *obj.Player) {
	switch args[0] {
	case "b":
		enemy := obj.Enemy[user.Area][rand.Intn(len(obj.Enemy))]
		game.FightEnemy(user, enemy, &isRunning)
	case "clear":
		game.CleanTerminal()
	case "cn":
		if len(args) < 2 {
			fmt.Println("[Game] Please input new name after command")
			return
		}
		if args[1] == user.Name {
			fmt.Println("[Game] Cannot change name with same name")
			return
		}
		fmt.Printf("[Game] Change your name to %s\n", args[1])
		user.SetName(args[1])
	case "h":

		fmt.Print(`[Game] All commands:
b		hunting for enemy
clear		clear terminal
cn		change nickname
h		for get all information
q, quit		quit game
resetgame	reset save data
s		save data
v		view user

`)
	case "q", "quit":
		fmt.Println("[Game] Exiting game")
		isRunning = false
	case "resetgame":
		if len(args) < 2 {
			fmt.Println("[Game] Please input second argument")
			return
		}
		if args[1] != "confirm" {
			fmt.Println("[Game] Please confirm the reset")
			return
		}

		*user = obj.Player{Name: user.Name, Money: 0, Health: obj.Health{Current: 10, Max: 10}}
		fmt.Println("[Game] Data has been reset")
	case "s":
		game.SaveToFile("data.json", *user)
		fmt.Println("[Game] Data has been saved")
	case "v":
		user.Display()
	}
}
