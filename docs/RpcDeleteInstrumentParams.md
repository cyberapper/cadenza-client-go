# RpcDeleteInstrumentParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstrumentId** | **string** | Instrument ID (format: VENUE:BASE/QUOTE) | 

## Methods

### NewRpcDeleteInstrumentParams

`func NewRpcDeleteInstrumentParams(instrumentId string, ) *RpcDeleteInstrumentParams`

NewRpcDeleteInstrumentParams instantiates a new RpcDeleteInstrumentParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcDeleteInstrumentParamsWithDefaults

`func NewRpcDeleteInstrumentParamsWithDefaults() *RpcDeleteInstrumentParams`

NewRpcDeleteInstrumentParamsWithDefaults instantiates a new RpcDeleteInstrumentParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstrumentId

`func (o *RpcDeleteInstrumentParams) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *RpcDeleteInstrumentParams) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *RpcDeleteInstrumentParams) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


