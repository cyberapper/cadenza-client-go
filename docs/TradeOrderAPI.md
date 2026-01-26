# \TradeOrderAPI

All URIs are relative to *https://cadenza-api-uat.algo724.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelTradeOrder**](TradeOrderAPI.md#CancelTradeOrder) | **Post** /api/v3/tradeOrder/cancel | Cancel trade order
[**ListTradeOrders**](TradeOrderAPI.md#ListTradeOrders) | **Get** /api/v3/tradeOrder/list | List trade orders
[**SubmitTradeOrder**](TradeOrderAPI.md#SubmitTradeOrder) | **Post** /api/v3/tradeOrder/submit | Submit trade order



## CancelTradeOrder

> CancelTradeOrder200Response CancelTradeOrder(ctx).CancelTradeOrderRequest(cancelTradeOrderRequest).Execute()

Cancel trade order



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
	cancelTradeOrderRequest := *openapiclient.NewCancelTradeOrderRequest("TradingAccountId_example", "TradeOrderId_example") // CancelTradeOrderRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeOrderAPI.CancelTradeOrder(context.Background()).CancelTradeOrderRequest(cancelTradeOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeOrderAPI.CancelTradeOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelTradeOrder`: CancelTradeOrder200Response
	fmt.Fprintf(os.Stdout, "Response from `TradeOrderAPI.CancelTradeOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelTradeOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cancelTradeOrderRequest** | [**CancelTradeOrderRequest**](CancelTradeOrderRequest.md) |  | 

### Return type

[**CancelTradeOrder200Response**](CancelTradeOrder200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTradeOrders

> ListTradeOrders200Response ListTradeOrders(ctx).TradeOrderId(tradeOrderId).OrderStatus(orderStatus).TradingAccountId(tradingAccountId).InstrumentId(instrumentId).StartTime(startTime).EndTime(endTime).Limit(limit).Offset(offset).Cursor(cursor).Ascending(ascending).Execute()

List trade orders



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
	tradeOrderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Trade order ID (optional)
	orderStatus := openapiclient.orderStatus("") // OrderStatus | Order status (optional)
	tradingAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Trading account ID (optional)
	instrumentId := "instrumentId_example" // string | Instrument ID (optional)
	startTime := int32(1622505600000) // int32 | Start time (in unix milliseconds), of the created at field (optional)
	endTime := int32(56) // int32 | End time (in unix milliseconds), of the created at field (optional)
	limit := int32(100) // int32 | Limit the number of returned results (optional) (default to 50)
	offset := int32(0) // int32 | Offset of the returned results (optional) (default to 0)
	cursor := "cursor_example" // string |  (optional)
	ascending := true // bool | Return records in ascending order (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeOrderAPI.ListTradeOrders(context.Background()).TradeOrderId(tradeOrderId).OrderStatus(orderStatus).TradingAccountId(tradingAccountId).InstrumentId(instrumentId).StartTime(startTime).EndTime(endTime).Limit(limit).Offset(offset).Cursor(cursor).Ascending(ascending).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeOrderAPI.ListTradeOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTradeOrders`: ListTradeOrders200Response
	fmt.Fprintf(os.Stdout, "Response from `TradeOrderAPI.ListTradeOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTradeOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tradeOrderId** | **string** | Trade order ID | 
 **orderStatus** | [**OrderStatus**](OrderStatus.md) | Order status | 
 **tradingAccountId** | **string** | Trading account ID | 
 **instrumentId** | **string** | Instrument ID | 
 **startTime** | **int32** | Start time (in unix milliseconds), of the created at field | 
 **endTime** | **int32** | End time (in unix milliseconds), of the created at field | 
 **limit** | **int32** | Limit the number of returned results | [default to 50]
 **offset** | **int32** | Offset of the returned results | [default to 0]
 **cursor** | **string** |  | 
 **ascending** | **bool** | Return records in ascending order | [default to false]

### Return type

[**ListTradeOrders200Response**](ListTradeOrders200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitTradeOrder

> SubmitTradeOrder200Response SubmitTradeOrder(ctx).SubmitTradeOrderRequest(submitTradeOrderRequest).Execute()

Submit trade order



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
	submitTradeOrderRequest := *openapiclient.NewSubmitTradeOrderRequest("TradingAccountId_example", "BINANCE:BTC/USDT", openapiclient.orderSide("BUY"), openapiclient.orderType(""), "1234.56789000") // SubmitTradeOrderRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeOrderAPI.SubmitTradeOrder(context.Background()).SubmitTradeOrderRequest(submitTradeOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeOrderAPI.SubmitTradeOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitTradeOrder`: SubmitTradeOrder200Response
	fmt.Fprintf(os.Stdout, "Response from `TradeOrderAPI.SubmitTradeOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitTradeOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitTradeOrderRequest** | [**SubmitTradeOrderRequest**](SubmitTradeOrderRequest.md) |  | 

### Return type

[**SubmitTradeOrder200Response**](SubmitTradeOrder200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

