# TradeOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradeOrderId** | **string** | UUID string | 
**TradingAccountId** | **string** | UUID string | 
**Venue** | [**Venue**](Venue.md) |  | 
**PositionId** | Pointer to **string** | UUID string | [optional] 
**InstrumentId** | **string** | Instrument ID in format {VENUE}:{BASE}/{QUOTE} | 
**QuoteId** | Pointer to **string** | UUID string | [optional] 
**BaseAsset** | **string** | Base asset in the trading pair | 
**QuoteAsset** | **string** | Quote asset in the trading pair | 
**OrderSide** | [**OrderSide**](OrderSide.md) |  | 
**OrderType** | [**OrderType**](OrderType.md) |  | 
**TimeInForce** | [**TimeInForce**](TimeInForce.md) |  | 
**Status** | [**OrderStatus**](OrderStatus.md) |  | 
**RejectReason** | Pointer to **string** | Reason for order rejection | [optional] 
**CancelReason** | Pointer to **string** | Reason for order cancellation | [optional] 
**LimitPrice** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**StopPrice** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**Quantity** | **string** | Decimal value as string to preserve precision | 
**OrderQuantityType** | [**OrderQuantityType**](OrderQuantityType.md) |  | 
**QuantityRounding** | Pointer to [**QuantityRounding**](QuantityRounding.md) |  | [optional] [default to QUANTITYROUNDING_EMPTY]
**ExecutedPrice** | **string** | Decimal value as string to preserve precision | 
**ExecutedQuantity** | **string** | Decimal value as string to preserve precision | 
**ExecutedCost** | **string** | Decimal value as string to preserve precision | 
**Fees** | [**[]SecurityQuantity**](SecurityQuantity.md) | Aggregated fees across all executions | 
**Executions** | Pointer to [**[]TradeExecution**](TradeExecution.md) | Detailed breakdown of executions across different venues | [optional] 
**CreatedAt** | **int64** | Unix timestamp in milliseconds | 
**CreatedAtDateTime** | Pointer to **time.Time** | Creation timestamp in ISO 8601 format | [optional] 
**UpdatedAt** | **int64** | Unix timestamp in milliseconds | 
**UpdatedAtDateTime** | Pointer to **time.Time** | Last update timestamp in ISO 8601 format | [optional] 
**ExpireAt** | Pointer to **int64** | Unix timestamp in milliseconds | [optional] 
**ExpireAtDateTime** | Pointer to **time.Time** | Expiration timestamp in ISO 8601 format | [optional] 
**CanceledAt** | Pointer to **int64** | Unix timestamp in milliseconds | [optional] 
**CanceledAtDateTime** | Pointer to **time.Time** | Cancellation timestamp in ISO 8601 format | [optional] 

## Methods

### NewTradeOrder

`func NewTradeOrder(tradeOrderId string, tradingAccountId string, venue Venue, instrumentId string, baseAsset string, quoteAsset string, orderSide OrderSide, orderType OrderType, timeInForce TimeInForce, status OrderStatus, quantity string, orderQuantityType OrderQuantityType, executedPrice string, executedQuantity string, executedCost string, fees []SecurityQuantity, createdAt int64, updatedAt int64, ) *TradeOrder`

NewTradeOrder instantiates a new TradeOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradeOrderWithDefaults

`func NewTradeOrderWithDefaults() *TradeOrder`

NewTradeOrderWithDefaults instantiates a new TradeOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradeOrderId

`func (o *TradeOrder) GetTradeOrderId() string`

GetTradeOrderId returns the TradeOrderId field if non-nil, zero value otherwise.

### GetTradeOrderIdOk

`func (o *TradeOrder) GetTradeOrderIdOk() (*string, bool)`

GetTradeOrderIdOk returns a tuple with the TradeOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeOrderId

`func (o *TradeOrder) SetTradeOrderId(v string)`

SetTradeOrderId sets TradeOrderId field to given value.


### GetTradingAccountId

`func (o *TradeOrder) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *TradeOrder) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *TradeOrder) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.


### GetVenue

`func (o *TradeOrder) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *TradeOrder) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *TradeOrder) SetVenue(v Venue)`

SetVenue sets Venue field to given value.


### GetPositionId

`func (o *TradeOrder) GetPositionId() string`

GetPositionId returns the PositionId field if non-nil, zero value otherwise.

### GetPositionIdOk

`func (o *TradeOrder) GetPositionIdOk() (*string, bool)`

GetPositionIdOk returns a tuple with the PositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionId

`func (o *TradeOrder) SetPositionId(v string)`

SetPositionId sets PositionId field to given value.

### HasPositionId

`func (o *TradeOrder) HasPositionId() bool`

HasPositionId returns a boolean if a field has been set.

### GetInstrumentId

`func (o *TradeOrder) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *TradeOrder) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *TradeOrder) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.


### GetQuoteId

`func (o *TradeOrder) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *TradeOrder) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *TradeOrder) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.

### HasQuoteId

`func (o *TradeOrder) HasQuoteId() bool`

HasQuoteId returns a boolean if a field has been set.

### GetBaseAsset

`func (o *TradeOrder) GetBaseAsset() string`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *TradeOrder) GetBaseAssetOk() (*string, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *TradeOrder) SetBaseAsset(v string)`

SetBaseAsset sets BaseAsset field to given value.


### GetQuoteAsset

`func (o *TradeOrder) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *TradeOrder) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *TradeOrder) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.


### GetOrderSide

`func (o *TradeOrder) GetOrderSide() OrderSide`

GetOrderSide returns the OrderSide field if non-nil, zero value otherwise.

### GetOrderSideOk

`func (o *TradeOrder) GetOrderSideOk() (*OrderSide, bool)`

GetOrderSideOk returns a tuple with the OrderSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSide

`func (o *TradeOrder) SetOrderSide(v OrderSide)`

SetOrderSide sets OrderSide field to given value.


### GetOrderType

`func (o *TradeOrder) GetOrderType() OrderType`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *TradeOrder) GetOrderTypeOk() (*OrderType, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *TradeOrder) SetOrderType(v OrderType)`

SetOrderType sets OrderType field to given value.


### GetTimeInForce

`func (o *TradeOrder) GetTimeInForce() TimeInForce`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *TradeOrder) GetTimeInForceOk() (*TimeInForce, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *TradeOrder) SetTimeInForce(v TimeInForce)`

SetTimeInForce sets TimeInForce field to given value.


### GetStatus

`func (o *TradeOrder) GetStatus() OrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TradeOrder) GetStatusOk() (*OrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TradeOrder) SetStatus(v OrderStatus)`

SetStatus sets Status field to given value.


### GetRejectReason

`func (o *TradeOrder) GetRejectReason() string`

GetRejectReason returns the RejectReason field if non-nil, zero value otherwise.

### GetRejectReasonOk

`func (o *TradeOrder) GetRejectReasonOk() (*string, bool)`

GetRejectReasonOk returns a tuple with the RejectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectReason

`func (o *TradeOrder) SetRejectReason(v string)`

SetRejectReason sets RejectReason field to given value.

### HasRejectReason

`func (o *TradeOrder) HasRejectReason() bool`

HasRejectReason returns a boolean if a field has been set.

### GetCancelReason

`func (o *TradeOrder) GetCancelReason() string`

GetCancelReason returns the CancelReason field if non-nil, zero value otherwise.

### GetCancelReasonOk

`func (o *TradeOrder) GetCancelReasonOk() (*string, bool)`

GetCancelReasonOk returns a tuple with the CancelReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelReason

`func (o *TradeOrder) SetCancelReason(v string)`

SetCancelReason sets CancelReason field to given value.

### HasCancelReason

`func (o *TradeOrder) HasCancelReason() bool`

HasCancelReason returns a boolean if a field has been set.

### GetLimitPrice

`func (o *TradeOrder) GetLimitPrice() string`

GetLimitPrice returns the LimitPrice field if non-nil, zero value otherwise.

### GetLimitPriceOk

`func (o *TradeOrder) GetLimitPriceOk() (*string, bool)`

GetLimitPriceOk returns a tuple with the LimitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPrice

`func (o *TradeOrder) SetLimitPrice(v string)`

SetLimitPrice sets LimitPrice field to given value.

### HasLimitPrice

`func (o *TradeOrder) HasLimitPrice() bool`

HasLimitPrice returns a boolean if a field has been set.

### GetStopPrice

`func (o *TradeOrder) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *TradeOrder) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *TradeOrder) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *TradeOrder) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *TradeOrder) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *TradeOrder) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *TradeOrder) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetOrderQuantityType

`func (o *TradeOrder) GetOrderQuantityType() OrderQuantityType`

GetOrderQuantityType returns the OrderQuantityType field if non-nil, zero value otherwise.

### GetOrderQuantityTypeOk

`func (o *TradeOrder) GetOrderQuantityTypeOk() (*OrderQuantityType, bool)`

GetOrderQuantityTypeOk returns a tuple with the OrderQuantityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderQuantityType

`func (o *TradeOrder) SetOrderQuantityType(v OrderQuantityType)`

SetOrderQuantityType sets OrderQuantityType field to given value.


### GetQuantityRounding

`func (o *TradeOrder) GetQuantityRounding() QuantityRounding`

GetQuantityRounding returns the QuantityRounding field if non-nil, zero value otherwise.

### GetQuantityRoundingOk

`func (o *TradeOrder) GetQuantityRoundingOk() (*QuantityRounding, bool)`

GetQuantityRoundingOk returns a tuple with the QuantityRounding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityRounding

`func (o *TradeOrder) SetQuantityRounding(v QuantityRounding)`

SetQuantityRounding sets QuantityRounding field to given value.

### HasQuantityRounding

`func (o *TradeOrder) HasQuantityRounding() bool`

HasQuantityRounding returns a boolean if a field has been set.

### GetExecutedPrice

`func (o *TradeOrder) GetExecutedPrice() string`

GetExecutedPrice returns the ExecutedPrice field if non-nil, zero value otherwise.

### GetExecutedPriceOk

`func (o *TradeOrder) GetExecutedPriceOk() (*string, bool)`

GetExecutedPriceOk returns a tuple with the ExecutedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedPrice

`func (o *TradeOrder) SetExecutedPrice(v string)`

SetExecutedPrice sets ExecutedPrice field to given value.


### GetExecutedQuantity

`func (o *TradeOrder) GetExecutedQuantity() string`

GetExecutedQuantity returns the ExecutedQuantity field if non-nil, zero value otherwise.

### GetExecutedQuantityOk

`func (o *TradeOrder) GetExecutedQuantityOk() (*string, bool)`

GetExecutedQuantityOk returns a tuple with the ExecutedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQuantity

`func (o *TradeOrder) SetExecutedQuantity(v string)`

SetExecutedQuantity sets ExecutedQuantity field to given value.


### GetExecutedCost

`func (o *TradeOrder) GetExecutedCost() string`

GetExecutedCost returns the ExecutedCost field if non-nil, zero value otherwise.

### GetExecutedCostOk

`func (o *TradeOrder) GetExecutedCostOk() (*string, bool)`

GetExecutedCostOk returns a tuple with the ExecutedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedCost

`func (o *TradeOrder) SetExecutedCost(v string)`

SetExecutedCost sets ExecutedCost field to given value.


### GetFees

`func (o *TradeOrder) GetFees() []SecurityQuantity`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *TradeOrder) GetFeesOk() (*[]SecurityQuantity, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *TradeOrder) SetFees(v []SecurityQuantity)`

SetFees sets Fees field to given value.


### GetExecutions

`func (o *TradeOrder) GetExecutions() []TradeExecution`

GetExecutions returns the Executions field if non-nil, zero value otherwise.

### GetExecutionsOk

`func (o *TradeOrder) GetExecutionsOk() (*[]TradeExecution, bool)`

GetExecutionsOk returns a tuple with the Executions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutions

`func (o *TradeOrder) SetExecutions(v []TradeExecution)`

SetExecutions sets Executions field to given value.

### HasExecutions

`func (o *TradeOrder) HasExecutions() bool`

HasExecutions returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TradeOrder) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TradeOrder) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TradeOrder) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedAtDateTime

`func (o *TradeOrder) GetCreatedAtDateTime() time.Time`

GetCreatedAtDateTime returns the CreatedAtDateTime field if non-nil, zero value otherwise.

### GetCreatedAtDateTimeOk

`func (o *TradeOrder) GetCreatedAtDateTimeOk() (*time.Time, bool)`

GetCreatedAtDateTimeOk returns a tuple with the CreatedAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtDateTime

`func (o *TradeOrder) SetCreatedAtDateTime(v time.Time)`

SetCreatedAtDateTime sets CreatedAtDateTime field to given value.

### HasCreatedAtDateTime

`func (o *TradeOrder) HasCreatedAtDateTime() bool`

HasCreatedAtDateTime returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TradeOrder) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TradeOrder) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TradeOrder) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpdatedAtDateTime

`func (o *TradeOrder) GetUpdatedAtDateTime() time.Time`

GetUpdatedAtDateTime returns the UpdatedAtDateTime field if non-nil, zero value otherwise.

### GetUpdatedAtDateTimeOk

`func (o *TradeOrder) GetUpdatedAtDateTimeOk() (*time.Time, bool)`

GetUpdatedAtDateTimeOk returns a tuple with the UpdatedAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtDateTime

`func (o *TradeOrder) SetUpdatedAtDateTime(v time.Time)`

SetUpdatedAtDateTime sets UpdatedAtDateTime field to given value.

### HasUpdatedAtDateTime

`func (o *TradeOrder) HasUpdatedAtDateTime() bool`

HasUpdatedAtDateTime returns a boolean if a field has been set.

### GetExpireAt

`func (o *TradeOrder) GetExpireAt() int64`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *TradeOrder) GetExpireAtOk() (*int64, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *TradeOrder) SetExpireAt(v int64)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *TradeOrder) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### GetExpireAtDateTime

`func (o *TradeOrder) GetExpireAtDateTime() time.Time`

GetExpireAtDateTime returns the ExpireAtDateTime field if non-nil, zero value otherwise.

### GetExpireAtDateTimeOk

`func (o *TradeOrder) GetExpireAtDateTimeOk() (*time.Time, bool)`

GetExpireAtDateTimeOk returns a tuple with the ExpireAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAtDateTime

`func (o *TradeOrder) SetExpireAtDateTime(v time.Time)`

SetExpireAtDateTime sets ExpireAtDateTime field to given value.

### HasExpireAtDateTime

`func (o *TradeOrder) HasExpireAtDateTime() bool`

HasExpireAtDateTime returns a boolean if a field has been set.

### GetCanceledAt

`func (o *TradeOrder) GetCanceledAt() int64`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *TradeOrder) GetCanceledAtOk() (*int64, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *TradeOrder) SetCanceledAt(v int64)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *TradeOrder) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.

### GetCanceledAtDateTime

`func (o *TradeOrder) GetCanceledAtDateTime() time.Time`

GetCanceledAtDateTime returns the CanceledAtDateTime field if non-nil, zero value otherwise.

### GetCanceledAtDateTimeOk

`func (o *TradeOrder) GetCanceledAtDateTimeOk() (*time.Time, bool)`

GetCanceledAtDateTimeOk returns a tuple with the CanceledAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAtDateTime

`func (o *TradeOrder) SetCanceledAtDateTime(v time.Time)`

SetCanceledAtDateTime sets CanceledAtDateTime field to given value.

### HasCanceledAtDateTime

`func (o *TradeOrder) HasCanceledAtDateTime() bool`

HasCanceledAtDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


