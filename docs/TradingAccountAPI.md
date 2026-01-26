# \TradingAccountAPI

All URIs are relative to *https://cadenza-api-uat.algo724.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectTradingAccount**](TradingAccountAPI.md#ConnectTradingAccount) | **Post** /api/v3/tradingAccount/connect | Connect trading account
[**DisableTradingAccount**](TradingAccountAPI.md#DisableTradingAccount) | **Post** /api/v3/tradingAccount/disable | Disable trading account
[**DisconnectTradingAccount**](TradingAccountAPI.md#DisconnectTradingAccount) | **Post** /api/v3/tradingAccount/disconnect | Disconnect trading account
[**EnableTradingAccount**](TradingAccountAPI.md#EnableTradingAccount) | **Post** /api/v3/tradingAccount/enable | Enable trading account
[**ListTradingAccountOperations**](TradingAccountAPI.md#ListTradingAccountOperations) | **Get** /api/v3/tradingAccount/operation/list | List trading account operations
[**ListTradingAccounts**](TradingAccountAPI.md#ListTradingAccounts) | **Get** /api/v3/tradingAccount/list | List trading accounts
[**UpdateTradingAccount**](TradingAccountAPI.md#UpdateTradingAccount) | **Post** /api/v3/tradingAccount/update | Update trading account



## ConnectTradingAccount

> ConnectTradingAccount200Response ConnectTradingAccount(ctx).ConnectTradingAccountRequest(connectTradingAccountRequest).Execute()

Connect trading account



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
	connectTradingAccountRequest := *openapiclient.NewConnectTradingAccountRequest([]string{"CredentialIds_example"}, "ExternalTradingAccountId_example") // ConnectTradingAccountRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.ConnectTradingAccount(context.Background()).ConnectTradingAccountRequest(connectTradingAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.ConnectTradingAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectTradingAccount`: ConnectTradingAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.ConnectTradingAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectTradingAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectTradingAccountRequest** | [**ConnectTradingAccountRequest**](ConnectTradingAccountRequest.md) |  | 

### Return type

[**ConnectTradingAccount200Response**](ConnectTradingAccount200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableTradingAccount

> ConnectTradingAccount200Response DisableTradingAccount(ctx).DisableTradingAccountRequest(disableTradingAccountRequest).Execute()

Disable trading account



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
	disableTradingAccountRequest := *openapiclient.NewDisableTradingAccountRequest("TradingAccountId_example") // DisableTradingAccountRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.DisableTradingAccount(context.Background()).DisableTradingAccountRequest(disableTradingAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.DisableTradingAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableTradingAccount`: ConnectTradingAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.DisableTradingAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableTradingAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableTradingAccountRequest** | [**DisableTradingAccountRequest**](DisableTradingAccountRequest.md) |  | 

### Return type

[**ConnectTradingAccount200Response**](ConnectTradingAccount200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisconnectTradingAccount

> ConnectTradingAccount200Response DisconnectTradingAccount(ctx).DisconnectTradingAccountRequest(disconnectTradingAccountRequest).Execute()

Disconnect trading account



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
	disconnectTradingAccountRequest := *openapiclient.NewDisconnectTradingAccountRequest("TradingAccountId_example") // DisconnectTradingAccountRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.DisconnectTradingAccount(context.Background()).DisconnectTradingAccountRequest(disconnectTradingAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.DisconnectTradingAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisconnectTradingAccount`: ConnectTradingAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.DisconnectTradingAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectTradingAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disconnectTradingAccountRequest** | [**DisconnectTradingAccountRequest**](DisconnectTradingAccountRequest.md) |  | 

### Return type

[**ConnectTradingAccount200Response**](ConnectTradingAccount200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableTradingAccount

> ConnectTradingAccount200Response EnableTradingAccount(ctx).EnableTradingAccountRequest(enableTradingAccountRequest).Execute()

Enable trading account



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
	enableTradingAccountRequest := *openapiclient.NewEnableTradingAccountRequest("TradingAccountId_example") // EnableTradingAccountRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.EnableTradingAccount(context.Background()).EnableTradingAccountRequest(enableTradingAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.EnableTradingAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableTradingAccount`: ConnectTradingAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.EnableTradingAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableTradingAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableTradingAccountRequest** | [**EnableTradingAccountRequest**](EnableTradingAccountRequest.md) |  | 

### Return type

[**ConnectTradingAccount200Response**](ConnectTradingAccount200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTradingAccountOperations

> ListTradingAccountOperations200Response ListTradingAccountOperations(ctx).TradingAccountId(tradingAccountId).OperationType(operationType).StartTime(startTime).EndTime(endTime).Limit(limit).Offset(offset).Cursor(cursor).Execute()

List trading account operations



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
	operationType := openapiclient.operationType("CONNECT") // OperationType | Operation type (optional)
	startTime := int32(1622505600000) // int32 | Start time (in unix milliseconds), of the created at field (optional)
	endTime := int32(56) // int32 | End time (in unix milliseconds), of the created at field (optional)
	limit := int32(100) // int32 | Limit the number of returned results (optional) (default to 50)
	offset := int32(0) // int32 | Offset of the returned results (optional) (default to 0)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.ListTradingAccountOperations(context.Background()).TradingAccountId(tradingAccountId).OperationType(operationType).StartTime(startTime).EndTime(endTime).Limit(limit).Offset(offset).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.ListTradingAccountOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTradingAccountOperations`: ListTradingAccountOperations200Response
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.ListTradingAccountOperations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTradingAccountOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tradingAccountId** | **string** | Trading account ID | 
 **operationType** | [**OperationType**](OperationType.md) | Operation type | 
 **startTime** | **int32** | Start time (in unix milliseconds), of the created at field | 
 **endTime** | **int32** | End time (in unix milliseconds), of the created at field | 
 **limit** | **int32** | Limit the number of returned results | [default to 50]
 **offset** | **int32** | Offset of the returned results | [default to 0]
 **cursor** | **string** |  | 

### Return type

[**ListTradingAccountOperations200Response**](ListTradingAccountOperations200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTradingAccounts

> ListTradingAccounts200Response ListTradingAccounts(ctx).TradingAccountId(tradingAccountId).Venue(venue).AccountStatus(accountStatus).UserId(userId).TenantId(tenantId).Limit(limit).Offset(offset).Execute()

List trading accounts



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
	accountStatus := openapiclient.tradingAccountStatus("SETUP") // TradingAccountStatus | Account status (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by user ID (optional)
	tenantId := "tenantId_example" // string | Filter by tenant ID (optional)
	limit := int32(100) // int32 | Limit the number of returned results (optional) (default to 50)
	offset := int32(0) // int32 | Offset of the returned results (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.ListTradingAccounts(context.Background()).TradingAccountId(tradingAccountId).Venue(venue).AccountStatus(accountStatus).UserId(userId).TenantId(tenantId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.ListTradingAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTradingAccounts`: ListTradingAccounts200Response
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.ListTradingAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTradingAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tradingAccountId** | **string** | Trading account ID | 
 **venue** | [**Venue**](Venue.md) | Exchange type | 
 **accountStatus** | [**TradingAccountStatus**](TradingAccountStatus.md) | Account status | 
 **userId** | **string** | Filter by user ID | 
 **tenantId** | **string** | Filter by tenant ID | 
 **limit** | **int32** | Limit the number of returned results | [default to 50]
 **offset** | **int32** | Offset of the returned results | [default to 0]

### Return type

[**ListTradingAccounts200Response**](ListTradingAccounts200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTradingAccount

> ConnectTradingAccount200Response UpdateTradingAccount(ctx).UpdateTradingAccountRequest(updateTradingAccountRequest).Execute()

Update trading account



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
	updateTradingAccountRequest := *openapiclient.NewUpdateTradingAccountRequest("TradingAccountId_example") // UpdateTradingAccountRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.UpdateTradingAccount(context.Background()).UpdateTradingAccountRequest(updateTradingAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.UpdateTradingAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTradingAccount`: ConnectTradingAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.UpdateTradingAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTradingAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTradingAccountRequest** | [**UpdateTradingAccountRequest**](UpdateTradingAccountRequest.md) |  | 

### Return type

[**ConnectTradingAccount200Response**](ConnectTradingAccount200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

