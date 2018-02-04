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
	v, err := strconv.ParseInt(strings.Split(action, "~")[0], 10, 8)
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

func getFieldAsDecimal(details []string, index int) decimal.Decimal {
	if len(details) <= index {
		return decimal.NewFromFloat(-1)
	}

	value, err := decimal.NewFromString(details[index])
	if err != nil {
		return decimal.NewFromFloat(-1)
	}

	return value
}

//{SubscriptionId}~{ExchangeName}~{CurrencySymbol}~{CurrencySymbol}~{Flag}~{TradeId}~{TimeStamp}~{Quantity}~{Price}~{Total}
func toTradeRecord(action string) *TradeRecord {
	details := strings.Split(action, "~")

	timeStr, err := strconv.ParseInt(details[6], 10, 64)
	time := time.Unix(timeStr, 0)
	if err != nil {
		return nil
	}

	quantity := getFieldAsDecimal(details, 7)

	price := getFieldAsDecimal(details, 8)

	totalPrice := getFieldAsDecimal(details, 9)

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

func toMarketRecord(action string) *MarketRecord {
	details := strings.Split(action, "~")

	dayOpen := getFieldAsDecimal(details, 17)
	dayHigh := getFieldAsDecimal(details, 18)
	dayLow := getFieldAsDecimal(details, 19)

	hourOpen := getFieldAsDecimal(details, 14)
	hourHigh := getFieldAsDecimal(details, 15)
	hourLow := getFieldAsDecimal(details, 16)

	CurrencyFrom := Currency{Name: details[2], Symbol: SYMBOL[details[2]]}
	CurrencyTo := Currency{Name: details[3], Symbol: SYMBOL[details[3]]}

	dayVolumeValueFrom := getFieldAsDecimal(details, 12)
	dayVolumeValueTo := getFieldAsDecimal(details, 13)

	hourVolumeValueFrom := getFieldAsDecimal(details, 10)
	hourVolumeValueTo := getFieldAsDecimal(details, 10)

	return &MarketRecord{
		Day: Market{
			Price: MarketPrice{
				Open: dayOpen,
				High: dayHigh,
				Low:  dayLow,
			},
			Volume: map[Currency]Price{
				CurrencyFrom: {
					Currency: CurrencyFrom,
					Value:    dayVolumeValueFrom,
				},
				CurrencyTo: {
					Currency: CurrencyTo,
					Value:    dayVolumeValueTo,
				},
			},
		},
		Hour: Market{
			Price: MarketPrice{
				Open: hourOpen,
				High: hourHigh,
				Low:  hourLow,
			},
			Volume: map[Currency]Price{
				CurrencyFrom: {
					Currency: CurrencyFrom,
					Value:    hourVolumeValueFrom,
				},
				CurrencyTo: {
					Currency: CurrencyTo,
					Value:    hourVolumeValueTo,
				},
			},
		},
	}
}
