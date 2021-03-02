package bitcoin

import (
	"encoding/json"
	"errors"
	"example/currencies"
	"net/http"
)

const (
	baseURL = "https://api.coindesk.com/v1/bpi/currentprice/"
)

var (
	errInvalidCurrency = errors.New("invalid currency")
)

type currencyDetails struct {
	Code        string  `json:"code"`
	Rate        string  `json:"rate"`
	Description string  `json:"description"`
	RateFloat   float64 `json:"rate_float"`
}

type bpiModel struct {
	AUD currencyDetails `json:"AUD,omitempty"`
	BRL currencyDetails `json:"BRL,omitempty"`
	CAD currencyDetails `json:"CAD,omitempty"`
	CNY currencyDetails `json:"CNY,omitempty"`
	COP currencyDetails `json:"COP,omitempty"`
	EUR currencyDetails `json:"EUR,omitempty"`
	HKD currencyDetails `json:"HKD,omitempty"`
	JPY currencyDetails `json:"JPY,omitempty"`
	PEN currencyDetails `json:"PEN,omitempty"`
	USD currencyDetails `json:"USD,omitempty"`
}

// BTC bitcoin representation
type BTC struct {
	Bpi      bpiModel            `json:"bpi"`
	Currency currencies.Currency `json:"Currency"`
}

func getBitcoin(currency currencies.Currency) (*BTC, error) {
	var btc = new(BTC)

	response, err := http.Get(baseURL + currency + ".json")
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(response.Body).Decode(btc)
	if err != nil {
		return nil, err
	}

	btc.Currency = currency

	return btc, nil
}

func (btc *BTC) profitPercentage(amount int) (float64, error) {
	var price float64

	switch btc.currency {
	case currencies.CurrencyAUD:
		price = btc.Bpi.AUD.RateFloat
	case currencies.CurrencyBRL:
		price = btc.Bpi.BRL.RateFloat
	case currencies.CurrencyCAD:
		price = btc.Bpi.CAD.RateFloat
	case currencies.CurrencyCNY:
		price = btc.Bpi.CNY.RateFloat
	case currencies.CurrencyCOP:
		price = btc.Bpi.COP.RateFloat
	case currencies.CurrencyEUR:
		price = btc.Bpi.EUR.RateFloat
	case currencies.CurrencyHKD:
		price = btc.Bpi.HKD.RateFloat
	case currencies.CurrencyJPY:
		price = btc.Bpi.JPY.RateFloat
	case currencies.CurrencyPEN:
		price = btc.Bpi.PEN.RateFloat
	case currencies.CurrencyUSD:
		price = btc.Bpi.USD.RateFloat
	default:
		return 0, errInvalidCurrency
	}

	return (float64(amount) / price * 100) - 100, nil
}
