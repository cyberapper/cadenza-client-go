# RpcPagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **int32** | Offset for pagination | [optional] 
**Limit** | Pointer to **int32** | Limit for pagination | [optional] 
**Total** | Pointer to **int32** | Total number of items | [optional] 
**Cursor** | Pointer to **string** | Cursor for next page | [optional] 
**HasNext** | Pointer to **bool** | Whether there are more items | [optional] 

## Methods

### NewRpcPagination

`func NewRpcPagination() *RpcPagination`

NewRpcPagination instantiates a new RpcPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcPaginationWithDefaults

`func NewRpcPaginationWithDefaults() *RpcPagination`

NewRpcPaginationWithDefaults instantiates a new RpcPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *RpcPagination) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *RpcPagination) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *RpcPagination) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *RpcPagination) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *RpcPagination) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RpcPagination) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RpcPagination) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *RpcPagination) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotal

`func (o *RpcPagination) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RpcPagination) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RpcPagination) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *RpcPagination) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetCursor

`func (o *RpcPagination) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *RpcPagination) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *RpcPagination) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *RpcPagination) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetHasNext

`func (o *RpcPagination) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *RpcPagination) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *RpcPagination) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *RpcPagination) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


