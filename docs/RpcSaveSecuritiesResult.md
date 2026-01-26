# RpcSaveSecuritiesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RpcSecurity**](RpcSecurity.md) |  | [optional] 
**Error** | Pointer to [**RpcError**](RpcError.md) |  | [optional] 

## Methods

### NewRpcSaveSecuritiesResult

`func NewRpcSaveSecuritiesResult() *RpcSaveSecuritiesResult`

NewRpcSaveSecuritiesResult instantiates a new RpcSaveSecuritiesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcSaveSecuritiesResultWithDefaults

`func NewRpcSaveSecuritiesResultWithDefaults() *RpcSaveSecuritiesResult`

NewRpcSaveSecuritiesResultWithDefaults instantiates a new RpcSaveSecuritiesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RpcSaveSecuritiesResult) GetData() []RpcSecurity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RpcSaveSecuritiesResult) GetDataOk() (*[]RpcSecurity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RpcSaveSecuritiesResult) SetData(v []RpcSecurity)`

SetData sets Data field to given value.

### HasData

`func (o *RpcSaveSecuritiesResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *RpcSaveSecuritiesResult) GetError() RpcError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RpcSaveSecuritiesResult) GetErrorOk() (*RpcError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RpcSaveSecuritiesResult) SetError(v RpcError)`

SetError sets Error field to given value.

### HasError

`func (o *RpcSaveSecuritiesResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


