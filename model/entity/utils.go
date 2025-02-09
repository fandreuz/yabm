package entity

import "encoding/json"

func EntityToString(entity any) string {
	b, jsonErr := json.Marshal(entity)
	if jsonErr != nil {
		panic(jsonErr)
	}
	return string(b)
}
