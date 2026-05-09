# Ticker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstrumentId** | Pointer to **string** | Instrument ID in format {VENUE}:{BASE}/{QUOTE} | [optional] 
**Venue** | [**Venue**](Venue.md) |  | 
**Symbol** | **string** | Trading pair symbol in format {BASE}/{QUOTE} | 
**LastPrice** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**LastQuantity** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**BidPrice** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**BidQuantity** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**AskPrice** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**AskQuantity** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**Timestamp** | **int64** | Unix timestamp in milliseconds | 

## Methods

### NewTicker

`func NewTicker(venue Venue, symbol string, timestamp int64, ) *Ticker`

NewTicker instantiates a new Ticker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTickerWithDefaults

`func NewTickerWithDefaults() *Ticker`

NewTickerWithDefaults instantiates a new Ticker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstrumentId

`func (o *Ticker) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *Ticker) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *Ticker) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *Ticker) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetVenue

`func (o *Ticker) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *Ticker) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *Ticker) SetVenue(v Venue)`

SetVenue sets Venue field to given value.


### GetSymbol

`func (o *Ticker) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Ticker) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Ticker) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetLastPrice

`func (o *Ticker) GetLastPrice() string`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *Ticker) GetLastPriceOk() (*string, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *Ticker) SetLastPrice(v string)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *Ticker) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLastQuantity

`func (o *Ticker) GetLastQuantity() string`

GetLastQuantity returns the LastQuantity field if non-nil, zero value otherwise.

### GetLastQuantityOk

`func (o *Ticker) GetLastQuantityOk() (*string, bool)`

GetLastQuantityOk returns a tuple with the LastQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQuantity

`func (o *Ticker) SetLastQuantity(v string)`

SetLastQuantity sets LastQuantity field to given value.

### HasLastQuantity

`func (o *Ticker) HasLastQuantity() bool`

HasLastQuantity returns a boolean if a field has been set.

### GetBidPrice

`func (o *Ticker) GetBidPrice() string`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *Ticker) GetBidPriceOk() (*string, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *Ticker) SetBidPrice(v string)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *Ticker) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetBidQuantity

`func (o *Ticker) GetBidQuantity() string`

GetBidQuantity returns the BidQuantity field if non-nil, zero value otherwise.

### GetBidQuantityOk

`func (o *Ticker) GetBidQuantityOk() (*string, bool)`

GetBidQuantityOk returns a tuple with the BidQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidQuantity

`func (o *Ticker) SetBidQuantity(v string)`

SetBidQuantity sets BidQuantity field to given value.

### HasBidQuantity

`func (o *Ticker) HasBidQuantity() bool`

HasBidQuantity returns a boolean if a field has been set.

### GetAskPrice

`func (o *Ticker) GetAskPrice() string`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *Ticker) GetAskPriceOk() (*string, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *Ticker) SetAskPrice(v string)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *Ticker) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetAskQuantity

`func (o *Ticker) GetAskQuantity() string`

GetAskQuantity returns the AskQuantity field if non-nil, zero value otherwise.

### GetAskQuantityOk

`func (o *Ticker) GetAskQuantityOk() (*string, bool)`

GetAskQuantityOk returns a tuple with the AskQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskQuantity

`func (o *Ticker) SetAskQuantity(v string)`

SetAskQuantity sets AskQuantity field to given value.

### HasAskQuantity

`func (o *Ticker) HasAskQuantity() bool`

HasAskQuantity returns a boolean if a field has been set.

### GetTimestamp

`func (o *Ticker) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Ticker) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Ticker) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


