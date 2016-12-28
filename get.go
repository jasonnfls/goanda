package goanda

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (client *Client) GetAccounts() Accounts {

	req, err := http.NewRequest("GET", client.GetBase("/accounts"), nil)
	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var accounts Accounts
	err = decoder.Decode(&accounts)
	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}
	return accounts
}

func (client *Client) GetAccount(ID string) Account {
	req, err := http.NewRequest("GET", client.GetBase("/accounts/"+ID), nil)
	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var account Account
	err = decoder.Decode(&account)
	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}
	return account
}

func (client *Client) GetAccountSummary(ID string) AccountSummary {

	req, err := http.NewRequest("GET", client.GetBase("/accounts/"+ID+"/summary"), nil)
	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var accountSummary AccountSummary
	err = decoder.Decode(&accountSummary)
	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}
	return accountSummary
}

func (client *Client) GetAccountInstruments(ID string) AccountInstruments {

	req, err := http.NewRequest("GET", client.GetBase("/accounts/"+ID+"/instruments"), nil)
	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var accountInstruments AccountInstruments
	err = decoder.Decode(&accountInstruments)
	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}
	return accountInstruments
}

func (client *Client) GetAccountChanges(ID string, transactionID string) AccountChanges {

	req, err := http.NewRequest("GET", client.GetBase("/accounts/"+ID+"/changes"), nil)
	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}

	query := req.URL.Query()
	query.Add("sinceTransactionID", transactionID)
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var accountChanges AccountChanges
	err = decoder.Decode(&accountChanges)
	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}
	return accountChanges
}

func (client *Client) PatchAccountConfiguration(ID string, alias string, marginRate string) ConfigurationConfirmation {

	var jsonString = []byte(`{"alias": "` + alias + `", "marginRate": "` + marginRate + `"}`)

	req, err := http.NewRequest("PATCH", client.GetBase("/accounts/"+ID+"/configuration"), bytes.NewBuffer(jsonString))

	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var configurationConfirmation ConfigurationConfirmation
	err = decoder.Decode(&configurationConfirmation)
	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}
	return configurationConfirmation
}

