# RpcPortfolio

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradingAccountId** | Pointer to **string** |  | [optional] 
**AccountInfo** | Pointer to [**RpcTradingAccount**](RpcTradingAccount.md) |  | [optional] 
**Balances** | Pointer to [**[]RpcBalanceEntry**](RpcBalanceEntry.md) |  | [optional] 
**Positions** | Pointer to [**[]RpcPositionEntry**](RpcPositionEntry.md) |  | [optional] 
**Summary** | Pointer to [**RpcPortfolioSummary**](RpcPortfolioSummary.md) |  | [optional] 
**UpdateMode** | Pointer to [**UpdateMode**](UpdateMode.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRpcPortfolio

`func NewRpcPortfolio() *RpcPortfolio`

NewRpcPortfolio instantiates a new RpcPortfolio object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcPortfolioWithDefaults

`func NewRpcPortfolioWithDefaults() *RpcPortfolio`

NewRpcPortfolioWithDefaults instantiates a new RpcPortfolio object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradingAccountId

`func (o *RpcPortfolio) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *RpcPortfolio) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *RpcPortfolio) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.

### HasTradingAccountId

`func (o *RpcPortfolio) HasTradingAccountId() bool`

HasTradingAccountId returns a boolean if a field has been set.

### GetAccountInfo

`func (o *RpcPortfolio) GetAccountInfo() RpcTradingAccount`

GetAccountInfo returns the AccountInfo field if non-nil, zero value otherwise.

### GetAccountInfoOk

`func (o *RpcPortfolio) GetAccountInfoOk() (*RpcTradingAccount, bool)`

GetAccountInfoOk returns a tuple with the AccountInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInfo

`func (o *RpcPortfolio) SetAccountInfo(v RpcTradingAccount)`

SetAccountInfo sets AccountInfo field to given value.

### HasAccountInfo

`func (o *RpcPortfolio) HasAccountInfo() bool`

HasAccountInfo returns a boolean if a field has been set.

### GetBalances

`func (o *RpcPortfolio) GetBalances() []RpcBalanceEntry`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *RpcPortfolio) GetBalancesOk() (*[]RpcBalanceEntry, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *RpcPortfolio) SetBalances(v []RpcBalanceEntry)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *RpcPortfolio) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetPositions

`func (o *RpcPortfolio) GetPositions() []RpcPositionEntry`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *RpcPortfolio) GetPositionsOk() (*[]RpcPositionEntry, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *RpcPortfolio) SetPositions(v []RpcPositionEntry)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *RpcPortfolio) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetSummary

`func (o *RpcPortfolio) GetSummary() RpcPortfolioSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *RpcPortfolio) GetSummaryOk() (*RpcPortfolioSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *RpcPortfolio) SetSummary(v RpcPortfolioSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *RpcPortfolio) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUpdateMode

`func (o *RpcPortfolio) GetUpdateMode() UpdateMode`

GetUpdateMode returns the UpdateMode field if non-nil, zero value otherwise.

### GetUpdateModeOk

`func (o *RpcPortfolio) GetUpdateModeOk() (*UpdateMode, bool)`

GetUpdateModeOk returns a tuple with the UpdateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateMode

`func (o *RpcPortfolio) SetUpdateMode(v UpdateMode)`

SetUpdateMode sets UpdateMode field to given value.

### HasUpdateMode

`func (o *RpcPortfolio) HasUpdateMode() bool`

HasUpdateMode returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RpcPortfolio) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RpcPortfolio) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RpcPortfolio) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RpcPortfolio) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


