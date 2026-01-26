# SecurityQuantity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Security** | **string** | Security ID in format venue:symbol (e.g., BINANCE:BTC) | 
**Asset** | **string** | Asset symbol (e.g., BTC, USDT, BNB) | 
**Venue** | [**Venue**](Venue.md) |  | 
**Quantity** | **string** | Decimal value as string to preserve precision | 

## Methods

### NewSecurityQuantity

`func NewSecurityQuantity(security string, asset string, venue Venue, quantity string, ) *SecurityQuantity`

NewSecurityQuantity instantiates a new SecurityQuantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityQuantityWithDefaults

`func NewSecurityQuantityWithDefaults() *SecurityQuantity`

NewSecurityQuantityWithDefaults instantiates a new SecurityQuantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurity

`func (o *SecurityQuantity) GetSecurity() string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *SecurityQuantity) GetSecurityOk() (*string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *SecurityQuantity) SetSecurity(v string)`

SetSecurity sets Security field to given value.


### GetAsset

`func (o *SecurityQuantity) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *SecurityQuantity) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *SecurityQuantity) SetAsset(v string)`

SetAsset sets Asset field to given value.


### GetVenue

`func (o *SecurityQuantity) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *SecurityQuantity) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *SecurityQuantity) SetVenue(v Venue)`

SetVenue sets Venue field to given value.


### GetQuantity

`func (o *SecurityQuantity) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SecurityQuantity) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SecurityQuantity) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


