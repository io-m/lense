package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Response is generic struct with message and status
// code fields. It is used for responding with nicely formatted
// error messages
type Response struct {
	Msg        string `json:"msg,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
}

// Back is generic function that returns error embedded
// in struct. It requires http status code and corresponding message
func Back(sc int, msg string) *Response {
	return &Response{
		Msg:        msg,
		StatusCode: sc,
	}
}

// ToJSON is wrapper for json Encode function
func ToJSON(w http.ResponseWriter, data interface{}) error {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return err
	}
	return nil
}

// FromJSON is wrapper for json Decode function
func FromJSON(r *http.Request, dest interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&dest); err != nil {
		return err
	}
	return nil
}

// Params is helper function for extracting parameter out of URI
func Params(r *http.Request, param string, ok ...bool) (interface{}, error) {
	if len(ok) == 0 {
		ok[0] = true
	}
	vars := mux.Vars(r)
	stringParam := vars[param]
	if !ok[0] {
		paramInt, err := strconv.Atoi(stringParam)
		if err != nil {
			return nil, err
		}
		return paramInt, nil
	}
	return stringParam, nil
}
