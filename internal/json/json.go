package json

import (
	"encoding/json"
	"net/http"
)

func Write(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)

}

func Read(r *http.Request, dst any) error {
	decorder := json.NewDecoder(r.Body)
	decorder.DisallowUnknownFields()
	if err := decorder.Decode(dst); err != nil {
		return err
	}

	return nil
}

