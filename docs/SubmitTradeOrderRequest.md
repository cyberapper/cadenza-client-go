# SubmitTradeOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradingAccountId** | **string** | UUID string | 
**InstrumentId** | **string** | Instrument ID in format {VENUE}:{BASE}/{QUOTE} | 
**IdempotencyKey** | Pointer to **string** | Idempotency key to prevent duplicate request processing | [optional] 
**ClientOrderId** | Pointer to **string** | Client-provided order ID, used as idempotency key | [optional] 
**OrderSide** | [**OrderSide**](OrderSide.md) |  | 
**OrderType** | [**NullableOrderType**](OrderType.md) |  | 
**LimitPrice** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**StopPrice** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**Quantity** | **string** | Decimal value as string to preserve precision | 
**QuantityType** | Pointer to [**OrderQuantityType**](OrderQuantityType.md) |  | [optional] 
**QuantityRounding** | Pointer to [**QuantityRounding**](QuantityRounding.md) |  | [optional] [default to QUANTITYROUNDING_EMPTY]
**PositionId** | Pointer to **string** | UUID string | [optional] 
**TimeInForce** | Pointer to [**NullableTimeInForce**](TimeInForce.md) |  | [optional] 
**ExpireAt** | Pointer to **int64** | Unix timestamp in milliseconds | [optional] 
**QuoteId** | Pointer to **string** | UUID string | [optional] 
**Leverage** | Pointer to **int32** | Leverage | [optional] 
**AwaitClosed** | Pointer to **bool** | If true, the API will wait up to 1 second for the order to reach a closed/finalized state (FILLED, REJECTED, EXPIRED, CANCELLED) before responding. If false or omitted, returns immediately with the initial order state. Useful for market orders that typically fill immediately.  | [optional] [default to false]
**TakeProfitPrice** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**TakeProfitLimitPrice** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**StopLossPrice** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**StopLossLimitPrice** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**TakeProfitTimeInForce** | Pointer to [**NullableTimeInForce**](TimeInForce.md) |  | [optional] 
**StopLossTimeInForce** | Pointer to [**NullableTimeInForce**](TimeInForce.md) |  | [optional] 

## Methods

### NewSubmitTradeOrderRequest

`func NewSubmitTradeOrderRequest(tradingAccountId string, instrumentId string, orderSide OrderSide, orderType NullableOrderType, quantity string, ) *SubmitTradeOrderRequest`

NewSubmitTradeOrderRequest instantiates a new SubmitTradeOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitTradeOrderRequestWithDefaults

`func NewSubmitTradeOrderRequestWithDefaults() *SubmitTradeOrderRequest`

NewSubmitTradeOrderRequestWithDefaults instantiates a new SubmitTradeOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradingAccountId

`func (o *SubmitTradeOrderRequest) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *SubmitTradeOrderRequest) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *SubmitTradeOrderRequest) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.


### GetInstrumentId

`func (o *SubmitTradeOrderRequest) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *SubmitTradeOrderRequest) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *SubmitTradeOrderRequest) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.


### GetIdempotencyKey

`func (o *SubmitTradeOrderRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *SubmitTradeOrderRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *SubmitTradeOrderRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *SubmitTradeOrderRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetClientOrderId

`func (o *SubmitTradeOrderRequest) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *SubmitTradeOrderRequest) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *SubmitTradeOrderRequest) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *SubmitTradeOrderRequest) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetOrderSide

`func (o *SubmitTradeOrderRequest) GetOrderSide() OrderSide`

GetOrderSide returns the OrderSide field if non-nil, zero value otherwise.

### GetOrderSideOk

`func (o *SubmitTradeOrderRequest) GetOrderSideOk() (*OrderSide, bool)`

GetOrderSideOk returns a tuple with the OrderSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSide

`func (o *SubmitTradeOrderRequest) SetOrderSide(v OrderSide)`

SetOrderSide sets OrderSide field to given value.


### GetOrderType

`func (o *SubmitTradeOrderRequest) GetOrderType() OrderType`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *SubmitTradeOrderRequest) GetOrderTypeOk() (*OrderType, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *SubmitTradeOrderRequest) SetOrderType(v OrderType)`

SetOrderType sets OrderType field to given value.


### SetOrderTypeNil

`func (o *SubmitTradeOrderRequest) SetOrderTypeNil(b bool)`

 SetOrderTypeNil sets the value for OrderType to be an explicit nil

### UnsetOrderType
`func (o *SubmitTradeOrderRequest) UnsetOrderType()`

UnsetOrderType ensures that no value is present for OrderType, not even an explicit nil
### GetLimitPrice

`func (o *SubmitTradeOrderRequest) GetLimitPrice() string`

GetLimitPrice returns the LimitPrice field if non-nil, zero value otherwise.

### GetLimitPriceOk

`func (o *SubmitTradeOrderRequest) GetLimitPriceOk() (*string, bool)`

GetLimitPriceOk returns a tuple with the LimitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPrice

`func (o *SubmitTradeOrderRequest) SetLimitPrice(v string)`

SetLimitPrice sets LimitPrice field to given value.

### HasLimitPrice

`func (o *SubmitTradeOrderRequest) HasLimitPrice() bool`

HasLimitPrice returns a boolean if a field has been set.

### GetStopPrice

`func (o *SubmitTradeOrderRequest) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *SubmitTradeOrderRequest) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *SubmitTradeOrderRequest) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *SubmitTradeOrderRequest) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *SubmitTradeOrderRequest) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SubmitTradeOrderRequest) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SubmitTradeOrderRequest) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetQuantityType

`func (o *SubmitTradeOrderRequest) GetQuantityType() OrderQuantityType`

GetQuantityType returns the QuantityType field if non-nil, zero value otherwise.

### GetQuantityTypeOk

`func (o *SubmitTradeOrderRequest) GetQuantityTypeOk() (*OrderQuantityType, bool)`

GetQuantityTypeOk returns a tuple with the QuantityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityType

`func (o *SubmitTradeOrderRequest) SetQuantityType(v OrderQuantityType)`

SetQuantityType sets QuantityType field to given value.

### HasQuantityType

`func (o *SubmitTradeOrderRequest) HasQuantityType() bool`

HasQuantityType returns a boolean if a field has been set.

### GetQuantityRounding

`func (o *SubmitTradeOrderRequest) GetQuantityRounding() QuantityRounding`

GetQuantityRounding returns the QuantityRounding field if non-nil, zero value otherwise.

### GetQuantityRoundingOk

`func (o *SubmitTradeOrderRequest) GetQuantityRoundingOk() (*QuantityRounding, bool)`

GetQuantityRoundingOk returns a tuple with the QuantityRounding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityRounding

`func (o *SubmitTradeOrderRequest) SetQuantityRounding(v QuantityRounding)`

SetQuantityRounding sets QuantityRounding field to given value.

### HasQuantityRounding

`func (o *SubmitTradeOrderRequest) HasQuantityRounding() bool`

HasQuantityRounding returns a boolean if a field has been set.

### GetPositionId

`func (o *SubmitTradeOrderRequest) GetPositionId() string`

GetPositionId returns the PositionId field if non-nil, zero value otherwise.

### GetPositionIdOk

`func (o *SubmitTradeOrderRequest) GetPositionIdOk() (*string, bool)`

GetPositionIdOk returns a tuple with the PositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionId

`func (o *SubmitTradeOrderRequest) SetPositionId(v string)`

SetPositionId sets PositionId field to given value.

### HasPositionId

`func (o *SubmitTradeOrderRequest) HasPositionId() bool`

HasPositionId returns a boolean if a field has been set.

### GetTimeInForce

`func (o *SubmitTradeOrderRequest) GetTimeInForce() TimeInForce`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *SubmitTradeOrderRequest) GetTimeInForceOk() (*TimeInForce, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *SubmitTradeOrderRequest) SetTimeInForce(v TimeInForce)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *SubmitTradeOrderRequest) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### SetTimeInForceNil

`func (o *SubmitTradeOrderRequest) SetTimeInForceNil(b bool)`

 SetTimeInForceNil sets the value for TimeInForce to be an explicit nil

### UnsetTimeInForce
`func (o *SubmitTradeOrderRequest) UnsetTimeInForce()`

UnsetTimeInForce ensures that no value is present for TimeInForce, not even an explicit nil
### GetExpireAt

`func (o *SubmitTradeOrderRequest) GetExpireAt() int64`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *SubmitTradeOrderRequest) GetExpireAtOk() (*int64, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *SubmitTradeOrderRequest) SetExpireAt(v int64)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *SubmitTradeOrderRequest) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### GetQuoteId

`func (o *SubmitTradeOrderRequest) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *SubmitTradeOrderRequest) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *SubmitTradeOrderRequest) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.

### HasQuoteId

`func (o *SubmitTradeOrderRequest) HasQuoteId() bool`

HasQuoteId returns a boolean if a field has been set.

### GetLeverage

`func (o *SubmitTradeOrderRequest) GetLeverage() int32`

GetLeverage returns the Leverage field if non-nil, zero value otherwise.

### GetLeverageOk

`func (o *SubmitTradeOrderRequest) GetLeverageOk() (*int32, bool)`

GetLeverageOk returns a tuple with the Leverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverage

`func (o *SubmitTradeOrderRequest) SetLeverage(v int32)`

SetLeverage sets Leverage field to given value.

### HasLeverage

`func (o *SubmitTradeOrderRequest) HasLeverage() bool`

HasLeverage returns a boolean if a field has been set.

### GetAwaitClosed

`func (o *SubmitTradeOrderRequest) GetAwaitClosed() bool`

GetAwaitClosed returns the AwaitClosed field if non-nil, zero value otherwise.

### GetAwaitClosedOk

`func (o *SubmitTradeOrderRequest) GetAwaitClosedOk() (*bool, bool)`

GetAwaitClosedOk returns a tuple with the AwaitClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaitClosed

`func (o *SubmitTradeOrderRequest) SetAwaitClosed(v bool)`

SetAwaitClosed sets AwaitClosed field to given value.

### HasAwaitClosed

`func (o *SubmitTradeOrderRequest) HasAwaitClosed() bool`

HasAwaitClosed returns a boolean if a field has been set.

### GetTakeProfitPrice

`func (o *SubmitTradeOrderRequest) GetTakeProfitPrice() string`

GetTakeProfitPrice returns the TakeProfitPrice field if non-nil, zero value otherwise.

### GetTakeProfitPriceOk

`func (o *SubmitTradeOrderRequest) GetTakeProfitPriceOk() (*string, bool)`

GetTakeProfitPriceOk returns a tuple with the TakeProfitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakeProfitPrice

`func (o *SubmitTradeOrderRequest) SetTakeProfitPrice(v string)`

SetTakeProfitPrice sets TakeProfitPrice field to given value.

### HasTakeProfitPrice

`func (o *SubmitTradeOrderRequest) HasTakeProfitPrice() bool`

HasTakeProfitPrice returns a boolean if a field has been set.

### GetTakeProfitLimitPrice

`func (o *SubmitTradeOrderRequest) GetTakeProfitLimitPrice() string`

GetTakeProfitLimitPrice returns the TakeProfitLimitPrice field if non-nil, zero value otherwise.

### GetTakeProfitLimitPriceOk

`func (o *SubmitTradeOrderRequest) GetTakeProfitLimitPriceOk() (*string, bool)`

GetTakeProfitLimitPriceOk returns a tuple with the TakeProfitLimitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakeProfitLimitPrice

`func (o *SubmitTradeOrderRequest) SetTakeProfitLimitPrice(v string)`

SetTakeProfitLimitPrice sets TakeProfitLimitPrice field to given value.

### HasTakeProfitLimitPrice

`func (o *SubmitTradeOrderRequest) HasTakeProfitLimitPrice() bool`

HasTakeProfitLimitPrice returns a boolean if a field has been set.

### GetStopLossPrice

`func (o *SubmitTradeOrderRequest) GetStopLossPrice() string`

GetStopLossPrice returns the StopLossPrice field if non-nil, zero value otherwise.

### GetStopLossPriceOk

`func (o *SubmitTradeOrderRequest) GetStopLossPriceOk() (*string, bool)`

GetStopLossPriceOk returns a tuple with the StopLossPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopLossPrice

`func (o *SubmitTradeOrderRequest) SetStopLossPrice(v string)`

SetStopLossPrice sets StopLossPrice field to given value.

### HasStopLossPrice

`func (o *SubmitTradeOrderRequest) HasStopLossPrice() bool`

HasStopLossPrice returns a boolean if a field has been set.

### GetStopLossLimitPrice

`func (o *SubmitTradeOrderRequest) GetStopLossLimitPrice() string`

GetStopLossLimitPrice returns the StopLossLimitPrice field if non-nil, zero value otherwise.

### GetStopLossLimitPriceOk

`func (o *SubmitTradeOrderRequest) GetStopLossLimitPriceOk() (*string, bool)`

GetStopLossLimitPriceOk returns a tuple with the StopLossLimitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopLossLimitPrice

`func (o *SubmitTradeOrderRequest) SetStopLossLimitPrice(v string)`

SetStopLossLimitPrice sets StopLossLimitPrice field to given value.

### HasStopLossLimitPrice

`func (o *SubmitTradeOrderRequest) HasStopLossLimitPrice() bool`

HasStopLossLimitPrice returns a boolean if a field has been set.

### GetTakeProfitTimeInForce

`func (o *SubmitTradeOrderRequest) GetTakeProfitTimeInForce() TimeInForce`

GetTakeProfitTimeInForce returns the TakeProfitTimeInForce field if non-nil, zero value otherwise.

### GetTakeProfitTimeInForceOk

`func (o *SubmitTradeOrderRequest) GetTakeProfitTimeInForceOk() (*TimeInForce, bool)`

GetTakeProfitTimeInForceOk returns a tuple with the TakeProfitTimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakeProfitTimeInForce

`func (o *SubmitTradeOrderRequest) SetTakeProfitTimeInForce(v TimeInForce)`

SetTakeProfitTimeInForce sets TakeProfitTimeInForce field to given value.

### HasTakeProfitTimeInForce

`func (o *SubmitTradeOrderRequest) HasTakeProfitTimeInForce() bool`

HasTakeProfitTimeInForce returns a boolean if a field has been set.

### SetTakeProfitTimeInForceNil

`func (o *SubmitTradeOrderRequest) SetTakeProfitTimeInForceNil(b bool)`

 SetTakeProfitTimeInForceNil sets the value for TakeProfitTimeInForce to be an explicit nil

### UnsetTakeProfitTimeInForce
`func (o *SubmitTradeOrderRequest) UnsetTakeProfitTimeInForce()`

UnsetTakeProfitTimeInForce ensures that no value is present for TakeProfitTimeInForce, not even an explicit nil
### GetStopLossTimeInForce

`func (o *SubmitTradeOrderRequest) GetStopLossTimeInForce() TimeInForce`

GetStopLossTimeInForce returns the StopLossTimeInForce field if non-nil, zero value otherwise.

### GetStopLossTimeInForceOk

`func (o *SubmitTradeOrderRequest) GetStopLossTimeInForceOk() (*TimeInForce, bool)`

GetStopLossTimeInForceOk returns a tuple with the StopLossTimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopLossTimeInForce

`func (o *SubmitTradeOrderRequest) SetStopLossTimeInForce(v TimeInForce)`

SetStopLossTimeInForce sets StopLossTimeInForce field to given value.

### HasStopLossTimeInForce

`func (o *SubmitTradeOrderRequest) HasStopLossTimeInForce() bool`

HasStopLossTimeInForce returns a boolean if a field has been set.

### SetStopLossTimeInForceNil

`func (o *SubmitTradeOrderRequest) SetStopLossTimeInForceNil(b bool)`

 SetStopLossTimeInForceNil sets the value for StopLossTimeInForce to be an explicit nil

### UnsetStopLossTimeInForce
`func (o *SubmitTradeOrderRequest) UnsetStopLossTimeInForce()`

UnsetStopLossTimeInForce ensures that no value is present for StopLossTimeInForce, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


