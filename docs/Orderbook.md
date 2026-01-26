# Orderbook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstrumentId** | Pointer to **string** | Instrument ID in format {VENUE}:{BASE}/{QUOTE} | [optional] 
**Venue** | [**Venue**](Venue.md) |  | 
**Symbol** | **string** |  | 
**Asks** | **[][]string** |  | 
**Bids** | **[][]string** |  | 
**Timestamp** | **int64** | Unix timestamp in milliseconds | 

## Methods

### NewOrderbook

`func NewOrderbook(venue Venue, symbol string, asks [][]string, bids [][]string, timestamp int64, ) *Orderbook`

NewOrderbook instantiates a new Orderbook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderbookWithDefaults

`func NewOrderbookWithDefaults() *Orderbook`

NewOrderbookWithDefaults instantiates a new Orderbook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstrumentId

`func (o *Orderbook) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *Orderbook) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *Orderbook) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *Orderbook) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetVenue

`func (o *Orderbook) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *Orderbook) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *Orderbook) SetVenue(v Venue)`

SetVenue sets Venue field to given value.


### GetSymbol

`func (o *Orderbook) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Orderbook) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Orderbook) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetAsks

`func (o *Orderbook) GetAsks() [][]string`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *Orderbook) GetAsksOk() (*[][]string, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *Orderbook) SetAsks(v [][]string)`

SetAsks sets Asks field to given value.


### GetBids

`func (o *Orderbook) GetBids() [][]string`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *Orderbook) GetBidsOk() (*[][]string, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *Orderbook) SetBids(v [][]string)`

SetBids sets Bids field to given value.


### GetTimestamp

`func (o *Orderbook) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Orderbook) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Orderbook) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


