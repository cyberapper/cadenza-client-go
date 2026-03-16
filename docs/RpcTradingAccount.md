# RpcTradingAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradingAccountId** | Pointer to **string** | Unique trading account ID | [optional] 
**UserId** | Pointer to **string** | User ID | [optional] 
**TenantId** | Pointer to **string** | Tenant identifier for multi-tenancy | [optional] 
**Nickname** | Pointer to **string** | Account nickname | [optional] 
**ExternalAccountId** | Pointer to **string** | External account ID at venue | [optional] 
**Venue** | Pointer to [**Venue**](Venue.md) |  | [optional] 
**Status** | Pointer to [**TradingAccountStatus**](TradingAccountStatus.md) |  | [optional] 
**AccountType** | Pointer to [**TradingAccountType**](TradingAccountType.md) |  | [optional] 
**ExternalAccountType** | Pointer to **string** | Type of account on the exchange (set by market connector) | [optional] 
**PositionMode** | Pointer to [**PositionMode**](PositionMode.md) |  | [optional] 
**CollateralMode** | Pointer to [**CollateralMode**](CollateralMode.md) |  | [optional] 
**MarginMode** | Pointer to [**MarginMode**](MarginMode.md) |  | [optional] 
**Credentials** | Pointer to [**[]RpcTradingAccountCredential**](RpcTradingAccountCredential.md) | Account credentials | [optional] 
**Config** | Pointer to [**RpcTradingAccountConfig**](RpcTradingAccountConfig.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRpcTradingAccount

`func NewRpcTradingAccount() *RpcTradingAccount`

NewRpcTradingAccount instantiates a new RpcTradingAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcTradingAccountWithDefaults

`func NewRpcTradingAccountWithDefaults() *RpcTradingAccount`

NewRpcTradingAccountWithDefaults instantiates a new RpcTradingAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradingAccountId

`func (o *RpcTradingAccount) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *RpcTradingAccount) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *RpcTradingAccount) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.

### HasTradingAccountId

`func (o *RpcTradingAccount) HasTradingAccountId() bool`

HasTradingAccountId returns a boolean if a field has been set.

### GetUserId

`func (o *RpcTradingAccount) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RpcTradingAccount) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RpcTradingAccount) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *RpcTradingAccount) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetTenantId

`func (o *RpcTradingAccount) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *RpcTradingAccount) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *RpcTradingAccount) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *RpcTradingAccount) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetNickname

`func (o *RpcTradingAccount) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *RpcTradingAccount) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *RpcTradingAccount) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *RpcTradingAccount) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetExternalAccountId

`func (o *RpcTradingAccount) GetExternalAccountId() string`

GetExternalAccountId returns the ExternalAccountId field if non-nil, zero value otherwise.

### GetExternalAccountIdOk

`func (o *RpcTradingAccount) GetExternalAccountIdOk() (*string, bool)`

GetExternalAccountIdOk returns a tuple with the ExternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAccountId

`func (o *RpcTradingAccount) SetExternalAccountId(v string)`

SetExternalAccountId sets ExternalAccountId field to given value.

### HasExternalAccountId

`func (o *RpcTradingAccount) HasExternalAccountId() bool`

HasExternalAccountId returns a boolean if a field has been set.

### GetVenue

`func (o *RpcTradingAccount) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *RpcTradingAccount) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *RpcTradingAccount) SetVenue(v Venue)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *RpcTradingAccount) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetStatus

`func (o *RpcTradingAccount) GetStatus() TradingAccountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RpcTradingAccount) GetStatusOk() (*TradingAccountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RpcTradingAccount) SetStatus(v TradingAccountStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RpcTradingAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAccountType

`func (o *RpcTradingAccount) GetAccountType() TradingAccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *RpcTradingAccount) GetAccountTypeOk() (*TradingAccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *RpcTradingAccount) SetAccountType(v TradingAccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *RpcTradingAccount) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetExternalAccountType

`func (o *RpcTradingAccount) GetExternalAccountType() string`

GetExternalAccountType returns the ExternalAccountType field if non-nil, zero value otherwise.

### GetExternalAccountTypeOk

`func (o *RpcTradingAccount) GetExternalAccountTypeOk() (*string, bool)`

GetExternalAccountTypeOk returns a tuple with the ExternalAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAccountType

`func (o *RpcTradingAccount) SetExternalAccountType(v string)`

SetExternalAccountType sets ExternalAccountType field to given value.

### HasExternalAccountType

`func (o *RpcTradingAccount) HasExternalAccountType() bool`

HasExternalAccountType returns a boolean if a field has been set.

### GetPositionMode

`func (o *RpcTradingAccount) GetPositionMode() PositionMode`

GetPositionMode returns the PositionMode field if non-nil, zero value otherwise.

### GetPositionModeOk

`func (o *RpcTradingAccount) GetPositionModeOk() (*PositionMode, bool)`

GetPositionModeOk returns a tuple with the PositionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionMode

`func (o *RpcTradingAccount) SetPositionMode(v PositionMode)`

SetPositionMode sets PositionMode field to given value.

### HasPositionMode

`func (o *RpcTradingAccount) HasPositionMode() bool`

HasPositionMode returns a boolean if a field has been set.

### GetCollateralMode

`func (o *RpcTradingAccount) GetCollateralMode() CollateralMode`

GetCollateralMode returns the CollateralMode field if non-nil, zero value otherwise.

### GetCollateralModeOk

`func (o *RpcTradingAccount) GetCollateralModeOk() (*CollateralMode, bool)`

GetCollateralModeOk returns a tuple with the CollateralMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralMode

`func (o *RpcTradingAccount) SetCollateralMode(v CollateralMode)`

SetCollateralMode sets CollateralMode field to given value.

### HasCollateralMode

`func (o *RpcTradingAccount) HasCollateralMode() bool`

HasCollateralMode returns a boolean if a field has been set.

### GetMarginMode

`func (o *RpcTradingAccount) GetMarginMode() MarginMode`

GetMarginMode returns the MarginMode field if non-nil, zero value otherwise.

### GetMarginModeOk

`func (o *RpcTradingAccount) GetMarginModeOk() (*MarginMode, bool)`

GetMarginModeOk returns a tuple with the MarginMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginMode

`func (o *RpcTradingAccount) SetMarginMode(v MarginMode)`

SetMarginMode sets MarginMode field to given value.

### HasMarginMode

`func (o *RpcTradingAccount) HasMarginMode() bool`

HasMarginMode returns a boolean if a field has been set.

### GetCredentials

`func (o *RpcTradingAccount) GetCredentials() []RpcTradingAccountCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *RpcTradingAccount) GetCredentialsOk() (*[]RpcTradingAccountCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *RpcTradingAccount) SetCredentials(v []RpcTradingAccountCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *RpcTradingAccount) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetConfig

`func (o *RpcTradingAccount) GetConfig() RpcTradingAccountConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *RpcTradingAccount) GetConfigOk() (*RpcTradingAccountConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *RpcTradingAccount) SetConfig(v RpcTradingAccountConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *RpcTradingAccount) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RpcTradingAccount) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RpcTradingAccount) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RpcTradingAccount) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RpcTradingAccount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RpcTradingAccount) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RpcTradingAccount) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RpcTradingAccount) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RpcTradingAccount) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


