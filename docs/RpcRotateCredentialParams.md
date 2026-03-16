# RpcRotateCredentialParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialId** | **string** |  | 
**CredentialType** | [**CredentialType**](CredentialType.md) |  | 
**ApiKey** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**SecretPassphrase** | Pointer to **string** |  | [optional] 

## Methods

### NewRpcRotateCredentialParams

`func NewRpcRotateCredentialParams(credentialId string, credentialType CredentialType, ) *RpcRotateCredentialParams`

NewRpcRotateCredentialParams instantiates a new RpcRotateCredentialParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcRotateCredentialParamsWithDefaults

`func NewRpcRotateCredentialParamsWithDefaults() *RpcRotateCredentialParams`

NewRpcRotateCredentialParamsWithDefaults instantiates a new RpcRotateCredentialParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialId

`func (o *RpcRotateCredentialParams) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *RpcRotateCredentialParams) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *RpcRotateCredentialParams) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.


### GetCredentialType

`func (o *RpcRotateCredentialParams) GetCredentialType() CredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *RpcRotateCredentialParams) GetCredentialTypeOk() (*CredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *RpcRotateCredentialParams) SetCredentialType(v CredentialType)`

SetCredentialType sets CredentialType field to given value.


### GetApiKey

`func (o *RpcRotateCredentialParams) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *RpcRotateCredentialParams) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *RpcRotateCredentialParams) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *RpcRotateCredentialParams) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *RpcRotateCredentialParams) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *RpcRotateCredentialParams) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *RpcRotateCredentialParams) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *RpcRotateCredentialParams) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetSecretPassphrase

`func (o *RpcRotateCredentialParams) GetSecretPassphrase() string`

GetSecretPassphrase returns the SecretPassphrase field if non-nil, zero value otherwise.

### GetSecretPassphraseOk

`func (o *RpcRotateCredentialParams) GetSecretPassphraseOk() (*string, bool)`

GetSecretPassphraseOk returns a tuple with the SecretPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretPassphrase

`func (o *RpcRotateCredentialParams) SetSecretPassphrase(v string)`

SetSecretPassphrase sets SecretPassphrase field to given value.

### HasSecretPassphrase

`func (o *RpcRotateCredentialParams) HasSecretPassphrase() bool`

HasSecretPassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


