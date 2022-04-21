package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/mail"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll((r.Body)); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func JSONReponse(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		res, _ := json.Marshal(struct { Message string `json:"message"`}{Message: fmt.Sprint(err)})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}
	res, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func IsValidEmail(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
}