package response

import (
	"net/http"
)

func ErrorJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	response := map[string]interface{}{}
	if data != nil {
		response["error"] = true

		switch r := data.(type) {
		case string:
			response["message"] = r
		case error:
			response["message"] = r.Error()
		case map[string]interface{}:
			response["data"] = r
		case struct{}:
			response["data"] = r
		}

		JSON(w, response, statusCode)
		return
	}
}
