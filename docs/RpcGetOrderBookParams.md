# RpcGetOrderBookParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstrumentId** | Pointer to **string** | Instrument ID (e.g., BINANCE:BTC/USDT) | [optional] 
**Venue** | Pointer to **string** | Venue (alternative to instrumentId) | [optional] 
**Symbol** | Pointer to **string** | Symbol (alternative to instrumentId) | [optional] 
**Depth** | Pointer to **int32** | Order book depth | [optional] [default to 10]

## Methods

### NewRpcGetOrderBookParams

`func NewRpcGetOrderBookParams() *RpcGetOrderBookParams`

NewRpcGetOrderBookParams instantiates a new RpcGetOrderBookParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcGetOrderBookParamsWithDefaults

`func NewRpcGetOrderBookParamsWithDefaults() *RpcGetOrderBookParams`

NewRpcGetOrderBookParamsWithDefaults instantiates a new RpcGetOrderBookParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstrumentId

`func (o *RpcGetOrderBookParams) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *RpcGetOrderBookParams) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *RpcGetOrderBookParams) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *RpcGetOrderBookParams) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetVenue

`func (o *RpcGetOrderBookParams) GetVenue() string`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *RpcGetOrderBookParams) GetVenueOk() (*string, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *RpcGetOrderBookParams) SetVenue(v string)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *RpcGetOrderBookParams) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetSymbol

`func (o *RpcGetOrderBookParams) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *RpcGetOrderBookParams) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *RpcGetOrderBookParams) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *RpcGetOrderBookParams) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDepth

`func (o *RpcGetOrderBookParams) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *RpcGetOrderBookParams) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *RpcGetOrderBookParams) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *RpcGetOrderBookParams) HasDepth() bool`

HasDepth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


