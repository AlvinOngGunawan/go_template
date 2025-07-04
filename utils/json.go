package utils

import (
	"encoding/json"
	"strings"
)

type Accessory struct {
	Name string `json:"name"`
}

func DecodeAccessoriesArrayToString(text string, contentIfNull string) string {
	if text == "" || text == "null" {
		return contentIfNull
	}

	var accs []Accessory
	err := json.Unmarshal([]byte(text), &accs)
	if err != nil {
		return contentIfNull
	}

	names := make([]string, 0, len(accs))
	for _, acc := range accs {
		names = append(names, acc.Name)
	}

	return strings.Join(names, ", ")
}
