package commons

import (
	"encoding/json"
)

func PrintStruct(payload interface{}) string {
	pbytes, _ := json.Marshal(payload)
	return string(pbytes)
}

func ApiErrorResponse(message string, additionalInfo map[string]interface{}) map[string]interface{} {
	response := map[string]interface{}{
		"status":  "Error",
		"message": message,
	}
	if len(additionalInfo) > 0 {
		response["additional_info"] = additionalInfo
	}
	return response
}
