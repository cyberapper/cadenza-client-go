# SyncMarketInstrumentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Venue** | [**Venue**](Venue.md) |  | 
**Symbols** | **[]string** | symbol list, example \\[\&quot;BTC/ETH\&quot;,\&quot;BTC/USDT\&quot;\\] | 

## Methods

### NewSyncMarketInstrumentsRequest

`func NewSyncMarketInstrumentsRequest(venue Venue, symbols []string, ) *SyncMarketInstrumentsRequest`

NewSyncMarketInstrumentsRequest instantiates a new SyncMarketInstrumentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncMarketInstrumentsRequestWithDefaults

`func NewSyncMarketInstrumentsRequestWithDefaults() *SyncMarketInstrumentsRequest`

NewSyncMarketInstrumentsRequestWithDefaults instantiates a new SyncMarketInstrumentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVenue

`func (o *SyncMarketInstrumentsRequest) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *SyncMarketInstrumentsRequest) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *SyncMarketInstrumentsRequest) SetVenue(v Venue)`

SetVenue sets Venue field to given value.


### GetSymbols

`func (o *SyncMarketInstrumentsRequest) GetSymbols() []string`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *SyncMarketInstrumentsRequest) GetSymbolsOk() (*[]string, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *SyncMarketInstrumentsRequest) SetSymbols(v []string)`

SetSymbols sets Symbols field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


