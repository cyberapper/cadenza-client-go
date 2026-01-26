# TradingAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradingAccountId** | **string** | UUID string | 
**ExternalTradingAccountId** | **string** | External trading account ID | 
**Venue** | [**Venue**](Venue.md) |  | 
**Nickname** | **string** | Nickname of the trading account | 
**AccountType** | [**AccountType**](AccountType.md) |  | 
**ExternalAccountType** | **string** | Type of account on the exchange (set by market connector) | 
**PositionMode** | Pointer to [**PositionMode**](PositionMode.md) |  | [optional] 
**CollateralMode** | Pointer to [**CollateralMode**](CollateralMode.md) |  | [optional] 
**MarginMode** | Pointer to [**MarginMode**](MarginMode.md) |  | [optional] 
**Credentials** | [**[]TradingAccountCredential**](TradingAccountCredential.md) |  | 
**Status** | [**TradingAccountStatus**](TradingAccountStatus.md) |  | 
**CreatedAt** | **int64** | Unix timestamp in milliseconds | 
**CreatedAtDateTime** | Pointer to **time.Time** | Creation timestamp in ISO 8601 format | [optional] 
**UpdatedAt** | **int64** | Unix timestamp in milliseconds | 
**UpdatedAtDateTime** | Pointer to **time.Time** | Last update timestamp in ISO 8601 format | [optional] 

## Methods

### NewTradingAccount

`func NewTradingAccount(tradingAccountId string, externalTradingAccountId string, venue Venue, nickname string, accountType AccountType, externalAccountType string, credentials []TradingAccountCredential, status TradingAccountStatus, createdAt int64, updatedAt int64, ) *TradingAccount`

NewTradingAccount instantiates a new TradingAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradingAccountWithDefaults

`func NewTradingAccountWithDefaults() *TradingAccount`

NewTradingAccountWithDefaults instantiates a new TradingAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradingAccountId

`func (o *TradingAccount) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *TradingAccount) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *TradingAccount) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.


### GetExternalTradingAccountId

`func (o *TradingAccount) GetExternalTradingAccountId() string`

GetExternalTradingAccountId returns the ExternalTradingAccountId field if non-nil, zero value otherwise.

### GetExternalTradingAccountIdOk

`func (o *TradingAccount) GetExternalTradingAccountIdOk() (*string, bool)`

GetExternalTradingAccountIdOk returns a tuple with the ExternalTradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTradingAccountId

`func (o *TradingAccount) SetExternalTradingAccountId(v string)`

SetExternalTradingAccountId sets ExternalTradingAccountId field to given value.


### GetVenue

`func (o *TradingAccount) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *TradingAccount) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *TradingAccount) SetVenue(v Venue)`

SetVenue sets Venue field to given value.


### GetNickname

`func (o *TradingAccount) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *TradingAccount) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *TradingAccount) SetNickname(v string)`

SetNickname sets Nickname field to given value.


### GetAccountType

`func (o *TradingAccount) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *TradingAccount) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *TradingAccount) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetExternalAccountType

`func (o *TradingAccount) GetExternalAccountType() string`

GetExternalAccountType returns the ExternalAccountType field if non-nil, zero value otherwise.

### GetExternalAccountTypeOk

`func (o *TradingAccount) GetExternalAccountTypeOk() (*string, bool)`

GetExternalAccountTypeOk returns a tuple with the ExternalAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAccountType

`func (o *TradingAccount) SetExternalAccountType(v string)`

SetExternalAccountType sets ExternalAccountType field to given value.


### GetPositionMode

`func (o *TradingAccount) GetPositionMode() PositionMode`

GetPositionMode returns the PositionMode field if non-nil, zero value otherwise.

### GetPositionModeOk

`func (o *TradingAccount) GetPositionModeOk() (*PositionMode, bool)`

GetPositionModeOk returns a tuple with the PositionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionMode

`func (o *TradingAccount) SetPositionMode(v PositionMode)`

SetPositionMode sets PositionMode field to given value.

### HasPositionMode

`func (o *TradingAccount) HasPositionMode() bool`

HasPositionMode returns a boolean if a field has been set.

### GetCollateralMode

`func (o *TradingAccount) GetCollateralMode() CollateralMode`

GetCollateralMode returns the CollateralMode field if non-nil, zero value otherwise.

### GetCollateralModeOk

`func (o *TradingAccount) GetCollateralModeOk() (*CollateralMode, bool)`

GetCollateralModeOk returns a tuple with the CollateralMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralMode

`func (o *TradingAccount) SetCollateralMode(v CollateralMode)`

SetCollateralMode sets CollateralMode field to given value.

### HasCollateralMode

`func (o *TradingAccount) HasCollateralMode() bool`

HasCollateralMode returns a boolean if a field has been set.

### GetMarginMode

`func (o *TradingAccount) GetMarginMode() MarginMode`

GetMarginMode returns the MarginMode field if non-nil, zero value otherwise.

### GetMarginModeOk

`func (o *TradingAccount) GetMarginModeOk() (*MarginMode, bool)`

GetMarginModeOk returns a tuple with the MarginMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginMode

`func (o *TradingAccount) SetMarginMode(v MarginMode)`

SetMarginMode sets MarginMode field to given value.

### HasMarginMode

`func (o *TradingAccount) HasMarginMode() bool`

HasMarginMode returns a boolean if a field has been set.

### GetCredentials

`func (o *TradingAccount) GetCredentials() []TradingAccountCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *TradingAccount) GetCredentialsOk() (*[]TradingAccountCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *TradingAccount) SetCredentials(v []TradingAccountCredential)`

SetCredentials sets Credentials field to given value.


### GetStatus

`func (o *TradingAccount) GetStatus() TradingAccountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TradingAccount) GetStatusOk() (*TradingAccountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TradingAccount) SetStatus(v TradingAccountStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *TradingAccount) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TradingAccount) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TradingAccount) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedAtDateTime

`func (o *TradingAccount) GetCreatedAtDateTime() time.Time`

GetCreatedAtDateTime returns the CreatedAtDateTime field if non-nil, zero value otherwise.

### GetCreatedAtDateTimeOk

`func (o *TradingAccount) GetCreatedAtDateTimeOk() (*time.Time, bool)`

GetCreatedAtDateTimeOk returns a tuple with the CreatedAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtDateTime

`func (o *TradingAccount) SetCreatedAtDateTime(v time.Time)`

SetCreatedAtDateTime sets CreatedAtDateTime field to given value.

### HasCreatedAtDateTime

`func (o *TradingAccount) HasCreatedAtDateTime() bool`

HasCreatedAtDateTime returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TradingAccount) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TradingAccount) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TradingAccount) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpdatedAtDateTime

`func (o *TradingAccount) GetUpdatedAtDateTime() time.Time`

GetUpdatedAtDateTime returns the UpdatedAtDateTime field if non-nil, zero value otherwise.

### GetUpdatedAtDateTimeOk

`func (o *TradingAccount) GetUpdatedAtDateTimeOk() (*time.Time, bool)`

GetUpdatedAtDateTimeOk returns a tuple with the UpdatedAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtDateTime

`func (o *TradingAccount) SetUpdatedAtDateTime(v time.Time)`

SetUpdatedAtDateTime sets UpdatedAtDateTime field to given value.

### HasUpdatedAtDateTime

`func (o *TradingAccount) HasUpdatedAtDateTime() bool`

HasUpdatedAtDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


