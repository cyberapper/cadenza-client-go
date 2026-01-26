# RpcTicker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstrumentId** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Exchange** | Pointer to **string** |  | [optional] 
**ExchangeTime** | Pointer to **time.Time** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**Ask** | Pointer to **string** | Ask price (decimal) | [optional] 
**AskQuantity** | Pointer to **string** | Ask quantity (decimal) | [optional] 
**Bid** | Pointer to **string** | Bid price (decimal) | [optional] 
**BidQuantity** | Pointer to **string** | Bid quantity (decimal) | [optional] 
**Last** | Pointer to **string** | Last price (decimal) | [optional] 
**LastQuantity** | Pointer to **string** | Last quantity (decimal) | [optional] 

## Methods

### NewRpcTicker

`func NewRpcTicker() *RpcTicker`

NewRpcTicker instantiates a new RpcTicker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcTickerWithDefaults

`func NewRpcTickerWithDefaults() *RpcTicker`

NewRpcTickerWithDefaults instantiates a new RpcTicker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstrumentId

`func (o *RpcTicker) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *RpcTicker) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *RpcTicker) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *RpcTicker) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetSymbol

`func (o *RpcTicker) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *RpcTicker) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *RpcTicker) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *RpcTicker) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetExchange

`func (o *RpcTicker) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *RpcTicker) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *RpcTicker) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *RpcTicker) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetExchangeTime

`func (o *RpcTicker) GetExchangeTime() time.Time`

GetExchangeTime returns the ExchangeTime field if non-nil, zero value otherwise.

### GetExchangeTimeOk

`func (o *RpcTicker) GetExchangeTimeOk() (*time.Time, bool)`

GetExchangeTimeOk returns a tuple with the ExchangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTime

`func (o *RpcTicker) SetExchangeTime(v time.Time)`

SetExchangeTime sets ExchangeTime field to given value.

### HasExchangeTime

`func (o *RpcTicker) HasExchangeTime() bool`

HasExchangeTime returns a boolean if a field has been set.

### GetTime

`func (o *RpcTicker) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *RpcTicker) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *RpcTicker) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *RpcTicker) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetAsk

`func (o *RpcTicker) GetAsk() string`

GetAsk returns the Ask field if non-nil, zero value otherwise.

### GetAskOk

`func (o *RpcTicker) GetAskOk() (*string, bool)`

GetAskOk returns a tuple with the Ask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsk

`func (o *RpcTicker) SetAsk(v string)`

SetAsk sets Ask field to given value.

### HasAsk

`func (o *RpcTicker) HasAsk() bool`

HasAsk returns a boolean if a field has been set.

### GetAskQuantity

`func (o *RpcTicker) GetAskQuantity() string`

GetAskQuantity returns the AskQuantity field if non-nil, zero value otherwise.

### GetAskQuantityOk

`func (o *RpcTicker) GetAskQuantityOk() (*string, bool)`

GetAskQuantityOk returns a tuple with the AskQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskQuantity

`func (o *RpcTicker) SetAskQuantity(v string)`

SetAskQuantity sets AskQuantity field to given value.

### HasAskQuantity

`func (o *RpcTicker) HasAskQuantity() bool`

HasAskQuantity returns a boolean if a field has been set.

### GetBid

`func (o *RpcTicker) GetBid() string`

GetBid returns the Bid field if non-nil, zero value otherwise.

### GetBidOk

`func (o *RpcTicker) GetBidOk() (*string, bool)`

GetBidOk returns a tuple with the Bid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBid

`func (o *RpcTicker) SetBid(v string)`

SetBid sets Bid field to given value.

### HasBid

`func (o *RpcTicker) HasBid() bool`

HasBid returns a boolean if a field has been set.

### GetBidQuantity

`func (o *RpcTicker) GetBidQuantity() string`

GetBidQuantity returns the BidQuantity field if non-nil, zero value otherwise.

### GetBidQuantityOk

`func (o *RpcTicker) GetBidQuantityOk() (*string, bool)`

GetBidQuantityOk returns a tuple with the BidQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidQuantity

`func (o *RpcTicker) SetBidQuantity(v string)`

SetBidQuantity sets BidQuantity field to given value.

### HasBidQuantity

`func (o *RpcTicker) HasBidQuantity() bool`

HasBidQuantity returns a boolean if a field has been set.

### GetLast

`func (o *RpcTicker) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *RpcTicker) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *RpcTicker) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *RpcTicker) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetLastQuantity

`func (o *RpcTicker) GetLastQuantity() string`

GetLastQuantity returns the LastQuantity field if non-nil, zero value otherwise.

### GetLastQuantityOk

`func (o *RpcTicker) GetLastQuantityOk() (*string, bool)`

GetLastQuantityOk returns a tuple with the LastQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQuantity

`func (o *RpcTicker) SetLastQuantity(v string)`

SetLastQuantity sets LastQuantity field to given value.

### HasLastQuantity

`func (o *RpcTicker) HasLastQuantity() bool`

HasLastQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


