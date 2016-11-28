package goanda

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Getter(url string, token string) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var Account AccountChanges
	err = decoder.Decode(&Account)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", Account.State)

}
