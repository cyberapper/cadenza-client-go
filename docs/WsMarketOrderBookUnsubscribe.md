# WsMarketOrderBookUnsubscribe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**Channel** | **string** |  | 
**Params** | [**WsMarketOrderBookUnsubscribeParams**](WsMarketOrderBookUnsubscribeParams.md) |  | 

## Methods

### NewWsMarketOrderBookUnsubscribe

`func NewWsMarketOrderBookUnsubscribe(action string, channel string, params WsMarketOrderBookUnsubscribeParams, ) *WsMarketOrderBookUnsubscribe`

NewWsMarketOrderBookUnsubscribe instantiates a new WsMarketOrderBookUnsubscribe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsMarketOrderBookUnsubscribeWithDefaults

`func NewWsMarketOrderBookUnsubscribeWithDefaults() *WsMarketOrderBookUnsubscribe`

NewWsMarketOrderBookUnsubscribeWithDefaults instantiates a new WsMarketOrderBookUnsubscribe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WsMarketOrderBookUnsubscribe) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WsMarketOrderBookUnsubscribe) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WsMarketOrderBookUnsubscribe) SetAction(v string)`

SetAction sets Action field to given value.


### GetChannel

`func (o *WsMarketOrderBookUnsubscribe) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *WsMarketOrderBookUnsubscribe) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *WsMarketOrderBookUnsubscribe) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetParams

`func (o *WsMarketOrderBookUnsubscribe) GetParams() WsMarketOrderBookUnsubscribeParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *WsMarketOrderBookUnsubscribe) GetParamsOk() (*WsMarketOrderBookUnsubscribeParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *WsMarketOrderBookUnsubscribe) SetParams(v WsMarketOrderBookUnsubscribeParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


