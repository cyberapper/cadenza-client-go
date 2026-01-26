# RpcGetSystemInfoResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**RpcSystemInfo**](RpcSystemInfo.md) |  | [optional] 
**Error** | Pointer to [**RpcError**](RpcError.md) |  | [optional] 

## Methods

### NewRpcGetSystemInfoResult

`func NewRpcGetSystemInfoResult() *RpcGetSystemInfoResult`

NewRpcGetSystemInfoResult instantiates a new RpcGetSystemInfoResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcGetSystemInfoResultWithDefaults

`func NewRpcGetSystemInfoResultWithDefaults() *RpcGetSystemInfoResult`

NewRpcGetSystemInfoResultWithDefaults instantiates a new RpcGetSystemInfoResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RpcGetSystemInfoResult) GetData() RpcSystemInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RpcGetSystemInfoResult) GetDataOk() (*RpcSystemInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RpcGetSystemInfoResult) SetData(v RpcSystemInfo)`

SetData sets Data field to given value.

### HasData

`func (o *RpcGetSystemInfoResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *RpcGetSystemInfoResult) GetError() RpcError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RpcGetSystemInfoResult) GetErrorOk() (*RpcError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RpcGetSystemInfoResult) SetError(v RpcError)`

SetError sets Error field to given value.

### HasError

`func (o *RpcGetSystemInfoResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


