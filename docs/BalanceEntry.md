# BalanceEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecuritySymbol** | **string** | Security symbol | 
**SecurityType** | [**SecurityType**](SecurityType.md) |  | 
**ExternalBalanceId** | Pointer to **NullableString** | External balance ID from the exchange | [optional] 
**TradingAccountId** | **string** | UUID string | 
**Status** | [**BalanceStatus**](BalanceStatus.md) |  | 
**PositionId** | Pointer to **string** | UUID string | [optional] 
**Free** | **string** | Decimal value as string to preserve precision | 
**Locked** | **string** | Decimal value as string to preserve precision | 
**Borrowed** | **string** | Decimal value as string to preserve precision | 
**Total** | **string** | Decimal value as string to preserve precision | 
**Net** | **string** | Decimal value as string to preserve precision | 
**UpdatedAt** | **int64** | Unix timestamp in milliseconds | 
**UpdatedAtDateTime** | Pointer to **time.Time** | Last update timestamp in ISO 8601 format | [optional] 

## Methods

### NewBalanceEntry

`func NewBalanceEntry(securitySymbol string, securityType SecurityType, tradingAccountId string, status BalanceStatus, free string, locked string, borrowed string, total string, net string, updatedAt int64, ) *BalanceEntry`

NewBalanceEntry instantiates a new BalanceEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceEntryWithDefaults

`func NewBalanceEntryWithDefaults() *BalanceEntry`

NewBalanceEntryWithDefaults instantiates a new BalanceEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecuritySymbol

`func (o *BalanceEntry) GetSecuritySymbol() string`

GetSecuritySymbol returns the SecuritySymbol field if non-nil, zero value otherwise.

### GetSecuritySymbolOk

`func (o *BalanceEntry) GetSecuritySymbolOk() (*string, bool)`

GetSecuritySymbolOk returns a tuple with the SecuritySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritySymbol

`func (o *BalanceEntry) SetSecuritySymbol(v string)`

SetSecuritySymbol sets SecuritySymbol field to given value.


### GetSecurityType

`func (o *BalanceEntry) GetSecurityType() SecurityType`

GetSecurityType returns the SecurityType field if non-nil, zero value otherwise.

### GetSecurityTypeOk

`func (o *BalanceEntry) GetSecurityTypeOk() (*SecurityType, bool)`

GetSecurityTypeOk returns a tuple with the SecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityType

`func (o *BalanceEntry) SetSecurityType(v SecurityType)`

SetSecurityType sets SecurityType field to given value.


### GetExternalBalanceId

`func (o *BalanceEntry) GetExternalBalanceId() string`

GetExternalBalanceId returns the ExternalBalanceId field if non-nil, zero value otherwise.

### GetExternalBalanceIdOk

`func (o *BalanceEntry) GetExternalBalanceIdOk() (*string, bool)`

GetExternalBalanceIdOk returns a tuple with the ExternalBalanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalBalanceId

`func (o *BalanceEntry) SetExternalBalanceId(v string)`

SetExternalBalanceId sets ExternalBalanceId field to given value.

### HasExternalBalanceId

`func (o *BalanceEntry) HasExternalBalanceId() bool`

HasExternalBalanceId returns a boolean if a field has been set.

### SetExternalBalanceIdNil

`func (o *BalanceEntry) SetExternalBalanceIdNil(b bool)`

 SetExternalBalanceIdNil sets the value for ExternalBalanceId to be an explicit nil

### UnsetExternalBalanceId
`func (o *BalanceEntry) UnsetExternalBalanceId()`

UnsetExternalBalanceId ensures that no value is present for ExternalBalanceId, not even an explicit nil
### GetTradingAccountId

`func (o *BalanceEntry) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *BalanceEntry) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *BalanceEntry) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.


### GetStatus

`func (o *BalanceEntry) GetStatus() BalanceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BalanceEntry) GetStatusOk() (*BalanceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BalanceEntry) SetStatus(v BalanceStatus)`

SetStatus sets Status field to given value.


### GetPositionId

`func (o *BalanceEntry) GetPositionId() string`

GetPositionId returns the PositionId field if non-nil, zero value otherwise.

### GetPositionIdOk

`func (o *BalanceEntry) GetPositionIdOk() (*string, bool)`

GetPositionIdOk returns a tuple with the PositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionId

`func (o *BalanceEntry) SetPositionId(v string)`

SetPositionId sets PositionId field to given value.

### HasPositionId

`func (o *BalanceEntry) HasPositionId() bool`

HasPositionId returns a boolean if a field has been set.

### GetFree

`func (o *BalanceEntry) GetFree() string`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *BalanceEntry) GetFreeOk() (*string, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *BalanceEntry) SetFree(v string)`

SetFree sets Free field to given value.


### GetLocked

`func (o *BalanceEntry) GetLocked() string`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *BalanceEntry) GetLockedOk() (*string, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *BalanceEntry) SetLocked(v string)`

SetLocked sets Locked field to given value.


### GetBorrowed

`func (o *BalanceEntry) GetBorrowed() string`

GetBorrowed returns the Borrowed field if non-nil, zero value otherwise.

### GetBorrowedOk

`func (o *BalanceEntry) GetBorrowedOk() (*string, bool)`

GetBorrowedOk returns a tuple with the Borrowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowed

`func (o *BalanceEntry) SetBorrowed(v string)`

SetBorrowed sets Borrowed field to given value.


### GetTotal

`func (o *BalanceEntry) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BalanceEntry) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BalanceEntry) SetTotal(v string)`

SetTotal sets Total field to given value.


### GetNet

`func (o *BalanceEntry) GetNet() string`

GetNet returns the Net field if non-nil, zero value otherwise.

### GetNetOk

`func (o *BalanceEntry) GetNetOk() (*string, bool)`

GetNetOk returns a tuple with the Net field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet

`func (o *BalanceEntry) SetNet(v string)`

SetNet sets Net field to given value.


### GetUpdatedAt

`func (o *BalanceEntry) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BalanceEntry) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BalanceEntry) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpdatedAtDateTime

`func (o *BalanceEntry) GetUpdatedAtDateTime() time.Time`

GetUpdatedAtDateTime returns the UpdatedAtDateTime field if non-nil, zero value otherwise.

### GetUpdatedAtDateTimeOk

`func (o *BalanceEntry) GetUpdatedAtDateTimeOk() (*time.Time, bool)`

GetUpdatedAtDateTimeOk returns a tuple with the UpdatedAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtDateTime

`func (o *BalanceEntry) SetUpdatedAtDateTime(v time.Time)`

SetUpdatedAtDateTime sets UpdatedAtDateTime field to given value.

### HasUpdatedAtDateTime

`func (o *BalanceEntry) HasUpdatedAtDateTime() bool`

HasUpdatedAtDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


