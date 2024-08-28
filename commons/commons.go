package commons

import (
	"encoding/json"
)

func PrintStruct(payload interface{}) string {
	pbytes, _ := json.Marshal(payload)
	return string(pbytes)
}
