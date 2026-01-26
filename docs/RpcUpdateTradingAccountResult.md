# RpcUpdateTradingAccountResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**RpcTradingAccount**](RpcTradingAccount.md) |  | [optional] 
**Error** | Pointer to [**RpcError**](RpcError.md) |  | [optional] 

## Methods

### NewRpcUpdateTradingAccountResult

`func NewRpcUpdateTradingAccountResult() *RpcUpdateTradingAccountResult`

NewRpcUpdateTradingAccountResult instantiates a new RpcUpdateTradingAccountResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcUpdateTradingAccountResultWithDefaults

`func NewRpcUpdateTradingAccountResultWithDefaults() *RpcUpdateTradingAccountResult`

NewRpcUpdateTradingAccountResultWithDefaults instantiates a new RpcUpdateTradingAccountResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RpcUpdateTradingAccountResult) GetData() RpcTradingAccount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RpcUpdateTradingAccountResult) GetDataOk() (*RpcTradingAccount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RpcUpdateTradingAccountResult) SetData(v RpcTradingAccount)`

SetData sets Data field to given value.

### HasData

`func (o *RpcUpdateTradingAccountResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *RpcUpdateTradingAccountResult) GetError() RpcError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RpcUpdateTradingAccountResult) GetErrorOk() (*RpcError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RpcUpdateTradingAccountResult) SetError(v RpcError)`

SetError sets Error field to given value.

### HasError

`func (o *RpcUpdateTradingAccountResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


