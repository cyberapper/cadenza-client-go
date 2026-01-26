# RpcOrderBookUpdatedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**RpcOrderBook**](RpcOrderBook.md) |  | [optional] 

## Methods

### NewRpcOrderBookUpdatedEvent

`func NewRpcOrderBookUpdatedEvent() *RpcOrderBookUpdatedEvent`

NewRpcOrderBookUpdatedEvent instantiates a new RpcOrderBookUpdatedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcOrderBookUpdatedEventWithDefaults

`func NewRpcOrderBookUpdatedEventWithDefaults() *RpcOrderBookUpdatedEvent`

NewRpcOrderBookUpdatedEventWithDefaults instantiates a new RpcOrderBookUpdatedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RpcOrderBookUpdatedEvent) GetData() RpcOrderBook`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RpcOrderBookUpdatedEvent) GetDataOk() (*RpcOrderBook, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RpcOrderBookUpdatedEvent) SetData(v RpcOrderBook)`

SetData sets Data field to given value.

### HasData

`func (o *RpcOrderBookUpdatedEvent) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


