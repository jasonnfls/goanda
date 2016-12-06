package goanda

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"net/http"
)

func GetAccounts() Accounts {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api-fxtrade.oanda.com/v3/accounts", nil)
	if err != nil {
		fmt.Println(err)
		//NEEDS BETTER ERROR HANDLING
	}

	BearerToken := "Bearer " + Token
	req.Header.Add("Authorization", BearerToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		//NEEDS BETTER ERROR HANDLING
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var accounts Accounts
	err = decoder.Decode(&accounts)
	if err != nil {
		fmt.Println(err)
		//NEEDS BETTER ERROR HANDLING
	}
	return accounts
}

func GetAccount(ID string) Account {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api-fxtrade.oanda.com/v3/accounts/"+ID, nil)
	if err != nil {
		fmt.Println(err)
		//NEEDS BETTER ERROR HANDLING
	}

	BearerToken := "Bearer " + Token
	req.Header.Add("Authorization", BearerToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		//NEEDS BETTER ERROR HANDLING
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var account Account
	err = decoder.Decode(&account)
	if err != nil {
		fmt.Println(err)
		//NEEDS BETTER ERROR HANDLING
	}
	return account
}

func GetAccountSummary(ID string) AccountSummary {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api-fxtrade.oanda.com/v3/accounts/"+ID+"/summary", nil)
	if err != nil {
		fmt.Println(err)
		//NEEDS BETTER ERROR HANDLING
	}

	BearerToken := "Bearer " + Token
	req.Header.Add("Authorization", BearerToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		//NEEDS BETTER ERROR HANDLING
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var accountSummary AccountSummary
	err = decoder.Decode(&accountSummary)
	if err != nil {
		fmt.Println(err)
		//NEEDS BETTER ERROR HANDLING
	}
	return accountSummary
}

func GetAccountInstruments(ID string) AccountInstruments {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api-fxtrade.oanda.com/v3/accounts/"+ID+"/instruments", nil)
	if err != nil {
		fmt.Println(err)
		//NEEDS BETTER ERROR HANDLING
	}

	BearerToken := "Bearer " + Token
	req.Header.Add("Authorization", BearerToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		//NEEDS BETTER ERROR HANDLING
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var accountInstruments AccountInstruments
	err = decoder.Decode(&accountInstruments)
	if err != nil {
		fmt.Println(err)
		//NEEDS BETTER ERROR HANDLING
	}
	return accountInstruments
}

func GetAccountChanges(ID string, transactionID string) AccountChanges {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api-fxtrade.oanda.com/v3/accounts/"+ID+"/changes", nil)
	if err != nil {
		fmt.Println(err)
		//NEEDS BETTER ERROR HANDLING
	}

	BearerToken := "Bearer " + Token
	req.Header.Add("Authorization", BearerToken)
	req.Header.Set("Content-Type", "application/json")

	query := req.URL.Query()
	query.Add("sinceTransactionID", transactionID)
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		//NEEDS BETTER ERROR HANDLING
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var accountChanges AccountChanges
	err = decoder.Decode(&accountChanges)
	if err != nil {
		fmt.Println(err)
		//NEEDS BETTER ERROR HANDLING
	}
	return accountChanges
}

func PatchAccountConfiguration(ID string, alias string, marginRate string) ConfigurationConfirmation {

	var jsonString = []byte(`{"alias": "` + alias + `", "marginRate": "` + marginRate + `"}`)

	client := &http.Client{}
	req, err := http.NewRequest("PATCH", "https://api-fxtrade.oanda.com/v3/accounts/"+ID+"/configuration", bytes.NewBuffer(jsonString))

	if err != nil {
		fmt.Println(err)
		//NEEDS BETTER ERROR HANDLING
	}

	BearerToken := "Bearer " + Token
	req.Header.Add("Authorization", BearerToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		//NEEDS BETTER ERROR HANDLING
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var configurationConfirmation ConfigurationConfirmation
	err = decoder.Decode(&configurationConfirmation)
	if err != nil {
		fmt.Println(err)
		//NEEDS BETTER ERROR HANDLING
	}
	return configurationConfirmation
}

func Getter(url string, token string) {

	//inlets: url, query params, body, struct
	//outlet: populated struct

	//var jsonString = []byte(`{"alias": "Set", "marginRate": "0.02"}`)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	//req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonString))
	if err != nil {
		fmt.Println(err)
		return
	}
	BearerToken := "Bearer " + token

	req.Header.Add("Authorization", BearerToken)
	req.Header.Set("Content-Type", "application/json")

	query := req.URL.Query()
	query.Add("count", "6")
	query.Add("granularity", "S5")
	query.Add("price", "M")
	req.URL.RawQuery = query.Encode()

	fmt.Println(req.URL.String())

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var Instrument Instrument
	err = decoder.Decode(&Instrument)
	if err != nil {
		fmt.Println(err)
		return
	}
	//color.Cyan("████████\t%+v", Instrument.Candles[0].Mid)
	fmt.Printf("%v %+v\n", color.CyanString("██"), Instrument.Candles[0].Mid)
	fmt.Printf("%v %+v\n", color.RedString("██"), Instrument.Candles[0].Mid)
	fmt.Printf("%v %+v\n", color.MagentaString("██"), Instrument.Candles[0].Mid)
	fmt.Printf("%v %+v\n", color.BlueString("██"), Instrument.Candles[0].Mid)
	fmt.Printf("%v %+v\n", color.GreenString("██"), Instrument.Candles[0].Mid)
	fmt.Printf("%v %+v\n", color.YellowString("██"), Instrument.Candles[0].Mid)

	// six color codes

}
