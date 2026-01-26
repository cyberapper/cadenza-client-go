# WsMarketOrderBookUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** |  | 
**Data** | [**Orderbook**](Orderbook.md) |  | 
**SubscriptionId** | **string** | UUID string | 
**Timestamp** | **int64** | Unix timestamp in milliseconds | 

## Methods

### NewWsMarketOrderBookUpdate

`func NewWsMarketOrderBookUpdate(channel string, data Orderbook, subscriptionId string, timestamp int64, ) *WsMarketOrderBookUpdate`

NewWsMarketOrderBookUpdate instantiates a new WsMarketOrderBookUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsMarketOrderBookUpdateWithDefaults

`func NewWsMarketOrderBookUpdateWithDefaults() *WsMarketOrderBookUpdate`

NewWsMarketOrderBookUpdateWithDefaults instantiates a new WsMarketOrderBookUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *WsMarketOrderBookUpdate) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *WsMarketOrderBookUpdate) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *WsMarketOrderBookUpdate) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetData

`func (o *WsMarketOrderBookUpdate) GetData() Orderbook`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WsMarketOrderBookUpdate) GetDataOk() (*Orderbook, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WsMarketOrderBookUpdate) SetData(v Orderbook)`

SetData sets Data field to given value.


### GetSubscriptionId

`func (o *WsMarketOrderBookUpdate) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *WsMarketOrderBookUpdate) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *WsMarketOrderBookUpdate) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTimestamp

`func (o *WsMarketOrderBookUpdate) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WsMarketOrderBookUpdate) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WsMarketOrderBookUpdate) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


