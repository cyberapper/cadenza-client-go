# RpcSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityId** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**ExternalSymbol** | Pointer to **string** |  | [optional] 
**Venue** | Pointer to [**Venue**](Venue.md) |  | [optional] 

## Methods

### NewRpcSecurity

`func NewRpcSecurity() *RpcSecurity`

NewRpcSecurity instantiates a new RpcSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcSecurityWithDefaults

`func NewRpcSecurityWithDefaults() *RpcSecurity`

NewRpcSecurityWithDefaults instantiates a new RpcSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityId

`func (o *RpcSecurity) GetSecurityId() string`

GetSecurityId returns the SecurityId field if non-nil, zero value otherwise.

### GetSecurityIdOk

`func (o *RpcSecurity) GetSecurityIdOk() (*string, bool)`

GetSecurityIdOk returns a tuple with the SecurityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityId

`func (o *RpcSecurity) SetSecurityId(v string)`

SetSecurityId sets SecurityId field to given value.

### HasSecurityId

`func (o *RpcSecurity) HasSecurityId() bool`

HasSecurityId returns a boolean if a field has been set.

### GetSymbol

`func (o *RpcSecurity) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *RpcSecurity) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *RpcSecurity) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *RpcSecurity) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetExternalSymbol

`func (o *RpcSecurity) GetExternalSymbol() string`

GetExternalSymbol returns the ExternalSymbol field if non-nil, zero value otherwise.

### GetExternalSymbolOk

`func (o *RpcSecurity) GetExternalSymbolOk() (*string, bool)`

GetExternalSymbolOk returns a tuple with the ExternalSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSymbol

`func (o *RpcSecurity) SetExternalSymbol(v string)`

SetExternalSymbol sets ExternalSymbol field to given value.

### HasExternalSymbol

`func (o *RpcSecurity) HasExternalSymbol() bool`

HasExternalSymbol returns a boolean if a field has been set.

### GetVenue

`func (o *RpcSecurity) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *RpcSecurity) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *RpcSecurity) SetVenue(v Venue)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *RpcSecurity) HasVenue() bool`

HasVenue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


