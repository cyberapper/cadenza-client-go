# CancelTradeOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradingAccountId** | **string** | UUID string | 
**TradeOrderId** | **string** | UUID string | 

## Methods

### NewCancelTradeOrderRequest

`func NewCancelTradeOrderRequest(tradingAccountId string, tradeOrderId string, ) *CancelTradeOrderRequest`

NewCancelTradeOrderRequest instantiates a new CancelTradeOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelTradeOrderRequestWithDefaults

`func NewCancelTradeOrderRequestWithDefaults() *CancelTradeOrderRequest`

NewCancelTradeOrderRequestWithDefaults instantiates a new CancelTradeOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradingAccountId

`func (o *CancelTradeOrderRequest) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *CancelTradeOrderRequest) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *CancelTradeOrderRequest) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.


### GetTradeOrderId

`func (o *CancelTradeOrderRequest) GetTradeOrderId() string`

GetTradeOrderId returns the TradeOrderId field if non-nil, zero value otherwise.

### GetTradeOrderIdOk

`func (o *CancelTradeOrderRequest) GetTradeOrderIdOk() (*string, bool)`

GetTradeOrderIdOk returns a tuple with the TradeOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeOrderId

`func (o *CancelTradeOrderRequest) SetTradeOrderId(v string)`

SetTradeOrderId sets TradeOrderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


