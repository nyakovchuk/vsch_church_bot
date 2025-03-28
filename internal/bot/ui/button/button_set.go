package button

type ButtonConfig struct {
	Label string
	Data  string
}

type ButtonSet struct {
	Order   []string
	Buttons map[string]ButtonConfig
}
