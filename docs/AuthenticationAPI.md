# \AuthenticationAPI

All URIs are relative to *https://cadenza-api-uat.algo724.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthGetUser**](AuthenticationAPI.md#AuthGetUser) | **Get** /api/v3/auth/user | Get current user
[**AuthLogin**](AuthenticationAPI.md#AuthLogin) | **Post** /api/v3/auth/login | Login with email and password
[**AuthLogout**](AuthenticationAPI.md#AuthLogout) | **Post** /api/v3/auth/logout | Logout user
[**AuthRecover**](AuthenticationAPI.md#AuthRecover) | **Post** /api/v3/auth/recover | Request password recovery
[**AuthRefreshToken**](AuthenticationAPI.md#AuthRefreshToken) | **Post** /api/v3/auth/token/refresh | Refresh access token
[**AuthSignup**](AuthenticationAPI.md#AuthSignup) | **Post** /api/v3/auth/signup | Sign up new user
[**AuthUpdateUser**](AuthenticationAPI.md#AuthUpdateUser) | **Put** /api/v3/auth/user | Update current user



## AuthGetUser

> AuthGetUser200Response AuthGetUser(ctx).Execute()

Get current user



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
	resp, r, err := apiClient.AuthenticationAPI.AuthGetUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AuthGetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthGetUser`: AuthGetUser200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AuthGetUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthGetUserRequest struct via the builder pattern


### Return type

[**AuthGetUser200Response**](AuthGetUser200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthLogin

> AuthSignup200Response AuthLogin(ctx).AuthLoginRequest(authLoginRequest).Execute()

Login with email and password



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
	authLoginRequest := *openapiclient.NewAuthLoginRequest("user@example.com", "securepassword123") // AuthLoginRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AuthLogin(context.Background()).AuthLoginRequest(authLoginRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AuthLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthLogin`: AuthSignup200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AuthLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authLoginRequest** | [**AuthLoginRequest**](AuthLoginRequest.md) |  | 

### Return type

[**AuthSignup200Response**](AuthSignup200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthLogout

> AuthLogout200Response AuthLogout(ctx).Execute()

Logout user



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
	resp, r, err := apiClient.AuthenticationAPI.AuthLogout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AuthLogout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthLogout`: AuthLogout200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AuthLogout`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthLogoutRequest struct via the builder pattern


### Return type

[**AuthLogout200Response**](AuthLogout200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthRecover

> AuthLogout200Response AuthRecover(ctx).AuthRecoverRequest(authRecoverRequest).Execute()

Request password recovery



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
	authRecoverRequest := *openapiclient.NewAuthRecoverRequest("user@example.com") // AuthRecoverRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AuthRecover(context.Background()).AuthRecoverRequest(authRecoverRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AuthRecover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthRecover`: AuthLogout200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AuthRecover`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthRecoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authRecoverRequest** | [**AuthRecoverRequest**](AuthRecoverRequest.md) |  | 

### Return type

[**AuthLogout200Response**](AuthLogout200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthRefreshToken

> AuthSignup200Response AuthRefreshToken(ctx).AuthRefreshTokenRequest(authRefreshTokenRequest).Execute()

Refresh access token



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
	authRefreshTokenRequest := *openapiclient.NewAuthRefreshTokenRequest("RefreshToken_example") // AuthRefreshTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AuthRefreshToken(context.Background()).AuthRefreshTokenRequest(authRefreshTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AuthRefreshToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthRefreshToken`: AuthSignup200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AuthRefreshToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthRefreshTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authRefreshTokenRequest** | [**AuthRefreshTokenRequest**](AuthRefreshTokenRequest.md) |  | 

### Return type

[**AuthSignup200Response**](AuthSignup200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthSignup

> AuthSignup200Response AuthSignup(ctx).AuthSignupRequest(authSignupRequest).Execute()

Sign up new user



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
	authSignupRequest := *openapiclient.NewAuthSignupRequest("user@example.com", "securepassword123") // AuthSignupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AuthSignup(context.Background()).AuthSignupRequest(authSignupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AuthSignup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSignup`: AuthSignup200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AuthSignup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authSignupRequest** | [**AuthSignupRequest**](AuthSignupRequest.md) |  | 

### Return type

[**AuthSignup200Response**](AuthSignup200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthUpdateUser

> AuthGetUser200Response AuthUpdateUser(ctx).AuthUpdateUserRequest(authUpdateUserRequest).Execute()

Update current user



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
	authUpdateUserRequest := *openapiclient.NewAuthUpdateUserRequest() // AuthUpdateUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AuthUpdateUser(context.Background()).AuthUpdateUserRequest(authUpdateUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AuthUpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthUpdateUser`: AuthGetUser200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AuthUpdateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authUpdateUserRequest** | [**AuthUpdateUserRequest**](AuthUpdateUserRequest.md) |  | 

### Return type

[**AuthGetUser200Response**](AuthGetUser200Response.md)

### Authorization

[SupabaseOAuth2BearerAuth](../README.md#SupabaseOAuth2BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

