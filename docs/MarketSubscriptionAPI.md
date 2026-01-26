# \MarketSubscriptionAPI

All URIs are relative to *https://cadenza-api-uat.algo724.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMarketOrderBookSubscriptions**](MarketSubscriptionAPI.md#ListMarketOrderBookSubscriptions) | **Get** /api/v3/market/subscription/orderBook/list | List order book subscriptions
[**SubscribeMarketOrderBook**](MarketSubscriptionAPI.md#SubscribeMarketOrderBook) | **Post** /api/v3/market/subscription/orderBook/subscribe | Subscribe to order book updates
[**UnsubscribeMarketOrderBook**](MarketSubscriptionAPI.md#UnsubscribeMarketOrderBook) | **Post** /api/v3/market/subscription/orderBook/unsubscribe | Unsubscribe from order book updates



## ListMarketOrderBookSubscriptions

> ListTradingAccountSubscriptions200Response ListMarketOrderBookSubscriptions(ctx).InstrumentId(instrumentId).Execute()

List order book subscriptions



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
	instrumentId := "instrumentId_example" // string | Instrument ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketSubscriptionAPI.ListMarketOrderBookSubscriptions(context.Background()).InstrumentId(instrumentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketSubscriptionAPI.ListMarketOrderBookSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMarketOrderBookSubscriptions`: ListTradingAccountSubscriptions200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketSubscriptionAPI.ListMarketOrderBookSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMarketOrderBookSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instrumentId** | **string** | Instrument ID | 

### Return type

[**ListTradingAccountSubscriptions200Response**](ListTradingAccountSubscriptions200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscribeMarketOrderBook

> SubscribeTradingAccount200Response SubscribeMarketOrderBook(ctx).SubscribeMarketOrderBookRequest(subscribeMarketOrderBookRequest).Execute()

Subscribe to order book updates



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
	subscribeMarketOrderBookRequest := *openapiclient.NewSubscribeMarketOrderBookRequest("BINANCE:BTC/USDT") // SubscribeMarketOrderBookRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketSubscriptionAPI.SubscribeMarketOrderBook(context.Background()).SubscribeMarketOrderBookRequest(subscribeMarketOrderBookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketSubscriptionAPI.SubscribeMarketOrderBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscribeMarketOrderBook`: SubscribeTradingAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketSubscriptionAPI.SubscribeMarketOrderBook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeMarketOrderBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscribeMarketOrderBookRequest** | [**SubscribeMarketOrderBookRequest**](SubscribeMarketOrderBookRequest.md) |  | 

### Return type

[**SubscribeTradingAccount200Response**](SubscribeTradingAccount200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsubscribeMarketOrderBook

> SubscribeTradingAccount200Response UnsubscribeMarketOrderBook(ctx).UnsubscribeRequest(unsubscribeRequest).Execute()

Unsubscribe from order book updates



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
	unsubscribeRequest := *openapiclient.NewUnsubscribeRequest() // UnsubscribeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketSubscriptionAPI.UnsubscribeMarketOrderBook(context.Background()).UnsubscribeRequest(unsubscribeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketSubscriptionAPI.UnsubscribeMarketOrderBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsubscribeMarketOrderBook`: SubscribeTradingAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketSubscriptionAPI.UnsubscribeMarketOrderBook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeMarketOrderBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unsubscribeRequest** | [**UnsubscribeRequest**](UnsubscribeRequest.md) |  | 

### Return type

[**SubscribeTradingAccount200Response**](SubscribeTradingAccount200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

