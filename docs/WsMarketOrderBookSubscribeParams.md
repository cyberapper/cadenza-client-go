# WsMarketOrderBookSubscribeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Venue** | [**Venue**](Venue.md) |  | 
**Symbol** | **string** |  | 
**Depth** | Pointer to **int32** |  | [optional] [default to 20]

## Methods

### NewWsMarketOrderBookSubscribeParams

`func NewWsMarketOrderBookSubscribeParams(venue Venue, symbol string, ) *WsMarketOrderBookSubscribeParams`

NewWsMarketOrderBookSubscribeParams instantiates a new WsMarketOrderBookSubscribeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsMarketOrderBookSubscribeParamsWithDefaults

`func NewWsMarketOrderBookSubscribeParamsWithDefaults() *WsMarketOrderBookSubscribeParams`

NewWsMarketOrderBookSubscribeParamsWithDefaults instantiates a new WsMarketOrderBookSubscribeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVenue

`func (o *WsMarketOrderBookSubscribeParams) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *WsMarketOrderBookSubscribeParams) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *WsMarketOrderBookSubscribeParams) SetVenue(v Venue)`

SetVenue sets Venue field to given value.


### GetSymbol

`func (o *WsMarketOrderBookSubscribeParams) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *WsMarketOrderBookSubscribeParams) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *WsMarketOrderBookSubscribeParams) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetDepth

`func (o *WsMarketOrderBookSubscribeParams) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *WsMarketOrderBookSubscribeParams) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *WsMarketOrderBookSubscribeParams) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *WsMarketOrderBookSubscribeParams) HasDepth() bool`

HasDepth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


