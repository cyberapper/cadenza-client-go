# RpcSaveCredentialParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | [**RpcTradingAccountCredential**](RpcTradingAccountCredential.md) |  | 

## Methods

### NewRpcSaveCredentialParams

`func NewRpcSaveCredentialParams(credential RpcTradingAccountCredential, ) *RpcSaveCredentialParams`

NewRpcSaveCredentialParams instantiates a new RpcSaveCredentialParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcSaveCredentialParamsWithDefaults

`func NewRpcSaveCredentialParamsWithDefaults() *RpcSaveCredentialParams`

NewRpcSaveCredentialParamsWithDefaults instantiates a new RpcSaveCredentialParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredential

`func (o *RpcSaveCredentialParams) GetCredential() RpcTradingAccountCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *RpcSaveCredentialParams) GetCredentialOk() (*RpcTradingAccountCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *RpcSaveCredentialParams) SetCredential(v RpcTradingAccountCredential)`

SetCredential sets Credential field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


