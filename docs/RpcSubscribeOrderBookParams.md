# RpcSubscribeOrderBookParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstrumentIds** | Pointer to **[]string** | List of instrument IDs to subscribe | [optional] 
**Venue** | Pointer to **string** | Venue for symbols | [optional] 
**Symbols** | Pointer to **[]string** | List of symbols to subscribe | [optional] 

## Methods

### NewRpcSubscribeOrderBookParams

`func NewRpcSubscribeOrderBookParams() *RpcSubscribeOrderBookParams`

NewRpcSubscribeOrderBookParams instantiates a new RpcSubscribeOrderBookParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcSubscribeOrderBookParamsWithDefaults

`func NewRpcSubscribeOrderBookParamsWithDefaults() *RpcSubscribeOrderBookParams`

NewRpcSubscribeOrderBookParamsWithDefaults instantiates a new RpcSubscribeOrderBookParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstrumentIds

`func (o *RpcSubscribeOrderBookParams) GetInstrumentIds() []string`

GetInstrumentIds returns the InstrumentIds field if non-nil, zero value otherwise.

### GetInstrumentIdsOk

`func (o *RpcSubscribeOrderBookParams) GetInstrumentIdsOk() (*[]string, bool)`

GetInstrumentIdsOk returns a tuple with the InstrumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentIds

`func (o *RpcSubscribeOrderBookParams) SetInstrumentIds(v []string)`

SetInstrumentIds sets InstrumentIds field to given value.

### HasInstrumentIds

`func (o *RpcSubscribeOrderBookParams) HasInstrumentIds() bool`

HasInstrumentIds returns a boolean if a field has been set.

### GetVenue

`func (o *RpcSubscribeOrderBookParams) GetVenue() string`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *RpcSubscribeOrderBookParams) GetVenueOk() (*string, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *RpcSubscribeOrderBookParams) SetVenue(v string)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *RpcSubscribeOrderBookParams) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetSymbols

`func (o *RpcSubscribeOrderBookParams) GetSymbols() []string`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *RpcSubscribeOrderBookParams) GetSymbolsOk() (*[]string, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *RpcSubscribeOrderBookParams) SetSymbols(v []string)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *RpcSubscribeOrderBookParams) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


