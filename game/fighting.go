package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"trpg/obj"
)

func FightEnemy(player *obj.Player, target obj.Actor, isRunning *bool) {
	*isRunning = false
	fmt.Println(player, target)

	isFighting := true

	tempEnemy := target

	CleanTerminal()
	fmt.Printf("You fighting with %s, press h for help\n\n", tempEnemy.Id)
	ShowStats(*player, target)
	for isFighting {
		Fighting(player, &tempEnemy, &isFighting)
	}
	*isRunning = true
}

func getPlayerStats(player *obj.Player) (int, int, int) {
	var atk, def, res int = 2, 0, 0

	if player.Weapon != nil {
		if player.Weapon.Stat.Atk > 0 {
			atk = player.Weapon.Stat.Atk
		}
		if player.Weapon.Stat.Def > 0 {
			def += player.Weapon.Stat.Def
		}
	}

	if player.Armor != nil {
		if player.Armor.Stat.Res > 0 {
			res = player.Armor.Stat.Res
		}

		if player.Armor.Stat.Def > 0 {
			def += player.Armor.Stat.Def
		}
	}

	return atk, def, res
}

func Fighting(player *obj.Player, target *obj.Actor, isFighting *bool) {
	if player.Health.Current == 0 {
		fmt.Println("[Fight] You lose the fight")
		*isFighting = false
		return
	}
	if target.Health.Current <= 0 {
		fmt.Printf("[Fight] You won the fight, you get %.2f\n\n", target.XP)
		Leveling(player, target.XP)
		*isFighting = false
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("@-> ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	args := strings.Split(input, " ")

	switch args[0] {
	case "a", "attack":
		AttackAction(player, target)
	case "h", "help":
		fmt.Print(`[Fight] Help command:
h		help
s		using shield, increase your def
v		display enemy & player stats
q, flee		quit battle

`)
	case "s":
		DefAction(player, target)
	case "v":
		CleanTerminal()
		ShowStats(*player, *target)
	case "q", "flee":
		fmt.Println("[Game] You flee the battle")
		*isFighting = false
	}
}

func ShowStats(player obj.Player, enemy obj.Actor) {
	var res, pRes string = "", ""

	var playerAtk, playerDef, playerRes int = getPlayerStats(&player)

	if enemy.Stats.Res > 0 {
		res = ", " + strconv.FormatInt(int64(enemy.Stats.Res), 7) + "% Res"
	}
	if playerRes > 0 {
		pRes = ", " + strconv.FormatInt(int64(playerRes), 7) + "% Res"
	}

	fmt.Printf("%s:\nHealth	%d/%d HP\n%d Atk, %d Def%s\n\n%s:\nHealth %d/%d HP\n%d Atk %d Def%s\n\n", enemy.Id, enemy.Health.Current, enemy.Health.Max, enemy.Stats.Atk, enemy.Stats.Def, res, player.Name, player.Health.Current, player.Health.Max, playerAtk, playerDef, pRes)
}

func AttackAction(player *obj.Player, enemy *obj.Actor) {
	var playerAtk, playerDef, _ int = getPlayerStats(player)

	enemyAtk, playerDmg := enemy.Stats.Atk-playerDef, playerAtk-enemy.Stats.Def
	if enemyAtk < 0 {
		enemyAtk = 1
	}
	if playerDmg < 0 {
		playerDmg = 1
	}

	player.Health.Current -= enemyAtk

	enemy.Health.Current -= playerDmg

	fmt.Printf("[Fight] You success do %d atk to enemy, but you get %d damage\n", playerDmg, enemyAtk)

	if player.Health.Current < 0 {
		player.Health.Current = 0
		player.SetCurrentHealth(0)
	}
	player.SetCurrentHealth(player.Health.Current)
}

func DefAction(player *obj.Player, enemy *obj.Actor) {
	var _, playerDef, _ int = getPlayerStats(player)

	enemyAtk, playerDmg := enemy.Stats.Atk-((playerDef+2)*2), 0
	if enemyAtk < 0 {
		enemyAtk = 1
	}
	if playerDmg < 0 {
		playerDmg = 1
	}

	player.Health.Current -= enemyAtk

	enemy.Health.Current -= playerDmg

	fmt.Printf("[Fight] You blocking, you get %d damage\n", enemyAtk)
	if player.Health.Current < 0 {
		player.Health.Current = 0
		player.SetCurrentHealth(0)
	}
	player.SetCurrentHealth(player.Health.Current)
}
