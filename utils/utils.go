package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil{
		return fmt.Errorf("Missing Request Body")
	}
	return json.NewDecoder(r.Body).Decode(payload);
}

func WriteJSON(w http.ResponseWriter, status int, content any)  error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(content)
}

func WriteError(w http.ResponseWriter, status int, err error){
	WriteJSON(w, status, map[string]string{"Error": err.Error()})
}