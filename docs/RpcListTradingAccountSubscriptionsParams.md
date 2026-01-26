# RpcListTradingAccountSubscriptionsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradingAccountId** | **string** |  | 
**SubscriptionType** | Pointer to [**SubscriptionType**](SubscriptionType.md) |  | [optional] 
**Status** | Pointer to [**SubscriptionStatus**](SubscriptionStatus.md) |  | [optional] 
**Pagination** | Pointer to [**RpcPagination**](RpcPagination.md) |  | [optional] 

## Methods

### NewRpcListTradingAccountSubscriptionsParams

`func NewRpcListTradingAccountSubscriptionsParams(tradingAccountId string, ) *RpcListTradingAccountSubscriptionsParams`

NewRpcListTradingAccountSubscriptionsParams instantiates a new RpcListTradingAccountSubscriptionsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcListTradingAccountSubscriptionsParamsWithDefaults

`func NewRpcListTradingAccountSubscriptionsParamsWithDefaults() *RpcListTradingAccountSubscriptionsParams`

NewRpcListTradingAccountSubscriptionsParamsWithDefaults instantiates a new RpcListTradingAccountSubscriptionsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradingAccountId

`func (o *RpcListTradingAccountSubscriptionsParams) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *RpcListTradingAccountSubscriptionsParams) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *RpcListTradingAccountSubscriptionsParams) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.


### GetSubscriptionType

`func (o *RpcListTradingAccountSubscriptionsParams) GetSubscriptionType() SubscriptionType`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *RpcListTradingAccountSubscriptionsParams) GetSubscriptionTypeOk() (*SubscriptionType, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *RpcListTradingAccountSubscriptionsParams) SetSubscriptionType(v SubscriptionType)`

SetSubscriptionType sets SubscriptionType field to given value.

### HasSubscriptionType

`func (o *RpcListTradingAccountSubscriptionsParams) HasSubscriptionType() bool`

HasSubscriptionType returns a boolean if a field has been set.

### GetStatus

`func (o *RpcListTradingAccountSubscriptionsParams) GetStatus() SubscriptionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RpcListTradingAccountSubscriptionsParams) GetStatusOk() (*SubscriptionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RpcListTradingAccountSubscriptionsParams) SetStatus(v SubscriptionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RpcListTradingAccountSubscriptionsParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPagination

`func (o *RpcListTradingAccountSubscriptionsParams) GetPagination() RpcPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RpcListTradingAccountSubscriptionsParams) GetPaginationOk() (*RpcPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RpcListTradingAccountSubscriptionsParams) SetPagination(v RpcPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RpcListTradingAccountSubscriptionsParams) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


