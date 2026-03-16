# PositionEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PositionId** | **string** | UUID string | 
**SecuritySymbol** | **string** | Security symbol | 
**ExternalPositionId** | Pointer to **NullableString** | External position ID from the exchange | [optional] 
**TradingAccountId** | **string** | UUID string | 
**InstrumentId** | Pointer to **string** | Instrument ID in format {VENUE}:{BASE}/{QUOTE} | [optional] 
**SecurityType** | [**SecurityType**](SecurityType.md) |  | 
**Status** | [**PositionStatus**](PositionStatus.md) |  | 
**Quantity** | **string** | Decimal value as string to preserve precision | 
**EntryPrice** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**ExitPrice** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**CurrentPrice** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**UnrealizedPnl** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**RealizedPnl** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**CreatedAt** | **int64** | Unix timestamp in milliseconds | 
**CreatedAtDateTime** | Pointer to **time.Time** | Creation timestamp in ISO 8601 format | [optional] 
**UpdatedAt** | **int64** | Unix timestamp in milliseconds | 
**UpdatedAtDateTime** | Pointer to **time.Time** | Last update timestamp in ISO 8601 format | [optional] 
**ClosedAt** | Pointer to **int64** | Unix timestamp in milliseconds | [optional] 
**ClosedAtDateTime** | Pointer to **NullableTime** | Position closure timestamp in ISO 8601 format | [optional] 

## Methods

### NewPositionEntry

`func NewPositionEntry(positionId string, securitySymbol string, tradingAccountId string, securityType SecurityType, status PositionStatus, quantity string, createdAt int64, updatedAt int64, ) *PositionEntry`

NewPositionEntry instantiates a new PositionEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositionEntryWithDefaults

`func NewPositionEntryWithDefaults() *PositionEntry`

NewPositionEntryWithDefaults instantiates a new PositionEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPositionId

`func (o *PositionEntry) GetPositionId() string`

GetPositionId returns the PositionId field if non-nil, zero value otherwise.

### GetPositionIdOk

`func (o *PositionEntry) GetPositionIdOk() (*string, bool)`

GetPositionIdOk returns a tuple with the PositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionId

`func (o *PositionEntry) SetPositionId(v string)`

SetPositionId sets PositionId field to given value.


### GetSecuritySymbol

`func (o *PositionEntry) GetSecuritySymbol() string`

GetSecuritySymbol returns the SecuritySymbol field if non-nil, zero value otherwise.

### GetSecuritySymbolOk

`func (o *PositionEntry) GetSecuritySymbolOk() (*string, bool)`

GetSecuritySymbolOk returns a tuple with the SecuritySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritySymbol

`func (o *PositionEntry) SetSecuritySymbol(v string)`

SetSecuritySymbol sets SecuritySymbol field to given value.


### GetExternalPositionId

`func (o *PositionEntry) GetExternalPositionId() string`

GetExternalPositionId returns the ExternalPositionId field if non-nil, zero value otherwise.

### GetExternalPositionIdOk

`func (o *PositionEntry) GetExternalPositionIdOk() (*string, bool)`

GetExternalPositionIdOk returns a tuple with the ExternalPositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPositionId

`func (o *PositionEntry) SetExternalPositionId(v string)`

SetExternalPositionId sets ExternalPositionId field to given value.

### HasExternalPositionId

`func (o *PositionEntry) HasExternalPositionId() bool`

HasExternalPositionId returns a boolean if a field has been set.

### SetExternalPositionIdNil

`func (o *PositionEntry) SetExternalPositionIdNil(b bool)`

 SetExternalPositionIdNil sets the value for ExternalPositionId to be an explicit nil

### UnsetExternalPositionId
`func (o *PositionEntry) UnsetExternalPositionId()`

UnsetExternalPositionId ensures that no value is present for ExternalPositionId, not even an explicit nil
### GetTradingAccountId

`func (o *PositionEntry) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *PositionEntry) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *PositionEntry) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.


### GetInstrumentId

`func (o *PositionEntry) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *PositionEntry) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *PositionEntry) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *PositionEntry) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetSecurityType

`func (o *PositionEntry) GetSecurityType() SecurityType`

GetSecurityType returns the SecurityType field if non-nil, zero value otherwise.

### GetSecurityTypeOk

`func (o *PositionEntry) GetSecurityTypeOk() (*SecurityType, bool)`

GetSecurityTypeOk returns a tuple with the SecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityType

`func (o *PositionEntry) SetSecurityType(v SecurityType)`

SetSecurityType sets SecurityType field to given value.


### GetStatus

`func (o *PositionEntry) GetStatus() PositionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PositionEntry) GetStatusOk() (*PositionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PositionEntry) SetStatus(v PositionStatus)`

SetStatus sets Status field to given value.


### GetQuantity

`func (o *PositionEntry) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PositionEntry) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PositionEntry) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetEntryPrice

`func (o *PositionEntry) GetEntryPrice() string`

GetEntryPrice returns the EntryPrice field if non-nil, zero value otherwise.

### GetEntryPriceOk

`func (o *PositionEntry) GetEntryPriceOk() (*string, bool)`

GetEntryPriceOk returns a tuple with the EntryPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPrice

`func (o *PositionEntry) SetEntryPrice(v string)`

SetEntryPrice sets EntryPrice field to given value.

### HasEntryPrice

`func (o *PositionEntry) HasEntryPrice() bool`

HasEntryPrice returns a boolean if a field has been set.

### GetExitPrice

`func (o *PositionEntry) GetExitPrice() string`

GetExitPrice returns the ExitPrice field if non-nil, zero value otherwise.

### GetExitPriceOk

`func (o *PositionEntry) GetExitPriceOk() (*string, bool)`

GetExitPriceOk returns a tuple with the ExitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitPrice

`func (o *PositionEntry) SetExitPrice(v string)`

SetExitPrice sets ExitPrice field to given value.

### HasExitPrice

`func (o *PositionEntry) HasExitPrice() bool`

HasExitPrice returns a boolean if a field has been set.

### GetCurrentPrice

`func (o *PositionEntry) GetCurrentPrice() string`

GetCurrentPrice returns the CurrentPrice field if non-nil, zero value otherwise.

### GetCurrentPriceOk

`func (o *PositionEntry) GetCurrentPriceOk() (*string, bool)`

GetCurrentPriceOk returns a tuple with the CurrentPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPrice

`func (o *PositionEntry) SetCurrentPrice(v string)`

SetCurrentPrice sets CurrentPrice field to given value.

### HasCurrentPrice

`func (o *PositionEntry) HasCurrentPrice() bool`

HasCurrentPrice returns a boolean if a field has been set.

### GetUnrealizedPnl

`func (o *PositionEntry) GetUnrealizedPnl() string`

GetUnrealizedPnl returns the UnrealizedPnl field if non-nil, zero value otherwise.

### GetUnrealizedPnlOk

`func (o *PositionEntry) GetUnrealizedPnlOk() (*string, bool)`

GetUnrealizedPnlOk returns a tuple with the UnrealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedPnl

`func (o *PositionEntry) SetUnrealizedPnl(v string)`

SetUnrealizedPnl sets UnrealizedPnl field to given value.

### HasUnrealizedPnl

`func (o *PositionEntry) HasUnrealizedPnl() bool`

HasUnrealizedPnl returns a boolean if a field has been set.

### GetRealizedPnl

`func (o *PositionEntry) GetRealizedPnl() string`

GetRealizedPnl returns the RealizedPnl field if non-nil, zero value otherwise.

### GetRealizedPnlOk

`func (o *PositionEntry) GetRealizedPnlOk() (*string, bool)`

GetRealizedPnlOk returns a tuple with the RealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedPnl

`func (o *PositionEntry) SetRealizedPnl(v string)`

SetRealizedPnl sets RealizedPnl field to given value.

### HasRealizedPnl

`func (o *PositionEntry) HasRealizedPnl() bool`

HasRealizedPnl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PositionEntry) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PositionEntry) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PositionEntry) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedAtDateTime

`func (o *PositionEntry) GetCreatedAtDateTime() time.Time`

GetCreatedAtDateTime returns the CreatedAtDateTime field if non-nil, zero value otherwise.

### GetCreatedAtDateTimeOk

`func (o *PositionEntry) GetCreatedAtDateTimeOk() (*time.Time, bool)`

GetCreatedAtDateTimeOk returns a tuple with the CreatedAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtDateTime

`func (o *PositionEntry) SetCreatedAtDateTime(v time.Time)`

SetCreatedAtDateTime sets CreatedAtDateTime field to given value.

### HasCreatedAtDateTime

`func (o *PositionEntry) HasCreatedAtDateTime() bool`

HasCreatedAtDateTime returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PositionEntry) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PositionEntry) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PositionEntry) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpdatedAtDateTime

`func (o *PositionEntry) GetUpdatedAtDateTime() time.Time`

GetUpdatedAtDateTime returns the UpdatedAtDateTime field if non-nil, zero value otherwise.

### GetUpdatedAtDateTimeOk

`func (o *PositionEntry) GetUpdatedAtDateTimeOk() (*time.Time, bool)`

GetUpdatedAtDateTimeOk returns a tuple with the UpdatedAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtDateTime

`func (o *PositionEntry) SetUpdatedAtDateTime(v time.Time)`

SetUpdatedAtDateTime sets UpdatedAtDateTime field to given value.

### HasUpdatedAtDateTime

`func (o *PositionEntry) HasUpdatedAtDateTime() bool`

HasUpdatedAtDateTime returns a boolean if a field has been set.

### GetClosedAt

`func (o *PositionEntry) GetClosedAt() int64`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *PositionEntry) GetClosedAtOk() (*int64, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *PositionEntry) SetClosedAt(v int64)`

SetClosedAt sets ClosedAt field to given value.

### HasClosedAt

`func (o *PositionEntry) HasClosedAt() bool`

HasClosedAt returns a boolean if a field has been set.

### GetClosedAtDateTime

`func (o *PositionEntry) GetClosedAtDateTime() time.Time`

GetClosedAtDateTime returns the ClosedAtDateTime field if non-nil, zero value otherwise.

### GetClosedAtDateTimeOk

`func (o *PositionEntry) GetClosedAtDateTimeOk() (*time.Time, bool)`

GetClosedAtDateTimeOk returns a tuple with the ClosedAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAtDateTime

`func (o *PositionEntry) SetClosedAtDateTime(v time.Time)`

SetClosedAtDateTime sets ClosedAtDateTime field to given value.

### HasClosedAtDateTime

`func (o *PositionEntry) HasClosedAtDateTime() bool`

HasClosedAtDateTime returns a boolean if a field has been set.

### SetClosedAtDateTimeNil

`func (o *PositionEntry) SetClosedAtDateTimeNil(b bool)`

 SetClosedAtDateTimeNil sets the value for ClosedAtDateTime to be an explicit nil

### UnsetClosedAtDateTime
`func (o *PositionEntry) UnsetClosedAtDateTime()`

UnsetClosedAtDateTime ensures that no value is present for ClosedAtDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


