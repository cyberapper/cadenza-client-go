# \MarketAPI

All URIs are relative to *https://cadenza-api-uat.algo724.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketInstrument**](MarketAPI.md#DeleteMarketInstrument) | **Delete** /api/v3/market/instrument/delete | Delete market instrument
[**DeleteMarketSecurity**](MarketAPI.md#DeleteMarketSecurity) | **Delete** /api/v3/market/security/delete | Delete market security
[**DisableMarketInstrument**](MarketAPI.md#DisableMarketInstrument) | **Post** /api/v3/market/instrument/disable | Disable market instrument
[**EnableMarketInstrument**](MarketAPI.md#EnableMarketInstrument) | **Post** /api/v3/market/instrument/enable | Enable market instrument
[**GetMarketOrderBook**](MarketAPI.md#GetMarketOrderBook) | **Get** /api/v3/market/orderBook/get | Get market order book
[**ListMarketInstruments**](MarketAPI.md#ListMarketInstruments) | **Get** /api/v3/market/instrument/list | List market instruments
[**ListMarketOrderBooks**](MarketAPI.md#ListMarketOrderBooks) | **Get** /api/v3/market/orderBook/list | List market order books
[**ListMarketSecurities**](MarketAPI.md#ListMarketSecurities) | **Get** /api/v3/market/security/list | List market securities
[**ListMarketVenues**](MarketAPI.md#ListMarketVenues) | **Get** /api/v3/market/venue/list | List market venues
[**SyncMarketInstruments**](MarketAPI.md#SyncMarketInstruments) | **Post** /api/v3/market/instrument/sync | Sync market instruments
[**SyncMarketSecurities**](MarketAPI.md#SyncMarketSecurities) | **Post** /api/v3/market/security/sync | Sync market securities



## DeleteMarketInstrument

> DeleteMarketInstrument200Response DeleteMarketInstrument(ctx).DeleteMarketInstrumentRequest(deleteMarketInstrumentRequest).Execute()

Delete market instrument



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
	deleteMarketInstrumentRequest := *openapiclient.NewDeleteMarketInstrumentRequest("BINANCE:BTC/USDT") // DeleteMarketInstrumentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.DeleteMarketInstrument(context.Background()).DeleteMarketInstrumentRequest(deleteMarketInstrumentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.DeleteMarketInstrument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMarketInstrument`: DeleteMarketInstrument200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.DeleteMarketInstrument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketInstrumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteMarketInstrumentRequest** | [**DeleteMarketInstrumentRequest**](DeleteMarketInstrumentRequest.md) |  | 

### Return type

[**DeleteMarketInstrument200Response**](DeleteMarketInstrument200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMarketSecurity

> DeleteMarketSecurity200Response DeleteMarketSecurity(ctx).DeleteMarketSecurityRequest(deleteMarketSecurityRequest).Execute()

Delete market security



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
	deleteMarketSecurityRequest := *openapiclient.NewDeleteMarketSecurityRequest("BINANCE:BTC") // DeleteMarketSecurityRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.DeleteMarketSecurity(context.Background()).DeleteMarketSecurityRequest(deleteMarketSecurityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.DeleteMarketSecurity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMarketSecurity`: DeleteMarketSecurity200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.DeleteMarketSecurity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteMarketSecurityRequest** | [**DeleteMarketSecurityRequest**](DeleteMarketSecurityRequest.md) |  | 

### Return type

[**DeleteMarketSecurity200Response**](DeleteMarketSecurity200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableMarketInstrument

> EnableMarketInstrument200Response DisableMarketInstrument(ctx).DisableMarketInstrumentRequest(disableMarketInstrumentRequest).Execute()

Disable market instrument



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
	disableMarketInstrumentRequest := *openapiclient.NewDisableMarketInstrumentRequest("BINANCE:BTC/USDT") // DisableMarketInstrumentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.DisableMarketInstrument(context.Background()).DisableMarketInstrumentRequest(disableMarketInstrumentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.DisableMarketInstrument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableMarketInstrument`: EnableMarketInstrument200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.DisableMarketInstrument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableMarketInstrumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableMarketInstrumentRequest** | [**DisableMarketInstrumentRequest**](DisableMarketInstrumentRequest.md) |  | 

### Return type

[**EnableMarketInstrument200Response**](EnableMarketInstrument200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableMarketInstrument

> EnableMarketInstrument200Response EnableMarketInstrument(ctx).EnableMarketInstrumentRequest(enableMarketInstrumentRequest).Execute()

Enable market instrument



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
	enableMarketInstrumentRequest := *openapiclient.NewEnableMarketInstrumentRequest("BINANCE:BTC/USDT") // EnableMarketInstrumentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.EnableMarketInstrument(context.Background()).EnableMarketInstrumentRequest(enableMarketInstrumentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.EnableMarketInstrument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableMarketInstrument`: EnableMarketInstrument200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.EnableMarketInstrument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableMarketInstrumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableMarketInstrumentRequest** | [**EnableMarketInstrumentRequest**](EnableMarketInstrumentRequest.md) |  | 

### Return type

[**EnableMarketInstrument200Response**](EnableMarketInstrument200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketOrderBook

> GetMarketOrderBook200Response GetMarketOrderBook(ctx).InstrumentId(instrumentId).Venue(venue).Symbol(symbol).Depth(depth).Execute()

Get market order book



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
	venue := openapiclient.venue("") // Venue | Exchange type (optional)
	symbol := "BTC/USDT" // string | Instrument Symbol (optional)
	depth := int32(56) // int32 | Order book depth (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.GetMarketOrderBook(context.Background()).InstrumentId(instrumentId).Venue(venue).Symbol(symbol).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.GetMarketOrderBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketOrderBook`: GetMarketOrderBook200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.GetMarketOrderBook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketOrderBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instrumentId** | **string** | Instrument ID | 
 **venue** | [**Venue**](Venue.md) | Exchange type | 
 **symbol** | **string** | Instrument Symbol | 
 **depth** | **int32** | Order book depth | [default to 10]

### Return type

[**GetMarketOrderBook200Response**](GetMarketOrderBook200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMarketInstruments

> ListMarketInstruments200Response ListMarketInstruments(ctx).Venue(venue).Symbols(symbols).SecurityType(securityType).InstrumentStatus(instrumentStatus).Limit(limit).Offset(offset).Execute()

List market instruments



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
	venue := openapiclient.venue("") // Venue | Exchange type (optional)
	symbols := []string{"BTC/USDT"} // []string | Instrument Symbols array (optional)
	securityType := openapiclient.securityType("") // SecurityType | Security type (optional)
	instrumentStatus := openapiclient.instrumentStatus("ACTIVE") // InstrumentStatus | Instrument status (optional)
	limit := int32(100) // int32 | Limit the number of returned results (optional) (default to 50)
	offset := int32(0) // int32 | Offset of the returned results (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.ListMarketInstruments(context.Background()).Venue(venue).Symbols(symbols).SecurityType(securityType).InstrumentStatus(instrumentStatus).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.ListMarketInstruments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMarketInstruments`: ListMarketInstruments200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.ListMarketInstruments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMarketInstrumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **venue** | [**Venue**](Venue.md) | Exchange type | 
 **symbols** | **[]string** | Instrument Symbols array | 
 **securityType** | [**SecurityType**](SecurityType.md) | Security type | 
 **instrumentStatus** | [**InstrumentStatus**](InstrumentStatus.md) | Instrument status | 
 **limit** | **int32** | Limit the number of returned results | [default to 50]
 **offset** | **int32** | Offset of the returned results | [default to 0]

### Return type

[**ListMarketInstruments200Response**](ListMarketInstruments200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMarketOrderBooks

> ListMarketOrderBooks200Response ListMarketOrderBooks(ctx).InstrumentIds(instrumentIds).Venue(venue).Symbols(symbols).Depth(depth).Execute()

List market order books



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
	instrumentIds := []string{"BINANCE:BTC/USDT"} // []string |  (optional)
	venue := openapiclient.venue("") // Venue | Exchange type (optional)
	symbols := []string{"BTC/USDT"} // []string | Instrument Symbols array (optional)
	depth := int32(56) // int32 | Order book depth (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.ListMarketOrderBooks(context.Background()).InstrumentIds(instrumentIds).Venue(venue).Symbols(symbols).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.ListMarketOrderBooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMarketOrderBooks`: ListMarketOrderBooks200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.ListMarketOrderBooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMarketOrderBooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instrumentIds** | **[]string** |  | 
 **venue** | [**Venue**](Venue.md) | Exchange type | 
 **symbols** | **[]string** | Instrument Symbols array | 
 **depth** | **int32** | Order book depth | [default to 10]

### Return type

[**ListMarketOrderBooks200Response**](ListMarketOrderBooks200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMarketSecurities

> ListMarketSecurities200Response ListMarketSecurities(ctx).Venue(venue).SecurityId(securityId).Limit(limit).Offset(offset).Execute()

List market securities



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
	venue := openapiclient.venue("") // Venue | Exchange type (optional)
	securityId := "securityId_example" // string | Security ID (optional)
	limit := int32(100) // int32 | Limit the number of returned results (optional) (default to 50)
	offset := int32(0) // int32 | Offset of the returned results (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.ListMarketSecurities(context.Background()).Venue(venue).SecurityId(securityId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.ListMarketSecurities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMarketSecurities`: ListMarketSecurities200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.ListMarketSecurities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMarketSecuritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **venue** | [**Venue**](Venue.md) | Exchange type | 
 **securityId** | **string** | Security ID | 
 **limit** | **int32** | Limit the number of returned results | [default to 50]
 **offset** | **int32** | Offset of the returned results | [default to 0]

### Return type

[**ListMarketSecurities200Response**](ListMarketSecurities200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMarketVenues

> ListMarketVenues200Response ListMarketVenues(ctx).Execute()

List market venues



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.ListMarketVenues(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.ListMarketVenues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMarketVenues`: ListMarketVenues200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.ListMarketVenues`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMarketVenuesRequest struct via the builder pattern


### Return type

[**ListMarketVenues200Response**](ListMarketVenues200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncMarketInstruments

> SyncMarketInstruments200Response SyncMarketInstruments(ctx).SyncMarketInstrumentsRequest(syncMarketInstrumentsRequest).Execute()

Sync market instruments



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
	syncMarketInstrumentsRequest := *openapiclient.NewSyncMarketInstrumentsRequest(openapiclient.venue(""), []string{"Symbols_example"}) // SyncMarketInstrumentsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.SyncMarketInstruments(context.Background()).SyncMarketInstrumentsRequest(syncMarketInstrumentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.SyncMarketInstruments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncMarketInstruments`: SyncMarketInstruments200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.SyncMarketInstruments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncMarketInstrumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncMarketInstrumentsRequest** | [**SyncMarketInstrumentsRequest**](SyncMarketInstrumentsRequest.md) |  | 

### Return type

[**SyncMarketInstruments200Response**](SyncMarketInstruments200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncMarketSecurities

> DeleteMarketInstrument200Response SyncMarketSecurities(ctx).SyncMarketSecuritiesRequest(syncMarketSecuritiesRequest).Execute()

Sync market securities



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
	syncMarketSecuritiesRequest := *openapiclient.NewSyncMarketSecuritiesRequest(openapiclient.venue("")) // SyncMarketSecuritiesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.SyncMarketSecurities(context.Background()).SyncMarketSecuritiesRequest(syncMarketSecuritiesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.SyncMarketSecurities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncMarketSecurities`: DeleteMarketInstrument200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.SyncMarketSecurities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncMarketSecuritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncMarketSecuritiesRequest** | [**SyncMarketSecuritiesRequest**](SyncMarketSecuritiesRequest.md) |  | 

### Return type

[**DeleteMarketInstrument200Response**](DeleteMarketInstrument200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

