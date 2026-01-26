# RpcSubscribeTradingAccountStreamParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradingAccount** | [**RpcTradingAccount**](RpcTradingAccount.md) |  | 
**SubscriptionType** | [**SubscriptionType**](SubscriptionType.md) |  | 

## Methods

### NewRpcSubscribeTradingAccountStreamParams

`func NewRpcSubscribeTradingAccountStreamParams(tradingAccount RpcTradingAccount, subscriptionType SubscriptionType, ) *RpcSubscribeTradingAccountStreamParams`

NewRpcSubscribeTradingAccountStreamParams instantiates a new RpcSubscribeTradingAccountStreamParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcSubscribeTradingAccountStreamParamsWithDefaults

`func NewRpcSubscribeTradingAccountStreamParamsWithDefaults() *RpcSubscribeTradingAccountStreamParams`

NewRpcSubscribeTradingAccountStreamParamsWithDefaults instantiates a new RpcSubscribeTradingAccountStreamParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradingAccount

`func (o *RpcSubscribeTradingAccountStreamParams) GetTradingAccount() RpcTradingAccount`

GetTradingAccount returns the TradingAccount field if non-nil, zero value otherwise.

### GetTradingAccountOk

`func (o *RpcSubscribeTradingAccountStreamParams) GetTradingAccountOk() (*RpcTradingAccount, bool)`

GetTradingAccountOk returns a tuple with the TradingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccount

`func (o *RpcSubscribeTradingAccountStreamParams) SetTradingAccount(v RpcTradingAccount)`

SetTradingAccount sets TradingAccount field to given value.


### GetSubscriptionType

`func (o *RpcSubscribeTradingAccountStreamParams) GetSubscriptionType() SubscriptionType`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *RpcSubscribeTradingAccountStreamParams) GetSubscriptionTypeOk() (*SubscriptionType, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *RpcSubscribeTradingAccountStreamParams) SetSubscriptionType(v SubscriptionType)`

SetSubscriptionType sets SubscriptionType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


