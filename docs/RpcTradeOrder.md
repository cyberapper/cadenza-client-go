# RpcTradeOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradeOrderId** | Pointer to **string** | Unique trade order ID | [optional] 
**ClientOrderId** | Pointer to **string** | Client-assigned order ID | [optional] 
**OriginalClientOrderId** | Pointer to **string** | Original client order ID (for cancel/replace) | [optional] 
**ExternalOrderId** | Pointer to **string** | Order ID from external venue | [optional] 
**IdempotencyKey** | Pointer to **string** | Idempotency key | [optional] 
**Venue** | Pointer to **string** | Trading venue | [optional] 
**TradingAccountId** | Pointer to **string** | Trading account ID | [optional] 
**ExternalTradingAccountId** | Pointer to **string** | External account ID at venue | [optional] 
**InstrumentId** | Pointer to **string** | Instrument ID (VENUE:BASE/QUOTE) | [optional] 
**ExternalSymbol** | Pointer to **string** | Symbol at external venue | [optional] 
**BaseAsset** | Pointer to **string** | Base asset | [optional] 
**QuoteAsset** | Pointer to **string** | Quote asset | [optional] 
**OrderType** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**OrderSide** | Pointer to [**OrderSide**](OrderSide.md) |  | [optional] 
**QuantityType** | Pointer to [**QuantityType**](QuantityType.md) |  | [optional] 
**Quantity** | Pointer to **string** | Order quantity (decimal string) | [optional] 
**QuoteCurrencyQuantity** | Pointer to **string** | Quantity in quote currency | [optional] 
**PositionPercentage** | Pointer to **string** | Quantity as position percentage | [optional] 
**QuantityRounding** | Pointer to [**QuantityRounding**](QuantityRounding.md) |  | [optional] [default to QUANTITYROUNDING_EMPTY]
**LimitPrice** | Pointer to **string** | Limit price (decimal string) | [optional] 
**StopPrice** | Pointer to **string** | Stop price (decimal string) | [optional] 
**TimeInForce** | Pointer to [**TimeInForce**](TimeInForce.md) |  | [optional] 
**Status** | Pointer to [**OrderStatus**](OrderStatus.md) |  | [optional] 
**ExecutedPrice** | Pointer to **string** | Average executed price | [optional] 
**ExecutedPercentage** | Pointer to **string** | Percentage of order filled | [optional] 
**ExecutedQuantity** | Pointer to **string** | Quantity executed | [optional] 
**ExecutedCost** | Pointer to **string** | Total cost in quote currency | [optional] 
**Fees** | Pointer to [**[]RpcSecurityQuantity**](RpcSecurityQuantity.md) | Fees charged | [optional] 
**CancelReason** | Pointer to **string** | Reason for cancellation | [optional] 
**RejectReason** | Pointer to **string** | Reason for rejection | [optional] 
**CreatedAt** | Pointer to **time.Time** | Order creation time | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Last update time | [optional] 
**ExpireAt** | Pointer to **time.Time** | Order expiration time (for GTD) | [optional] 
**LastExecutionAt** | Pointer to **time.Time** | Time of last execution | [optional] 
**CanceledAt** | Pointer to **time.Time** | Time of cancellation | [optional] 

## Methods

### NewRpcTradeOrder

`func NewRpcTradeOrder() *RpcTradeOrder`

NewRpcTradeOrder instantiates a new RpcTradeOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcTradeOrderWithDefaults

`func NewRpcTradeOrderWithDefaults() *RpcTradeOrder`

NewRpcTradeOrderWithDefaults instantiates a new RpcTradeOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradeOrderId

`func (o *RpcTradeOrder) GetTradeOrderId() string`

GetTradeOrderId returns the TradeOrderId field if non-nil, zero value otherwise.

### GetTradeOrderIdOk

`func (o *RpcTradeOrder) GetTradeOrderIdOk() (*string, bool)`

GetTradeOrderIdOk returns a tuple with the TradeOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeOrderId

`func (o *RpcTradeOrder) SetTradeOrderId(v string)`

SetTradeOrderId sets TradeOrderId field to given value.

### HasTradeOrderId

`func (o *RpcTradeOrder) HasTradeOrderId() bool`

HasTradeOrderId returns a boolean if a field has been set.

### GetClientOrderId

`func (o *RpcTradeOrder) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *RpcTradeOrder) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *RpcTradeOrder) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *RpcTradeOrder) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetOriginalClientOrderId

`func (o *RpcTradeOrder) GetOriginalClientOrderId() string`

GetOriginalClientOrderId returns the OriginalClientOrderId field if non-nil, zero value otherwise.

### GetOriginalClientOrderIdOk

`func (o *RpcTradeOrder) GetOriginalClientOrderIdOk() (*string, bool)`

GetOriginalClientOrderIdOk returns a tuple with the OriginalClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalClientOrderId

`func (o *RpcTradeOrder) SetOriginalClientOrderId(v string)`

SetOriginalClientOrderId sets OriginalClientOrderId field to given value.

### HasOriginalClientOrderId

`func (o *RpcTradeOrder) HasOriginalClientOrderId() bool`

HasOriginalClientOrderId returns a boolean if a field has been set.

### GetExternalOrderId

`func (o *RpcTradeOrder) GetExternalOrderId() string`

GetExternalOrderId returns the ExternalOrderId field if non-nil, zero value otherwise.

### GetExternalOrderIdOk

`func (o *RpcTradeOrder) GetExternalOrderIdOk() (*string, bool)`

GetExternalOrderIdOk returns a tuple with the ExternalOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOrderId

`func (o *RpcTradeOrder) SetExternalOrderId(v string)`

SetExternalOrderId sets ExternalOrderId field to given value.

### HasExternalOrderId

`func (o *RpcTradeOrder) HasExternalOrderId() bool`

HasExternalOrderId returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *RpcTradeOrder) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *RpcTradeOrder) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *RpcTradeOrder) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *RpcTradeOrder) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetVenue

`func (o *RpcTradeOrder) GetVenue() string`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *RpcTradeOrder) GetVenueOk() (*string, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *RpcTradeOrder) SetVenue(v string)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *RpcTradeOrder) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetTradingAccountId

`func (o *RpcTradeOrder) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *RpcTradeOrder) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *RpcTradeOrder) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.

### HasTradingAccountId

`func (o *RpcTradeOrder) HasTradingAccountId() bool`

HasTradingAccountId returns a boolean if a field has been set.

### GetExternalTradingAccountId

`func (o *RpcTradeOrder) GetExternalTradingAccountId() string`

GetExternalTradingAccountId returns the ExternalTradingAccountId field if non-nil, zero value otherwise.

### GetExternalTradingAccountIdOk

`func (o *RpcTradeOrder) GetExternalTradingAccountIdOk() (*string, bool)`

GetExternalTradingAccountIdOk returns a tuple with the ExternalTradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTradingAccountId

`func (o *RpcTradeOrder) SetExternalTradingAccountId(v string)`

SetExternalTradingAccountId sets ExternalTradingAccountId field to given value.

### HasExternalTradingAccountId

`func (o *RpcTradeOrder) HasExternalTradingAccountId() bool`

HasExternalTradingAccountId returns a boolean if a field has been set.

### GetInstrumentId

`func (o *RpcTradeOrder) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *RpcTradeOrder) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *RpcTradeOrder) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *RpcTradeOrder) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetExternalSymbol

`func (o *RpcTradeOrder) GetExternalSymbol() string`

GetExternalSymbol returns the ExternalSymbol field if non-nil, zero value otherwise.

### GetExternalSymbolOk

`func (o *RpcTradeOrder) GetExternalSymbolOk() (*string, bool)`

GetExternalSymbolOk returns a tuple with the ExternalSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSymbol

`func (o *RpcTradeOrder) SetExternalSymbol(v string)`

SetExternalSymbol sets ExternalSymbol field to given value.

### HasExternalSymbol

`func (o *RpcTradeOrder) HasExternalSymbol() bool`

HasExternalSymbol returns a boolean if a field has been set.

### GetBaseAsset

`func (o *RpcTradeOrder) GetBaseAsset() string`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *RpcTradeOrder) GetBaseAssetOk() (*string, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *RpcTradeOrder) SetBaseAsset(v string)`

SetBaseAsset sets BaseAsset field to given value.

### HasBaseAsset

`func (o *RpcTradeOrder) HasBaseAsset() bool`

HasBaseAsset returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *RpcTradeOrder) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *RpcTradeOrder) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *RpcTradeOrder) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *RpcTradeOrder) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetOrderType

`func (o *RpcTradeOrder) GetOrderType() OrderType`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *RpcTradeOrder) GetOrderTypeOk() (*OrderType, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *RpcTradeOrder) SetOrderType(v OrderType)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *RpcTradeOrder) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetOrderSide

`func (o *RpcTradeOrder) GetOrderSide() OrderSide`

GetOrderSide returns the OrderSide field if non-nil, zero value otherwise.

### GetOrderSideOk

`func (o *RpcTradeOrder) GetOrderSideOk() (*OrderSide, bool)`

GetOrderSideOk returns a tuple with the OrderSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSide

`func (o *RpcTradeOrder) SetOrderSide(v OrderSide)`

SetOrderSide sets OrderSide field to given value.

### HasOrderSide

`func (o *RpcTradeOrder) HasOrderSide() bool`

HasOrderSide returns a boolean if a field has been set.

### GetQuantityType

`func (o *RpcTradeOrder) GetQuantityType() QuantityType`

GetQuantityType returns the QuantityType field if non-nil, zero value otherwise.

### GetQuantityTypeOk

`func (o *RpcTradeOrder) GetQuantityTypeOk() (*QuantityType, bool)`

GetQuantityTypeOk returns a tuple with the QuantityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityType

`func (o *RpcTradeOrder) SetQuantityType(v QuantityType)`

SetQuantityType sets QuantityType field to given value.

### HasQuantityType

`func (o *RpcTradeOrder) HasQuantityType() bool`

HasQuantityType returns a boolean if a field has been set.

### GetQuantity

`func (o *RpcTradeOrder) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *RpcTradeOrder) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *RpcTradeOrder) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *RpcTradeOrder) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetQuoteCurrencyQuantity

`func (o *RpcTradeOrder) GetQuoteCurrencyQuantity() string`

GetQuoteCurrencyQuantity returns the QuoteCurrencyQuantity field if non-nil, zero value otherwise.

### GetQuoteCurrencyQuantityOk

`func (o *RpcTradeOrder) GetQuoteCurrencyQuantityOk() (*string, bool)`

GetQuoteCurrencyQuantityOk returns a tuple with the QuoteCurrencyQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCurrencyQuantity

`func (o *RpcTradeOrder) SetQuoteCurrencyQuantity(v string)`

SetQuoteCurrencyQuantity sets QuoteCurrencyQuantity field to given value.

### HasQuoteCurrencyQuantity

`func (o *RpcTradeOrder) HasQuoteCurrencyQuantity() bool`

HasQuoteCurrencyQuantity returns a boolean if a field has been set.

### GetPositionPercentage

`func (o *RpcTradeOrder) GetPositionPercentage() string`

GetPositionPercentage returns the PositionPercentage field if non-nil, zero value otherwise.

### GetPositionPercentageOk

`func (o *RpcTradeOrder) GetPositionPercentageOk() (*string, bool)`

GetPositionPercentageOk returns a tuple with the PositionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionPercentage

`func (o *RpcTradeOrder) SetPositionPercentage(v string)`

SetPositionPercentage sets PositionPercentage field to given value.

### HasPositionPercentage

`func (o *RpcTradeOrder) HasPositionPercentage() bool`

HasPositionPercentage returns a boolean if a field has been set.

### GetQuantityRounding

`func (o *RpcTradeOrder) GetQuantityRounding() QuantityRounding`

GetQuantityRounding returns the QuantityRounding field if non-nil, zero value otherwise.

### GetQuantityRoundingOk

`func (o *RpcTradeOrder) GetQuantityRoundingOk() (*QuantityRounding, bool)`

GetQuantityRoundingOk returns a tuple with the QuantityRounding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityRounding

`func (o *RpcTradeOrder) SetQuantityRounding(v QuantityRounding)`

SetQuantityRounding sets QuantityRounding field to given value.

### HasQuantityRounding

`func (o *RpcTradeOrder) HasQuantityRounding() bool`

HasQuantityRounding returns a boolean if a field has been set.

### GetLimitPrice

`func (o *RpcTradeOrder) GetLimitPrice() string`

GetLimitPrice returns the LimitPrice field if non-nil, zero value otherwise.

### GetLimitPriceOk

`func (o *RpcTradeOrder) GetLimitPriceOk() (*string, bool)`

GetLimitPriceOk returns a tuple with the LimitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPrice

`func (o *RpcTradeOrder) SetLimitPrice(v string)`

SetLimitPrice sets LimitPrice field to given value.

### HasLimitPrice

`func (o *RpcTradeOrder) HasLimitPrice() bool`

HasLimitPrice returns a boolean if a field has been set.

### GetStopPrice

`func (o *RpcTradeOrder) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *RpcTradeOrder) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *RpcTradeOrder) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *RpcTradeOrder) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetTimeInForce

`func (o *RpcTradeOrder) GetTimeInForce() TimeInForce`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *RpcTradeOrder) GetTimeInForceOk() (*TimeInForce, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *RpcTradeOrder) SetTimeInForce(v TimeInForce)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *RpcTradeOrder) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetStatus

`func (o *RpcTradeOrder) GetStatus() OrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RpcTradeOrder) GetStatusOk() (*OrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RpcTradeOrder) SetStatus(v OrderStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RpcTradeOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExecutedPrice

`func (o *RpcTradeOrder) GetExecutedPrice() string`

GetExecutedPrice returns the ExecutedPrice field if non-nil, zero value otherwise.

### GetExecutedPriceOk

`func (o *RpcTradeOrder) GetExecutedPriceOk() (*string, bool)`

GetExecutedPriceOk returns a tuple with the ExecutedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedPrice

`func (o *RpcTradeOrder) SetExecutedPrice(v string)`

SetExecutedPrice sets ExecutedPrice field to given value.

### HasExecutedPrice

`func (o *RpcTradeOrder) HasExecutedPrice() bool`

HasExecutedPrice returns a boolean if a field has been set.

### GetExecutedPercentage

`func (o *RpcTradeOrder) GetExecutedPercentage() string`

GetExecutedPercentage returns the ExecutedPercentage field if non-nil, zero value otherwise.

### GetExecutedPercentageOk

`func (o *RpcTradeOrder) GetExecutedPercentageOk() (*string, bool)`

GetExecutedPercentageOk returns a tuple with the ExecutedPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedPercentage

`func (o *RpcTradeOrder) SetExecutedPercentage(v string)`

SetExecutedPercentage sets ExecutedPercentage field to given value.

### HasExecutedPercentage

`func (o *RpcTradeOrder) HasExecutedPercentage() bool`

HasExecutedPercentage returns a boolean if a field has been set.

### GetExecutedQuantity

`func (o *RpcTradeOrder) GetExecutedQuantity() string`

GetExecutedQuantity returns the ExecutedQuantity field if non-nil, zero value otherwise.

### GetExecutedQuantityOk

`func (o *RpcTradeOrder) GetExecutedQuantityOk() (*string, bool)`

GetExecutedQuantityOk returns a tuple with the ExecutedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQuantity

`func (o *RpcTradeOrder) SetExecutedQuantity(v string)`

SetExecutedQuantity sets ExecutedQuantity field to given value.

### HasExecutedQuantity

`func (o *RpcTradeOrder) HasExecutedQuantity() bool`

HasExecutedQuantity returns a boolean if a field has been set.

### GetExecutedCost

`func (o *RpcTradeOrder) GetExecutedCost() string`

GetExecutedCost returns the ExecutedCost field if non-nil, zero value otherwise.

### GetExecutedCostOk

`func (o *RpcTradeOrder) GetExecutedCostOk() (*string, bool)`

GetExecutedCostOk returns a tuple with the ExecutedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedCost

`func (o *RpcTradeOrder) SetExecutedCost(v string)`

SetExecutedCost sets ExecutedCost field to given value.

### HasExecutedCost

`func (o *RpcTradeOrder) HasExecutedCost() bool`

HasExecutedCost returns a boolean if a field has been set.

### GetFees

`func (o *RpcTradeOrder) GetFees() []RpcSecurityQuantity`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *RpcTradeOrder) GetFeesOk() (*[]RpcSecurityQuantity, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *RpcTradeOrder) SetFees(v []RpcSecurityQuantity)`

SetFees sets Fees field to given value.

### HasFees

`func (o *RpcTradeOrder) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetCancelReason

`func (o *RpcTradeOrder) GetCancelReason() string`

GetCancelReason returns the CancelReason field if non-nil, zero value otherwise.

### GetCancelReasonOk

`func (o *RpcTradeOrder) GetCancelReasonOk() (*string, bool)`

GetCancelReasonOk returns a tuple with the CancelReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelReason

`func (o *RpcTradeOrder) SetCancelReason(v string)`

SetCancelReason sets CancelReason field to given value.

### HasCancelReason

`func (o *RpcTradeOrder) HasCancelReason() bool`

HasCancelReason returns a boolean if a field has been set.

### GetRejectReason

`func (o *RpcTradeOrder) GetRejectReason() string`

GetRejectReason returns the RejectReason field if non-nil, zero value otherwise.

### GetRejectReasonOk

`func (o *RpcTradeOrder) GetRejectReasonOk() (*string, bool)`

GetRejectReasonOk returns a tuple with the RejectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectReason

`func (o *RpcTradeOrder) SetRejectReason(v string)`

SetRejectReason sets RejectReason field to given value.

### HasRejectReason

`func (o *RpcTradeOrder) HasRejectReason() bool`

HasRejectReason returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RpcTradeOrder) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RpcTradeOrder) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RpcTradeOrder) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RpcTradeOrder) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RpcTradeOrder) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RpcTradeOrder) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RpcTradeOrder) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RpcTradeOrder) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetExpireAt

`func (o *RpcTradeOrder) GetExpireAt() time.Time`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *RpcTradeOrder) GetExpireAtOk() (*time.Time, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *RpcTradeOrder) SetExpireAt(v time.Time)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *RpcTradeOrder) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### GetLastExecutionAt

`func (o *RpcTradeOrder) GetLastExecutionAt() time.Time`

GetLastExecutionAt returns the LastExecutionAt field if non-nil, zero value otherwise.

### GetLastExecutionAtOk

`func (o *RpcTradeOrder) GetLastExecutionAtOk() (*time.Time, bool)`

GetLastExecutionAtOk returns a tuple with the LastExecutionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionAt

`func (o *RpcTradeOrder) SetLastExecutionAt(v time.Time)`

SetLastExecutionAt sets LastExecutionAt field to given value.

### HasLastExecutionAt

`func (o *RpcTradeOrder) HasLastExecutionAt() bool`

HasLastExecutionAt returns a boolean if a field has been set.

### GetCanceledAt

`func (o *RpcTradeOrder) GetCanceledAt() time.Time`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *RpcTradeOrder) GetCanceledAtOk() (*time.Time, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *RpcTradeOrder) SetCanceledAt(v time.Time)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *RpcTradeOrder) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


