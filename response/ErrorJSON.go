package response

import (
	"net/http"
)

func ErrorJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	response := map[string]interface{}{}
	if data != nil {
		response["error"] = true

		response["message"] = data
		if e, ok := data.(error); ok {
			response["message"] = e.Error()
		}

		JSON(w, response, statusCode)
		return
	}
}
