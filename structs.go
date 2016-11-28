package goanda

type Accounts struct {
	ID   string      `json:"id"`
	Tags interface{} `json:"tags"`
}

type Account struct {
	Details           *AccountDetails `json:"account"`
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

type AccountInstruments struct {
	Instruments       []*Instruments `json:"instruments"`
	LastTransactionID string         `json:"lastTransactionID"`
}

type Instruments struct {
	DisplayName                 string `json:"displayName"`
	DisplayPrecision            int    `json:"displayPrecision"`
	MarginRate                  string `json:"marginRate"`
	MaximumOrderUnits           string `json:"maximumOrderUnits"`
	MaximumPositionSize         string `json:"maximumPositionSize"`
	MaximumTrailingStopDistance string `json:"maximumTrailingStopDistance"`
	MinimumTradeSize            string `json:"minimumTradeSize"`
	MinimumTrailingStopDistance string `json:"minimumTrailingStopDistance"`
	Name                        string `json:"name"`
	PipLocation                 int    `json:"pipLocation"`
	TradeUnitsPrecision         int    `json:"tradeUnitsPrecision"`
	Type                        string `json:"type"`
}

type AccountChanges struct {
	Changes           *Changes `json:"changes"`
	LastTransactionID string   `json:"lastTransactionID"`
	State             *State   `json:"state"`
}

type Changes struct {
	OrdersCancelled interface{} `json:"ordersCancelled"`
	OrdersCreated   interface{} `json:"ordersCreated"`
	OrdersFilled    interface{} `json:"ordersFilled"`
	OrdersTriggered interface{} `json:"ordersTriggered"`
	Positions       interface{} `json:"positions"`
	TradesClosed    interface{} `json:"tradesClosed"`
	TradesOpened    interface{} `json:"tradesOpened"`
	TradesReduced   interface{} `json:"tradesReduced"`
	Transactions    interface{} `json:"transactions"`
}

type State struct {
	NAV                        string      `json:"NAV"`
	MarginAvailable            string      `json:"marginAvailable"`
	MarginCallMarginUsed       string      `json:"marginCallMarginUsed"`
	MarginCallPercent          string      `json:"marginCallPercent"`
	MarginCloseoutMarginUsed   string      `json:"marginCloseoutMarginUsed"`
	MarginCloseoutNAV          string      `json:"marginCloseoutNAV"`
	MarginCloseoutPercent      string      `json:"marginCloseoutPercent"`
	MarginCloseoutUnrealizedPL string      `json:"marginCloseoutUnrealizedPL"`
	MarginUsed                 string      `json:"marginUsed"`
	Orders                     interface{} `json:"orders"`
	PositionValue              string      `json:"positionValue"`
	Positions                  interface{} `json:"positions"`
	Trades                     interface{} `json:"trades"`
	UnrealizedPL               string      `json:"unrealizedPL"`
	WithdrawalLimit            string      `json:"withdrawalLimit"`
}
