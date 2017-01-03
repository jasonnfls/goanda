package goanda

import (
	"fmt"
	"bytes"
	"net/http"
	"encoding/json"
)

type OrderResponse struct {
	OrderCreateTransaction *Transaction
	OrderFillTransaction *OrderFillTransaction
	OrderCancelTransaction *OrderCancelTransaction
	LastTransactionID string
	RelatedTransactionIDs []string
}

type OrderRequest struct {
	Type string `json:"type"`
	Units int `json:"units,string"`
	Instrument string `json:"instrument"`
	TimeInForce  string `json:"timeInForce"`
	ClientExtensions *ClientExtensions `json:"clientExtensions,omitempty"`
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions,omitempty"`
}

func (client *Client) NewOrder(ID string, orderRequest *OrderRequest) OrderResponse {
	jsonbytes, err := json.Marshal(&struct {
		Order *OrderRequest `json:"order"`
	}{
		Order: orderRequest,
	})
	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}

	req, err := http.NewRequest("POST", client.GetBase("/accounts/"+ID+"/orders"), bytes.NewBuffer(jsonbytes))
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
	fmt.Println(resp.Status)

	decoder := json.NewDecoder(resp.Body)
	var responseObj OrderResponse
	err = decoder.Decode(&responseObj)
	if err != nil {
		panic(err)
		//NEEDS BETTER ERROR HANDLING
	}
	return responseObj
}

