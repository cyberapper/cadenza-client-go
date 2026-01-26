# RpcSubmitTradeOrderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradeOrder** | [**RpcTradeOrder**](RpcTradeOrder.md) |  | 
**TradingAccountId** | **string** | Trading account ID to place order on | 
**IdempotencyKey** | Pointer to **string** | Idempotency key to prevent duplicate orders | [optional] 
**AwaitClosed** | Pointer to **bool** | Wait for order to reach terminal state before responding | [optional] [default to false]

## Methods

### NewRpcSubmitTradeOrderParams

`func NewRpcSubmitTradeOrderParams(tradeOrder RpcTradeOrder, tradingAccountId string, ) *RpcSubmitTradeOrderParams`

NewRpcSubmitTradeOrderParams instantiates a new RpcSubmitTradeOrderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcSubmitTradeOrderParamsWithDefaults

`func NewRpcSubmitTradeOrderParamsWithDefaults() *RpcSubmitTradeOrderParams`

NewRpcSubmitTradeOrderParamsWithDefaults instantiates a new RpcSubmitTradeOrderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradeOrder

`func (o *RpcSubmitTradeOrderParams) GetTradeOrder() RpcTradeOrder`

GetTradeOrder returns the TradeOrder field if non-nil, zero value otherwise.

### GetTradeOrderOk

`func (o *RpcSubmitTradeOrderParams) GetTradeOrderOk() (*RpcTradeOrder, bool)`

GetTradeOrderOk returns a tuple with the TradeOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeOrder

`func (o *RpcSubmitTradeOrderParams) SetTradeOrder(v RpcTradeOrder)`

SetTradeOrder sets TradeOrder field to given value.


### GetTradingAccountId

`func (o *RpcSubmitTradeOrderParams) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *RpcSubmitTradeOrderParams) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *RpcSubmitTradeOrderParams) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.


### GetIdempotencyKey

`func (o *RpcSubmitTradeOrderParams) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *RpcSubmitTradeOrderParams) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *RpcSubmitTradeOrderParams) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *RpcSubmitTradeOrderParams) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetAwaitClosed

`func (o *RpcSubmitTradeOrderParams) GetAwaitClosed() bool`

GetAwaitClosed returns the AwaitClosed field if non-nil, zero value otherwise.

### GetAwaitClosedOk

`func (o *RpcSubmitTradeOrderParams) GetAwaitClosedOk() (*bool, bool)`

GetAwaitClosedOk returns a tuple with the AwaitClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaitClosed

`func (o *RpcSubmitTradeOrderParams) SetAwaitClosed(v bool)`

SetAwaitClosed sets AwaitClosed field to given value.

### HasAwaitClosed

`func (o *RpcSubmitTradeOrderParams) HasAwaitClosed() bool`

HasAwaitClosed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


