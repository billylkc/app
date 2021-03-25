package app

import (
	"encoding/json"
	"fmt"

	"github.com/billylkc/app/calc"
)

func Dev() string {
	res, err := calc.GetDailySales("2021-03-25", 7)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(PrettyPrint(res))
	return "Dev"
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
