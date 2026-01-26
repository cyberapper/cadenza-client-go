# RpcSaveTradingAccountsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RpcTradingAccount**](RpcTradingAccount.md) |  | [optional] 
**Error** | Pointer to [**RpcError**](RpcError.md) |  | [optional] 

## Methods

### NewRpcSaveTradingAccountsResult

`func NewRpcSaveTradingAccountsResult() *RpcSaveTradingAccountsResult`

NewRpcSaveTradingAccountsResult instantiates a new RpcSaveTradingAccountsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcSaveTradingAccountsResultWithDefaults

`func NewRpcSaveTradingAccountsResultWithDefaults() *RpcSaveTradingAccountsResult`

NewRpcSaveTradingAccountsResultWithDefaults instantiates a new RpcSaveTradingAccountsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RpcSaveTradingAccountsResult) GetData() []RpcTradingAccount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RpcSaveTradingAccountsResult) GetDataOk() (*[]RpcTradingAccount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RpcSaveTradingAccountsResult) SetData(v []RpcTradingAccount)`

SetData sets Data field to given value.

### HasData

`func (o *RpcSaveTradingAccountsResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *RpcSaveTradingAccountsResult) GetError() RpcError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RpcSaveTradingAccountsResult) GetErrorOk() (*RpcError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RpcSaveTradingAccountsResult) SetError(v RpcError)`

SetError sets Error field to given value.

### HasError

`func (o *RpcSaveTradingAccountsResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


