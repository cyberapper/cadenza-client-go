# Instrument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstrumentId** | **string** | Instrument ID in format {VENUE}:{BASE}/{QUOTE} | 
**Venue** | [**Venue**](Venue.md) |  | 
**Symbol** | **string** | Human-readable symbol format | 
**ExternalSymbol** | **string** | Symbol format used by the exchange | 
**Description** | Pointer to **string** | Symbol description, human readable description of the instrument | [optional] 
**InstrumentType** | [**InstrumentType**](InstrumentType.md) |  | 
**Status** | [**InstrumentStatus**](InstrumentStatus.md) |  | 
**BaseAsset** | **string** |  | 
**QuoteAsset** | **string** |  | 
**BaseSecurityType** | [**SecurityType**](SecurityType.md) |  | 
**QuoteSecurityType** | [**SecurityType**](SecurityType.md) |  | 
**BasePrecision** | **int32** | Base asset precision | 
**QuotePrecision** | **int32** | Quote asset precision | 
**BaseMaxSignificant** | **NullableInt32** | Maximum significant digits for base asset | 
**QuoteMaxSignificant** | **NullableInt32** | Maximum significant digits for quote asset | 
**LotSize** | **string** | Decimal value as string to preserve precision | 
**PipSize** | **string** | Decimal value as string to preserve precision | 
**BaseScale** | **NullableInt32** | Base asset scale factor | 
**QuoteScale** | **NullableInt32** | Quote asset scale factor | 
**MinQuantity** | **string** | Decimal value as string to preserve precision | 
**MaxQuantity** | **string** | Decimal value as string to preserve precision | 
**MinNotional** | **string** | Decimal value as string to preserve precision | 
**MaxNotional** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**OrderFilters** | Pointer to **map[string]interface{}** | Exchange-specific order filters | [optional] 
**OrderTypes** | [**[]OrderType**](OrderType.md) |  | 
**TimeInForceOptions** | [**[]TimeInForce**](TimeInForce.md) |  | 
**TradingHours** | Pointer to **map[string]interface{}** | Trading hours and schedule information | [optional] 
**IsIcebergAllowed** | Pointer to **bool** | Whether iceberg orders are allowed | [optional] 
**IcebergMinQuantity** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**DeliveryDate** | Pointer to **int64** | Unix timestamp in milliseconds | [optional] 
**DeliveryDateTime** | Pointer to **time.Time** | Delivery date in ISO 8601 format for derivatives | [optional] 
**ExerciseStyle** | Pointer to **string** | Exercise style for options | [optional] 
**StrikePrice** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 

## Methods

### NewInstrument

`func NewInstrument(instrumentId string, venue Venue, symbol string, externalSymbol string, instrumentType InstrumentType, status InstrumentStatus, baseAsset string, quoteAsset string, baseSecurityType SecurityType, quoteSecurityType SecurityType, basePrecision int32, quotePrecision int32, baseMaxSignificant NullableInt32, quoteMaxSignificant NullableInt32, lotSize string, pipSize string, baseScale NullableInt32, quoteScale NullableInt32, minQuantity string, maxQuantity string, minNotional string, orderTypes []OrderType, timeInForceOptions []TimeInForce, ) *Instrument`

NewInstrument instantiates a new Instrument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstrumentWithDefaults

`func NewInstrumentWithDefaults() *Instrument`

NewInstrumentWithDefaults instantiates a new Instrument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstrumentId

`func (o *Instrument) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *Instrument) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *Instrument) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.


### GetVenue

`func (o *Instrument) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *Instrument) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *Instrument) SetVenue(v Venue)`

SetVenue sets Venue field to given value.


### GetSymbol

`func (o *Instrument) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Instrument) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Instrument) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetExternalSymbol

`func (o *Instrument) GetExternalSymbol() string`

GetExternalSymbol returns the ExternalSymbol field if non-nil, zero value otherwise.

### GetExternalSymbolOk

`func (o *Instrument) GetExternalSymbolOk() (*string, bool)`

GetExternalSymbolOk returns a tuple with the ExternalSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSymbol

`func (o *Instrument) SetExternalSymbol(v string)`

SetExternalSymbol sets ExternalSymbol field to given value.


### GetDescription

`func (o *Instrument) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Instrument) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Instrument) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Instrument) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstrumentType

`func (o *Instrument) GetInstrumentType() InstrumentType`

GetInstrumentType returns the InstrumentType field if non-nil, zero value otherwise.

### GetInstrumentTypeOk

`func (o *Instrument) GetInstrumentTypeOk() (*InstrumentType, bool)`

GetInstrumentTypeOk returns a tuple with the InstrumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentType

`func (o *Instrument) SetInstrumentType(v InstrumentType)`

SetInstrumentType sets InstrumentType field to given value.


### GetStatus

`func (o *Instrument) GetStatus() InstrumentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Instrument) GetStatusOk() (*InstrumentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Instrument) SetStatus(v InstrumentStatus)`

SetStatus sets Status field to given value.


### GetBaseAsset

`func (o *Instrument) GetBaseAsset() string`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *Instrument) GetBaseAssetOk() (*string, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *Instrument) SetBaseAsset(v string)`

SetBaseAsset sets BaseAsset field to given value.


### GetQuoteAsset

`func (o *Instrument) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *Instrument) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *Instrument) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.


### GetBaseSecurityType

`func (o *Instrument) GetBaseSecurityType() SecurityType`

GetBaseSecurityType returns the BaseSecurityType field if non-nil, zero value otherwise.

### GetBaseSecurityTypeOk

`func (o *Instrument) GetBaseSecurityTypeOk() (*SecurityType, bool)`

GetBaseSecurityTypeOk returns a tuple with the BaseSecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSecurityType

`func (o *Instrument) SetBaseSecurityType(v SecurityType)`

SetBaseSecurityType sets BaseSecurityType field to given value.


### GetQuoteSecurityType

`func (o *Instrument) GetQuoteSecurityType() SecurityType`

GetQuoteSecurityType returns the QuoteSecurityType field if non-nil, zero value otherwise.

### GetQuoteSecurityTypeOk

`func (o *Instrument) GetQuoteSecurityTypeOk() (*SecurityType, bool)`

GetQuoteSecurityTypeOk returns a tuple with the QuoteSecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteSecurityType

`func (o *Instrument) SetQuoteSecurityType(v SecurityType)`

SetQuoteSecurityType sets QuoteSecurityType field to given value.


### GetBasePrecision

`func (o *Instrument) GetBasePrecision() int32`

GetBasePrecision returns the BasePrecision field if non-nil, zero value otherwise.

### GetBasePrecisionOk

`func (o *Instrument) GetBasePrecisionOk() (*int32, bool)`

GetBasePrecisionOk returns a tuple with the BasePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePrecision

`func (o *Instrument) SetBasePrecision(v int32)`

SetBasePrecision sets BasePrecision field to given value.


### GetQuotePrecision

`func (o *Instrument) GetQuotePrecision() int32`

GetQuotePrecision returns the QuotePrecision field if non-nil, zero value otherwise.

### GetQuotePrecisionOk

`func (o *Instrument) GetQuotePrecisionOk() (*int32, bool)`

GetQuotePrecisionOk returns a tuple with the QuotePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotePrecision

`func (o *Instrument) SetQuotePrecision(v int32)`

SetQuotePrecision sets QuotePrecision field to given value.


### GetBaseMaxSignificant

`func (o *Instrument) GetBaseMaxSignificant() int32`

GetBaseMaxSignificant returns the BaseMaxSignificant field if non-nil, zero value otherwise.

### GetBaseMaxSignificantOk

`func (o *Instrument) GetBaseMaxSignificantOk() (*int32, bool)`

GetBaseMaxSignificantOk returns a tuple with the BaseMaxSignificant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseMaxSignificant

`func (o *Instrument) SetBaseMaxSignificant(v int32)`

SetBaseMaxSignificant sets BaseMaxSignificant field to given value.


### SetBaseMaxSignificantNil

`func (o *Instrument) SetBaseMaxSignificantNil(b bool)`

 SetBaseMaxSignificantNil sets the value for BaseMaxSignificant to be an explicit nil

### UnsetBaseMaxSignificant
`func (o *Instrument) UnsetBaseMaxSignificant()`

UnsetBaseMaxSignificant ensures that no value is present for BaseMaxSignificant, not even an explicit nil
### GetQuoteMaxSignificant

`func (o *Instrument) GetQuoteMaxSignificant() int32`

GetQuoteMaxSignificant returns the QuoteMaxSignificant field if non-nil, zero value otherwise.

### GetQuoteMaxSignificantOk

`func (o *Instrument) GetQuoteMaxSignificantOk() (*int32, bool)`

GetQuoteMaxSignificantOk returns a tuple with the QuoteMaxSignificant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteMaxSignificant

`func (o *Instrument) SetQuoteMaxSignificant(v int32)`

SetQuoteMaxSignificant sets QuoteMaxSignificant field to given value.


### SetQuoteMaxSignificantNil

`func (o *Instrument) SetQuoteMaxSignificantNil(b bool)`

 SetQuoteMaxSignificantNil sets the value for QuoteMaxSignificant to be an explicit nil

### UnsetQuoteMaxSignificant
`func (o *Instrument) UnsetQuoteMaxSignificant()`

UnsetQuoteMaxSignificant ensures that no value is present for QuoteMaxSignificant, not even an explicit nil
### GetLotSize

`func (o *Instrument) GetLotSize() string`

GetLotSize returns the LotSize field if non-nil, zero value otherwise.

### GetLotSizeOk

`func (o *Instrument) GetLotSizeOk() (*string, bool)`

GetLotSizeOk returns a tuple with the LotSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotSize

`func (o *Instrument) SetLotSize(v string)`

SetLotSize sets LotSize field to given value.


### GetPipSize

`func (o *Instrument) GetPipSize() string`

GetPipSize returns the PipSize field if non-nil, zero value otherwise.

### GetPipSizeOk

`func (o *Instrument) GetPipSizeOk() (*string, bool)`

GetPipSizeOk returns a tuple with the PipSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipSize

`func (o *Instrument) SetPipSize(v string)`

SetPipSize sets PipSize field to given value.


### GetBaseScale

`func (o *Instrument) GetBaseScale() int32`

GetBaseScale returns the BaseScale field if non-nil, zero value otherwise.

### GetBaseScaleOk

`func (o *Instrument) GetBaseScaleOk() (*int32, bool)`

GetBaseScaleOk returns a tuple with the BaseScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScale

`func (o *Instrument) SetBaseScale(v int32)`

SetBaseScale sets BaseScale field to given value.


### SetBaseScaleNil

`func (o *Instrument) SetBaseScaleNil(b bool)`

 SetBaseScaleNil sets the value for BaseScale to be an explicit nil

### UnsetBaseScale
`func (o *Instrument) UnsetBaseScale()`

UnsetBaseScale ensures that no value is present for BaseScale, not even an explicit nil
### GetQuoteScale

`func (o *Instrument) GetQuoteScale() int32`

GetQuoteScale returns the QuoteScale field if non-nil, zero value otherwise.

### GetQuoteScaleOk

`func (o *Instrument) GetQuoteScaleOk() (*int32, bool)`

GetQuoteScaleOk returns a tuple with the QuoteScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteScale

`func (o *Instrument) SetQuoteScale(v int32)`

SetQuoteScale sets QuoteScale field to given value.


### SetQuoteScaleNil

`func (o *Instrument) SetQuoteScaleNil(b bool)`

 SetQuoteScaleNil sets the value for QuoteScale to be an explicit nil

### UnsetQuoteScale
`func (o *Instrument) UnsetQuoteScale()`

UnsetQuoteScale ensures that no value is present for QuoteScale, not even an explicit nil
### GetMinQuantity

`func (o *Instrument) GetMinQuantity() string`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *Instrument) GetMinQuantityOk() (*string, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *Instrument) SetMinQuantity(v string)`

SetMinQuantity sets MinQuantity field to given value.


### GetMaxQuantity

`func (o *Instrument) GetMaxQuantity() string`

GetMaxQuantity returns the MaxQuantity field if non-nil, zero value otherwise.

### GetMaxQuantityOk

`func (o *Instrument) GetMaxQuantityOk() (*string, bool)`

GetMaxQuantityOk returns a tuple with the MaxQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQuantity

`func (o *Instrument) SetMaxQuantity(v string)`

SetMaxQuantity sets MaxQuantity field to given value.


### GetMinNotional

`func (o *Instrument) GetMinNotional() string`

GetMinNotional returns the MinNotional field if non-nil, zero value otherwise.

### GetMinNotionalOk

`func (o *Instrument) GetMinNotionalOk() (*string, bool)`

GetMinNotionalOk returns a tuple with the MinNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNotional

`func (o *Instrument) SetMinNotional(v string)`

SetMinNotional sets MinNotional field to given value.


### GetMaxNotional

`func (o *Instrument) GetMaxNotional() string`

GetMaxNotional returns the MaxNotional field if non-nil, zero value otherwise.

### GetMaxNotionalOk

`func (o *Instrument) GetMaxNotionalOk() (*string, bool)`

GetMaxNotionalOk returns a tuple with the MaxNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNotional

`func (o *Instrument) SetMaxNotional(v string)`

SetMaxNotional sets MaxNotional field to given value.

### HasMaxNotional

`func (o *Instrument) HasMaxNotional() bool`

HasMaxNotional returns a boolean if a field has been set.

### GetOrderFilters

`func (o *Instrument) GetOrderFilters() map[string]interface{}`

GetOrderFilters returns the OrderFilters field if non-nil, zero value otherwise.

### GetOrderFiltersOk

`func (o *Instrument) GetOrderFiltersOk() (*map[string]interface{}, bool)`

GetOrderFiltersOk returns a tuple with the OrderFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderFilters

`func (o *Instrument) SetOrderFilters(v map[string]interface{})`

SetOrderFilters sets OrderFilters field to given value.

### HasOrderFilters

`func (o *Instrument) HasOrderFilters() bool`

HasOrderFilters returns a boolean if a field has been set.

### GetOrderTypes

`func (o *Instrument) GetOrderTypes() []OrderType`

GetOrderTypes returns the OrderTypes field if non-nil, zero value otherwise.

### GetOrderTypesOk

`func (o *Instrument) GetOrderTypesOk() (*[]OrderType, bool)`

GetOrderTypesOk returns a tuple with the OrderTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTypes

`func (o *Instrument) SetOrderTypes(v []OrderType)`

SetOrderTypes sets OrderTypes field to given value.


### GetTimeInForceOptions

`func (o *Instrument) GetTimeInForceOptions() []TimeInForce`

GetTimeInForceOptions returns the TimeInForceOptions field if non-nil, zero value otherwise.

### GetTimeInForceOptionsOk

`func (o *Instrument) GetTimeInForceOptionsOk() (*[]TimeInForce, bool)`

GetTimeInForceOptionsOk returns a tuple with the TimeInForceOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForceOptions

`func (o *Instrument) SetTimeInForceOptions(v []TimeInForce)`

SetTimeInForceOptions sets TimeInForceOptions field to given value.


### GetTradingHours

`func (o *Instrument) GetTradingHours() map[string]interface{}`

GetTradingHours returns the TradingHours field if non-nil, zero value otherwise.

### GetTradingHoursOk

`func (o *Instrument) GetTradingHoursOk() (*map[string]interface{}, bool)`

GetTradingHoursOk returns a tuple with the TradingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingHours

`func (o *Instrument) SetTradingHours(v map[string]interface{})`

SetTradingHours sets TradingHours field to given value.

### HasTradingHours

`func (o *Instrument) HasTradingHours() bool`

HasTradingHours returns a boolean if a field has been set.

### GetIsIcebergAllowed

`func (o *Instrument) GetIsIcebergAllowed() bool`

GetIsIcebergAllowed returns the IsIcebergAllowed field if non-nil, zero value otherwise.

### GetIsIcebergAllowedOk

`func (o *Instrument) GetIsIcebergAllowedOk() (*bool, bool)`

GetIsIcebergAllowedOk returns a tuple with the IsIcebergAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIcebergAllowed

`func (o *Instrument) SetIsIcebergAllowed(v bool)`

SetIsIcebergAllowed sets IsIcebergAllowed field to given value.

### HasIsIcebergAllowed

`func (o *Instrument) HasIsIcebergAllowed() bool`

HasIsIcebergAllowed returns a boolean if a field has been set.

### GetIcebergMinQuantity

`func (o *Instrument) GetIcebergMinQuantity() string`

GetIcebergMinQuantity returns the IcebergMinQuantity field if non-nil, zero value otherwise.

### GetIcebergMinQuantityOk

`func (o *Instrument) GetIcebergMinQuantityOk() (*string, bool)`

GetIcebergMinQuantityOk returns a tuple with the IcebergMinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergMinQuantity

`func (o *Instrument) SetIcebergMinQuantity(v string)`

SetIcebergMinQuantity sets IcebergMinQuantity field to given value.

### HasIcebergMinQuantity

`func (o *Instrument) HasIcebergMinQuantity() bool`

HasIcebergMinQuantity returns a boolean if a field has been set.

### GetDeliveryDate

`func (o *Instrument) GetDeliveryDate() int64`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *Instrument) GetDeliveryDateOk() (*int64, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *Instrument) SetDeliveryDate(v int64)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *Instrument) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetDeliveryDateTime

`func (o *Instrument) GetDeliveryDateTime() time.Time`

GetDeliveryDateTime returns the DeliveryDateTime field if non-nil, zero value otherwise.

### GetDeliveryDateTimeOk

`func (o *Instrument) GetDeliveryDateTimeOk() (*time.Time, bool)`

GetDeliveryDateTimeOk returns a tuple with the DeliveryDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDateTime

`func (o *Instrument) SetDeliveryDateTime(v time.Time)`

SetDeliveryDateTime sets DeliveryDateTime field to given value.

### HasDeliveryDateTime

`func (o *Instrument) HasDeliveryDateTime() bool`

HasDeliveryDateTime returns a boolean if a field has been set.

### GetExerciseStyle

`func (o *Instrument) GetExerciseStyle() string`

GetExerciseStyle returns the ExerciseStyle field if non-nil, zero value otherwise.

### GetExerciseStyleOk

`func (o *Instrument) GetExerciseStyleOk() (*string, bool)`

GetExerciseStyleOk returns a tuple with the ExerciseStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExerciseStyle

`func (o *Instrument) SetExerciseStyle(v string)`

SetExerciseStyle sets ExerciseStyle field to given value.

### HasExerciseStyle

`func (o *Instrument) HasExerciseStyle() bool`

HasExerciseStyle returns a boolean if a field has been set.

### GetStrikePrice

`func (o *Instrument) GetStrikePrice() string`

GetStrikePrice returns the StrikePrice field if non-nil, zero value otherwise.

### GetStrikePriceOk

`func (o *Instrument) GetStrikePriceOk() (*string, bool)`

GetStrikePriceOk returns a tuple with the StrikePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikePrice

`func (o *Instrument) SetStrikePrice(v string)`

SetStrikePrice sets StrikePrice field to given value.

### HasStrikePrice

`func (o *Instrument) HasStrikePrice() bool`

HasStrikePrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


