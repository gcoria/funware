package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func HandleCreateAccount(w http.ResponseWriter, req *http.Request) {
	var account Account
	if err := json.NewDecoder(req.Body).Decode(&account); err != nil {
		slog.Error("account creation failed: ", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func notifyAccountCreated() {

}

type Account struct {
	Username string
	Email    string
}

func main() {

}
