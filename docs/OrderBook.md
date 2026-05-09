# OrderBook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateType** | [**UpdateType**](UpdateType.md) |  | 
**InstrumentId** | Pointer to **string** | Instrument ID in format {VENUE}:{BASE}/{QUOTE} | [optional] 
**Venue** | [**Venue**](Venue.md) |  | 
**Symbol** | **string** | Trading pair symbol in format {BASE}/{QUOTE} | 
**OrderBookType** | [**OrderBookType**](OrderBookType.md) |  | 
**Columns** | **[]string** | Positional layout of each entry in &#x60;bids&#x60; / &#x60;asks&#x60;. Length must equal each row tuple length. Determined by &#x60;orderBookType&#x60; (e.g. &#x60;L2_COUNTED&#x60; → &#x60;[\&quot;price\&quot;, \&quot;quantity\&quot;, \&quot;orderCount\&quot;]&#x60;).  | 
**Bids** | **[][]string** | Bid rows (sorted by price descending) | 
**Asks** | **[][]string** | Ask rows (sorted by price ascending) | 
**Timestamp** | **int64** | Unix timestamp in milliseconds | 

## Methods

### NewOrderBook

`func NewOrderBook(updateType UpdateType, venue Venue, symbol string, orderBookType OrderBookType, columns []string, bids [][]string, asks [][]string, timestamp int64, ) *OrderBook`

NewOrderBook instantiates a new OrderBook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderBookWithDefaults

`func NewOrderBookWithDefaults() *OrderBook`

NewOrderBookWithDefaults instantiates a new OrderBook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateType

`func (o *OrderBook) GetUpdateType() UpdateType`

GetUpdateType returns the UpdateType field if non-nil, zero value otherwise.

### GetUpdateTypeOk

`func (o *OrderBook) GetUpdateTypeOk() (*UpdateType, bool)`

GetUpdateTypeOk returns a tuple with the UpdateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateType

`func (o *OrderBook) SetUpdateType(v UpdateType)`

SetUpdateType sets UpdateType field to given value.


### GetInstrumentId

`func (o *OrderBook) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *OrderBook) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *OrderBook) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *OrderBook) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetVenue

`func (o *OrderBook) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *OrderBook) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *OrderBook) SetVenue(v Venue)`

SetVenue sets Venue field to given value.


### GetSymbol

`func (o *OrderBook) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OrderBook) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OrderBook) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetOrderBookType

`func (o *OrderBook) GetOrderBookType() OrderBookType`

GetOrderBookType returns the OrderBookType field if non-nil, zero value otherwise.

### GetOrderBookTypeOk

`func (o *OrderBook) GetOrderBookTypeOk() (*OrderBookType, bool)`

GetOrderBookTypeOk returns a tuple with the OrderBookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBookType

`func (o *OrderBook) SetOrderBookType(v OrderBookType)`

SetOrderBookType sets OrderBookType field to given value.


### GetColumns

`func (o *OrderBook) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *OrderBook) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *OrderBook) SetColumns(v []string)`

SetColumns sets Columns field to given value.


### GetBids

`func (o *OrderBook) GetBids() [][]string`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *OrderBook) GetBidsOk() (*[][]string, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *OrderBook) SetBids(v [][]string)`

SetBids sets Bids field to given value.


### GetAsks

`func (o *OrderBook) GetAsks() [][]string`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *OrderBook) GetAsksOk() (*[][]string, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *OrderBook) SetAsks(v [][]string)`

SetAsks sets Asks field to given value.


### GetTimestamp

`func (o *OrderBook) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OrderBook) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OrderBook) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


