# FermataWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | UUID string | 
**WalletType** | [**WalletType**](WalletType.md) |  | 
**Status** | [**WalletStatus**](WalletStatus.md) |  | 
**AllowNegative** | Pointer to **bool** | Whether this wallet allows negative balances (e.g., dealer short positions) | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional attributes (name, created_by, etc.) | [optional] 
**CreatedAt** | **int64** | Unix timestamp in milliseconds | 
**CreatedAtDateTime** | Pointer to **time.Time** | Wallet creation timestamp in ISO 8601 format | [optional] 
**UpdatedAt** | **int64** | Unix timestamp in milliseconds | 
**UpdatedAtDateTime** | Pointer to **time.Time** | Last update timestamp in ISO 8601 format | [optional] 

## Methods

### NewFermataWallet

`func NewFermataWallet(walletId string, walletType WalletType, status WalletStatus, createdAt int64, updatedAt int64, ) *FermataWallet`

NewFermataWallet instantiates a new FermataWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFermataWalletWithDefaults

`func NewFermataWalletWithDefaults() *FermataWallet`

NewFermataWalletWithDefaults instantiates a new FermataWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *FermataWallet) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *FermataWallet) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *FermataWallet) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetWalletType

`func (o *FermataWallet) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *FermataWallet) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *FermataWallet) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.


### GetStatus

`func (o *FermataWallet) GetStatus() WalletStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FermataWallet) GetStatusOk() (*WalletStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FermataWallet) SetStatus(v WalletStatus)`

SetStatus sets Status field to given value.


### GetAllowNegative

`func (o *FermataWallet) GetAllowNegative() bool`

GetAllowNegative returns the AllowNegative field if non-nil, zero value otherwise.

### GetAllowNegativeOk

`func (o *FermataWallet) GetAllowNegativeOk() (*bool, bool)`

GetAllowNegativeOk returns a tuple with the AllowNegative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNegative

`func (o *FermataWallet) SetAllowNegative(v bool)`

SetAllowNegative sets AllowNegative field to given value.

### HasAllowNegative

`func (o *FermataWallet) HasAllowNegative() bool`

HasAllowNegative returns a boolean if a field has been set.

### GetMetadata

`func (o *FermataWallet) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FermataWallet) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FermataWallet) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FermataWallet) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FermataWallet) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FermataWallet) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FermataWallet) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedAtDateTime

`func (o *FermataWallet) GetCreatedAtDateTime() time.Time`

GetCreatedAtDateTime returns the CreatedAtDateTime field if non-nil, zero value otherwise.

### GetCreatedAtDateTimeOk

`func (o *FermataWallet) GetCreatedAtDateTimeOk() (*time.Time, bool)`

GetCreatedAtDateTimeOk returns a tuple with the CreatedAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtDateTime

`func (o *FermataWallet) SetCreatedAtDateTime(v time.Time)`

SetCreatedAtDateTime sets CreatedAtDateTime field to given value.

### HasCreatedAtDateTime

`func (o *FermataWallet) HasCreatedAtDateTime() bool`

HasCreatedAtDateTime returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FermataWallet) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FermataWallet) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FermataWallet) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpdatedAtDateTime

`func (o *FermataWallet) GetUpdatedAtDateTime() time.Time`

GetUpdatedAtDateTime returns the UpdatedAtDateTime field if non-nil, zero value otherwise.

### GetUpdatedAtDateTimeOk

`func (o *FermataWallet) GetUpdatedAtDateTimeOk() (*time.Time, bool)`

GetUpdatedAtDateTimeOk returns a tuple with the UpdatedAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtDateTime

`func (o *FermataWallet) SetUpdatedAtDateTime(v time.Time)`

SetUpdatedAtDateTime sets UpdatedAtDateTime field to given value.

### HasUpdatedAtDateTime

`func (o *FermataWallet) HasUpdatedAtDateTime() bool`

HasUpdatedAtDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


