package response

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, data interface{}, statusCode int) (err error) {

	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	err = jsonEncoder.Encode(data)
	if err != nil {
		ErrorJSON(w, err, http.StatusBadGateway)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)
	w.Write(bf.Bytes())
	return nil
}
