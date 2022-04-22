package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/MonduCareers/-Nwuguru-Sunday-Coding-Challenge/pkg/config"
)

func ParseContent(r *http.Response, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func TruncateDBTables () {
	db := config.GetDB()
	db.Exec("TRUNCATE TABLE bank_accounts")
	db.Exec("TRUNCATE TABLE bank_account_transactions")
}
