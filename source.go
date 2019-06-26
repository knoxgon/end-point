package main

import (
	"encoding/json"
	"net/http"

	"github.com/knoxgon/codetest/ibanpkg"
)

/*
	Author: Volkan Güven
	Description: Endpoint for IBAN verification
*/

func ibanHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	iban := req.URL.Query().Get("iban")

	if ibanpkg.ControlIban(iban) {
		json.NewEncoder(res).Encode(map[string]string{"status": "Success"})
	} else {
		json.NewEncoder(res).Encode(map[string]string{"status": "Failure"})
	}
}

func main() {

	http.HandleFunc("/ibanverifier", ibanHandler)
	err := http.ListenAndServe(":3333", nil)

	if err != nil {
		panic(err.Error())
	}
}
