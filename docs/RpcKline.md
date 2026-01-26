# RpcKline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstrumentId** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**Candles** | Pointer to [**[]RpcOhlcv**](RpcOhlcv.md) |  | [optional] 

## Methods

### NewRpcKline

`func NewRpcKline() *RpcKline`

NewRpcKline instantiates a new RpcKline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcKlineWithDefaults

`func NewRpcKlineWithDefaults() *RpcKline`

NewRpcKlineWithDefaults instantiates a new RpcKline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstrumentId

`func (o *RpcKline) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *RpcKline) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *RpcKline) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *RpcKline) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetSymbol

`func (o *RpcKline) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *RpcKline) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *RpcKline) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *RpcKline) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetInterval

`func (o *RpcKline) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *RpcKline) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *RpcKline) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *RpcKline) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetCandles

`func (o *RpcKline) GetCandles() []RpcOhlcv`

GetCandles returns the Candles field if non-nil, zero value otherwise.

### GetCandlesOk

`func (o *RpcKline) GetCandlesOk() (*[]RpcOhlcv, bool)`

GetCandlesOk returns a tuple with the Candles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandles

`func (o *RpcKline) SetCandles(v []RpcOhlcv)`

SetCandles sets Candles field to given value.

### HasCandles

`func (o *RpcKline) HasCandles() bool`

HasCandles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


