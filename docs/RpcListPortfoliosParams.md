# RpcListPortfoliosParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradingAccountId** | Pointer to **string** | Filter by trading account ID | [optional] 
**Venue** | Pointer to **string** | Filter by venue | [optional] 
**Currency** | Pointer to **string** | Filter by currency | [optional] 
**Pagination** | Pointer to [**RpcPagination**](RpcPagination.md) |  | [optional] 

## Methods

### NewRpcListPortfoliosParams

`func NewRpcListPortfoliosParams() *RpcListPortfoliosParams`

NewRpcListPortfoliosParams instantiates a new RpcListPortfoliosParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcListPortfoliosParamsWithDefaults

`func NewRpcListPortfoliosParamsWithDefaults() *RpcListPortfoliosParams`

NewRpcListPortfoliosParamsWithDefaults instantiates a new RpcListPortfoliosParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradingAccountId

`func (o *RpcListPortfoliosParams) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *RpcListPortfoliosParams) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *RpcListPortfoliosParams) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.

### HasTradingAccountId

`func (o *RpcListPortfoliosParams) HasTradingAccountId() bool`

HasTradingAccountId returns a boolean if a field has been set.

### GetVenue

`func (o *RpcListPortfoliosParams) GetVenue() string`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *RpcListPortfoliosParams) GetVenueOk() (*string, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *RpcListPortfoliosParams) SetVenue(v string)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *RpcListPortfoliosParams) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetCurrency

`func (o *RpcListPortfoliosParams) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RpcListPortfoliosParams) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RpcListPortfoliosParams) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *RpcListPortfoliosParams) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPagination

`func (o *RpcListPortfoliosParams) GetPagination() RpcPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RpcListPortfoliosParams) GetPaginationOk() (*RpcPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RpcListPortfoliosParams) SetPagination(v RpcPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RpcListPortfoliosParams) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


