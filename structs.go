package goanda

type Accounts struct {
	ID   string      `json:"id"`
	Tags interface{} `json:"tags"`
}

type Account struct {
	Details           *AccountDetails `json:"account"`
	LastTransactionID string          `json:"lastTransactionID"`
}

type AccountSummary struct {
	Details           *AccountSummary `json:"account"`
	LastTransactionID string          `json:"lastTransactionID"`
}

type AccountDetails struct {
	NAV                         string      `json:"NAV"`
	Alias                       string      `json:"alias"`
	Balance                     string      `json:"balance"`
	CreatedByUserID             int         `json:"createdByUserID"`
	CreatedTime                 string      `json:"createdTime"`
	Currency                    string      `json:"currency"`
	HedgingEnabled              bool        `json:"hedgingEnabled"`
	ID                          string      `json:"id"`
	LastTransactionID           string      `json:"lastTransactionID"`
	MarginAvailable             string      `json:"marginAvailable"`
	MarginCallMarginUsed        string      `json:"marginCallMarginUsed"`
	MarginCallPercent           string      `json:"marginCallPercent"`
	MarginCloseoutMarginUsed    string      `json:"marginCloseoutMarginUsed"`
	MarginCloseoutNAV           string      `json:"marginCloseoutNAV"`
	MarginCloseoutPercent       string      `json:"marginCloseoutPercent"`
	MarginCloseoutPositionValue string      `json:"marginCloseoutPositionValue"`
	MarginCloseoutUnrealizedPL  string      `json:"marginCloseoutUnrealizedPL"`
	MarginRate                  string      `json:"marginRate"`
	MarginUsed                  string      `json:"marginUsed"`
	OpenPositionCount           int         `json:"openPositionCount"`
	OpenTradeCount              int         `json:"openTradeCount"`
	Orders                      interface{} `json:"orders"`
	PendingOrderCount           int         `json:"pendingOrderCount"`
	Pl                          string      `json:"pl"`
	PositionValue               string      `json:"positionValue"`
	Positions                   interface{} `json:"positions"`
	ResettablePL                string      `json:"resettablePL"`
	Trades                      interface{} `json:"trades"`
	UnrealizedPL                string      `json:"unrealizedPL"`
	WithdrawalLimit             string      `json:"withdrawalLimit"`
}

type AccountSummary struct {
	NAV                         string `json:"NAV"`
	Alias                       string `json:"alias"`
	Balance                     string `json:"balance"`
	CreatedByUserID             int    `json:"createdByUserID"`
	CreatedTime                 string `json:"createdTime"`
	Currency                    string `json:"currency"`
	HedgingEnabled              bool   `json:"hedgingEnabled"`
	ID                          string `json:"id"`
	LastTransactionID           string `json:"lastTransactionID"`
	MarginAvailable             string `json:"marginAvailable"`
	MarginCallMarginUsed        string `json:"marginCallMarginUsed"`
	MarginCallPercent           string `json:"marginCallPercent"`
	MarginCloseoutMarginUsed    string `json:"marginCloseoutMarginUsed"`
	MarginCloseoutNAV           string `json:"marginCloseoutNAV"`
	MarginCloseoutPercent       string `json:"marginCloseoutPercent"`
	MarginCloseoutPositionValue string `json:"marginCloseoutPositionValue"`
	MarginCloseoutUnrealizedPL  string `json:"marginCloseoutUnrealizedPL"`
	MarginRate                  string `json:"marginRate"`
	MarginUsed                  string `json:"marginUsed"`
	OpenPositionCount           int    `json:"openPositionCount"`
	OpenTradeCount              int    `json:"openTradeCount"`
	PendingOrderCount           int    `json:"pendingOrderCount"`
	Pl                          string `json:"pl"`
	PositionValue               string `json:"positionValue"`
	ResettablePL                string `json:"resettablePL"`
	UnrealizedPL                string `json:"unrealizedPL"`
	WithdrawalLimit             string `json:"withdrawalLimit"`
}
