# AuthUserIdentitiesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identity ID | [optional] 
**UserId** | Pointer to **string** | User ID | [optional] 
**IdentityData** | Pointer to **map[string]interface{}** | Identity provider data | [optional] 
**Provider** | Pointer to **string** | Identity provider name (email, google, etc.) | [optional] 
**LastSignInAt** | Pointer to **NullableTime** | Last sign in timestamp for this identity | [optional] 
**CreatedAt** | Pointer to **time.Time** | Identity creation timestamp | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Identity last update timestamp | [optional] 

## Methods

### NewAuthUserIdentitiesInner

`func NewAuthUserIdentitiesInner() *AuthUserIdentitiesInner`

NewAuthUserIdentitiesInner instantiates a new AuthUserIdentitiesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthUserIdentitiesInnerWithDefaults

`func NewAuthUserIdentitiesInnerWithDefaults() *AuthUserIdentitiesInner`

NewAuthUserIdentitiesInnerWithDefaults instantiates a new AuthUserIdentitiesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthUserIdentitiesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthUserIdentitiesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthUserIdentitiesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthUserIdentitiesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *AuthUserIdentitiesInner) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuthUserIdentitiesInner) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuthUserIdentitiesInner) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AuthUserIdentitiesInner) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetIdentityData

`func (o *AuthUserIdentitiesInner) GetIdentityData() map[string]interface{}`

GetIdentityData returns the IdentityData field if non-nil, zero value otherwise.

### GetIdentityDataOk

`func (o *AuthUserIdentitiesInner) GetIdentityDataOk() (*map[string]interface{}, bool)`

GetIdentityDataOk returns a tuple with the IdentityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityData

`func (o *AuthUserIdentitiesInner) SetIdentityData(v map[string]interface{})`

SetIdentityData sets IdentityData field to given value.

### HasIdentityData

`func (o *AuthUserIdentitiesInner) HasIdentityData() bool`

HasIdentityData returns a boolean if a field has been set.

### GetProvider

`func (o *AuthUserIdentitiesInner) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AuthUserIdentitiesInner) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AuthUserIdentitiesInner) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AuthUserIdentitiesInner) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetLastSignInAt

`func (o *AuthUserIdentitiesInner) GetLastSignInAt() time.Time`

GetLastSignInAt returns the LastSignInAt field if non-nil, zero value otherwise.

### GetLastSignInAtOk

`func (o *AuthUserIdentitiesInner) GetLastSignInAtOk() (*time.Time, bool)`

GetLastSignInAtOk returns a tuple with the LastSignInAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSignInAt

`func (o *AuthUserIdentitiesInner) SetLastSignInAt(v time.Time)`

SetLastSignInAt sets LastSignInAt field to given value.

### HasLastSignInAt

`func (o *AuthUserIdentitiesInner) HasLastSignInAt() bool`

HasLastSignInAt returns a boolean if a field has been set.

### SetLastSignInAtNil

`func (o *AuthUserIdentitiesInner) SetLastSignInAtNil(b bool)`

 SetLastSignInAtNil sets the value for LastSignInAt to be an explicit nil

### UnsetLastSignInAt
`func (o *AuthUserIdentitiesInner) UnsetLastSignInAt()`

UnsetLastSignInAt ensures that no value is present for LastSignInAt, not even an explicit nil
### GetCreatedAt

`func (o *AuthUserIdentitiesInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthUserIdentitiesInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthUserIdentitiesInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthUserIdentitiesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuthUserIdentitiesInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuthUserIdentitiesInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuthUserIdentitiesInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuthUserIdentitiesInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


