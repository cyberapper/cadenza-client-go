# CreateTradingAccountCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Venue** | [**Venue**](Venue.md) |  | 
**Nickname** | Pointer to **string** | Nickname of the credential | [optional] 
**CredentialType** | [**CredentialType**](CredentialType.md) |  | 
**ApiKey** | Pointer to **string** |  | [optional] 
**ApiSecret** | Pointer to **string** |  | [optional] 
**ApiPassphrase** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateTradingAccountCredentialRequest

`func NewCreateTradingAccountCredentialRequest(venue Venue, credentialType CredentialType, ) *CreateTradingAccountCredentialRequest`

NewCreateTradingAccountCredentialRequest instantiates a new CreateTradingAccountCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradingAccountCredentialRequestWithDefaults

`func NewCreateTradingAccountCredentialRequestWithDefaults() *CreateTradingAccountCredentialRequest`

NewCreateTradingAccountCredentialRequestWithDefaults instantiates a new CreateTradingAccountCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVenue

`func (o *CreateTradingAccountCredentialRequest) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *CreateTradingAccountCredentialRequest) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *CreateTradingAccountCredentialRequest) SetVenue(v Venue)`

SetVenue sets Venue field to given value.


### GetNickname

`func (o *CreateTradingAccountCredentialRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *CreateTradingAccountCredentialRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *CreateTradingAccountCredentialRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *CreateTradingAccountCredentialRequest) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetCredentialType

`func (o *CreateTradingAccountCredentialRequest) GetCredentialType() CredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *CreateTradingAccountCredentialRequest) GetCredentialTypeOk() (*CredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *CreateTradingAccountCredentialRequest) SetCredentialType(v CredentialType)`

SetCredentialType sets CredentialType field to given value.


### GetApiKey

`func (o *CreateTradingAccountCredentialRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CreateTradingAccountCredentialRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CreateTradingAccountCredentialRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *CreateTradingAccountCredentialRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetApiSecret

`func (o *CreateTradingAccountCredentialRequest) GetApiSecret() string`

GetApiSecret returns the ApiSecret field if non-nil, zero value otherwise.

### GetApiSecretOk

`func (o *CreateTradingAccountCredentialRequest) GetApiSecretOk() (*string, bool)`

GetApiSecretOk returns a tuple with the ApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSecret

`func (o *CreateTradingAccountCredentialRequest) SetApiSecret(v string)`

SetApiSecret sets ApiSecret field to given value.

### HasApiSecret

`func (o *CreateTradingAccountCredentialRequest) HasApiSecret() bool`

HasApiSecret returns a boolean if a field has been set.

### GetApiPassphrase

`func (o *CreateTradingAccountCredentialRequest) GetApiPassphrase() string`

GetApiPassphrase returns the ApiPassphrase field if non-nil, zero value otherwise.

### GetApiPassphraseOk

`func (o *CreateTradingAccountCredentialRequest) GetApiPassphraseOk() (*string, bool)`

GetApiPassphraseOk returns a tuple with the ApiPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPassphrase

`func (o *CreateTradingAccountCredentialRequest) SetApiPassphrase(v string)`

SetApiPassphrase sets ApiPassphrase field to given value.

### HasApiPassphrase

`func (o *CreateTradingAccountCredentialRequest) HasApiPassphrase() bool`

HasApiPassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


