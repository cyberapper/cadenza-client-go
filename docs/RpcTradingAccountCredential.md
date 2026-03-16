# RpcTradingAccountCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialId** | Pointer to **string** |  | [optional] 
**Nickname** | Pointer to **string** |  | [optional] 
**CredentialType** | Pointer to [**CredentialType**](CredentialType.md) |  | [optional] 
**Status** | Pointer to [**CredentialStatus**](CredentialStatus.md) |  | [optional] 
**Venue** | Pointer to [**Venue**](Venue.md) |  | [optional] 
**ApiKey** | Pointer to **string** | API key (only in responses where withSecret&#x3D;true) | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**RevokedAt** | Pointer to **time.Time** |  | [optional] 
**WithInfo** | Pointer to **bool** | Whether metadata fields are populated | [optional] 
**WithSecret** | Pointer to **bool** | Whether secret fields are populated | [optional] 

## Methods

### NewRpcTradingAccountCredential

`func NewRpcTradingAccountCredential() *RpcTradingAccountCredential`

NewRpcTradingAccountCredential instantiates a new RpcTradingAccountCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcTradingAccountCredentialWithDefaults

`func NewRpcTradingAccountCredentialWithDefaults() *RpcTradingAccountCredential`

NewRpcTradingAccountCredentialWithDefaults instantiates a new RpcTradingAccountCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialId

`func (o *RpcTradingAccountCredential) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *RpcTradingAccountCredential) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *RpcTradingAccountCredential) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.

### HasCredentialId

`func (o *RpcTradingAccountCredential) HasCredentialId() bool`

HasCredentialId returns a boolean if a field has been set.

### GetNickname

`func (o *RpcTradingAccountCredential) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *RpcTradingAccountCredential) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *RpcTradingAccountCredential) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *RpcTradingAccountCredential) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetCredentialType

`func (o *RpcTradingAccountCredential) GetCredentialType() CredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *RpcTradingAccountCredential) GetCredentialTypeOk() (*CredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *RpcTradingAccountCredential) SetCredentialType(v CredentialType)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *RpcTradingAccountCredential) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### GetStatus

`func (o *RpcTradingAccountCredential) GetStatus() CredentialStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RpcTradingAccountCredential) GetStatusOk() (*CredentialStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RpcTradingAccountCredential) SetStatus(v CredentialStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RpcTradingAccountCredential) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVenue

`func (o *RpcTradingAccountCredential) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *RpcTradingAccountCredential) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *RpcTradingAccountCredential) SetVenue(v Venue)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *RpcTradingAccountCredential) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetApiKey

`func (o *RpcTradingAccountCredential) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *RpcTradingAccountCredential) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *RpcTradingAccountCredential) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *RpcTradingAccountCredential) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RpcTradingAccountCredential) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RpcTradingAccountCredential) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RpcTradingAccountCredential) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RpcTradingAccountCredential) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RpcTradingAccountCredential) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RpcTradingAccountCredential) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RpcTradingAccountCredential) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RpcTradingAccountCredential) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRevokedAt

`func (o *RpcTradingAccountCredential) GetRevokedAt() time.Time`

GetRevokedAt returns the RevokedAt field if non-nil, zero value otherwise.

### GetRevokedAtOk

`func (o *RpcTradingAccountCredential) GetRevokedAtOk() (*time.Time, bool)`

GetRevokedAtOk returns a tuple with the RevokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAt

`func (o *RpcTradingAccountCredential) SetRevokedAt(v time.Time)`

SetRevokedAt sets RevokedAt field to given value.

### HasRevokedAt

`func (o *RpcTradingAccountCredential) HasRevokedAt() bool`

HasRevokedAt returns a boolean if a field has been set.

### GetWithInfo

`func (o *RpcTradingAccountCredential) GetWithInfo() bool`

GetWithInfo returns the WithInfo field if non-nil, zero value otherwise.

### GetWithInfoOk

`func (o *RpcTradingAccountCredential) GetWithInfoOk() (*bool, bool)`

GetWithInfoOk returns a tuple with the WithInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithInfo

`func (o *RpcTradingAccountCredential) SetWithInfo(v bool)`

SetWithInfo sets WithInfo field to given value.

### HasWithInfo

`func (o *RpcTradingAccountCredential) HasWithInfo() bool`

HasWithInfo returns a boolean if a field has been set.

### GetWithSecret

`func (o *RpcTradingAccountCredential) GetWithSecret() bool`

GetWithSecret returns the WithSecret field if non-nil, zero value otherwise.

### GetWithSecretOk

`func (o *RpcTradingAccountCredential) GetWithSecretOk() (*bool, bool)`

GetWithSecretOk returns a tuple with the WithSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithSecret

`func (o *RpcTradingAccountCredential) SetWithSecret(v bool)`

SetWithSecret sets WithSecret field to given value.

### HasWithSecret

`func (o *RpcTradingAccountCredential) HasWithSecret() bool`

HasWithSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


