package app

import (
	"encoding/json"
)

func Dev() string {
	return "Dev"
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
