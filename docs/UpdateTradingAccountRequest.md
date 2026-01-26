# UpdateTradingAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradingAccountId** | **string** | UUID string | 
**Nickname** | Pointer to **string** | New nickname for the trading account | [optional] 

## Methods

### NewUpdateTradingAccountRequest

`func NewUpdateTradingAccountRequest(tradingAccountId string, ) *UpdateTradingAccountRequest`

NewUpdateTradingAccountRequest instantiates a new UpdateTradingAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTradingAccountRequestWithDefaults

`func NewUpdateTradingAccountRequestWithDefaults() *UpdateTradingAccountRequest`

NewUpdateTradingAccountRequestWithDefaults instantiates a new UpdateTradingAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradingAccountId

`func (o *UpdateTradingAccountRequest) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *UpdateTradingAccountRequest) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *UpdateTradingAccountRequest) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.


### GetNickname

`func (o *UpdateTradingAccountRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *UpdateTradingAccountRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *UpdateTradingAccountRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *UpdateTradingAccountRequest) HasNickname() bool`

HasNickname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


