# Stock Market Data Functions

**IMPORTANT**: This implementation is not affiliated with or endorsed by Yahoo Finance. It is based on the open-source library [github.com/shoenig/yahoo-finance](https://github.com/shoenig/yahoo-finance).

## Overview

This package provides a set of functions to retrieve real-time and historical stock market data. It wraps the functionality provided by the `github.com/shoenig/yahoo-finance` Go package into individual, easy-to-use functions.

## Features

The package includes functions to retrieve:

- Current stock prices
- Previous closing prices
- Price changes (absolute and percentage)
- Exchange information
- Currency information
- Trading timestamps
- Instrument types

## Available Functions

### Price Information

- `YfStkPrice(sym string)` - Returns the current stock price
- `YfStkPrevClose(sym string)` - Returns the previous day's closing price
- `YfStkDelta(sym string)` - Returns the absolute price change
- `YfStkDeltaPerc(sym string)` - Returns the percentage price change
- `YfStkChange(sym string)` - Returns both absolute and percentage changes

### Market Information

- `YfStkSymbol(sym string)` - Returns the stock symbol
- `YfStkExchange(sym string)` - Returns the full exchange name
- `YfStkExchangeSymbol(sym string)` - Returns the exchange symbol
- `YfStkCurrency(sym string)` - Returns the trading currency

### Temporal Information

- `YfStkFirstTrade(sym string)` - Returns the first trading date/time
- `YfStkMarketTime(sym string)` - Returns the current market time

### Other Information

- `YfStkInstrument(sym string)` - Returns the instrument type (Note: Currently not implemented)

## Return Formats

All functions return strings in the following formats:

- Price values: `$XX.XX`
- Percentage values: `XX.XX%`
- Combined changes: `$XX.XX (XX.XX%)`
- Dates: `YYYY-MM-DD HH:mm:ss`
- Other values: Plain strings

## Error Handling

All functions include error handling and will return an error message in string format if the request fails or the symbol is invalid.

## Implementation Details

- Uses a custom HTTP client with a 10-second timeout
- All functions use the same client configuration through `getClient()`
- Built on top of the `github.com/shoenig/yahoo-finance` package

## Example Usage

```go
symbol := "AAPL"
fmt.Printf("Symbol: %s\n", YfStkSymbol(symbol))
fmt.Printf("Price: %s\n", YfStkPrice(symbol))
fmt.Printf("Previous Close: %s\n", YfStkPrevClose(symbol))
fmt.Printf("Change: %s\n", YfStkChange(symbol))
```

## Function Definitions

The package includes a JSON configuration file (`anthropicfunc.json`) that provides detailed function definitions. This file includes:
- Function names and descriptions
- Parameter specifications
- Return type information
- Return format specifications

Example JSON structure:
```json
{
  "functions": [
    {
      "name": "YfStkPrice",
      "description": "Gets the current stock price for a given symbol",
      "parameters": {
        "type": "object",
        "properties": {
          "sym": {
            "type": "string",
            "description": "A valid stock market symbol (e.g., \"AAPL\" for Apple Inc.)"
          }
        },
        "required": ["sym"]
      },
      "returnType": "string",
      "returnFormat": "$XX.XX"
    }
  ]
}
```

## Dependencies

- `github.com/shoenig/yahoo-finance`
- Standard Go libraries (`fmt`, `net/http`, `time`)

## Important Notes

1. This is not an official Yahoo Finance product
2. Data availability and accuracy depend on the underlying `github.com/shoenig/yahoo-finance` package
3. Respect rate limits and terms of service when using this package
4. The instrument type function is currently not implemented

## License

Please refer to the license terms of the `github.com/shoenig/yahoo-finance` package for usage rights and restrictions.
