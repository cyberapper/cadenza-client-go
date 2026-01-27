# AuthUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | User ID | [optional] 
**Email** | Pointer to **string** | User email address | [optional] 
**Phone** | Pointer to **NullableString** | User phone number | [optional] 
**EmailConfirmedAt** | Pointer to **NullableTime** | Email confirmation timestamp | [optional] 
**PhoneConfirmedAt** | Pointer to **NullableTime** | Phone confirmation timestamp | [optional] 
**LastSignInAt** | Pointer to **NullableTime** | Last sign in timestamp | [optional] 
**Role** | Pointer to **string** | User role | [optional] 
**CreatedAt** | Pointer to **time.Time** | Account creation timestamp | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Account last update timestamp | [optional] 
**AppMetadata** | Pointer to [**AuthUserAppMetadata**](AuthUserAppMetadata.md) |  | [optional] 
**UserMetadata** | Pointer to **map[string]interface{}** | User-defined metadata | [optional] 

## Methods

### NewAuthUser

`func NewAuthUser() *AuthUser`

NewAuthUser instantiates a new AuthUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthUserWithDefaults

`func NewAuthUserWithDefaults() *AuthUser`

NewAuthUserWithDefaults instantiates a new AuthUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *AuthUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AuthUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *AuthUser) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AuthUser) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AuthUser) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *AuthUser) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *AuthUser) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *AuthUser) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetEmailConfirmedAt

`func (o *AuthUser) GetEmailConfirmedAt() time.Time`

GetEmailConfirmedAt returns the EmailConfirmedAt field if non-nil, zero value otherwise.

### GetEmailConfirmedAtOk

`func (o *AuthUser) GetEmailConfirmedAtOk() (*time.Time, bool)`

GetEmailConfirmedAtOk returns a tuple with the EmailConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConfirmedAt

`func (o *AuthUser) SetEmailConfirmedAt(v time.Time)`

SetEmailConfirmedAt sets EmailConfirmedAt field to given value.

### HasEmailConfirmedAt

`func (o *AuthUser) HasEmailConfirmedAt() bool`

HasEmailConfirmedAt returns a boolean if a field has been set.

### SetEmailConfirmedAtNil

`func (o *AuthUser) SetEmailConfirmedAtNil(b bool)`

 SetEmailConfirmedAtNil sets the value for EmailConfirmedAt to be an explicit nil

### UnsetEmailConfirmedAt
`func (o *AuthUser) UnsetEmailConfirmedAt()`

UnsetEmailConfirmedAt ensures that no value is present for EmailConfirmedAt, not even an explicit nil
### GetPhoneConfirmedAt

`func (o *AuthUser) GetPhoneConfirmedAt() time.Time`

GetPhoneConfirmedAt returns the PhoneConfirmedAt field if non-nil, zero value otherwise.

### GetPhoneConfirmedAtOk

`func (o *AuthUser) GetPhoneConfirmedAtOk() (*time.Time, bool)`

GetPhoneConfirmedAtOk returns a tuple with the PhoneConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneConfirmedAt

`func (o *AuthUser) SetPhoneConfirmedAt(v time.Time)`

SetPhoneConfirmedAt sets PhoneConfirmedAt field to given value.

### HasPhoneConfirmedAt

`func (o *AuthUser) HasPhoneConfirmedAt() bool`

HasPhoneConfirmedAt returns a boolean if a field has been set.

### SetPhoneConfirmedAtNil

`func (o *AuthUser) SetPhoneConfirmedAtNil(b bool)`

 SetPhoneConfirmedAtNil sets the value for PhoneConfirmedAt to be an explicit nil

### UnsetPhoneConfirmedAt
`func (o *AuthUser) UnsetPhoneConfirmedAt()`

UnsetPhoneConfirmedAt ensures that no value is present for PhoneConfirmedAt, not even an explicit nil
### GetLastSignInAt

`func (o *AuthUser) GetLastSignInAt() time.Time`

GetLastSignInAt returns the LastSignInAt field if non-nil, zero value otherwise.

### GetLastSignInAtOk

`func (o *AuthUser) GetLastSignInAtOk() (*time.Time, bool)`

GetLastSignInAtOk returns a tuple with the LastSignInAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSignInAt

`func (o *AuthUser) SetLastSignInAt(v time.Time)`

SetLastSignInAt sets LastSignInAt field to given value.

### HasLastSignInAt

`func (o *AuthUser) HasLastSignInAt() bool`

HasLastSignInAt returns a boolean if a field has been set.

### SetLastSignInAtNil

`func (o *AuthUser) SetLastSignInAtNil(b bool)`

 SetLastSignInAtNil sets the value for LastSignInAt to be an explicit nil

### UnsetLastSignInAt
`func (o *AuthUser) UnsetLastSignInAt()`

UnsetLastSignInAt ensures that no value is present for LastSignInAt, not even an explicit nil
### GetRole

`func (o *AuthUser) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AuthUser) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AuthUser) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AuthUser) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuthUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuthUser) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuthUser) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuthUser) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuthUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAppMetadata

`func (o *AuthUser) GetAppMetadata() AuthUserAppMetadata`

GetAppMetadata returns the AppMetadata field if non-nil, zero value otherwise.

### GetAppMetadataOk

`func (o *AuthUser) GetAppMetadataOk() (*AuthUserAppMetadata, bool)`

GetAppMetadataOk returns a tuple with the AppMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppMetadata

`func (o *AuthUser) SetAppMetadata(v AuthUserAppMetadata)`

SetAppMetadata sets AppMetadata field to given value.

### HasAppMetadata

`func (o *AuthUser) HasAppMetadata() bool`

HasAppMetadata returns a boolean if a field has been set.

### GetUserMetadata

`func (o *AuthUser) GetUserMetadata() map[string]interface{}`

GetUserMetadata returns the UserMetadata field if non-nil, zero value otherwise.

### GetUserMetadataOk

`func (o *AuthUser) GetUserMetadataOk() (*map[string]interface{}, bool)`

GetUserMetadataOk returns a tuple with the UserMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMetadata

`func (o *AuthUser) SetUserMetadata(v map[string]interface{})`

SetUserMetadata sets UserMetadata field to given value.

### HasUserMetadata

`func (o *AuthUser) HasUserMetadata() bool`

HasUserMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


