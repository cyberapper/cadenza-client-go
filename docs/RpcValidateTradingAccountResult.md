# RpcValidateTradingAccountResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**RpcValidateTradingAccountResultData**](RpcValidateTradingAccountResultData.md) |  | [optional] 
**Error** | Pointer to [**RpcError**](RpcError.md) |  | [optional] 

## Methods

### NewRpcValidateTradingAccountResult

`func NewRpcValidateTradingAccountResult() *RpcValidateTradingAccountResult`

NewRpcValidateTradingAccountResult instantiates a new RpcValidateTradingAccountResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcValidateTradingAccountResultWithDefaults

`func NewRpcValidateTradingAccountResultWithDefaults() *RpcValidateTradingAccountResult`

NewRpcValidateTradingAccountResultWithDefaults instantiates a new RpcValidateTradingAccountResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RpcValidateTradingAccountResult) GetData() RpcValidateTradingAccountResultData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RpcValidateTradingAccountResult) GetDataOk() (*RpcValidateTradingAccountResultData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RpcValidateTradingAccountResult) SetData(v RpcValidateTradingAccountResultData)`

SetData sets Data field to given value.

### HasData

`func (o *RpcValidateTradingAccountResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *RpcValidateTradingAccountResult) GetError() RpcError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RpcValidateTradingAccountResult) GetErrorOk() (*RpcError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RpcValidateTradingAccountResult) SetError(v RpcError)`

SetError sets Error field to given value.

### HasError

`func (o *RpcValidateTradingAccountResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


