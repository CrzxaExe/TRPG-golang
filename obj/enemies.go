package obj

var Enemy = append([][]Actor{},
	[]Actor{
		{Id: "Slime", Health: Health{Max: 5, Current: 5}, Stats: ActorStat{Atk: 1, Def: 0, Res: 0}, XP: 0.3},
		{Id: "Harded Slime", Health: Health{Max: 9, Current: 9}, Stats: ActorStat{Atk: 1, Def: 1, Res: 0}, XP: 0.7},
	},
	[]Actor{
		{Id: "Wolf", Health: Health{Max: 20, Current: 20}, Stats: ActorStat{Atk: 5, Def: 2, Res: 0}, XP: 7.9},
		{Id: "Alpha Wolf", Health: Health{Max: 28, Current: 28}, Stats: ActorStat{Atk: 7, Def: 4, Res: 0}, XP: 12.5},
	},
)
