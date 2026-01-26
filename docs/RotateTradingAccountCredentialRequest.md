# RotateTradingAccountCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialId** | **string** | UUID string | 
**ApiKey** | **string** |  | 
**ApiSecret** | Pointer to **string** |  | [optional] 
**ApiPassphrase** | Pointer to **string** |  | [optional] 

## Methods

### NewRotateTradingAccountCredentialRequest

`func NewRotateTradingAccountCredentialRequest(credentialId string, apiKey string, ) *RotateTradingAccountCredentialRequest`

NewRotateTradingAccountCredentialRequest instantiates a new RotateTradingAccountCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotateTradingAccountCredentialRequestWithDefaults

`func NewRotateTradingAccountCredentialRequestWithDefaults() *RotateTradingAccountCredentialRequest`

NewRotateTradingAccountCredentialRequestWithDefaults instantiates a new RotateTradingAccountCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialId

`func (o *RotateTradingAccountCredentialRequest) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *RotateTradingAccountCredentialRequest) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *RotateTradingAccountCredentialRequest) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.


### GetApiKey

`func (o *RotateTradingAccountCredentialRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *RotateTradingAccountCredentialRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *RotateTradingAccountCredentialRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetApiSecret

`func (o *RotateTradingAccountCredentialRequest) GetApiSecret() string`

GetApiSecret returns the ApiSecret field if non-nil, zero value otherwise.

### GetApiSecretOk

`func (o *RotateTradingAccountCredentialRequest) GetApiSecretOk() (*string, bool)`

GetApiSecretOk returns a tuple with the ApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSecret

`func (o *RotateTradingAccountCredentialRequest) SetApiSecret(v string)`

SetApiSecret sets ApiSecret field to given value.

### HasApiSecret

`func (o *RotateTradingAccountCredentialRequest) HasApiSecret() bool`

HasApiSecret returns a boolean if a field has been set.

### GetApiPassphrase

`func (o *RotateTradingAccountCredentialRequest) GetApiPassphrase() string`

GetApiPassphrase returns the ApiPassphrase field if non-nil, zero value otherwise.

### GetApiPassphraseOk

`func (o *RotateTradingAccountCredentialRequest) GetApiPassphraseOk() (*string, bool)`

GetApiPassphraseOk returns a tuple with the ApiPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPassphrase

`func (o *RotateTradingAccountCredentialRequest) SetApiPassphrase(v string)`

SetApiPassphrase sets ApiPassphrase field to given value.

### HasApiPassphrase

`func (o *RotateTradingAccountCredentialRequest) HasApiPassphrase() bool`

HasApiPassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


