package CrawlerEngine

import (
	"strings"
	"strconv"
	"github.com/shopspring/decimal"
	"time"
)

func validSubsResp(resp SubsResp, to Currency) bool {
	return len(resp[to.Name].Trades) > 0 && len(resp[to.Name].Current) > 0 && resp[to.Name].CurrentAgg != ""
}

func getType(action string) int {
	v, err := strconv.ParseInt(strings.Split(action, "~")[0], 10, 1)
	if err != nil {
		return -1
	}
	return int(v)
}

func extractTradeAction(flag string) TradeAction {
	switch flag {
	case "1":
		return BUY
	case "2":
		return SELL
	case "4":
	default:
		break
	}
	return UNKNOWN
}

//{SubscriptionId}~{ExchangeName}~{CurrencySymbol}~{CurrencySymbol}~{Flag}~{TradeId}~{TimeStamp}~{Quantity}~{Price}~{Total}
func toTradeRecord(action string) *TradeRecord {
	details := strings.Split(action, "~")

	timeStr, err := strconv.ParseInt(details[6], 10, 64)
	time := time.Unix(timeStr, 0)
	if err != nil {
		return nil
	}

	quantity, err := decimal.NewFromString(details[7])
	if err != nil {
		return nil
	}
	price, err := decimal.NewFromString(details[8])
	if err != nil {
		return nil
	}

	totalPrice, err := decimal.NewFromString(details[9])
	if err != nil {
		return nil
	}

	return &TradeRecord{
		Market:    details[1],
		Id:        details[5],
		Action:    extractTradeAction(details[4]),
		Timestamp: time,
		Price: Price{
			Currency: Currency{Name: details[2], Symbol: SYMBOL[details[2]]},
			Value:    price,
		},
		Quantity: quantity,
		TotalPrice: Price{
			Currency: Currency{Name: details[2], Symbol: SYMBOL[details[2]]},
			Value:    totalPrice,
		},
	}
}
