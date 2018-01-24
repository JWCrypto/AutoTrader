package CrawlerEngine

var (
	TYPE = map[string]int{
		"TRADE":                0,
		"FEEDNEWS":             1,
		"CURRENT":              2,
		"LOADCOMPLATE":         3,
		"COINPAIRS":            4,
		"CURRENTAGG":           5,
		"TOPLIST":              6,
		"TOPLISTCHANGE":        7,
		"ORDERBOOK":            8,
		"FULLORDERBOOK":        9,
		"ACTIVATION":           10,
		"TRADECATCHUP":         100,
		"NEWSCATCHUP":          101,
		"TRADECATCHUPCOMPLETE": 300,
		"NEWSCATCHUPCOMPLETE":  301,
	}

	SYMBOL = map[string]string{
		"BTC":  "Ƀ",
		"XRP":  "℞",
		"LTC":  "Ł",
		"DAO":  "Ð",
		"USD":  "$",
		"CNY":  "¥",
		"EUR":  "€",
		"GBP":  "£",
		"JPY":  "¥",
		"PLN":  "zł",
		"RUB":  "₽",
		"ETH":  "Ξ",
		"GOLD": "Gold g",
		"INR":  "₹",
		"BRL":  "R$",
	}
)
