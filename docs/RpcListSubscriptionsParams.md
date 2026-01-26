# RpcListSubscriptionsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstrumentId** | Pointer to **[]string** |  | [optional] 
**SubscriptionType** | Pointer to [**SubscriptionType**](SubscriptionType.md) |  | [optional] 
**Venue** | Pointer to [**Venue**](Venue.md) |  | [optional] 

## Methods

### NewRpcListSubscriptionsParams

`func NewRpcListSubscriptionsParams() *RpcListSubscriptionsParams`

NewRpcListSubscriptionsParams instantiates a new RpcListSubscriptionsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcListSubscriptionsParamsWithDefaults

`func NewRpcListSubscriptionsParamsWithDefaults() *RpcListSubscriptionsParams`

NewRpcListSubscriptionsParamsWithDefaults instantiates a new RpcListSubscriptionsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstrumentId

`func (o *RpcListSubscriptionsParams) GetInstrumentId() []string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *RpcListSubscriptionsParams) GetInstrumentIdOk() (*[]string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *RpcListSubscriptionsParams) SetInstrumentId(v []string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *RpcListSubscriptionsParams) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetSubscriptionType

`func (o *RpcListSubscriptionsParams) GetSubscriptionType() SubscriptionType`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *RpcListSubscriptionsParams) GetSubscriptionTypeOk() (*SubscriptionType, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *RpcListSubscriptionsParams) SetSubscriptionType(v SubscriptionType)`

SetSubscriptionType sets SubscriptionType field to given value.

### HasSubscriptionType

`func (o *RpcListSubscriptionsParams) HasSubscriptionType() bool`

HasSubscriptionType returns a boolean if a field has been set.

### GetVenue

`func (o *RpcListSubscriptionsParams) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *RpcListSubscriptionsParams) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *RpcListSubscriptionsParams) SetVenue(v Venue)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *RpcListSubscriptionsParams) HasVenue() bool`

HasVenue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


