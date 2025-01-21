package obj

import "fmt"

type Player struct {
	Name       string     `json:"name"`
	Experience Experience `json:"experience"`
	Money      int32      `json:"money"`
	Health     Health     `json:"health"`
	Weapon     *Weapon    `json:"weapon"`
	Armor      *Armor     `json:"armor"`
	Items      []Item     `json:"items"`
}

type Item struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type Health struct {
	Current int `json:"current"`
	Max     int `json:"max"`
}

type Experience struct {
	Exp float32 `json:"exp"`
	Lvl int     `json:"lvl"`
}

type WeaponStat struct {
	Atk  int     `json:"atk"`
	Def  int     `json:"def"`
	Aspd float32 `json:"aspd"`
}

type Weapon struct {
	ID   string     `json:"id"`
	Name string     `json:"name"`
	Stat WeaponStat `json:"stat"`
}

type ArmorStat struct {
	Def int `json:"def"`
	Res int `json:"res"`
}

type Armor struct {
	ID   string    `json:"id"`
	Name string    `json:"name"`
	Stat ArmorStat `json:"stat"`
}

// Displaying user stats
func (t Player) Display() {
	fmt.Printf("Name	: %q\n", t.Name)
	fmt.Printf("Lvl	: %d (%.2f XP) \n", t.Experience.Lvl, t.Experience.Exp)
	if t.Weapon != nil {
		fmt.Printf("Weapon	: %s, %d Atk,  %d Def %.f Aspd\n", t.Weapon.Name, t.Weapon.Stat.Atk, t.Weapon.Stat.Def, t.Weapon.Stat.Aspd)
	}
	if t.Armor != nil {
		fmt.Printf("Weapon	: %s, %d Def %d Res\n", t.Armor.Name, t.Armor.Stat.Def, t.Armor.Stat.Res)
	}
	fmt.Printf("Health 	: %d/%d HP\n", t.Health.Current, t.Health.Max)
	fmt.Printf("Money 	: $%d\n", t.Money)
	if len(t.Items) != 0 {
		fmt.Printf("Items	: %v\n", t.Items)
	}
}

// set player name
func (t *Player) SetName(name string) {
	t.Name = name
}
