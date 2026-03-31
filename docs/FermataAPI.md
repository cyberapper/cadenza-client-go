# \FermataAPI

All URIs are relative to *https://cadenza-api-uat.algo724.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FermataCloseWallet**](FermataAPI.md#FermataCloseWallet) | **Post** /api/v3/fermata/wallet/close | Close wallet
[**FermataCreateWallet**](FermataAPI.md#FermataCreateWallet) | **Post** /api/v3/fermata/wallet/create | Create wallet
[**FermataFreezeWallet**](FermataAPI.md#FermataFreezeWallet) | **Post** /api/v3/fermata/wallet/freeze | Freeze wallet
[**FermataLinkDealer**](FermataAPI.md#FermataLinkDealer) | **Post** /api/v3/fermata/dealer/link | Link exchange account to dealer
[**FermataListWallets**](FermataAPI.md#FermataListWallets) | **Get** /api/v3/fermata/wallet/list | List wallets
[**FermataTransferWallet**](FermataAPI.md#FermataTransferWallet) | **Post** /api/v3/fermata/wallet/transfer | Transfer between wallets
[**FermataUnfreezeWallet**](FermataAPI.md#FermataUnfreezeWallet) | **Post** /api/v3/fermata/wallet/unfreeze | Unfreeze wallet
[**FermataUnlinkDealer**](FermataAPI.md#FermataUnlinkDealer) | **Post** /api/v3/fermata/dealer/unlink | Unlink exchange account from dealer



## FermataCloseWallet

> FermataCreateWallet200Response FermataCloseWallet(ctx).FermataCloseWalletRequest(fermataCloseWalletRequest).Execute()

Close wallet



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
	fermataCloseWalletRequest := *openapiclient.NewFermataCloseWalletRequest("WalletId_example") // FermataCloseWalletRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FermataAPI.FermataCloseWallet(context.Background()).FermataCloseWalletRequest(fermataCloseWalletRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FermataAPI.FermataCloseWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FermataCloseWallet`: FermataCreateWallet200Response
	fmt.Fprintf(os.Stdout, "Response from `FermataAPI.FermataCloseWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFermataCloseWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fermataCloseWalletRequest** | [**FermataCloseWalletRequest**](FermataCloseWalletRequest.md) |  | 

### Return type

[**FermataCreateWallet200Response**](FermataCreateWallet200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FermataCreateWallet

> FermataCreateWallet200Response FermataCreateWallet(ctx).FermataCreateWalletRequest(fermataCreateWalletRequest).Execute()

Create wallet



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
	fermataCreateWalletRequest := *openapiclient.NewFermataCreateWalletRequest() // FermataCreateWalletRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FermataAPI.FermataCreateWallet(context.Background()).FermataCreateWalletRequest(fermataCreateWalletRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FermataAPI.FermataCreateWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FermataCreateWallet`: FermataCreateWallet200Response
	fmt.Fprintf(os.Stdout, "Response from `FermataAPI.FermataCreateWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFermataCreateWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fermataCreateWalletRequest** | [**FermataCreateWalletRequest**](FermataCreateWalletRequest.md) |  | 

### Return type

[**FermataCreateWallet200Response**](FermataCreateWallet200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FermataFreezeWallet

> FermataCreateWallet200Response FermataFreezeWallet(ctx).FermataFreezeWalletRequest(fermataFreezeWalletRequest).Execute()

Freeze wallet



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
	fermataFreezeWalletRequest := *openapiclient.NewFermataFreezeWalletRequest("WalletId_example") // FermataFreezeWalletRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FermataAPI.FermataFreezeWallet(context.Background()).FermataFreezeWalletRequest(fermataFreezeWalletRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FermataAPI.FermataFreezeWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FermataFreezeWallet`: FermataCreateWallet200Response
	fmt.Fprintf(os.Stdout, "Response from `FermataAPI.FermataFreezeWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFermataFreezeWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fermataFreezeWalletRequest** | [**FermataFreezeWalletRequest**](FermataFreezeWalletRequest.md) |  | 

### Return type

[**FermataCreateWallet200Response**](FermataCreateWallet200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FermataLinkDealer

> DeleteMarketInstrument200Response FermataLinkDealer(ctx).FermataLinkDealerRequest(fermataLinkDealerRequest).Execute()

Link exchange account to dealer



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
	fermataLinkDealerRequest := *openapiclient.NewFermataLinkDealerRequest("DealerAccountId_example", "TradingAccountId_example") // FermataLinkDealerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FermataAPI.FermataLinkDealer(context.Background()).FermataLinkDealerRequest(fermataLinkDealerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FermataAPI.FermataLinkDealer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FermataLinkDealer`: DeleteMarketInstrument200Response
	fmt.Fprintf(os.Stdout, "Response from `FermataAPI.FermataLinkDealer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFermataLinkDealerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fermataLinkDealerRequest** | [**FermataLinkDealerRequest**](FermataLinkDealerRequest.md) |  | 

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


## FermataListWallets

> FermataListWallets200Response FermataListWallets(ctx).WalletType(walletType).Limit(limit).Offset(offset).Cursor(cursor).Execute()

List wallets



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
	walletType := openapiclient.walletType("DEALER") // WalletType | Filter by wallet type (optional)
	limit := int32(100) // int32 | Limit the number of returned results (optional) (default to 50)
	offset := int32(0) // int32 | Offset of the returned results (optional) (default to 0)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FermataAPI.FermataListWallets(context.Background()).WalletType(walletType).Limit(limit).Offset(offset).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FermataAPI.FermataListWallets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FermataListWallets`: FermataListWallets200Response
	fmt.Fprintf(os.Stdout, "Response from `FermataAPI.FermataListWallets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFermataListWalletsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletType** | [**WalletType**](WalletType.md) | Filter by wallet type | 
 **limit** | **int32** | Limit the number of returned results | [default to 50]
 **offset** | **int32** | Offset of the returned results | [default to 0]
 **cursor** | **string** |  | 

### Return type

[**FermataListWallets200Response**](FermataListWallets200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FermataTransferWallet

> DeleteMarketInstrument200Response FermataTransferWallet(ctx).FermataTransferWalletRequest(fermataTransferWalletRequest).Execute()

Transfer between wallets



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
	fermataTransferWalletRequest := *openapiclient.NewFermataTransferWalletRequest("SourceWalletId_example", "DestWalletId_example", "USDT", "100.50") // FermataTransferWalletRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FermataAPI.FermataTransferWallet(context.Background()).FermataTransferWalletRequest(fermataTransferWalletRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FermataAPI.FermataTransferWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FermataTransferWallet`: DeleteMarketInstrument200Response
	fmt.Fprintf(os.Stdout, "Response from `FermataAPI.FermataTransferWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFermataTransferWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fermataTransferWalletRequest** | [**FermataTransferWalletRequest**](FermataTransferWalletRequest.md) |  | 

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


## FermataUnfreezeWallet

> FermataCreateWallet200Response FermataUnfreezeWallet(ctx).FermataUnfreezeWalletRequest(fermataUnfreezeWalletRequest).Execute()

Unfreeze wallet



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
	fermataUnfreezeWalletRequest := *openapiclient.NewFermataUnfreezeWalletRequest("WalletId_example") // FermataUnfreezeWalletRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FermataAPI.FermataUnfreezeWallet(context.Background()).FermataUnfreezeWalletRequest(fermataUnfreezeWalletRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FermataAPI.FermataUnfreezeWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FermataUnfreezeWallet`: FermataCreateWallet200Response
	fmt.Fprintf(os.Stdout, "Response from `FermataAPI.FermataUnfreezeWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFermataUnfreezeWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fermataUnfreezeWalletRequest** | [**FermataUnfreezeWalletRequest**](FermataUnfreezeWalletRequest.md) |  | 

### Return type

[**FermataCreateWallet200Response**](FermataCreateWallet200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FermataUnlinkDealer

> DeleteMarketInstrument200Response FermataUnlinkDealer(ctx).FermataUnlinkDealerRequest(fermataUnlinkDealerRequest).Execute()

Unlink exchange account from dealer



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
	fermataUnlinkDealerRequest := *openapiclient.NewFermataUnlinkDealerRequest("DealerAccountId_example", "TradingAccountId_example") // FermataUnlinkDealerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FermataAPI.FermataUnlinkDealer(context.Background()).FermataUnlinkDealerRequest(fermataUnlinkDealerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FermataAPI.FermataUnlinkDealer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FermataUnlinkDealer`: DeleteMarketInstrument200Response
	fmt.Fprintf(os.Stdout, "Response from `FermataAPI.FermataUnlinkDealer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFermataUnlinkDealerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fermataUnlinkDealerRequest** | [**FermataUnlinkDealerRequest**](FermataUnlinkDealerRequest.md) |  | 

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

