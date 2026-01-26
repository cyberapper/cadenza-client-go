# RpcListTradingAccountsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RpcTradingAccount**](RpcTradingAccount.md) |  | [optional] 
**Pagination** | Pointer to [**RpcPagination**](RpcPagination.md) |  | [optional] 
**Error** | Pointer to [**RpcError**](RpcError.md) |  | [optional] 

## Methods

### NewRpcListTradingAccountsResult

`func NewRpcListTradingAccountsResult() *RpcListTradingAccountsResult`

NewRpcListTradingAccountsResult instantiates a new RpcListTradingAccountsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcListTradingAccountsResultWithDefaults

`func NewRpcListTradingAccountsResultWithDefaults() *RpcListTradingAccountsResult`

NewRpcListTradingAccountsResultWithDefaults instantiates a new RpcListTradingAccountsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RpcListTradingAccountsResult) GetData() []RpcTradingAccount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RpcListTradingAccountsResult) GetDataOk() (*[]RpcTradingAccount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RpcListTradingAccountsResult) SetData(v []RpcTradingAccount)`

SetData sets Data field to given value.

### HasData

`func (o *RpcListTradingAccountsResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *RpcListTradingAccountsResult) GetPagination() RpcPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RpcListTradingAccountsResult) GetPaginationOk() (*RpcPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RpcListTradingAccountsResult) SetPagination(v RpcPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RpcListTradingAccountsResult) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetError

`func (o *RpcListTradingAccountsResult) GetError() RpcError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RpcListTradingAccountsResult) GetErrorOk() (*RpcError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RpcListTradingAccountsResult) SetError(v RpcError)`

SetError sets Error field to given value.

### HasError

`func (o *RpcListTradingAccountsResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


