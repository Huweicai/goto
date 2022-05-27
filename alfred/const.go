package alfred

type Item struct {
	Title        string            `json:"title"`
	Valid        bool              `json:"valid,omitempty"`
	UID          string            `json:"uid,omitempty"`
	Type         string            `json:"type,omitempty"`
	Subtitle     string            `json:"subtitle,omitempty"`
	Arg          string            `json:"arg,omitempty"`
	Autocomplete string            `json:"autocomplete,omitempty"`
	QuickLookURL string            `json:"quicklookurl,omitempty"`
	Icon         *Icon             `json:"icon,omitempty"`
	Text         *Text             `json:"text,omitempty"`
	Mods         map[string]*Mod   `json:"mods,omitempty"`
	Variables    map[string]string `json:"variables,omitempty"` //From Alfred 3.4.1

	Rank int64 `json:"-"` // for internal sorting
}

type Icon struct {
	Type string `json:"type,omitempty"`
	Path string `json:"path"`
}

type Text struct {
	Copy      string `json:"copy,omitempty"`
	LargeType string `json:"largetype,omitempty"`
}

type Mod struct {
	Valid     bool              `json:"valid"`
	Arg       string            `json:"arg"`
	Subtitle  string            `json:"subtitle"`
	Icon      *Icon             `json:"icon,omitempty"`      //From Alfred 3.4.1
	Variables map[string]string `json:"variables,omitempty"` //From Alfred 3.4.1
}

type Output struct {
	Items []*Item `json:"items"`
}
