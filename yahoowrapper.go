package yahoowrapper
//package main

import (
	"fmt"
	//"anthropicfunc/yahoofinance"
	yahoofinance "github.com/shoenig/yahoo-finance"
	"net/http"
	"time"
)

func getClient() yahoofinance.Client {
	customClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	return yahoofinance.New(&yahoofinance.Options{
		HTTPClient: customClient,
	})
}

// Description: Gets the current stock price for a given symbol
// Name: YfStkPrice
// Parameter: sym - A valid stock market symbol (e.g., "AAPL" for Apple Inc.)
func YfStkPrice(sym string) string {
	yclient := getClient()
	c, err := yclient.Lookup(sym)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return fmt.Sprintf("$%.2f", c.Price())
}

// Description: Retrieves the stock symbol, useful for verification or display purposes
// Name: YfStkSymbol
// Parameter: sym - A valid stock market symbol (e.g., "MSFT" for Microsoft)
func YfStkSymbol(sym string) string {
	yclient := getClient()
	c, err := yclient.Lookup(sym)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return c.Symbol()
}

// Description: Gets the previous day's closing price for a stock
// Name: YfStkPrevClose
// Parameter: sym - A valid stock market symbol (e.g., "GOOGL" for Google)
func YfStkPrevClose(sym string) string {
	yclient := getClient()
	c, err := yclient.Lookup(sym)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return fmt.Sprintf("$%.2f", c.PrevClosePrice())
}

// Description: Gets the absolute price change (in dollars) from previous close
// Name: YfStkDelta
// Parameter: sym - A valid stock market symbol (e.g., "AMZN" for Amazon)
func YfStkDelta(sym string) string {
	yclient := getClient()
	c, err := yclient.Lookup(sym)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return fmt.Sprintf("$%.2f", c.Delta())
}

// Description: Gets the percentage price change from previous close
// Name: YfStkDeltaPerc
// Parameter: sym - A valid stock market symbol (e.g., "META" for Meta)
func YfStkDeltaPerc(sym string) string {
	yclient := getClient()
	c, err := yclient.Lookup(sym)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return fmt.Sprintf("%.2f%%", c.DeltaPerc())
}

// Description: Gets both the price change and percentage change in a combined format
// Name: YfStkChange
// Parameter: sym - A valid stock market symbol (e.g., "TSLA" for Tesla)
func YfStkChange(sym string) string {
	yclient := getClient()
	c, err := yclient.Lookup(sym)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return fmt.Sprintf("$%.2f (%.2f%%)", c.Delta(), c.DeltaPerc())
}

// Description: Gets the full name of the exchange where the stock is traded
// Name: YfStkExchange
// Parameter: sym - A valid stock market symbol (e.g., "NFLX" for Netflix)
func YfStkExchange(sym string) string {
	yclient := getClient()
	c, err := yclient.Lookup(sym)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return c.Exchange(false)
}

// Description: Gets the symbol/abbreviation of the exchange where the stock is traded
// Name: YfStkExchangeSymbol
// Parameter: sym - A valid stock market symbol (e.g., "IBM" for IBM)
func YfStkExchangeSymbol(sym string) string {
	yclient := getClient()
	c, err := yclient.Lookup(sym)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return c.Exchange(true)
}

// Description: Gets the currency in which the stock is traded
// Name: YfStkCurrency
// Parameter: sym - A valid stock market symbol (e.g., "NVDA" for NVIDIA)
func YfStkCurrency(sym string) string {
	yclient := getClient()
	c, err := yclient.Lookup(sym)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return c.Currency()
}

// Description: Gets the date and time when the stock was first traded
// Name: YfStkFirstTrade
// Parameter: sym - A valid stock market symbol (e.g., "AMD" for Advanced Micro Devices)
func YfStkFirstTrade(sym string) string {
	yclient := getClient()
	c, err := yclient.Lookup(sym)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return c.FirstTrade().Format("2006-01-02 15:04:05")
}

// Description: Gets the current market time for the stock
// Name: YfStkMarketTime
// Parameter: sym - A valid stock market symbol (e.g., "INTC" for Intel)
func YfStkMarketTime(sym string) string {
	yclient := getClient()
	c, err := yclient.Lookup(sym)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return c.MarketTime().Format("2006-01-02 15:04:05")
}

// Description: Gets the type of financial instrument (e.g., stock, ETF, etc.)
// Name: YfStkInstrument
// Parameter: sym - A valid stock market symbol (e.g., "SPY" for S&P 500 ETF)
func YfStkInstrument(sym string) string {
	return "Not IMplemented yet"
	yclient := getClient()
	c, err := yclient.Lookup(sym)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return c.InstrumentType()
}

func Testmain() {
	// Example usage of all functions
	symbol := "AAPL"
	fmt.Printf("Symbol: %s\n", YfStkSymbol(symbol))
	fmt.Printf("Price: %s\n", YfStkPrice(symbol))
	fmt.Printf("Previous Close: %s\n", YfStkPrevClose(symbol))
	fmt.Printf("Change: %s\n", YfStkChange(symbol))
	fmt.Printf("Delta: %s\n", YfStkDelta(symbol))
	fmt.Printf("Delta Percentage: %s\n", YfStkDeltaPerc(symbol))
	fmt.Printf("Exchange: %s\n", YfStkExchange(symbol))
	fmt.Printf("Exchange Symbol: %s\n", YfStkExchangeSymbol(symbol))
	fmt.Printf("Currency: %s\n", YfStkCurrency(symbol))
	fmt.Printf("First Trade: %s\n", YfStkFirstTrade(symbol))
	fmt.Printf("Market Time: %s\n", YfStkMarketTime(symbol))
	fmt.Println("")
	fmt.Printf("Instrument Type: %s\n", YfStkInstrument(symbol))
}
