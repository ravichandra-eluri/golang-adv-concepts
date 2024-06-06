package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int    `json:"Age"`
}

func main() {
	jsonData := `[{"FirstName":"James","LastName":"Bond","Age":32},{"FirstName":"Miss","LastName":"MoneyPenny","Age":27}]`
	bs := []byte(jsonData)

	fmt.Printf("%T\n", jsonData)
	fmt.Printf("%T\n", bs)

	var people []person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all the data", people)

	for i, v := range people {
		fmt.Println("\nperson number:", i)
		fmt.Println(v.FirstName, v.LastName, v.Age)
	}
}
