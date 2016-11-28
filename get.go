package goanda

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Getter(url string, token string) {

	var jsonString = []byte(`{"alias": "Set", "marginRate": "0.02"}`)

	client := &http.Client{}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonString))
	//	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var Account ConfigurationConfirmation
	err = decoder.Decode(&Account)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", Account.ClientConfigureTransaction)

}
