package goanda

/////////////////////////////////////////////////////////////////////////////////////////////////////////
// ACCOUNT BELOW
/////////////////////////////////////////////////////////////////////////////////////////////////////////

type Accounts struct {
	Accounts []AccountsDetails `json:"accounts"`
}

type AccountsDetails struct {
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
	CreatedTime                 DateTime    `json:"createdTime"`
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
	Orders                      []Order     `json:"orders"`
	PendingOrderCount           int         `json:"pendingOrderCount"`
	Pl                          string      `json:"pl"`
	PositionValue               string      `json:"positionValue"`
	Positions                   []Position  `json:"positions"`
	ResettablePL                string      `json:"resettablePL"`
	Trades                      []Trade     `json:"trades"`
	UnrealizedPL                string      `json:"unrealizedPL"`
	WithdrawalLimit             string      `json:"withdrawalLimit"`
}

type AccountSummary struct {
	Details           *AccountSummaryDetails `json:"account"`
	LastTransactionID string                 `json:"lastTransactionID"`
}

type AccountSummaryDetails struct {
	NAV                         string `json:"NAV"`
	Alias                       string `json:"alias"`
	Balance                     string `json:"balance"`
	CreatedByUserID             int    `json:"createdByUserID"`
	CreatedTime                 DateTime `json:"createdTime"`
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
	OrdersCancelled []Order `json:"ordersCancelled"`
	OrdersCreated   []Order `json:"ordersCreated"`
	OrdersFilled    []Order `json:"ordersFilled"`
	OrdersTriggered []Order `json:"ordersTriggered"`
	Positions       []Position `json:"positions"`
	TradesClosed    []Trade `json:"tradesClosed"`
	TradesOpened    []Trade `json:"tradesOpened"`
	TradesReduced   []Trade `json:"tradesReduced"`
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
	Positions                  []Position  `json:"positions"`
	Trades                     []Trade     `json:"trades"`
	UnrealizedPL               string      `json:"unrealizedPL"`
	WithdrawalLimit            string      `json:"withdrawalLimit"`
}

//PATCH
type Configuration struct {
	Alias      string `json:"alias"`
	MarginRate string `json:"marginRate"`
}

type ConfigurationConfirmation struct {
	ClientConfigureTransaction *ClientConfigureTransaction `json:"clientConfigureTransaction"`
	LastTransactionID          string                      `json:"lastTransactionID"`
}

type ClientConfigureTransaction struct {
	AccountID  string `json:"accountID"`
	Alias      string `json:"alias"`
	BatchID    string `json:"batchID"`
	ID         string `json:"id"`
	MarginRate string `json:"marginRate"`
	Time       DateTime `json:"time"`
	Type       string `json:"type"`
	UserID     int    `json:"userID"`
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////
// INSTRUMENT BELOW
/////////////////////////////////////////////////////////////////////////////////////////////////////////

type Instrument struct {
	Candles     []*Candles `json:"candles"`
	Granularity string     `json:"granularity"`
	Instrument  string     `json:"instrument"`
}

type Candles struct {
	Complete bool   `json:"complete"`
	Ask      *OHLC   `json:"ask"`
	Bid      *OHLC   `json:"bid"`
	Mid      *OHLC   `json:"mid"`
	Time     DateTime `json:"time"`
	Volume   int    `json:"volume"`
}

type OHLC struct {
	C string `json:"c"`
	H string `json:"h"`
	L string `json:"l"`
	O string `json:"o"`
}

type ClientExtensions struct {
	Comment string     `json:"comment"`
	Id      string     `json:"id"`
	Tag     string     `json:"tag"`
}

type Order struct {
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
	CreatedTime        DateTime   `json:"createdTime"`
	Id                 int      `json:"id,string"`
	Instrument         string   `json:"instrument"`
	PartialFill        string   `json:"partialFill"`
	PositionFill       string   `json:"positionFill"`
	Price              float64  `json:"price,string"`
	ReplacesOrderId    int      `json:"replacesOrderId,string"`
	State              string   `json:"state"`
	TimeInForce        string   `json:"timeInForce"`
	TriggerCondition   string   `json:"triggerCondition"`
	Type               string   `json:"type"`
	Units              int      `json:"units,string"`
}

type Position struct {
	Instrument      string   `json:"instrument"`
	PL              float64  `json:"pl,string"`
	UnrealizedPL    float64  `json:"unrealizedPL,string"`
	ResetttablePL   float64  `json:"resettablePL,string"`
	Long            *PositionDetail `json:"long"`
	Short           *PositionDetail `json:"short"`
}

type PositionDetail struct {
	PL              float64  `json:"pl,string"`
	ResetttablePL   float64  `json:"resettablePL,string"`
	Units           int      `json:"units,string"`
	UnrealizedPL    float64  `json:"unrealizedPL,string"`
}

type Trade struct {
	CurrentUnits    int      `json:"currentUnits,string"`
	Financing       float64  `json:"financing,string"`
	Id              int      `json:"Id,string"`
	InitialUnits    int      `json:"initialUnits,string"`
	Instrument      string   `json:"instrument"`
	OpenTime        DateTime `json:"openTime"`
	Price           float64  `json:"price,string"`
	RealizedPL      float64  `json:"realizedPL,string"`
	State           string   `json:"state"`
	UnrealizedPL    float64  `json:"unrealizedPL,string"`
}

type TakeProfitDetails struct {
	Price              float64  `json:"price,string"`
	TimeInForce        string   `json:"timeInForce"`
	GtdTime            DateTime `json:"gtdTime"`
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
}

type StopLossDetails TakeProfitDetails

type TrailingStopLossDetails struct {
	Distance           float64  `json:"distance,string"`
	TimeInForce        string   `json:"timeInForce"`
	GtdTime            DateTime `json:"gtdTime"`
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
}


type Transaction struct {
	Id         string `json:"id"`
	Time       DateTime `json:"time"`
	UserID     int    `json:"userID"`
	AccountID  string `json:"accountID"`
	BatchID    string `json:"batchID"`
}

type OrderFillTransaction struct {
	Transaction
	Type               string   `json:"type"`
	OrderID   string `json:"orderID"`
	ClientOrderID   string `json:"clientOrderID"`
	Instrument  string     `json:"instrument"`
	Units              int      `json:"units,string"`
	Price              float64  `json:"price,string"`
	Reason    string `json:"reason"`
	PL        float64 `json:"pl,string"`
	Financing float64 `json:"financing,string"`
	AccountBalance float64 `json:"accountBalance,string"`
}

type OrderCancelTransaction struct {
	Transaction
	Type               string   `json:"type"`
	OrderID   string `json:"orderID"`
	ClientOrderID   string `json:"clientOrderID"`
	Reason    string `json:"reason"`
	ReplacedByOrderID   string `json:"replacedByOrderID"`
}

