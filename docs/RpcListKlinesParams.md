# RpcListKlinesParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstrumentIds** | Pointer to **[]string** |  | [optional] 
**Venue** | Pointer to [**Venue**](Venue.md) |  | [optional] 
**Symbols** | Pointer to **[]string** |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 

## Methods

### NewRpcListKlinesParams

`func NewRpcListKlinesParams() *RpcListKlinesParams`

NewRpcListKlinesParams instantiates a new RpcListKlinesParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcListKlinesParamsWithDefaults

`func NewRpcListKlinesParamsWithDefaults() *RpcListKlinesParams`

NewRpcListKlinesParamsWithDefaults instantiates a new RpcListKlinesParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstrumentIds

`func (o *RpcListKlinesParams) GetInstrumentIds() []string`

GetInstrumentIds returns the InstrumentIds field if non-nil, zero value otherwise.

### GetInstrumentIdsOk

`func (o *RpcListKlinesParams) GetInstrumentIdsOk() (*[]string, bool)`

GetInstrumentIdsOk returns a tuple with the InstrumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentIds

`func (o *RpcListKlinesParams) SetInstrumentIds(v []string)`

SetInstrumentIds sets InstrumentIds field to given value.

### HasInstrumentIds

`func (o *RpcListKlinesParams) HasInstrumentIds() bool`

HasInstrumentIds returns a boolean if a field has been set.

### GetVenue

`func (o *RpcListKlinesParams) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *RpcListKlinesParams) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *RpcListKlinesParams) SetVenue(v Venue)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *RpcListKlinesParams) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetSymbols

`func (o *RpcListKlinesParams) GetSymbols() []string`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *RpcListKlinesParams) GetSymbolsOk() (*[]string, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *RpcListKlinesParams) SetSymbols(v []string)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *RpcListKlinesParams) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.

### GetInterval

`func (o *RpcListKlinesParams) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *RpcListKlinesParams) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *RpcListKlinesParams) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *RpcListKlinesParams) HasInterval() bool`

HasInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


