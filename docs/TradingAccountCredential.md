# TradingAccountCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialId** | **string** | UUID string | 
**Venue** | [**Venue**](Venue.md) |  | 
**CredentialType** | [**CredentialType**](CredentialType.md) |  | 
**Nickname** | Pointer to **string** |  | [optional] 
**Status** | [**TradingAccountCredentialStatus**](TradingAccountCredentialStatus.md) |  | 
**CreatedAt** | **int64** | Unix timestamp in milliseconds | 
**CreatedAtDateTime** | Pointer to **time.Time** | Creation timestamp in ISO 8601 format | [optional] 
**UpdatedAt** | **int64** | Unix timestamp in milliseconds | 
**UpdatedAtDateTime** | Pointer to **time.Time** | Last update timestamp in ISO 8601 format | [optional] 
**RevokedAt** | Pointer to **int64** | Unix timestamp in milliseconds | [optional] 
**RevokedAtDateTime** | Pointer to **time.Time** | Revocation timestamp in ISO 8601 format | [optional] 

## Methods

### NewTradingAccountCredential

`func NewTradingAccountCredential(credentialId string, venue Venue, credentialType CredentialType, status TradingAccountCredentialStatus, createdAt int64, updatedAt int64, ) *TradingAccountCredential`

NewTradingAccountCredential instantiates a new TradingAccountCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradingAccountCredentialWithDefaults

`func NewTradingAccountCredentialWithDefaults() *TradingAccountCredential`

NewTradingAccountCredentialWithDefaults instantiates a new TradingAccountCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialId

`func (o *TradingAccountCredential) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *TradingAccountCredential) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *TradingAccountCredential) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.


### GetVenue

`func (o *TradingAccountCredential) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *TradingAccountCredential) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *TradingAccountCredential) SetVenue(v Venue)`

SetVenue sets Venue field to given value.


### GetCredentialType

`func (o *TradingAccountCredential) GetCredentialType() CredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *TradingAccountCredential) GetCredentialTypeOk() (*CredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *TradingAccountCredential) SetCredentialType(v CredentialType)`

SetCredentialType sets CredentialType field to given value.


### GetNickname

`func (o *TradingAccountCredential) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *TradingAccountCredential) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *TradingAccountCredential) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *TradingAccountCredential) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetStatus

`func (o *TradingAccountCredential) GetStatus() TradingAccountCredentialStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TradingAccountCredential) GetStatusOk() (*TradingAccountCredentialStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TradingAccountCredential) SetStatus(v TradingAccountCredentialStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *TradingAccountCredential) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TradingAccountCredential) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TradingAccountCredential) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedAtDateTime

`func (o *TradingAccountCredential) GetCreatedAtDateTime() time.Time`

GetCreatedAtDateTime returns the CreatedAtDateTime field if non-nil, zero value otherwise.

### GetCreatedAtDateTimeOk

`func (o *TradingAccountCredential) GetCreatedAtDateTimeOk() (*time.Time, bool)`

GetCreatedAtDateTimeOk returns a tuple with the CreatedAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtDateTime

`func (o *TradingAccountCredential) SetCreatedAtDateTime(v time.Time)`

SetCreatedAtDateTime sets CreatedAtDateTime field to given value.

### HasCreatedAtDateTime

`func (o *TradingAccountCredential) HasCreatedAtDateTime() bool`

HasCreatedAtDateTime returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TradingAccountCredential) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TradingAccountCredential) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TradingAccountCredential) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpdatedAtDateTime

`func (o *TradingAccountCredential) GetUpdatedAtDateTime() time.Time`

GetUpdatedAtDateTime returns the UpdatedAtDateTime field if non-nil, zero value otherwise.

### GetUpdatedAtDateTimeOk

`func (o *TradingAccountCredential) GetUpdatedAtDateTimeOk() (*time.Time, bool)`

GetUpdatedAtDateTimeOk returns a tuple with the UpdatedAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtDateTime

`func (o *TradingAccountCredential) SetUpdatedAtDateTime(v time.Time)`

SetUpdatedAtDateTime sets UpdatedAtDateTime field to given value.

### HasUpdatedAtDateTime

`func (o *TradingAccountCredential) HasUpdatedAtDateTime() bool`

HasUpdatedAtDateTime returns a boolean if a field has been set.

### GetRevokedAt

`func (o *TradingAccountCredential) GetRevokedAt() int64`

GetRevokedAt returns the RevokedAt field if non-nil, zero value otherwise.

### GetRevokedAtOk

`func (o *TradingAccountCredential) GetRevokedAtOk() (*int64, bool)`

GetRevokedAtOk returns a tuple with the RevokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAt

`func (o *TradingAccountCredential) SetRevokedAt(v int64)`

SetRevokedAt sets RevokedAt field to given value.

### HasRevokedAt

`func (o *TradingAccountCredential) HasRevokedAt() bool`

HasRevokedAt returns a boolean if a field has been set.

### GetRevokedAtDateTime

`func (o *TradingAccountCredential) GetRevokedAtDateTime() time.Time`

GetRevokedAtDateTime returns the RevokedAtDateTime field if non-nil, zero value otherwise.

### GetRevokedAtDateTimeOk

`func (o *TradingAccountCredential) GetRevokedAtDateTimeOk() (*time.Time, bool)`

GetRevokedAtDateTimeOk returns a tuple with the RevokedAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAtDateTime

`func (o *TradingAccountCredential) SetRevokedAtDateTime(v time.Time)`

SetRevokedAtDateTime sets RevokedAtDateTime field to given value.

### HasRevokedAtDateTime

`func (o *TradingAccountCredential) HasRevokedAtDateTime() bool`

HasRevokedAtDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


