# RpcListTradeOrdersResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RpcTradeOrder**](RpcTradeOrder.md) |  | [optional] 
**Pagination** | Pointer to [**RpcPagination**](RpcPagination.md) |  | [optional] 
**Error** | Pointer to [**RpcError**](RpcError.md) |  | [optional] 

## Methods

### NewRpcListTradeOrdersResult

`func NewRpcListTradeOrdersResult() *RpcListTradeOrdersResult`

NewRpcListTradeOrdersResult instantiates a new RpcListTradeOrdersResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcListTradeOrdersResultWithDefaults

`func NewRpcListTradeOrdersResultWithDefaults() *RpcListTradeOrdersResult`

NewRpcListTradeOrdersResultWithDefaults instantiates a new RpcListTradeOrdersResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RpcListTradeOrdersResult) GetData() []RpcTradeOrder`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RpcListTradeOrdersResult) GetDataOk() (*[]RpcTradeOrder, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RpcListTradeOrdersResult) SetData(v []RpcTradeOrder)`

SetData sets Data field to given value.

### HasData

`func (o *RpcListTradeOrdersResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *RpcListTradeOrdersResult) GetPagination() RpcPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RpcListTradeOrdersResult) GetPaginationOk() (*RpcPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RpcListTradeOrdersResult) SetPagination(v RpcPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RpcListTradeOrdersResult) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetError

`func (o *RpcListTradeOrdersResult) GetError() RpcError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RpcListTradeOrdersResult) GetErrorOk() (*RpcError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RpcListTradeOrdersResult) SetError(v RpcError)`

SetError sets Error field to given value.

### HasError

`func (o *RpcListTradeOrdersResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


