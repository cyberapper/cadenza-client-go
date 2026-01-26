# RpcVerifyCredentialParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialIds** | Pointer to **[]string** |  | [optional] 
**Credentials** | Pointer to [**[]RpcTradingAccountCredential**](RpcTradingAccountCredential.md) |  | [optional] 

## Methods

### NewRpcVerifyCredentialParams

`func NewRpcVerifyCredentialParams() *RpcVerifyCredentialParams`

NewRpcVerifyCredentialParams instantiates a new RpcVerifyCredentialParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcVerifyCredentialParamsWithDefaults

`func NewRpcVerifyCredentialParamsWithDefaults() *RpcVerifyCredentialParams`

NewRpcVerifyCredentialParamsWithDefaults instantiates a new RpcVerifyCredentialParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialIds

`func (o *RpcVerifyCredentialParams) GetCredentialIds() []string`

GetCredentialIds returns the CredentialIds field if non-nil, zero value otherwise.

### GetCredentialIdsOk

`func (o *RpcVerifyCredentialParams) GetCredentialIdsOk() (*[]string, bool)`

GetCredentialIdsOk returns a tuple with the CredentialIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialIds

`func (o *RpcVerifyCredentialParams) SetCredentialIds(v []string)`

SetCredentialIds sets CredentialIds field to given value.

### HasCredentialIds

`func (o *RpcVerifyCredentialParams) HasCredentialIds() bool`

HasCredentialIds returns a boolean if a field has been set.

### GetCredentials

`func (o *RpcVerifyCredentialParams) GetCredentials() []RpcTradingAccountCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *RpcVerifyCredentialParams) GetCredentialsOk() (*[]RpcTradingAccountCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *RpcVerifyCredentialParams) SetCredentials(v []RpcTradingAccountCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *RpcVerifyCredentialParams) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


