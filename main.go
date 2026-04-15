package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
	commands "trpg/src/constant"
	"trpg/src/game"
	"trpg/src/obj"
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
	fmt.Print("Welcome to TRPG (Terminal RPG)\nUse command help or start to get some information\n\n")

	if user.Name == "" {
		fmt.Print("Input Username: ")
		fmt.Scan(&user.Name)
	}

	// Reading Input
	fmt.Println("Input your command: ")
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
	case "inventory", "i":
		w := tabwriter.NewWriter(os.Stdout, 10, 1, 10, ' ', 0)
		fmt.Fprintf(w, "Item^\tAmount^\tAtributes^\n \t \t \n")

		items := user.Items

		sort.Slice(items, func(i, j int) bool {
			return items[i].Name < items[j].Name
		})

		for _, item := range items {
			attributes := []string{}

			for _, att := range item.Attributes {
				attributes = append(attributes, fmt.Sprintf("%s: %d", att.Name, att.Value))
			}

			fmt.Fprintf(w, "%s\t%d\t%s\n", item.Name, item.Amount, strings.Join(attributes, ", "))
		}

		w.Flush()
	case "hunt":
		fmt.Println("[Game] Hunting an enemy")
		time.Sleep(time.Duration(rand.Intn(6)) * time.Second)

		enemy := obj.Enemy[user.Area][rand.Intn(len(obj.Enemy))]
		game.FightEnemy(user, enemy, &isRunning)
	case "clear", "cls":
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
	case "h", "help":

		w := tabwriter.NewWriter(os.Stdout, 2, 1, 4, ' ', 0)
		fmt.Fprintf(w, "Command^\tAliases^\tDescription^\n \t \t \n")

		for _, command := range commands.Commands {
			name := command.Name

			params := []string{}
			for _, param := range command.Params {
				var req string = ""

				if param.Required {
					req = "*"
				}

				params = append(params, fmt.Sprintf("<%s%s:%s>", param.Name, req, param.Type))
			}

			name = name + " " + strings.Join(params, " ")
			fmt.Fprintf(w, "%s\t%s\t%s\n", name, strings.Join(command.Alias, ", "), command.Description)
		}

		w.Flush()
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
	case "s", "save":
		game.SaveToFile("data.json", *user)
		fmt.Println("[Game] Data has been saved")
	case "v", "view":
		user.Display()
	}
}
