package utils

import "encoding/json"

func ToJSON(data interface{}) []byte {
	bytes, err := json.Marshal(data)

	if err != nil {
		return []byte("null")
	}
	return bytes
}
