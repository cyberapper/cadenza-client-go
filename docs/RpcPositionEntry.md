# RpcPositionEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PositionId** | Pointer to **string** |  | [optional] 
**ExternalPositionId** | Pointer to **string** |  | [optional] 
**TradingAccountId** | Pointer to **string** |  | [optional] 
**SecuritySymbol** | Pointer to **string** |  | [optional] 
**InstrumentId** | Pointer to **string** | Instrument ID | [optional] 
**SecurityType** | Pointer to [**SecurityType**](SecurityType.md) |  | [optional] 
**Status** | Pointer to [**PositionStatus**](PositionStatus.md) |  | [optional] 
**Quantity** | Pointer to **string** | Position quantity | [optional] 
**EntryPrice** | Pointer to **string** |  | [optional] 
**ExitPrice** | Pointer to **string** |  | [optional] 
**CurrentPrice** | Pointer to **string** |  | [optional] 
**UnrealizedPnl** | Pointer to **string** |  | [optional] 
**RealizedPnl** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRpcPositionEntry

`func NewRpcPositionEntry() *RpcPositionEntry`

NewRpcPositionEntry instantiates a new RpcPositionEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcPositionEntryWithDefaults

`func NewRpcPositionEntryWithDefaults() *RpcPositionEntry`

NewRpcPositionEntryWithDefaults instantiates a new RpcPositionEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPositionId

`func (o *RpcPositionEntry) GetPositionId() string`

GetPositionId returns the PositionId field if non-nil, zero value otherwise.

### GetPositionIdOk

`func (o *RpcPositionEntry) GetPositionIdOk() (*string, bool)`

GetPositionIdOk returns a tuple with the PositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionId

`func (o *RpcPositionEntry) SetPositionId(v string)`

SetPositionId sets PositionId field to given value.

### HasPositionId

`func (o *RpcPositionEntry) HasPositionId() bool`

HasPositionId returns a boolean if a field has been set.

### GetExternalPositionId

`func (o *RpcPositionEntry) GetExternalPositionId() string`

GetExternalPositionId returns the ExternalPositionId field if non-nil, zero value otherwise.

### GetExternalPositionIdOk

`func (o *RpcPositionEntry) GetExternalPositionIdOk() (*string, bool)`

GetExternalPositionIdOk returns a tuple with the ExternalPositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPositionId

`func (o *RpcPositionEntry) SetExternalPositionId(v string)`

SetExternalPositionId sets ExternalPositionId field to given value.

### HasExternalPositionId

`func (o *RpcPositionEntry) HasExternalPositionId() bool`

HasExternalPositionId returns a boolean if a field has been set.

### GetTradingAccountId

`func (o *RpcPositionEntry) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *RpcPositionEntry) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *RpcPositionEntry) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.

### HasTradingAccountId

`func (o *RpcPositionEntry) HasTradingAccountId() bool`

HasTradingAccountId returns a boolean if a field has been set.

### GetSecuritySymbol

`func (o *RpcPositionEntry) GetSecuritySymbol() string`

GetSecuritySymbol returns the SecuritySymbol field if non-nil, zero value otherwise.

### GetSecuritySymbolOk

`func (o *RpcPositionEntry) GetSecuritySymbolOk() (*string, bool)`

GetSecuritySymbolOk returns a tuple with the SecuritySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritySymbol

`func (o *RpcPositionEntry) SetSecuritySymbol(v string)`

SetSecuritySymbol sets SecuritySymbol field to given value.

### HasSecuritySymbol

`func (o *RpcPositionEntry) HasSecuritySymbol() bool`

HasSecuritySymbol returns a boolean if a field has been set.

### GetInstrumentId

`func (o *RpcPositionEntry) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *RpcPositionEntry) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *RpcPositionEntry) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *RpcPositionEntry) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetSecurityType

`func (o *RpcPositionEntry) GetSecurityType() SecurityType`

GetSecurityType returns the SecurityType field if non-nil, zero value otherwise.

### GetSecurityTypeOk

`func (o *RpcPositionEntry) GetSecurityTypeOk() (*SecurityType, bool)`

GetSecurityTypeOk returns a tuple with the SecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityType

`func (o *RpcPositionEntry) SetSecurityType(v SecurityType)`

SetSecurityType sets SecurityType field to given value.

### HasSecurityType

`func (o *RpcPositionEntry) HasSecurityType() bool`

HasSecurityType returns a boolean if a field has been set.

### GetStatus

`func (o *RpcPositionEntry) GetStatus() PositionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RpcPositionEntry) GetStatusOk() (*PositionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RpcPositionEntry) SetStatus(v PositionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RpcPositionEntry) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetQuantity

`func (o *RpcPositionEntry) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *RpcPositionEntry) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *RpcPositionEntry) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *RpcPositionEntry) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetEntryPrice

`func (o *RpcPositionEntry) GetEntryPrice() string`

GetEntryPrice returns the EntryPrice field if non-nil, zero value otherwise.

### GetEntryPriceOk

`func (o *RpcPositionEntry) GetEntryPriceOk() (*string, bool)`

GetEntryPriceOk returns a tuple with the EntryPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPrice

`func (o *RpcPositionEntry) SetEntryPrice(v string)`

SetEntryPrice sets EntryPrice field to given value.

### HasEntryPrice

`func (o *RpcPositionEntry) HasEntryPrice() bool`

HasEntryPrice returns a boolean if a field has been set.

### GetExitPrice

`func (o *RpcPositionEntry) GetExitPrice() string`

GetExitPrice returns the ExitPrice field if non-nil, zero value otherwise.

### GetExitPriceOk

`func (o *RpcPositionEntry) GetExitPriceOk() (*string, bool)`

GetExitPriceOk returns a tuple with the ExitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitPrice

`func (o *RpcPositionEntry) SetExitPrice(v string)`

SetExitPrice sets ExitPrice field to given value.

### HasExitPrice

`func (o *RpcPositionEntry) HasExitPrice() bool`

HasExitPrice returns a boolean if a field has been set.

### GetCurrentPrice

`func (o *RpcPositionEntry) GetCurrentPrice() string`

GetCurrentPrice returns the CurrentPrice field if non-nil, zero value otherwise.

### GetCurrentPriceOk

`func (o *RpcPositionEntry) GetCurrentPriceOk() (*string, bool)`

GetCurrentPriceOk returns a tuple with the CurrentPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPrice

`func (o *RpcPositionEntry) SetCurrentPrice(v string)`

SetCurrentPrice sets CurrentPrice field to given value.

### HasCurrentPrice

`func (o *RpcPositionEntry) HasCurrentPrice() bool`

HasCurrentPrice returns a boolean if a field has been set.

### GetUnrealizedPnl

`func (o *RpcPositionEntry) GetUnrealizedPnl() string`

GetUnrealizedPnl returns the UnrealizedPnl field if non-nil, zero value otherwise.

### GetUnrealizedPnlOk

`func (o *RpcPositionEntry) GetUnrealizedPnlOk() (*string, bool)`

GetUnrealizedPnlOk returns a tuple with the UnrealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedPnl

`func (o *RpcPositionEntry) SetUnrealizedPnl(v string)`

SetUnrealizedPnl sets UnrealizedPnl field to given value.

### HasUnrealizedPnl

`func (o *RpcPositionEntry) HasUnrealizedPnl() bool`

HasUnrealizedPnl returns a boolean if a field has been set.

### GetRealizedPnl

`func (o *RpcPositionEntry) GetRealizedPnl() string`

GetRealizedPnl returns the RealizedPnl field if non-nil, zero value otherwise.

### GetRealizedPnlOk

`func (o *RpcPositionEntry) GetRealizedPnlOk() (*string, bool)`

GetRealizedPnlOk returns a tuple with the RealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedPnl

`func (o *RpcPositionEntry) SetRealizedPnl(v string)`

SetRealizedPnl sets RealizedPnl field to given value.

### HasRealizedPnl

`func (o *RpcPositionEntry) HasRealizedPnl() bool`

HasRealizedPnl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RpcPositionEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RpcPositionEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RpcPositionEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RpcPositionEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RpcPositionEntry) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RpcPositionEntry) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RpcPositionEntry) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RpcPositionEntry) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


