package goanda

import (
	"bytes"
	"net/http"
	"encoding/json"
)

type UpdateTradeResponse struct {
	LastTransactionID string
	RelatedTransactionIDs []string
}

type UpdateTradeRequest struct {
	TakeProfit *TakeProfitDetails `json:"takeProfit,omitempty"`
	StopLoss *StopLossDetails `json:"stopLoss,omitempty"`
	TrailingStopLoss *TrailingStopLossDetails `json:"trailingStopLoss,omitempty"`
}

func (client *Client) UpdateTrade(ID string, tradeID string, updateTradeRequest *UpdateTradeRequest) UpdateTradeResponse {
	jsonbytes, err := json.Marshal(updateTradeRequest)
	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}

	req, err := http.NewRequest("PUT", client.GetBase("/accounts/"+ID+"/trades/"+tradeID+"/orders"), bytes.NewBuffer(jsonbytes))
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
	var responseObj UpdateTradeResponse
	err = decoder.Decode(&responseObj)
	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}
	return responseObj
}

