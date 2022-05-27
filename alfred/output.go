package alfred

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

func SHOW(out *Output) {
	if out != nil {
		out.Show()
	}
}

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

func (o *Output) AddSimpleTip(title, subTitle, arg, autoComplete string) *Item {
	item := NewSimpleItem(title, subTitle, arg, autoComplete)
	o.AddTip(item)

	return item
}

func (o *Output) AddTip(item *Item) {
	o.Items = append(o.Items, item)
}

func (o *Output) Show() {
	sort.Slice(o.Items, func(i, j int) bool {
		return o.Items[i].Rank < o.Items[i].Rank
	})

	text, err := json.Marshal(o)
	if err != nil {
		log.Fatalf("json marshal err:%s raw:%+v ", err.Error(), o)
		return
	}

	fmt.Println(string(text))
}
