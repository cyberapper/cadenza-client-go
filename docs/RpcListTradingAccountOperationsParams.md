# RpcListTradingAccountOperationsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradingAccountId** | **string** |  | 
**Pagination** | Pointer to [**RpcPagination**](RpcPagination.md) |  | [optional] 

## Methods

### NewRpcListTradingAccountOperationsParams

`func NewRpcListTradingAccountOperationsParams(tradingAccountId string, ) *RpcListTradingAccountOperationsParams`

NewRpcListTradingAccountOperationsParams instantiates a new RpcListTradingAccountOperationsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcListTradingAccountOperationsParamsWithDefaults

`func NewRpcListTradingAccountOperationsParamsWithDefaults() *RpcListTradingAccountOperationsParams`

NewRpcListTradingAccountOperationsParamsWithDefaults instantiates a new RpcListTradingAccountOperationsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradingAccountId

`func (o *RpcListTradingAccountOperationsParams) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *RpcListTradingAccountOperationsParams) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *RpcListTradingAccountOperationsParams) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.


### GetPagination

`func (o *RpcListTradingAccountOperationsParams) GetPagination() RpcPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RpcListTradingAccountOperationsParams) GetPaginationOk() (*RpcPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RpcListTradingAccountOperationsParams) SetPagination(v RpcPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RpcListTradingAccountOperationsParams) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


