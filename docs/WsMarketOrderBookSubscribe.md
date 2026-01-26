# WsMarketOrderBookSubscribe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**Channel** | **string** |  | 
**Params** | [**WsMarketOrderBookSubscribeParams**](WsMarketOrderBookSubscribeParams.md) |  | 

## Methods

### NewWsMarketOrderBookSubscribe

`func NewWsMarketOrderBookSubscribe(action string, channel string, params WsMarketOrderBookSubscribeParams, ) *WsMarketOrderBookSubscribe`

NewWsMarketOrderBookSubscribe instantiates a new WsMarketOrderBookSubscribe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsMarketOrderBookSubscribeWithDefaults

`func NewWsMarketOrderBookSubscribeWithDefaults() *WsMarketOrderBookSubscribe`

NewWsMarketOrderBookSubscribeWithDefaults instantiates a new WsMarketOrderBookSubscribe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WsMarketOrderBookSubscribe) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WsMarketOrderBookSubscribe) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WsMarketOrderBookSubscribe) SetAction(v string)`

SetAction sets Action field to given value.


### GetChannel

`func (o *WsMarketOrderBookSubscribe) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *WsMarketOrderBookSubscribe) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *WsMarketOrderBookSubscribe) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetParams

`func (o *WsMarketOrderBookSubscribe) GetParams() WsMarketOrderBookSubscribeParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *WsMarketOrderBookSubscribe) GetParamsOk() (*WsMarketOrderBookSubscribeParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *WsMarketOrderBookSubscribe) SetParams(v WsMarketOrderBookSubscribeParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


