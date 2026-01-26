# RpcListSecuritiesParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Venue** | Pointer to [**Venue**](Venue.md) |  | [optional] 
**Securities** | Pointer to **[]string** |  | [optional] 
**ExternalSymbols** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRpcListSecuritiesParams

`func NewRpcListSecuritiesParams() *RpcListSecuritiesParams`

NewRpcListSecuritiesParams instantiates a new RpcListSecuritiesParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcListSecuritiesParamsWithDefaults

`func NewRpcListSecuritiesParamsWithDefaults() *RpcListSecuritiesParams`

NewRpcListSecuritiesParamsWithDefaults instantiates a new RpcListSecuritiesParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVenue

`func (o *RpcListSecuritiesParams) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *RpcListSecuritiesParams) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *RpcListSecuritiesParams) SetVenue(v Venue)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *RpcListSecuritiesParams) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetSecurities

`func (o *RpcListSecuritiesParams) GetSecurities() []string`

GetSecurities returns the Securities field if non-nil, zero value otherwise.

### GetSecuritiesOk

`func (o *RpcListSecuritiesParams) GetSecuritiesOk() (*[]string, bool)`

GetSecuritiesOk returns a tuple with the Securities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurities

`func (o *RpcListSecuritiesParams) SetSecurities(v []string)`

SetSecurities sets Securities field to given value.

### HasSecurities

`func (o *RpcListSecuritiesParams) HasSecurities() bool`

HasSecurities returns a boolean if a field has been set.

### GetExternalSymbols

`func (o *RpcListSecuritiesParams) GetExternalSymbols() []string`

GetExternalSymbols returns the ExternalSymbols field if non-nil, zero value otherwise.

### GetExternalSymbolsOk

`func (o *RpcListSecuritiesParams) GetExternalSymbolsOk() (*[]string, bool)`

GetExternalSymbolsOk returns a tuple with the ExternalSymbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSymbols

`func (o *RpcListSecuritiesParams) SetExternalSymbols(v []string)`

SetExternalSymbols sets ExternalSymbols field to given value.

### HasExternalSymbols

`func (o *RpcListSecuritiesParams) HasExternalSymbols() bool`

HasExternalSymbols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


