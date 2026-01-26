# RpcOrderBook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstrumentId** | Pointer to **string** | Instrument ID | [optional] 
**Venue** | Pointer to **string** | Venue | [optional] 
**Symbol** | Pointer to **string** | Trading pair symbol | [optional] 
**Bids** | Pointer to [**[]RpcOrderBookLevel**](RpcOrderBookLevel.md) | Bid orders (sorted by price descending) | [optional] 
**Asks** | Pointer to [**[]RpcOrderBookLevel**](RpcOrderBookLevel.md) | Ask orders (sorted by price ascending) | [optional] 
**Timestamp** | Pointer to **time.Time** | Timestamp in ISO 8601 format (RFC3339). This is the native format used by Go&#39;s time.Time. | [optional] 
**SequenceNumber** | Pointer to **int64** | Sequence number for ordering updates | [optional] 

## Methods

### NewRpcOrderBook

`func NewRpcOrderBook() *RpcOrderBook`

NewRpcOrderBook instantiates a new RpcOrderBook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcOrderBookWithDefaults

`func NewRpcOrderBookWithDefaults() *RpcOrderBook`

NewRpcOrderBookWithDefaults instantiates a new RpcOrderBook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstrumentId

`func (o *RpcOrderBook) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *RpcOrderBook) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *RpcOrderBook) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *RpcOrderBook) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetVenue

`func (o *RpcOrderBook) GetVenue() string`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *RpcOrderBook) GetVenueOk() (*string, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *RpcOrderBook) SetVenue(v string)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *RpcOrderBook) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetSymbol

`func (o *RpcOrderBook) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *RpcOrderBook) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *RpcOrderBook) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *RpcOrderBook) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetBids

`func (o *RpcOrderBook) GetBids() []RpcOrderBookLevel`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *RpcOrderBook) GetBidsOk() (*[]RpcOrderBookLevel, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *RpcOrderBook) SetBids(v []RpcOrderBookLevel)`

SetBids sets Bids field to given value.

### HasBids

`func (o *RpcOrderBook) HasBids() bool`

HasBids returns a boolean if a field has been set.

### GetAsks

`func (o *RpcOrderBook) GetAsks() []RpcOrderBookLevel`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *RpcOrderBook) GetAsksOk() (*[]RpcOrderBookLevel, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *RpcOrderBook) SetAsks(v []RpcOrderBookLevel)`

SetAsks sets Asks field to given value.

### HasAsks

`func (o *RpcOrderBook) HasAsks() bool`

HasAsks returns a boolean if a field has been set.

### GetTimestamp

`func (o *RpcOrderBook) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RpcOrderBook) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RpcOrderBook) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RpcOrderBook) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *RpcOrderBook) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *RpcOrderBook) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *RpcOrderBook) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *RpcOrderBook) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


