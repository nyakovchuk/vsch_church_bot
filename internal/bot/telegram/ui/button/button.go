package button

type ButtonConfig struct {
	Label string
	Data  string
}

type Button struct {
	Config ButtonConfig
}

type ButtonsMap struct {
	Order   []string
	Buttons map[string]ButtonConfig
}
