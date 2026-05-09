# Kline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstrumentId** | Pointer to **string** | Instrument ID in format {VENUE}:{BASE}/{QUOTE} | [optional] 
**Venue** | [**Venue**](Venue.md) |  | 
**Symbol** | **string** | Trading pair symbol in format {BASE}/{QUOTE} | 
**Interval** | [**KlineInterval**](KlineInterval.md) |  | 
**Candles** | **[][]interface{}** | OHLCV tuples ordered ascending by &#x60;openTime&#x60;. | 
**IsClosed** | **bool** | Whether the **last** candle in &#x60;candles&#x60; is finalized. Historical candles before the last are always closed.  | 
**Timestamp** | **int64** | Unix timestamp in milliseconds | 

## Methods

### NewKline

`func NewKline(venue Venue, symbol string, interval KlineInterval, candles [][]interface{}, isClosed bool, timestamp int64, ) *Kline`

NewKline instantiates a new Kline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKlineWithDefaults

`func NewKlineWithDefaults() *Kline`

NewKlineWithDefaults instantiates a new Kline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstrumentId

`func (o *Kline) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *Kline) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *Kline) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *Kline) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetVenue

`func (o *Kline) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *Kline) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *Kline) SetVenue(v Venue)`

SetVenue sets Venue field to given value.


### GetSymbol

`func (o *Kline) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Kline) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Kline) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetInterval

`func (o *Kline) GetInterval() KlineInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Kline) GetIntervalOk() (*KlineInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Kline) SetInterval(v KlineInterval)`

SetInterval sets Interval field to given value.


### GetCandles

`func (o *Kline) GetCandles() [][]interface{}`

GetCandles returns the Candles field if non-nil, zero value otherwise.

### GetCandlesOk

`func (o *Kline) GetCandlesOk() (*[][]interface{}, bool)`

GetCandlesOk returns a tuple with the Candles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandles

`func (o *Kline) SetCandles(v [][]interface{})`

SetCandles sets Candles field to given value.


### GetIsClosed

`func (o *Kline) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *Kline) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *Kline) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.


### GetTimestamp

`func (o *Kline) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Kline) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Kline) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


