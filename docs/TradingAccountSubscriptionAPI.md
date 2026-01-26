# \TradingAccountSubscriptionAPI

All URIs are relative to *https://cadenza-api-uat.algo724.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListTradingAccountSubscriptions**](TradingAccountSubscriptionAPI.md#ListTradingAccountSubscriptions) | **Get** /api/v3/tradingAccount/subscription/list | List trading account subscriptions
[**SubscribeTradingAccount**](TradingAccountSubscriptionAPI.md#SubscribeTradingAccount) | **Post** /api/v3/tradingAccount/subscription/subscribe | Subscribe to trading account updates
[**UnsubscribeTradingAccount**](TradingAccountSubscriptionAPI.md#UnsubscribeTradingAccount) | **Post** /api/v3/tradingAccount/subscription/unsubscribe | Unsubscribe from trading account updates



## ListTradingAccountSubscriptions

> ListTradingAccountSubscriptions200Response ListTradingAccountSubscriptions(ctx).TradingAccountId(tradingAccountId).Limit(limit).Offset(offset).Cursor(cursor).Execute()

List trading account subscriptions



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
	limit := int32(100) // int32 | Limit the number of returned results (optional) (default to 50)
	offset := int32(0) // int32 | Offset of the returned results (optional) (default to 0)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountSubscriptionAPI.ListTradingAccountSubscriptions(context.Background()).TradingAccountId(tradingAccountId).Limit(limit).Offset(offset).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountSubscriptionAPI.ListTradingAccountSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTradingAccountSubscriptions`: ListTradingAccountSubscriptions200Response
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountSubscriptionAPI.ListTradingAccountSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTradingAccountSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tradingAccountId** | **string** | Trading account ID | 
 **limit** | **int32** | Limit the number of returned results | [default to 50]
 **offset** | **int32** | Offset of the returned results | [default to 0]
 **cursor** | **string** |  | 

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


## SubscribeTradingAccount

> SubscribeTradingAccount200Response SubscribeTradingAccount(ctx).SubscribeTradingAccountRequest(subscribeTradingAccountRequest).Execute()

Subscribe to trading account updates



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
	subscribeTradingAccountRequest := *openapiclient.NewSubscribeTradingAccountRequest("TradingAccountId_example") // SubscribeTradingAccountRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountSubscriptionAPI.SubscribeTradingAccount(context.Background()).SubscribeTradingAccountRequest(subscribeTradingAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountSubscriptionAPI.SubscribeTradingAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscribeTradingAccount`: SubscribeTradingAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountSubscriptionAPI.SubscribeTradingAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeTradingAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscribeTradingAccountRequest** | [**SubscribeTradingAccountRequest**](SubscribeTradingAccountRequest.md) |  | 

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


## UnsubscribeTradingAccount

> SubscribeTradingAccount200Response UnsubscribeTradingAccount(ctx).UnsubscribeRequest(unsubscribeRequest).Execute()

Unsubscribe from trading account updates



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
	resp, r, err := apiClient.TradingAccountSubscriptionAPI.UnsubscribeTradingAccount(context.Background()).UnsubscribeRequest(unsubscribeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountSubscriptionAPI.UnsubscribeTradingAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsubscribeTradingAccount`: SubscribeTradingAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountSubscriptionAPI.UnsubscribeTradingAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeTradingAccountRequest struct via the builder pattern


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

