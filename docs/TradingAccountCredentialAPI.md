# \TradingAccountCredentialAPI

All URIs are relative to *https://cadenza-api-uat.algo724.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTradingAccountCredential**](TradingAccountCredentialAPI.md#CreateTradingAccountCredential) | **Post** /api/v3/credential/create | Create trading account credential
[**ListTradingAccountCredentials**](TradingAccountCredentialAPI.md#ListTradingAccountCredentials) | **Get** /api/v3/credential/list | List trading account credentials
[**RevokeTradingAccountCredential**](TradingAccountCredentialAPI.md#RevokeTradingAccountCredential) | **Post** /api/v3/credential/revoke | Revoke trading account credential
[**RotateTradingAccountCredential**](TradingAccountCredentialAPI.md#RotateTradingAccountCredential) | **Post** /api/v3/credential/rotate | Rotate trading account credential
[**VerifyTradingAccountCredential**](TradingAccountCredentialAPI.md#VerifyTradingAccountCredential) | **Post** /api/v3/credential/verify | Verify trading account credential



## CreateTradingAccountCredential

> CreateTradingAccountCredential200Response CreateTradingAccountCredential(ctx).CreateTradingAccountCredentialRequest(createTradingAccountCredentialRequest).Execute()

Create trading account credential



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
	createTradingAccountCredentialRequest := *openapiclient.NewCreateTradingAccountCredentialRequest(openapiclient.venue(""), openapiclient.credentialType("")) // CreateTradingAccountCredentialRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountCredentialAPI.CreateTradingAccountCredential(context.Background()).CreateTradingAccountCredentialRequest(createTradingAccountCredentialRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountCredentialAPI.CreateTradingAccountCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingAccountCredential`: CreateTradingAccountCredential200Response
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountCredentialAPI.CreateTradingAccountCredential`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingAccountCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingAccountCredentialRequest** | [**CreateTradingAccountCredentialRequest**](CreateTradingAccountCredentialRequest.md) |  | 

### Return type

[**CreateTradingAccountCredential200Response**](CreateTradingAccountCredential200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTradingAccountCredentials

> ListTradingAccountCredentials200Response ListTradingAccountCredentials(ctx).CredentialType(credentialType).CredentialStatus(credentialStatus).CredentialIds(credentialIds).Execute()

List trading account credentials



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
	credentialType := openapiclient.credentialType("") // CredentialType | Credential type (optional)
	credentialStatus := openapiclient.tradingAccountStatus("SETUP") // TradingAccountStatus | Credential status (optional)
	credentialIds := []string{"Inner_example"} // []string | credentialId array (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountCredentialAPI.ListTradingAccountCredentials(context.Background()).CredentialType(credentialType).CredentialStatus(credentialStatus).CredentialIds(credentialIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountCredentialAPI.ListTradingAccountCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTradingAccountCredentials`: ListTradingAccountCredentials200Response
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountCredentialAPI.ListTradingAccountCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTradingAccountCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentialType** | [**CredentialType**](CredentialType.md) | Credential type | 
 **credentialStatus** | [**TradingAccountStatus**](TradingAccountStatus.md) | Credential status | 
 **credentialIds** | **[]string** | credentialId array | 

### Return type

[**ListTradingAccountCredentials200Response**](ListTradingAccountCredentials200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeTradingAccountCredential

> CreateTradingAccountCredential200Response RevokeTradingAccountCredential(ctx).RevokeTradingAccountCredentialRequest(revokeTradingAccountCredentialRequest).Execute()

Revoke trading account credential



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
	revokeTradingAccountCredentialRequest := *openapiclient.NewRevokeTradingAccountCredentialRequest("CredentialId_example") // RevokeTradingAccountCredentialRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountCredentialAPI.RevokeTradingAccountCredential(context.Background()).RevokeTradingAccountCredentialRequest(revokeTradingAccountCredentialRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountCredentialAPI.RevokeTradingAccountCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevokeTradingAccountCredential`: CreateTradingAccountCredential200Response
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountCredentialAPI.RevokeTradingAccountCredential`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeTradingAccountCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revokeTradingAccountCredentialRequest** | [**RevokeTradingAccountCredentialRequest**](RevokeTradingAccountCredentialRequest.md) |  | 

### Return type

[**CreateTradingAccountCredential200Response**](CreateTradingAccountCredential200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateTradingAccountCredential

> CreateTradingAccountCredential200Response RotateTradingAccountCredential(ctx).RotateTradingAccountCredentialRequest(rotateTradingAccountCredentialRequest).Execute()

Rotate trading account credential



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
	rotateTradingAccountCredentialRequest := *openapiclient.NewRotateTradingAccountCredentialRequest("CredentialId_example", "my_api_key") // RotateTradingAccountCredentialRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountCredentialAPI.RotateTradingAccountCredential(context.Background()).RotateTradingAccountCredentialRequest(rotateTradingAccountCredentialRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountCredentialAPI.RotateTradingAccountCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateTradingAccountCredential`: CreateTradingAccountCredential200Response
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountCredentialAPI.RotateTradingAccountCredential`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRotateTradingAccountCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rotateTradingAccountCredentialRequest** | [**RotateTradingAccountCredentialRequest**](RotateTradingAccountCredentialRequest.md) |  | 

### Return type

[**CreateTradingAccountCredential200Response**](CreateTradingAccountCredential200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyTradingAccountCredential

> VerifyTradingAccountCredential200Response VerifyTradingAccountCredential(ctx).VerifyTradingAccountCredentialRequest(verifyTradingAccountCredentialRequest).Execute()

Verify trading account credential



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
	verifyTradingAccountCredentialRequest := *openapiclient.NewVerifyTradingAccountCredentialRequest([]string{"CredentialIds_example"}) // VerifyTradingAccountCredentialRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountCredentialAPI.VerifyTradingAccountCredential(context.Background()).VerifyTradingAccountCredentialRequest(verifyTradingAccountCredentialRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountCredentialAPI.VerifyTradingAccountCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyTradingAccountCredential`: VerifyTradingAccountCredential200Response
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountCredentialAPI.VerifyTradingAccountCredential`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyTradingAccountCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyTradingAccountCredentialRequest** | [**VerifyTradingAccountCredentialRequest**](VerifyTradingAccountCredentialRequest.md) |  | 

### Return type

[**VerifyTradingAccountCredential200Response**](VerifyTradingAccountCredential200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

