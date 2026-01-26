# RpcBalanceEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceId** | Pointer to **string** |  | [optional] 
**ExternalBalanceId** | Pointer to **string** |  | [optional] 
**TradingAccountId** | Pointer to **string** |  | [optional] 
**SecuritySymbol** | Pointer to **string** | Asset symbol | [optional] 
**SecurityType** | Pointer to [**SecurityType**](SecurityType.md) |  | [optional] 
**Status** | Pointer to [**BalanceStatus**](BalanceStatus.md) |  | [optional] 
**Free** | Pointer to **string** | Available balance | [optional] 
**Locked** | Pointer to **string** | Locked balance | [optional] 
**Frozen** | Pointer to **string** | Frozen balance | [optional] 
**Borrowed** | Pointer to **string** | Borrowed amount | [optional] 
**InterestOwed** | Pointer to **string** | Interest owed | [optional] 
**Total** | Pointer to **string** | Total balance | [optional] 
**Net** | Pointer to **string** | Net equity | [optional] 
**CollateralWeight** | Pointer to **string** |  | [optional] 
**CollateralValue** | Pointer to **string** |  | [optional] 
**CollateralEnabled** | Pointer to **bool** |  | [optional] 
**CrossMargin** | Pointer to **string** |  | [optional] 
**IsolatedMargin** | Pointer to **string** |  | [optional] 
**MarginRatio** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRpcBalanceEntry

`func NewRpcBalanceEntry() *RpcBalanceEntry`

NewRpcBalanceEntry instantiates a new RpcBalanceEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcBalanceEntryWithDefaults

`func NewRpcBalanceEntryWithDefaults() *RpcBalanceEntry`

NewRpcBalanceEntryWithDefaults instantiates a new RpcBalanceEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceId

`func (o *RpcBalanceEntry) GetBalanceId() string`

GetBalanceId returns the BalanceId field if non-nil, zero value otherwise.

### GetBalanceIdOk

`func (o *RpcBalanceEntry) GetBalanceIdOk() (*string, bool)`

GetBalanceIdOk returns a tuple with the BalanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceId

`func (o *RpcBalanceEntry) SetBalanceId(v string)`

SetBalanceId sets BalanceId field to given value.

### HasBalanceId

`func (o *RpcBalanceEntry) HasBalanceId() bool`

HasBalanceId returns a boolean if a field has been set.

### GetExternalBalanceId

`func (o *RpcBalanceEntry) GetExternalBalanceId() string`

GetExternalBalanceId returns the ExternalBalanceId field if non-nil, zero value otherwise.

### GetExternalBalanceIdOk

`func (o *RpcBalanceEntry) GetExternalBalanceIdOk() (*string, bool)`

GetExternalBalanceIdOk returns a tuple with the ExternalBalanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalBalanceId

`func (o *RpcBalanceEntry) SetExternalBalanceId(v string)`

SetExternalBalanceId sets ExternalBalanceId field to given value.

### HasExternalBalanceId

`func (o *RpcBalanceEntry) HasExternalBalanceId() bool`

HasExternalBalanceId returns a boolean if a field has been set.

### GetTradingAccountId

`func (o *RpcBalanceEntry) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *RpcBalanceEntry) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *RpcBalanceEntry) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.

### HasTradingAccountId

`func (o *RpcBalanceEntry) HasTradingAccountId() bool`

HasTradingAccountId returns a boolean if a field has been set.

### GetSecuritySymbol

`func (o *RpcBalanceEntry) GetSecuritySymbol() string`

GetSecuritySymbol returns the SecuritySymbol field if non-nil, zero value otherwise.

### GetSecuritySymbolOk

`func (o *RpcBalanceEntry) GetSecuritySymbolOk() (*string, bool)`

GetSecuritySymbolOk returns a tuple with the SecuritySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritySymbol

`func (o *RpcBalanceEntry) SetSecuritySymbol(v string)`

SetSecuritySymbol sets SecuritySymbol field to given value.

### HasSecuritySymbol

`func (o *RpcBalanceEntry) HasSecuritySymbol() bool`

HasSecuritySymbol returns a boolean if a field has been set.

### GetSecurityType

`func (o *RpcBalanceEntry) GetSecurityType() SecurityType`

GetSecurityType returns the SecurityType field if non-nil, zero value otherwise.

### GetSecurityTypeOk

`func (o *RpcBalanceEntry) GetSecurityTypeOk() (*SecurityType, bool)`

GetSecurityTypeOk returns a tuple with the SecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityType

`func (o *RpcBalanceEntry) SetSecurityType(v SecurityType)`

SetSecurityType sets SecurityType field to given value.

### HasSecurityType

`func (o *RpcBalanceEntry) HasSecurityType() bool`

HasSecurityType returns a boolean if a field has been set.

### GetStatus

`func (o *RpcBalanceEntry) GetStatus() BalanceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RpcBalanceEntry) GetStatusOk() (*BalanceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RpcBalanceEntry) SetStatus(v BalanceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RpcBalanceEntry) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFree

`func (o *RpcBalanceEntry) GetFree() string`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *RpcBalanceEntry) GetFreeOk() (*string, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *RpcBalanceEntry) SetFree(v string)`

SetFree sets Free field to given value.

### HasFree

`func (o *RpcBalanceEntry) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetLocked

`func (o *RpcBalanceEntry) GetLocked() string`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *RpcBalanceEntry) GetLockedOk() (*string, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *RpcBalanceEntry) SetLocked(v string)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *RpcBalanceEntry) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetFrozen

`func (o *RpcBalanceEntry) GetFrozen() string`

GetFrozen returns the Frozen field if non-nil, zero value otherwise.

### GetFrozenOk

`func (o *RpcBalanceEntry) GetFrozenOk() (*string, bool)`

GetFrozenOk returns a tuple with the Frozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozen

`func (o *RpcBalanceEntry) SetFrozen(v string)`

SetFrozen sets Frozen field to given value.

### HasFrozen

`func (o *RpcBalanceEntry) HasFrozen() bool`

HasFrozen returns a boolean if a field has been set.

### GetBorrowed

`func (o *RpcBalanceEntry) GetBorrowed() string`

GetBorrowed returns the Borrowed field if non-nil, zero value otherwise.

### GetBorrowedOk

`func (o *RpcBalanceEntry) GetBorrowedOk() (*string, bool)`

GetBorrowedOk returns a tuple with the Borrowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowed

`func (o *RpcBalanceEntry) SetBorrowed(v string)`

SetBorrowed sets Borrowed field to given value.

### HasBorrowed

`func (o *RpcBalanceEntry) HasBorrowed() bool`

HasBorrowed returns a boolean if a field has been set.

### GetInterestOwed

`func (o *RpcBalanceEntry) GetInterestOwed() string`

GetInterestOwed returns the InterestOwed field if non-nil, zero value otherwise.

### GetInterestOwedOk

`func (o *RpcBalanceEntry) GetInterestOwedOk() (*string, bool)`

GetInterestOwedOk returns a tuple with the InterestOwed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestOwed

`func (o *RpcBalanceEntry) SetInterestOwed(v string)`

SetInterestOwed sets InterestOwed field to given value.

### HasInterestOwed

`func (o *RpcBalanceEntry) HasInterestOwed() bool`

HasInterestOwed returns a boolean if a field has been set.

### GetTotal

`func (o *RpcBalanceEntry) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RpcBalanceEntry) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RpcBalanceEntry) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *RpcBalanceEntry) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetNet

`func (o *RpcBalanceEntry) GetNet() string`

GetNet returns the Net field if non-nil, zero value otherwise.

### GetNetOk

`func (o *RpcBalanceEntry) GetNetOk() (*string, bool)`

GetNetOk returns a tuple with the Net field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet

`func (o *RpcBalanceEntry) SetNet(v string)`

SetNet sets Net field to given value.

### HasNet

`func (o *RpcBalanceEntry) HasNet() bool`

HasNet returns a boolean if a field has been set.

### GetCollateralWeight

`func (o *RpcBalanceEntry) GetCollateralWeight() string`

GetCollateralWeight returns the CollateralWeight field if non-nil, zero value otherwise.

### GetCollateralWeightOk

`func (o *RpcBalanceEntry) GetCollateralWeightOk() (*string, bool)`

GetCollateralWeightOk returns a tuple with the CollateralWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralWeight

`func (o *RpcBalanceEntry) SetCollateralWeight(v string)`

SetCollateralWeight sets CollateralWeight field to given value.

### HasCollateralWeight

`func (o *RpcBalanceEntry) HasCollateralWeight() bool`

HasCollateralWeight returns a boolean if a field has been set.

### GetCollateralValue

`func (o *RpcBalanceEntry) GetCollateralValue() string`

GetCollateralValue returns the CollateralValue field if non-nil, zero value otherwise.

### GetCollateralValueOk

`func (o *RpcBalanceEntry) GetCollateralValueOk() (*string, bool)`

GetCollateralValueOk returns a tuple with the CollateralValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralValue

`func (o *RpcBalanceEntry) SetCollateralValue(v string)`

SetCollateralValue sets CollateralValue field to given value.

### HasCollateralValue

`func (o *RpcBalanceEntry) HasCollateralValue() bool`

HasCollateralValue returns a boolean if a field has been set.

### GetCollateralEnabled

`func (o *RpcBalanceEntry) GetCollateralEnabled() bool`

GetCollateralEnabled returns the CollateralEnabled field if non-nil, zero value otherwise.

### GetCollateralEnabledOk

`func (o *RpcBalanceEntry) GetCollateralEnabledOk() (*bool, bool)`

GetCollateralEnabledOk returns a tuple with the CollateralEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralEnabled

`func (o *RpcBalanceEntry) SetCollateralEnabled(v bool)`

SetCollateralEnabled sets CollateralEnabled field to given value.

### HasCollateralEnabled

`func (o *RpcBalanceEntry) HasCollateralEnabled() bool`

HasCollateralEnabled returns a boolean if a field has been set.

### GetCrossMargin

`func (o *RpcBalanceEntry) GetCrossMargin() string`

GetCrossMargin returns the CrossMargin field if non-nil, zero value otherwise.

### GetCrossMarginOk

`func (o *RpcBalanceEntry) GetCrossMarginOk() (*string, bool)`

GetCrossMarginOk returns a tuple with the CrossMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossMargin

`func (o *RpcBalanceEntry) SetCrossMargin(v string)`

SetCrossMargin sets CrossMargin field to given value.

### HasCrossMargin

`func (o *RpcBalanceEntry) HasCrossMargin() bool`

HasCrossMargin returns a boolean if a field has been set.

### GetIsolatedMargin

`func (o *RpcBalanceEntry) GetIsolatedMargin() string`

GetIsolatedMargin returns the IsolatedMargin field if non-nil, zero value otherwise.

### GetIsolatedMarginOk

`func (o *RpcBalanceEntry) GetIsolatedMarginOk() (*string, bool)`

GetIsolatedMarginOk returns a tuple with the IsolatedMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolatedMargin

`func (o *RpcBalanceEntry) SetIsolatedMargin(v string)`

SetIsolatedMargin sets IsolatedMargin field to given value.

### HasIsolatedMargin

`func (o *RpcBalanceEntry) HasIsolatedMargin() bool`

HasIsolatedMargin returns a boolean if a field has been set.

### GetMarginRatio

`func (o *RpcBalanceEntry) GetMarginRatio() string`

GetMarginRatio returns the MarginRatio field if non-nil, zero value otherwise.

### GetMarginRatioOk

`func (o *RpcBalanceEntry) GetMarginRatioOk() (*string, bool)`

GetMarginRatioOk returns a tuple with the MarginRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginRatio

`func (o *RpcBalanceEntry) SetMarginRatio(v string)`

SetMarginRatio sets MarginRatio field to given value.

### HasMarginRatio

`func (o *RpcBalanceEntry) HasMarginRatio() bool`

HasMarginRatio returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RpcBalanceEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RpcBalanceEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RpcBalanceEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RpcBalanceEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RpcBalanceEntry) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RpcBalanceEntry) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RpcBalanceEntry) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RpcBalanceEntry) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


