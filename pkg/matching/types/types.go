package types

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type OrderType string
type OrderSide string
type TradeBy int

const (
	OrderTypeLimit          OrderType = "limit"
	OrderTypeMarket         OrderType = "market"
	OrderTypeMarketQuantity OrderType = "marketQty"
	OrderTypeMarketAmount   OrderType = "marketAmount"
	OrderSideBuy            OrderSide = "bid"
	OrderSideSell           OrderSide = "ask"
	TradeBySeller           TradeBy   = 1
	TradeByBuyer            TradeBy   = 2
)

func (os OrderSide) String() string {
	if os == OrderSideSell {
		return "ask"
	}
	return "bid"
}

func (ot OrderType) String() string {
	switch ot {
	case OrderTypeMarket:
		return "market"
	case OrderTypeMarketAmount:
		return "market_amount"
	case OrderTypeMarketQuantity:
		return "market_qty"
	default:
		return "limit"
	}
}

type TradeResult struct {
	Symbol                 string          `json:"symbol"`
	AskOrderId             string          `json:"ask"`
	BidOrderId             string          `json:"bid"`
	TradeQuantity          decimal.Decimal `json:"trade_quantity"`
	TradePrice             decimal.Decimal `json:"trade_price"`
	TradeBy                TradeBy         `json:"trade_by"`
	TradeTime              int64           `json:"trade_time"`                //纳秒级
	RemainderMarketOrderId string          `json:"remainder_market_order_id"` //市价订单标记，用于结算时取消市价单剩余未成交的部分
}

func (t *TradeResult) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(&t)
	return
}

func (t *TradeResult) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, t)
}

type RemoveResult struct {
	Symbol   string     `json:"symbol"`
	UniqueId string     `json:"unique_id"`
	Type     RemoveType `json:"type"`
}
