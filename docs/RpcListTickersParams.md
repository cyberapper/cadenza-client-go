# RpcListTickersParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstrumentIds** | Pointer to **[]string** |  | [optional] 
**Venue** | Pointer to [**Venue**](Venue.md) |  | [optional] 
**Symbols** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRpcListTickersParams

`func NewRpcListTickersParams() *RpcListTickersParams`

NewRpcListTickersParams instantiates a new RpcListTickersParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcListTickersParamsWithDefaults

`func NewRpcListTickersParamsWithDefaults() *RpcListTickersParams`

NewRpcListTickersParamsWithDefaults instantiates a new RpcListTickersParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstrumentIds

`func (o *RpcListTickersParams) GetInstrumentIds() []string`

GetInstrumentIds returns the InstrumentIds field if non-nil, zero value otherwise.

### GetInstrumentIdsOk

`func (o *RpcListTickersParams) GetInstrumentIdsOk() (*[]string, bool)`

GetInstrumentIdsOk returns a tuple with the InstrumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentIds

`func (o *RpcListTickersParams) SetInstrumentIds(v []string)`

SetInstrumentIds sets InstrumentIds field to given value.

### HasInstrumentIds

`func (o *RpcListTickersParams) HasInstrumentIds() bool`

HasInstrumentIds returns a boolean if a field has been set.

### GetVenue

`func (o *RpcListTickersParams) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *RpcListTickersParams) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *RpcListTickersParams) SetVenue(v Venue)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *RpcListTickersParams) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetSymbols

`func (o *RpcListTickersParams) GetSymbols() []string`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *RpcListTickersParams) GetSymbolsOk() (*[]string, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *RpcListTickersParams) SetSymbols(v []string)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *RpcListTickersParams) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


