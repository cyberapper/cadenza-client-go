# RpcCancelTradeOrderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradeOrderId** | Pointer to **string** | Trade order ID to cancel | [optional] 
**TradingAccountId** | **string** | Trading account ID | 
**ClientOrderId** | Pointer to **string** | Client order ID (alternative to tradeOrderId) | [optional] 
**ExternalOrderId** | Pointer to **string** | External order ID from venue | [optional] 

## Methods

### NewRpcCancelTradeOrderParams

`func NewRpcCancelTradeOrderParams(tradingAccountId string, ) *RpcCancelTradeOrderParams`

NewRpcCancelTradeOrderParams instantiates a new RpcCancelTradeOrderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcCancelTradeOrderParamsWithDefaults

`func NewRpcCancelTradeOrderParamsWithDefaults() *RpcCancelTradeOrderParams`

NewRpcCancelTradeOrderParamsWithDefaults instantiates a new RpcCancelTradeOrderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradeOrderId

`func (o *RpcCancelTradeOrderParams) GetTradeOrderId() string`

GetTradeOrderId returns the TradeOrderId field if non-nil, zero value otherwise.

### GetTradeOrderIdOk

`func (o *RpcCancelTradeOrderParams) GetTradeOrderIdOk() (*string, bool)`

GetTradeOrderIdOk returns a tuple with the TradeOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeOrderId

`func (o *RpcCancelTradeOrderParams) SetTradeOrderId(v string)`

SetTradeOrderId sets TradeOrderId field to given value.

### HasTradeOrderId

`func (o *RpcCancelTradeOrderParams) HasTradeOrderId() bool`

HasTradeOrderId returns a boolean if a field has been set.

### GetTradingAccountId

`func (o *RpcCancelTradeOrderParams) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *RpcCancelTradeOrderParams) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *RpcCancelTradeOrderParams) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.


### GetClientOrderId

`func (o *RpcCancelTradeOrderParams) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *RpcCancelTradeOrderParams) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *RpcCancelTradeOrderParams) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *RpcCancelTradeOrderParams) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetExternalOrderId

`func (o *RpcCancelTradeOrderParams) GetExternalOrderId() string`

GetExternalOrderId returns the ExternalOrderId field if non-nil, zero value otherwise.

### GetExternalOrderIdOk

`func (o *RpcCancelTradeOrderParams) GetExternalOrderIdOk() (*string, bool)`

GetExternalOrderIdOk returns a tuple with the ExternalOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOrderId

`func (o *RpcCancelTradeOrderParams) SetExternalOrderId(v string)`

SetExternalOrderId sets ExternalOrderId field to given value.

### HasExternalOrderId

`func (o *RpcCancelTradeOrderParams) HasExternalOrderId() bool`

HasExternalOrderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


