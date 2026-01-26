# RpcSyncSecuritiesParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Venue** | [**Venue**](Venue.md) |  | 
**Securities** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRpcSyncSecuritiesParams

`func NewRpcSyncSecuritiesParams(venue Venue, ) *RpcSyncSecuritiesParams`

NewRpcSyncSecuritiesParams instantiates a new RpcSyncSecuritiesParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcSyncSecuritiesParamsWithDefaults

`func NewRpcSyncSecuritiesParamsWithDefaults() *RpcSyncSecuritiesParams`

NewRpcSyncSecuritiesParamsWithDefaults instantiates a new RpcSyncSecuritiesParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVenue

`func (o *RpcSyncSecuritiesParams) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *RpcSyncSecuritiesParams) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *RpcSyncSecuritiesParams) SetVenue(v Venue)`

SetVenue sets Venue field to given value.


### GetSecurities

`func (o *RpcSyncSecuritiesParams) GetSecurities() []string`

GetSecurities returns the Securities field if non-nil, zero value otherwise.

### GetSecuritiesOk

`func (o *RpcSyncSecuritiesParams) GetSecuritiesOk() (*[]string, bool)`

GetSecuritiesOk returns a tuple with the Securities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurities

`func (o *RpcSyncSecuritiesParams) SetSecurities(v []string)`

SetSecurities sets Securities field to given value.

### HasSecurities

`func (o *RpcSyncSecuritiesParams) HasSecurities() bool`

HasSecurities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


