# TradeExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionId** | **string** | UUID string | 
**ExternalTradeId** | Pointer to **string** | Exchange&#39;s execution identifier | [optional] 
**Venue** | [**Venue**](Venue.md) |  | 
**InstrumentId** | **string** | Instrument ID in format {VENUE}:{BASE}/{QUOTE} | 
**OrderSide** | Pointer to [**OrderSide**](OrderSide.md) |  | [optional] 
**ExecutedQuantity** | **string** | Decimal value as string to preserve precision | 
**ExecutedPrice** | **string** | Decimal value as string to preserve precision | 
**ExecutedCost** | **string** | Decimal value as string to preserve precision | 
**Fees** | Pointer to [**[]SecurityQuantity**](SecurityQuantity.md) | Fees charged for this execution | [optional] 
**ExecutedAt** | **int64** | Unix timestamp in milliseconds | 
**ExecutedAtDateTime** | Pointer to **time.Time** | Execution timestamp in ISO 8601 format | [optional] 

## Methods

### NewTradeExecution

`func NewTradeExecution(executionId string, venue Venue, instrumentId string, executedQuantity string, executedPrice string, executedCost string, executedAt int64, ) *TradeExecution`

NewTradeExecution instantiates a new TradeExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradeExecutionWithDefaults

`func NewTradeExecutionWithDefaults() *TradeExecution`

NewTradeExecutionWithDefaults instantiates a new TradeExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionId

`func (o *TradeExecution) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *TradeExecution) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *TradeExecution) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetExternalTradeId

`func (o *TradeExecution) GetExternalTradeId() string`

GetExternalTradeId returns the ExternalTradeId field if non-nil, zero value otherwise.

### GetExternalTradeIdOk

`func (o *TradeExecution) GetExternalTradeIdOk() (*string, bool)`

GetExternalTradeIdOk returns a tuple with the ExternalTradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTradeId

`func (o *TradeExecution) SetExternalTradeId(v string)`

SetExternalTradeId sets ExternalTradeId field to given value.

### HasExternalTradeId

`func (o *TradeExecution) HasExternalTradeId() bool`

HasExternalTradeId returns a boolean if a field has been set.

### GetVenue

`func (o *TradeExecution) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *TradeExecution) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *TradeExecution) SetVenue(v Venue)`

SetVenue sets Venue field to given value.


### GetInstrumentId

`func (o *TradeExecution) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *TradeExecution) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *TradeExecution) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.


### GetOrderSide

`func (o *TradeExecution) GetOrderSide() OrderSide`

GetOrderSide returns the OrderSide field if non-nil, zero value otherwise.

### GetOrderSideOk

`func (o *TradeExecution) GetOrderSideOk() (*OrderSide, bool)`

GetOrderSideOk returns a tuple with the OrderSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSide

`func (o *TradeExecution) SetOrderSide(v OrderSide)`

SetOrderSide sets OrderSide field to given value.

### HasOrderSide

`func (o *TradeExecution) HasOrderSide() bool`

HasOrderSide returns a boolean if a field has been set.

### GetExecutedQuantity

`func (o *TradeExecution) GetExecutedQuantity() string`

GetExecutedQuantity returns the ExecutedQuantity field if non-nil, zero value otherwise.

### GetExecutedQuantityOk

`func (o *TradeExecution) GetExecutedQuantityOk() (*string, bool)`

GetExecutedQuantityOk returns a tuple with the ExecutedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQuantity

`func (o *TradeExecution) SetExecutedQuantity(v string)`

SetExecutedQuantity sets ExecutedQuantity field to given value.


### GetExecutedPrice

`func (o *TradeExecution) GetExecutedPrice() string`

GetExecutedPrice returns the ExecutedPrice field if non-nil, zero value otherwise.

### GetExecutedPriceOk

`func (o *TradeExecution) GetExecutedPriceOk() (*string, bool)`

GetExecutedPriceOk returns a tuple with the ExecutedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedPrice

`func (o *TradeExecution) SetExecutedPrice(v string)`

SetExecutedPrice sets ExecutedPrice field to given value.


### GetExecutedCost

`func (o *TradeExecution) GetExecutedCost() string`

GetExecutedCost returns the ExecutedCost field if non-nil, zero value otherwise.

### GetExecutedCostOk

`func (o *TradeExecution) GetExecutedCostOk() (*string, bool)`

GetExecutedCostOk returns a tuple with the ExecutedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedCost

`func (o *TradeExecution) SetExecutedCost(v string)`

SetExecutedCost sets ExecutedCost field to given value.


### GetFees

`func (o *TradeExecution) GetFees() []SecurityQuantity`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *TradeExecution) GetFeesOk() (*[]SecurityQuantity, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *TradeExecution) SetFees(v []SecurityQuantity)`

SetFees sets Fees field to given value.

### HasFees

`func (o *TradeExecution) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetExecutedAt

`func (o *TradeExecution) GetExecutedAt() int64`

GetExecutedAt returns the ExecutedAt field if non-nil, zero value otherwise.

### GetExecutedAtOk

`func (o *TradeExecution) GetExecutedAtOk() (*int64, bool)`

GetExecutedAtOk returns a tuple with the ExecutedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAt

`func (o *TradeExecution) SetExecutedAt(v int64)`

SetExecutedAt sets ExecutedAt field to given value.


### GetExecutedAtDateTime

`func (o *TradeExecution) GetExecutedAtDateTime() time.Time`

GetExecutedAtDateTime returns the ExecutedAtDateTime field if non-nil, zero value otherwise.

### GetExecutedAtDateTimeOk

`func (o *TradeExecution) GetExecutedAtDateTimeOk() (*time.Time, bool)`

GetExecutedAtDateTimeOk returns a tuple with the ExecutedAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAtDateTime

`func (o *TradeExecution) SetExecutedAtDateTime(v time.Time)`

SetExecutedAtDateTime sets ExecutedAtDateTime field to given value.

### HasExecutedAtDateTime

`func (o *TradeExecution) HasExecutedAtDateTime() bool`

HasExecutedAtDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


