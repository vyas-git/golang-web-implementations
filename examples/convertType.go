package examples

import (
	"encoding/json"
	"fmt"
)

type TypeMapping struct{}

func (TypeMapping) Run() {
	data := getItems()
	fmt.Println(data.Items)
}

type List1 struct {
	Ids []int
}

type List2 struct {
	Items []int `json:"Ids"`
}

func getItems() List2 {
	t1 := List1{Ids: []int{1, 2, 3}}
	var t2 List2
	Convert(t1, &t2)
	return t2
}
func Convert(in interface{}, out interface{}) error {
	j, err := json.Marshal(in)
	if err != nil {
		return err
	}
	err = json.Unmarshal(j, &out)
	if err != nil {
		return err
	}
	return nil
}
