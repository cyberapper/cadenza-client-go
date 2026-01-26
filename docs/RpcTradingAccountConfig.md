# RpcTradingAccountConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Leverage** | Pointer to **int32** | Account leverage | [optional] 

## Methods

### NewRpcTradingAccountConfig

`func NewRpcTradingAccountConfig() *RpcTradingAccountConfig`

NewRpcTradingAccountConfig instantiates a new RpcTradingAccountConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcTradingAccountConfigWithDefaults

`func NewRpcTradingAccountConfigWithDefaults() *RpcTradingAccountConfig`

NewRpcTradingAccountConfigWithDefaults instantiates a new RpcTradingAccountConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeverage

`func (o *RpcTradingAccountConfig) GetLeverage() int32`

GetLeverage returns the Leverage field if non-nil, zero value otherwise.

### GetLeverageOk

`func (o *RpcTradingAccountConfig) GetLeverageOk() (*int32, bool)`

GetLeverageOk returns a tuple with the Leverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverage

`func (o *RpcTradingAccountConfig) SetLeverage(v int32)`

SetLeverage sets Leverage field to given value.

### HasLeverage

`func (o *RpcTradingAccountConfig) HasLeverage() bool`

HasLeverage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


