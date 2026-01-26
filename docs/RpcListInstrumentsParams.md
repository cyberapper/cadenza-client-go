# RpcListInstrumentsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Venue** | Pointer to [**Venue**](Venue.md) |  | [optional] 
**Symbols** | Pointer to **[]string** |  | [optional] 
**SecurityType** | Pointer to **string** |  | [optional] 
**InstrumentStatus** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 

## Methods

### NewRpcListInstrumentsParams

`func NewRpcListInstrumentsParams() *RpcListInstrumentsParams`

NewRpcListInstrumentsParams instantiates a new RpcListInstrumentsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcListInstrumentsParamsWithDefaults

`func NewRpcListInstrumentsParamsWithDefaults() *RpcListInstrumentsParams`

NewRpcListInstrumentsParamsWithDefaults instantiates a new RpcListInstrumentsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVenue

`func (o *RpcListInstrumentsParams) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *RpcListInstrumentsParams) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *RpcListInstrumentsParams) SetVenue(v Venue)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *RpcListInstrumentsParams) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetSymbols

`func (o *RpcListInstrumentsParams) GetSymbols() []string`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *RpcListInstrumentsParams) GetSymbolsOk() (*[]string, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *RpcListInstrumentsParams) SetSymbols(v []string)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *RpcListInstrumentsParams) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.

### GetSecurityType

`func (o *RpcListInstrumentsParams) GetSecurityType() string`

GetSecurityType returns the SecurityType field if non-nil, zero value otherwise.

### GetSecurityTypeOk

`func (o *RpcListInstrumentsParams) GetSecurityTypeOk() (*string, bool)`

GetSecurityTypeOk returns a tuple with the SecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityType

`func (o *RpcListInstrumentsParams) SetSecurityType(v string)`

SetSecurityType sets SecurityType field to given value.

### HasSecurityType

`func (o *RpcListInstrumentsParams) HasSecurityType() bool`

HasSecurityType returns a boolean if a field has been set.

### GetInstrumentStatus

`func (o *RpcListInstrumentsParams) GetInstrumentStatus() string`

GetInstrumentStatus returns the InstrumentStatus field if non-nil, zero value otherwise.

### GetInstrumentStatusOk

`func (o *RpcListInstrumentsParams) GetInstrumentStatusOk() (*string, bool)`

GetInstrumentStatusOk returns a tuple with the InstrumentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentStatus

`func (o *RpcListInstrumentsParams) SetInstrumentStatus(v string)`

SetInstrumentStatus sets InstrumentStatus field to given value.

### HasInstrumentStatus

`func (o *RpcListInstrumentsParams) HasInstrumentStatus() bool`

HasInstrumentStatus returns a boolean if a field has been set.

### GetLimit

`func (o *RpcListInstrumentsParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RpcListInstrumentsParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RpcListInstrumentsParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *RpcListInstrumentsParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *RpcListInstrumentsParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *RpcListInstrumentsParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *RpcListInstrumentsParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *RpcListInstrumentsParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


