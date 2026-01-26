# RpcInstrumentUpdatedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RpcInstrument**](RpcInstrument.md) |  | [optional] 

## Methods

### NewRpcInstrumentUpdatedEvent

`func NewRpcInstrumentUpdatedEvent() *RpcInstrumentUpdatedEvent`

NewRpcInstrumentUpdatedEvent instantiates a new RpcInstrumentUpdatedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcInstrumentUpdatedEventWithDefaults

`func NewRpcInstrumentUpdatedEventWithDefaults() *RpcInstrumentUpdatedEvent`

NewRpcInstrumentUpdatedEventWithDefaults instantiates a new RpcInstrumentUpdatedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RpcInstrumentUpdatedEvent) GetData() []RpcInstrument`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RpcInstrumentUpdatedEvent) GetDataOk() (*[]RpcInstrument, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RpcInstrumentUpdatedEvent) SetData(v []RpcInstrument)`

SetData sets Data field to given value.

### HasData

`func (o *RpcInstrumentUpdatedEvent) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


