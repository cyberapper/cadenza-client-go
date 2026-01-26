# FinancialSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityId** | Pointer to **string** | Security ID, id in the format of venue:symbol | [optional] 
**Symbol** | Pointer to **string** | Symbol | [optional] 
**Venue** | Pointer to [**Venue**](Venue.md) |  | [optional] 
**SecurityType** | Pointer to [**SecurityType**](SecurityType.md) |  | [optional] 
**Precision** | Pointer to **int32** | Precision | [optional] 
**Scale** | Pointer to **int32** | Scale | [optional] 
**MinQuantity** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**LotSize** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 

## Methods

### NewFinancialSecurity

`func NewFinancialSecurity() *FinancialSecurity`

NewFinancialSecurity instantiates a new FinancialSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialSecurityWithDefaults

`func NewFinancialSecurityWithDefaults() *FinancialSecurity`

NewFinancialSecurityWithDefaults instantiates a new FinancialSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityId

`func (o *FinancialSecurity) GetSecurityId() string`

GetSecurityId returns the SecurityId field if non-nil, zero value otherwise.

### GetSecurityIdOk

`func (o *FinancialSecurity) GetSecurityIdOk() (*string, bool)`

GetSecurityIdOk returns a tuple with the SecurityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityId

`func (o *FinancialSecurity) SetSecurityId(v string)`

SetSecurityId sets SecurityId field to given value.

### HasSecurityId

`func (o *FinancialSecurity) HasSecurityId() bool`

HasSecurityId returns a boolean if a field has been set.

### GetSymbol

`func (o *FinancialSecurity) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *FinancialSecurity) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *FinancialSecurity) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *FinancialSecurity) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetVenue

`func (o *FinancialSecurity) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *FinancialSecurity) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *FinancialSecurity) SetVenue(v Venue)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *FinancialSecurity) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetSecurityType

`func (o *FinancialSecurity) GetSecurityType() SecurityType`

GetSecurityType returns the SecurityType field if non-nil, zero value otherwise.

### GetSecurityTypeOk

`func (o *FinancialSecurity) GetSecurityTypeOk() (*SecurityType, bool)`

GetSecurityTypeOk returns a tuple with the SecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityType

`func (o *FinancialSecurity) SetSecurityType(v SecurityType)`

SetSecurityType sets SecurityType field to given value.

### HasSecurityType

`func (o *FinancialSecurity) HasSecurityType() bool`

HasSecurityType returns a boolean if a field has been set.

### GetPrecision

`func (o *FinancialSecurity) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *FinancialSecurity) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *FinancialSecurity) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *FinancialSecurity) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetScale

`func (o *FinancialSecurity) GetScale() int32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *FinancialSecurity) GetScaleOk() (*int32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *FinancialSecurity) SetScale(v int32)`

SetScale sets Scale field to given value.

### HasScale

`func (o *FinancialSecurity) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetMinQuantity

`func (o *FinancialSecurity) GetMinQuantity() string`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *FinancialSecurity) GetMinQuantityOk() (*string, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *FinancialSecurity) SetMinQuantity(v string)`

SetMinQuantity sets MinQuantity field to given value.

### HasMinQuantity

`func (o *FinancialSecurity) HasMinQuantity() bool`

HasMinQuantity returns a boolean if a field has been set.

### GetLotSize

`func (o *FinancialSecurity) GetLotSize() string`

GetLotSize returns the LotSize field if non-nil, zero value otherwise.

### GetLotSizeOk

`func (o *FinancialSecurity) GetLotSizeOk() (*string, bool)`

GetLotSizeOk returns a tuple with the LotSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotSize

`func (o *FinancialSecurity) SetLotSize(v string)`

SetLotSize sets LotSize field to given value.

### HasLotSize

`func (o *FinancialSecurity) HasLotSize() bool`

HasLotSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


