# RpcUnsubscribeOrderBookParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | Pointer to **string** | Subscription ID to cancel | [optional] 
**InstrumentId** | Pointer to **string** | Instrument ID to unsubscribe | [optional] 

## Methods

### NewRpcUnsubscribeOrderBookParams

`func NewRpcUnsubscribeOrderBookParams() *RpcUnsubscribeOrderBookParams`

NewRpcUnsubscribeOrderBookParams instantiates a new RpcUnsubscribeOrderBookParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcUnsubscribeOrderBookParamsWithDefaults

`func NewRpcUnsubscribeOrderBookParamsWithDefaults() *RpcUnsubscribeOrderBookParams`

NewRpcUnsubscribeOrderBookParamsWithDefaults instantiates a new RpcUnsubscribeOrderBookParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *RpcUnsubscribeOrderBookParams) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *RpcUnsubscribeOrderBookParams) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *RpcUnsubscribeOrderBookParams) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *RpcUnsubscribeOrderBookParams) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetInstrumentId

`func (o *RpcUnsubscribeOrderBookParams) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *RpcUnsubscribeOrderBookParams) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *RpcUnsubscribeOrderBookParams) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *RpcUnsubscribeOrderBookParams) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


