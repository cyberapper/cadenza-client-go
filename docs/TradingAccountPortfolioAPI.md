# \TradingAccountPortfolioAPI

All URIs are relative to *https://cadenza-api-uat.algo724.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListTradingAccountPortfolios**](TradingAccountPortfolioAPI.md#ListTradingAccountPortfolios) | **Get** /api/v3/tradingAccount/portfolio/list | List trading account portfolios



## ListTradingAccountPortfolios

> ListTradingAccountPortfolios200Response ListTradingAccountPortfolios(ctx).TradingAccountId(tradingAccountId).Venue(venue).Currency(currency).Limit(limit).Offset(offset).Execute()

List trading account portfolios



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cyberapper/cadenza-client-go"
)

func main() {
	tradingAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Trading account ID (optional)
	venue := openapiclient.venue("") // Venue | Exchange type (optional)
	currency := "USDT" // string | Currency (optional)
	limit := int32(100) // int32 | Limit the number of returned results (optional) (default to 50)
	offset := int32(0) // int32 | Offset of the returned results (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountPortfolioAPI.ListTradingAccountPortfolios(context.Background()).TradingAccountId(tradingAccountId).Venue(venue).Currency(currency).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountPortfolioAPI.ListTradingAccountPortfolios``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTradingAccountPortfolios`: ListTradingAccountPortfolios200Response
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountPortfolioAPI.ListTradingAccountPortfolios`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTradingAccountPortfoliosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tradingAccountId** | **string** | Trading account ID | 
 **venue** | [**Venue**](Venue.md) | Exchange type | 
 **currency** | **string** | Currency | 
 **limit** | **int32** | Limit the number of returned results | [default to 50]
 **offset** | **int32** | Offset of the returned results | [default to 0]

### Return type

[**ListTradingAccountPortfolios200Response**](ListTradingAccountPortfolios200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

