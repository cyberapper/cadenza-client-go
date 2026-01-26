# RpcListTradingAccountsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradingAccountId** | Pointer to **string** | Filter by specific account ID | [optional] 
**Venue** | Pointer to **string** | Filter by venue | [optional] 
**Status** | Pointer to [**TradingAccountStatus**](TradingAccountStatus.md) |  | [optional] 
**Pagination** | Pointer to [**RpcPagination**](RpcPagination.md) |  | [optional] 

## Methods

### NewRpcListTradingAccountsParams

`func NewRpcListTradingAccountsParams() *RpcListTradingAccountsParams`

NewRpcListTradingAccountsParams instantiates a new RpcListTradingAccountsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcListTradingAccountsParamsWithDefaults

`func NewRpcListTradingAccountsParamsWithDefaults() *RpcListTradingAccountsParams`

NewRpcListTradingAccountsParamsWithDefaults instantiates a new RpcListTradingAccountsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradingAccountId

`func (o *RpcListTradingAccountsParams) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *RpcListTradingAccountsParams) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *RpcListTradingAccountsParams) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.

### HasTradingAccountId

`func (o *RpcListTradingAccountsParams) HasTradingAccountId() bool`

HasTradingAccountId returns a boolean if a field has been set.

### GetVenue

`func (o *RpcListTradingAccountsParams) GetVenue() string`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *RpcListTradingAccountsParams) GetVenueOk() (*string, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *RpcListTradingAccountsParams) SetVenue(v string)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *RpcListTradingAccountsParams) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetStatus

`func (o *RpcListTradingAccountsParams) GetStatus() TradingAccountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RpcListTradingAccountsParams) GetStatusOk() (*TradingAccountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RpcListTradingAccountsParams) SetStatus(v TradingAccountStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RpcListTradingAccountsParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPagination

`func (o *RpcListTradingAccountsParams) GetPagination() RpcPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RpcListTradingAccountsParams) GetPaginationOk() (*RpcPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RpcListTradingAccountsParams) SetPagination(v RpcPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RpcListTradingAccountsParams) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


