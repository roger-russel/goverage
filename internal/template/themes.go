package template

func getThemeBGColor(theme string) string {

	var color string

	switch theme {
	case "dracula":
		color = "#282a36"
	default:
		color = "#282a36"
	}

	return color

}
