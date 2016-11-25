package goanda

import (
	"strconv"
)

func ListInstruments(a int) (string, string) {
	token, err := ReturnToken(a)
	if err != nil {
		return "", ""
	}
	acc := strconv.Itoa(a)

	return "https://api-fxtrade.oanda.com/v1/instruments?accountId=" + acc, token
}

func StreamRates(a int) (string, string) {
	token, err := ReturnToken(a)
	if err != nil {
		return "", ""
	}
	acc := strconv.Itoa(a)

	return "https://stream-fxtrade.oanda.com/v1/prices?accountId=" + acc + "&instruments=AUD_CAD%2CAUD_CHF", token
}

func AccountData(a int) (string, string) {
	token, err := ReturnToken(a)
	if err != nil {
		return "", ""
	}
	acc := strconv.Itoa(a)

	return "https://api-fxtrade.oanda.com/v1/accounts/" + acc, token
}
func AllAccounts(a int) (string, string) {
	token, err := ReturnToken(a)
	if err != nil {
		return "", ""
	}

	return "https://api-fxtrade.oanda.com/v1/accounts", token
}
