package alfred

import (
	"encoding/json"
	"fmt"
	"log"
)

func NewSimpleItem(title, subTitle, arg, autoComplete string) *Item {
	return &Item{
		Title:        title,
		Subtitle:     subTitle,
		Arg:          arg,
		Autocomplete: autoComplete,
	}
}

func NewOutput() *Output {
	return &Output{}
}

func (o *Output) AddSimpleTip(title, subTitle, arg, autoComplete string) {
	o.AddTip(NewSimpleItem(title, subTitle, arg, autoComplete))
}

func (o *Output) AddTip(item *Item) {
	o.Items = append(o.Items, item)
}

func (o *Output) Show() {
	text, err := json.Marshal(o)
	if err != nil {
		log.Fatalf("json marshal err:%s raw:%+v ", err.Error(), o)
		return
	}
	fmt.Println(string(text))
}
