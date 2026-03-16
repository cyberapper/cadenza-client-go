# RpcListTradeOrdersParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradeOrderId** | Pointer to **string** | Filter by specific trade order ID | [optional] 
**TradingAccountId** | Pointer to **string** | Filter by trading account ID | [optional] 
**InstrumentId** | Pointer to **string** | Filter by instrument ID (e.g., BINANCE:BTC/USDT) | [optional] 
**Side** | Pointer to [**OrderSide**](OrderSide.md) |  | [optional] 
**OrderType** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**Status** | Pointer to [**OrderStatus**](OrderStatus.md) |  | [optional] 
**StartTime** | Pointer to **time.Time** | Filter orders created after this time | [optional] 
**EndTime** | Pointer to **time.Time** | Filter orders created before this time | [optional] 
**Pagination** | Pointer to [**RpcPagination**](RpcPagination.md) |  | [optional] 

## Methods

### NewRpcListTradeOrdersParams

`func NewRpcListTradeOrdersParams() *RpcListTradeOrdersParams`

NewRpcListTradeOrdersParams instantiates a new RpcListTradeOrdersParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcListTradeOrdersParamsWithDefaults

`func NewRpcListTradeOrdersParamsWithDefaults() *RpcListTradeOrdersParams`

NewRpcListTradeOrdersParamsWithDefaults instantiates a new RpcListTradeOrdersParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradeOrderId

`func (o *RpcListTradeOrdersParams) GetTradeOrderId() string`

GetTradeOrderId returns the TradeOrderId field if non-nil, zero value otherwise.

### GetTradeOrderIdOk

`func (o *RpcListTradeOrdersParams) GetTradeOrderIdOk() (*string, bool)`

GetTradeOrderIdOk returns a tuple with the TradeOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeOrderId

`func (o *RpcListTradeOrdersParams) SetTradeOrderId(v string)`

SetTradeOrderId sets TradeOrderId field to given value.

### HasTradeOrderId

`func (o *RpcListTradeOrdersParams) HasTradeOrderId() bool`

HasTradeOrderId returns a boolean if a field has been set.

### GetTradingAccountId

`func (o *RpcListTradeOrdersParams) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *RpcListTradeOrdersParams) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *RpcListTradeOrdersParams) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.

### HasTradingAccountId

`func (o *RpcListTradeOrdersParams) HasTradingAccountId() bool`

HasTradingAccountId returns a boolean if a field has been set.

### GetInstrumentId

`func (o *RpcListTradeOrdersParams) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *RpcListTradeOrdersParams) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *RpcListTradeOrdersParams) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *RpcListTradeOrdersParams) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetSide

`func (o *RpcListTradeOrdersParams) GetSide() OrderSide`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *RpcListTradeOrdersParams) GetSideOk() (*OrderSide, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *RpcListTradeOrdersParams) SetSide(v OrderSide)`

SetSide sets Side field to given value.

### HasSide

`func (o *RpcListTradeOrdersParams) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetOrderType

`func (o *RpcListTradeOrdersParams) GetOrderType() OrderType`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *RpcListTradeOrdersParams) GetOrderTypeOk() (*OrderType, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *RpcListTradeOrdersParams) SetOrderType(v OrderType)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *RpcListTradeOrdersParams) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetStatus

`func (o *RpcListTradeOrdersParams) GetStatus() OrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RpcListTradeOrdersParams) GetStatusOk() (*OrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RpcListTradeOrdersParams) SetStatus(v OrderStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RpcListTradeOrdersParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartTime

`func (o *RpcListTradeOrdersParams) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RpcListTradeOrdersParams) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RpcListTradeOrdersParams) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *RpcListTradeOrdersParams) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *RpcListTradeOrdersParams) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *RpcListTradeOrdersParams) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *RpcListTradeOrdersParams) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *RpcListTradeOrdersParams) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetPagination

`func (o *RpcListTradeOrdersParams) GetPagination() RpcPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RpcListTradeOrdersParams) GetPaginationOk() (*RpcPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RpcListTradeOrdersParams) SetPagination(v RpcPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RpcListTradeOrdersParams) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


