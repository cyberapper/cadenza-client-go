# RpcListSecuritiesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RpcSecurity**](RpcSecurity.md) |  | [optional] 
**Error** | Pointer to [**RpcError**](RpcError.md) |  | [optional] 

## Methods

### NewRpcListSecuritiesResult

`func NewRpcListSecuritiesResult() *RpcListSecuritiesResult`

NewRpcListSecuritiesResult instantiates a new RpcListSecuritiesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcListSecuritiesResultWithDefaults

`func NewRpcListSecuritiesResultWithDefaults() *RpcListSecuritiesResult`

NewRpcListSecuritiesResultWithDefaults instantiates a new RpcListSecuritiesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RpcListSecuritiesResult) GetData() []RpcSecurity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RpcListSecuritiesResult) GetDataOk() (*[]RpcSecurity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RpcListSecuritiesResult) SetData(v []RpcSecurity)`

SetData sets Data field to given value.

### HasData

`func (o *RpcListSecuritiesResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *RpcListSecuritiesResult) GetError() RpcError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RpcListSecuritiesResult) GetErrorOk() (*RpcError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RpcListSecuritiesResult) SetError(v RpcError)`

SetError sets Error field to given value.

### HasError

`func (o *RpcListSecuritiesResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


