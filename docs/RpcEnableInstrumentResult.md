# RpcEnableInstrumentResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**RpcInstrument**](RpcInstrument.md) |  | [optional] 
**Error** | Pointer to [**RpcError**](RpcError.md) |  | [optional] 

## Methods

### NewRpcEnableInstrumentResult

`func NewRpcEnableInstrumentResult() *RpcEnableInstrumentResult`

NewRpcEnableInstrumentResult instantiates a new RpcEnableInstrumentResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcEnableInstrumentResultWithDefaults

`func NewRpcEnableInstrumentResultWithDefaults() *RpcEnableInstrumentResult`

NewRpcEnableInstrumentResultWithDefaults instantiates a new RpcEnableInstrumentResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RpcEnableInstrumentResult) GetData() RpcInstrument`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RpcEnableInstrumentResult) GetDataOk() (*RpcInstrument, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RpcEnableInstrumentResult) SetData(v RpcInstrument)`

SetData sets Data field to given value.

### HasData

`func (o *RpcEnableInstrumentResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *RpcEnableInstrumentResult) GetError() RpcError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RpcEnableInstrumentResult) GetErrorOk() (*RpcError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RpcEnableInstrumentResult) SetError(v RpcError)`

SetError sets Error field to given value.

### HasError

`func (o *RpcEnableInstrumentResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


