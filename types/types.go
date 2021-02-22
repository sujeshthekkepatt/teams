package types

//Facts -struct
type Facts struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

//Sections -struct
type Sections struct {
	ActivityTitle    string  `json:"activityTitle"`
	ActivitySubtitle string  `json:"activitySubtitle"`
	Facts            []Facts `json:"facts"`
	Markdown         bool    `json:"markdown"`
}

//Payload -struct
type Payload struct {
	Type       string     `json:"@type"`
	ThemeColor string     `json:"themeColor"`
	Summary    string     `json:"summary"`
	Sections   []Sections `json:"sections"`
}
