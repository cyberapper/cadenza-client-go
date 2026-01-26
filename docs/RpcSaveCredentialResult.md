# RpcSaveCredentialResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**RpcTradingAccountCredential**](RpcTradingAccountCredential.md) |  | [optional] 
**Error** | Pointer to [**RpcError**](RpcError.md) |  | [optional] 

## Methods

### NewRpcSaveCredentialResult

`func NewRpcSaveCredentialResult() *RpcSaveCredentialResult`

NewRpcSaveCredentialResult instantiates a new RpcSaveCredentialResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcSaveCredentialResultWithDefaults

`func NewRpcSaveCredentialResultWithDefaults() *RpcSaveCredentialResult`

NewRpcSaveCredentialResultWithDefaults instantiates a new RpcSaveCredentialResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RpcSaveCredentialResult) GetData() RpcTradingAccountCredential`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RpcSaveCredentialResult) GetDataOk() (*RpcTradingAccountCredential, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RpcSaveCredentialResult) SetData(v RpcTradingAccountCredential)`

SetData sets Data field to given value.

### HasData

`func (o *RpcSaveCredentialResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *RpcSaveCredentialResult) GetError() RpcError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RpcSaveCredentialResult) GetErrorOk() (*RpcError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RpcSaveCredentialResult) SetError(v RpcError)`

SetError sets Error field to given value.

### HasError

`func (o *RpcSaveCredentialResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


