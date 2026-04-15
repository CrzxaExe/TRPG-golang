package constant

type CommandParams struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
}

type CommandData struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Alias       []string        `json:"alias"`
	Params      []CommandParams `json:"params"`
}

// Commands list
var Commands = append([]CommandData{},
	CommandData{
		Name:        "clear",
		Description: "clear terminal view",
		Alias:       []string{"cls"},
		Params:      []CommandParams{},
	},
	CommandData{
		Name:        "cn",
		Description: "change your nickname",
		Alias:       []string{},
		Params: []CommandParams{
			{
				Name:     "new",
				Type:     "string",
				Required: false,
			},
		},
	},
	CommandData{
		Name:        "help",
		Description: "get all information you may needed",
		Alias:       []string{"h"},
		Params:      []CommandParams{},
	},
	CommandData{
		Name:        "hunt",
		Description: "Hunting enemy from the current area",
		Alias:       []string{},
		Params:      []CommandParams{},
	},
	CommandData{
		Name:        "inventory",
		Description: "Display a list of your inventory",
		Alias:       []string{"i"},
		Params:      []CommandParams{},
	},
	CommandData{
		Name:        "quit",
		Description: "quit the game",
		Alias:       []string{"q"},
		Params:      []CommandParams{},
	},
	CommandData{
		Name:        "resetgame",
		Description: "reset your game data back to default",
		Alias:       []string{},
		Params: []CommandParams{
			{
				Name:     "confirm",
				Type:     "confirm",
				Required: true,
			},
		},
	},
	CommandData{
		Name:        "save",
		Description: "save data",
		Alias:       []string{"s"},
		Params:      []CommandParams{},
	},
	CommandData{
		Name:        "view",
		Description: "view user stats and info",
		Alias:       []string{"v"},
		Params:      []CommandParams{},
	},
)
