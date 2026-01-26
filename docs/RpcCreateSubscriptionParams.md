# RpcCreateSubscriptionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Venue** | [**Venue**](Venue.md) |  | 
**Instruments** | Pointer to **[]string** |  | [optional] 
**SubscriptionType** | [**SubscriptionType**](SubscriptionType.md) |  | 

## Methods

### NewRpcCreateSubscriptionParams

`func NewRpcCreateSubscriptionParams(venue Venue, subscriptionType SubscriptionType, ) *RpcCreateSubscriptionParams`

NewRpcCreateSubscriptionParams instantiates a new RpcCreateSubscriptionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcCreateSubscriptionParamsWithDefaults

`func NewRpcCreateSubscriptionParamsWithDefaults() *RpcCreateSubscriptionParams`

NewRpcCreateSubscriptionParamsWithDefaults instantiates a new RpcCreateSubscriptionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVenue

`func (o *RpcCreateSubscriptionParams) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *RpcCreateSubscriptionParams) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *RpcCreateSubscriptionParams) SetVenue(v Venue)`

SetVenue sets Venue field to given value.


### GetInstruments

`func (o *RpcCreateSubscriptionParams) GetInstruments() []string`

GetInstruments returns the Instruments field if non-nil, zero value otherwise.

### GetInstrumentsOk

`func (o *RpcCreateSubscriptionParams) GetInstrumentsOk() (*[]string, bool)`

GetInstrumentsOk returns a tuple with the Instruments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruments

`func (o *RpcCreateSubscriptionParams) SetInstruments(v []string)`

SetInstruments sets Instruments field to given value.

### HasInstruments

`func (o *RpcCreateSubscriptionParams) HasInstruments() bool`

HasInstruments returns a boolean if a field has been set.

### GetSubscriptionType

`func (o *RpcCreateSubscriptionParams) GetSubscriptionType() SubscriptionType`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *RpcCreateSubscriptionParams) GetSubscriptionTypeOk() (*SubscriptionType, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *RpcCreateSubscriptionParams) SetSubscriptionType(v SubscriptionType)`

SetSubscriptionType sets SubscriptionType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


