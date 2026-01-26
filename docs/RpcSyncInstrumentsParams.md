# RpcSyncInstrumentsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Venue** | [**Venue**](Venue.md) |  | 
**Symbols** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRpcSyncInstrumentsParams

`func NewRpcSyncInstrumentsParams(venue Venue, ) *RpcSyncInstrumentsParams`

NewRpcSyncInstrumentsParams instantiates a new RpcSyncInstrumentsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcSyncInstrumentsParamsWithDefaults

`func NewRpcSyncInstrumentsParamsWithDefaults() *RpcSyncInstrumentsParams`

NewRpcSyncInstrumentsParamsWithDefaults instantiates a new RpcSyncInstrumentsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVenue

`func (o *RpcSyncInstrumentsParams) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *RpcSyncInstrumentsParams) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *RpcSyncInstrumentsParams) SetVenue(v Venue)`

SetVenue sets Venue field to given value.


### GetSymbols

`func (o *RpcSyncInstrumentsParams) GetSymbols() []string`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *RpcSyncInstrumentsParams) GetSymbolsOk() (*[]string, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *RpcSyncInstrumentsParams) SetSymbols(v []string)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *RpcSyncInstrumentsParams) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


