package alfred

func NewSimpleItem(title, subTitle, arg, autoComplete string) *Item {
	return &Item{
		Title:        title,
		Subtitle:     subTitle,
		Arg:          arg,
		Autocomplete: autoComplete,
	}
}
