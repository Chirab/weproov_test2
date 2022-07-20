package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

func ParseIncomingInput(w http.ResponseWriter, r *http.Request, data interface{}) error {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		return nil
	}
	
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&data); err != nil {
		log.Println(err)
		return err
	}

	return nil 
}