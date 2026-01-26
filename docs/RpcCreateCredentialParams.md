# RpcCreateCredentialParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradingAccountId** | **string** |  | 
**Venue** | [**Venue**](Venue.md) |  | 
**CredentialType** | [**CredentialType**](CredentialType.md) |  | 
**ApiKey** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**SecretPassphrase** | Pointer to **string** |  | [optional] 
**Nickname** | Pointer to **string** |  | [optional] 

## Methods

### NewRpcCreateCredentialParams

`func NewRpcCreateCredentialParams(tradingAccountId string, venue Venue, credentialType CredentialType, ) *RpcCreateCredentialParams`

NewRpcCreateCredentialParams instantiates a new RpcCreateCredentialParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcCreateCredentialParamsWithDefaults

`func NewRpcCreateCredentialParamsWithDefaults() *RpcCreateCredentialParams`

NewRpcCreateCredentialParamsWithDefaults instantiates a new RpcCreateCredentialParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradingAccountId

`func (o *RpcCreateCredentialParams) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *RpcCreateCredentialParams) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *RpcCreateCredentialParams) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.


### GetVenue

`func (o *RpcCreateCredentialParams) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *RpcCreateCredentialParams) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *RpcCreateCredentialParams) SetVenue(v Venue)`

SetVenue sets Venue field to given value.


### GetCredentialType

`func (o *RpcCreateCredentialParams) GetCredentialType() CredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *RpcCreateCredentialParams) GetCredentialTypeOk() (*CredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *RpcCreateCredentialParams) SetCredentialType(v CredentialType)`

SetCredentialType sets CredentialType field to given value.


### GetApiKey

`func (o *RpcCreateCredentialParams) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *RpcCreateCredentialParams) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *RpcCreateCredentialParams) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *RpcCreateCredentialParams) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *RpcCreateCredentialParams) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *RpcCreateCredentialParams) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *RpcCreateCredentialParams) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *RpcCreateCredentialParams) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetSecretPassphrase

`func (o *RpcCreateCredentialParams) GetSecretPassphrase() string`

GetSecretPassphrase returns the SecretPassphrase field if non-nil, zero value otherwise.

### GetSecretPassphraseOk

`func (o *RpcCreateCredentialParams) GetSecretPassphraseOk() (*string, bool)`

GetSecretPassphraseOk returns a tuple with the SecretPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretPassphrase

`func (o *RpcCreateCredentialParams) SetSecretPassphrase(v string)`

SetSecretPassphrase sets SecretPassphrase field to given value.

### HasSecretPassphrase

`func (o *RpcCreateCredentialParams) HasSecretPassphrase() bool`

HasSecretPassphrase returns a boolean if a field has been set.

### GetNickname

`func (o *RpcCreateCredentialParams) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *RpcCreateCredentialParams) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *RpcCreateCredentialParams) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *RpcCreateCredentialParams) HasNickname() bool`

HasNickname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


