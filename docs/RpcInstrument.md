# RpcInstrument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstrumentId** | Pointer to **string** | Unique instrument identifier (format: VENUE:BASE/QUOTE) | [optional] 
**Venue** | Pointer to [**Venue**](Venue.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**ExternalSymbol** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InstrumentType** | Pointer to [**InstrumentType**](InstrumentType.md) |  | [optional] 
**Status** | Pointer to [**InstrumentStatus**](InstrumentStatus.md) |  | [optional] 
**BaseAsset** | Pointer to **string** |  | [optional] 
**QuoteAsset** | Pointer to **string** |  | [optional] 
**BaseSecurityType** | Pointer to [**SecurityType**](SecurityType.md) |  | [optional] 
**QuoteSecurityType** | Pointer to [**SecurityType**](SecurityType.md) |  | [optional] 
**BasePrecision** | Pointer to **int32** |  | [optional] 
**QuotePrecision** | Pointer to **int32** |  | [optional] 
**LotSize** | Pointer to **string** | Decimal value as string | [optional] 
**PipSize** | Pointer to **string** | Decimal value as string | [optional] 
**MinQuantity** | Pointer to **string** | Decimal value as string | [optional] 
**MaxQuantity** | Pointer to **string** | Decimal value as string | [optional] 
**MinNotional** | Pointer to **string** | Decimal value as string | [optional] 
**MaxNotional** | Pointer to **string** | Decimal value as string | [optional] 
**OrderTypes** | Pointer to **[]string** |  | [optional] 
**TimeInForceOptions** | Pointer to **[]string** |  | [optional] 
**Fee** | Pointer to **string** | Decimal value as string | [optional] 
**IsIcebergAllowed** | Pointer to **bool** |  | [optional] 
**ContractSize** | Pointer to **string** | Decimal value as string | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRpcInstrument

`func NewRpcInstrument() *RpcInstrument`

NewRpcInstrument instantiates a new RpcInstrument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcInstrumentWithDefaults

`func NewRpcInstrumentWithDefaults() *RpcInstrument`

NewRpcInstrumentWithDefaults instantiates a new RpcInstrument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstrumentId

`func (o *RpcInstrument) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *RpcInstrument) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *RpcInstrument) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *RpcInstrument) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetVenue

`func (o *RpcInstrument) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *RpcInstrument) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *RpcInstrument) SetVenue(v Venue)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *RpcInstrument) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetSymbol

`func (o *RpcInstrument) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *RpcInstrument) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *RpcInstrument) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *RpcInstrument) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetExternalSymbol

`func (o *RpcInstrument) GetExternalSymbol() string`

GetExternalSymbol returns the ExternalSymbol field if non-nil, zero value otherwise.

### GetExternalSymbolOk

`func (o *RpcInstrument) GetExternalSymbolOk() (*string, bool)`

GetExternalSymbolOk returns a tuple with the ExternalSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSymbol

`func (o *RpcInstrument) SetExternalSymbol(v string)`

SetExternalSymbol sets ExternalSymbol field to given value.

### HasExternalSymbol

`func (o *RpcInstrument) HasExternalSymbol() bool`

HasExternalSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *RpcInstrument) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RpcInstrument) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RpcInstrument) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RpcInstrument) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstrumentType

`func (o *RpcInstrument) GetInstrumentType() InstrumentType`

GetInstrumentType returns the InstrumentType field if non-nil, zero value otherwise.

### GetInstrumentTypeOk

`func (o *RpcInstrument) GetInstrumentTypeOk() (*InstrumentType, bool)`

GetInstrumentTypeOk returns a tuple with the InstrumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentType

`func (o *RpcInstrument) SetInstrumentType(v InstrumentType)`

SetInstrumentType sets InstrumentType field to given value.

### HasInstrumentType

`func (o *RpcInstrument) HasInstrumentType() bool`

HasInstrumentType returns a boolean if a field has been set.

### GetStatus

`func (o *RpcInstrument) GetStatus() InstrumentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RpcInstrument) GetStatusOk() (*InstrumentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RpcInstrument) SetStatus(v InstrumentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RpcInstrument) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBaseAsset

`func (o *RpcInstrument) GetBaseAsset() string`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *RpcInstrument) GetBaseAssetOk() (*string, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *RpcInstrument) SetBaseAsset(v string)`

SetBaseAsset sets BaseAsset field to given value.

### HasBaseAsset

`func (o *RpcInstrument) HasBaseAsset() bool`

HasBaseAsset returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *RpcInstrument) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *RpcInstrument) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *RpcInstrument) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *RpcInstrument) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetBaseSecurityType

`func (o *RpcInstrument) GetBaseSecurityType() SecurityType`

GetBaseSecurityType returns the BaseSecurityType field if non-nil, zero value otherwise.

### GetBaseSecurityTypeOk

`func (o *RpcInstrument) GetBaseSecurityTypeOk() (*SecurityType, bool)`

GetBaseSecurityTypeOk returns a tuple with the BaseSecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSecurityType

`func (o *RpcInstrument) SetBaseSecurityType(v SecurityType)`

SetBaseSecurityType sets BaseSecurityType field to given value.

### HasBaseSecurityType

`func (o *RpcInstrument) HasBaseSecurityType() bool`

HasBaseSecurityType returns a boolean if a field has been set.

### GetQuoteSecurityType

`func (o *RpcInstrument) GetQuoteSecurityType() SecurityType`

GetQuoteSecurityType returns the QuoteSecurityType field if non-nil, zero value otherwise.

### GetQuoteSecurityTypeOk

`func (o *RpcInstrument) GetQuoteSecurityTypeOk() (*SecurityType, bool)`

GetQuoteSecurityTypeOk returns a tuple with the QuoteSecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteSecurityType

`func (o *RpcInstrument) SetQuoteSecurityType(v SecurityType)`

SetQuoteSecurityType sets QuoteSecurityType field to given value.

### HasQuoteSecurityType

`func (o *RpcInstrument) HasQuoteSecurityType() bool`

HasQuoteSecurityType returns a boolean if a field has been set.

### GetBasePrecision

`func (o *RpcInstrument) GetBasePrecision() int32`

GetBasePrecision returns the BasePrecision field if non-nil, zero value otherwise.

### GetBasePrecisionOk

`func (o *RpcInstrument) GetBasePrecisionOk() (*int32, bool)`

GetBasePrecisionOk returns a tuple with the BasePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePrecision

`func (o *RpcInstrument) SetBasePrecision(v int32)`

SetBasePrecision sets BasePrecision field to given value.

### HasBasePrecision

`func (o *RpcInstrument) HasBasePrecision() bool`

HasBasePrecision returns a boolean if a field has been set.

### GetQuotePrecision

`func (o *RpcInstrument) GetQuotePrecision() int32`

GetQuotePrecision returns the QuotePrecision field if non-nil, zero value otherwise.

### GetQuotePrecisionOk

`func (o *RpcInstrument) GetQuotePrecisionOk() (*int32, bool)`

GetQuotePrecisionOk returns a tuple with the QuotePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotePrecision

`func (o *RpcInstrument) SetQuotePrecision(v int32)`

SetQuotePrecision sets QuotePrecision field to given value.

### HasQuotePrecision

`func (o *RpcInstrument) HasQuotePrecision() bool`

HasQuotePrecision returns a boolean if a field has been set.

### GetLotSize

`func (o *RpcInstrument) GetLotSize() string`

GetLotSize returns the LotSize field if non-nil, zero value otherwise.

### GetLotSizeOk

`func (o *RpcInstrument) GetLotSizeOk() (*string, bool)`

GetLotSizeOk returns a tuple with the LotSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotSize

`func (o *RpcInstrument) SetLotSize(v string)`

SetLotSize sets LotSize field to given value.

### HasLotSize

`func (o *RpcInstrument) HasLotSize() bool`

HasLotSize returns a boolean if a field has been set.

### GetPipSize

`func (o *RpcInstrument) GetPipSize() string`

GetPipSize returns the PipSize field if non-nil, zero value otherwise.

### GetPipSizeOk

`func (o *RpcInstrument) GetPipSizeOk() (*string, bool)`

GetPipSizeOk returns a tuple with the PipSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipSize

`func (o *RpcInstrument) SetPipSize(v string)`

SetPipSize sets PipSize field to given value.

### HasPipSize

`func (o *RpcInstrument) HasPipSize() bool`

HasPipSize returns a boolean if a field has been set.

### GetMinQuantity

`func (o *RpcInstrument) GetMinQuantity() string`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *RpcInstrument) GetMinQuantityOk() (*string, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *RpcInstrument) SetMinQuantity(v string)`

SetMinQuantity sets MinQuantity field to given value.

### HasMinQuantity

`func (o *RpcInstrument) HasMinQuantity() bool`

HasMinQuantity returns a boolean if a field has been set.

### GetMaxQuantity

`func (o *RpcInstrument) GetMaxQuantity() string`

GetMaxQuantity returns the MaxQuantity field if non-nil, zero value otherwise.

### GetMaxQuantityOk

`func (o *RpcInstrument) GetMaxQuantityOk() (*string, bool)`

GetMaxQuantityOk returns a tuple with the MaxQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQuantity

`func (o *RpcInstrument) SetMaxQuantity(v string)`

SetMaxQuantity sets MaxQuantity field to given value.

### HasMaxQuantity

`func (o *RpcInstrument) HasMaxQuantity() bool`

HasMaxQuantity returns a boolean if a field has been set.

### GetMinNotional

`func (o *RpcInstrument) GetMinNotional() string`

GetMinNotional returns the MinNotional field if non-nil, zero value otherwise.

### GetMinNotionalOk

`func (o *RpcInstrument) GetMinNotionalOk() (*string, bool)`

GetMinNotionalOk returns a tuple with the MinNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNotional

`func (o *RpcInstrument) SetMinNotional(v string)`

SetMinNotional sets MinNotional field to given value.

### HasMinNotional

`func (o *RpcInstrument) HasMinNotional() bool`

HasMinNotional returns a boolean if a field has been set.

### GetMaxNotional

`func (o *RpcInstrument) GetMaxNotional() string`

GetMaxNotional returns the MaxNotional field if non-nil, zero value otherwise.

### GetMaxNotionalOk

`func (o *RpcInstrument) GetMaxNotionalOk() (*string, bool)`

GetMaxNotionalOk returns a tuple with the MaxNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNotional

`func (o *RpcInstrument) SetMaxNotional(v string)`

SetMaxNotional sets MaxNotional field to given value.

### HasMaxNotional

`func (o *RpcInstrument) HasMaxNotional() bool`

HasMaxNotional returns a boolean if a field has been set.

### GetOrderTypes

`func (o *RpcInstrument) GetOrderTypes() []string`

GetOrderTypes returns the OrderTypes field if non-nil, zero value otherwise.

### GetOrderTypesOk

`func (o *RpcInstrument) GetOrderTypesOk() (*[]string, bool)`

GetOrderTypesOk returns a tuple with the OrderTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTypes

`func (o *RpcInstrument) SetOrderTypes(v []string)`

SetOrderTypes sets OrderTypes field to given value.

### HasOrderTypes

`func (o *RpcInstrument) HasOrderTypes() bool`

HasOrderTypes returns a boolean if a field has been set.

### GetTimeInForceOptions

`func (o *RpcInstrument) GetTimeInForceOptions() []string`

GetTimeInForceOptions returns the TimeInForceOptions field if non-nil, zero value otherwise.

### GetTimeInForceOptionsOk

`func (o *RpcInstrument) GetTimeInForceOptionsOk() (*[]string, bool)`

GetTimeInForceOptionsOk returns a tuple with the TimeInForceOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForceOptions

`func (o *RpcInstrument) SetTimeInForceOptions(v []string)`

SetTimeInForceOptions sets TimeInForceOptions field to given value.

### HasTimeInForceOptions

`func (o *RpcInstrument) HasTimeInForceOptions() bool`

HasTimeInForceOptions returns a boolean if a field has been set.

### GetFee

`func (o *RpcInstrument) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *RpcInstrument) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *RpcInstrument) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *RpcInstrument) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetIsIcebergAllowed

`func (o *RpcInstrument) GetIsIcebergAllowed() bool`

GetIsIcebergAllowed returns the IsIcebergAllowed field if non-nil, zero value otherwise.

### GetIsIcebergAllowedOk

`func (o *RpcInstrument) GetIsIcebergAllowedOk() (*bool, bool)`

GetIsIcebergAllowedOk returns a tuple with the IsIcebergAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIcebergAllowed

`func (o *RpcInstrument) SetIsIcebergAllowed(v bool)`

SetIsIcebergAllowed sets IsIcebergAllowed field to given value.

### HasIsIcebergAllowed

`func (o *RpcInstrument) HasIsIcebergAllowed() bool`

HasIsIcebergAllowed returns a boolean if a field has been set.

### GetContractSize

`func (o *RpcInstrument) GetContractSize() string`

GetContractSize returns the ContractSize field if non-nil, zero value otherwise.

### GetContractSizeOk

`func (o *RpcInstrument) GetContractSizeOk() (*string, bool)`

GetContractSizeOk returns a tuple with the ContractSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractSize

`func (o *RpcInstrument) SetContractSize(v string)`

SetContractSize sets ContractSize field to given value.

### HasContractSize

`func (o *RpcInstrument) HasContractSize() bool`

HasContractSize returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RpcInstrument) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RpcInstrument) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RpcInstrument) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RpcInstrument) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RpcInstrument) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RpcInstrument) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RpcInstrument) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RpcInstrument) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


