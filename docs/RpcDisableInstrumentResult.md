# RpcDisableInstrumentResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**RpcInstrument**](RpcInstrument.md) |  | [optional] 
**Error** | Pointer to [**RpcError**](RpcError.md) |  | [optional] 

## Methods

### NewRpcDisableInstrumentResult

`func NewRpcDisableInstrumentResult() *RpcDisableInstrumentResult`

NewRpcDisableInstrumentResult instantiates a new RpcDisableInstrumentResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcDisableInstrumentResultWithDefaults

`func NewRpcDisableInstrumentResultWithDefaults() *RpcDisableInstrumentResult`

NewRpcDisableInstrumentResultWithDefaults instantiates a new RpcDisableInstrumentResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RpcDisableInstrumentResult) GetData() RpcInstrument`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RpcDisableInstrumentResult) GetDataOk() (*RpcInstrument, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RpcDisableInstrumentResult) SetData(v RpcInstrument)`

SetData sets Data field to given value.

### HasData

`func (o *RpcDisableInstrumentResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *RpcDisableInstrumentResult) GetError() RpcError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RpcDisableInstrumentResult) GetErrorOk() (*RpcError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RpcDisableInstrumentResult) SetError(v RpcError)`

SetError sets Error field to given value.

### HasError

`func (o *RpcDisableInstrumentResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


