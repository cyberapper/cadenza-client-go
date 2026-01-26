# WsRPCResultData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RpcSubscription**](RpcSubscription.md) |  | [optional] 
**Error** | Pointer to [**RpcError**](RpcError.md) |  | [optional] 
**Pagination** | Pointer to [**RpcPagination**](RpcPagination.md) |  | [optional] 

## Methods

### NewWsRPCResultData

`func NewWsRPCResultData() *WsRPCResultData`

NewWsRPCResultData instantiates a new WsRPCResultData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsRPCResultDataWithDefaults

`func NewWsRPCResultDataWithDefaults() *WsRPCResultData`

NewWsRPCResultDataWithDefaults instantiates a new WsRPCResultData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WsRPCResultData) GetData() []RpcSubscription`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WsRPCResultData) GetDataOk() (*[]RpcSubscription, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WsRPCResultData) SetData(v []RpcSubscription)`

SetData sets Data field to given value.

### HasData

`func (o *WsRPCResultData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *WsRPCResultData) GetError() RpcError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WsRPCResultData) GetErrorOk() (*RpcError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WsRPCResultData) SetError(v RpcError)`

SetError sets Error field to given value.

### HasError

`func (o *WsRPCResultData) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPagination

`func (o *WsRPCResultData) GetPagination() RpcPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *WsRPCResultData) GetPaginationOk() (*RpcPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *WsRPCResultData) SetPagination(v RpcPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *WsRPCResultData) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


