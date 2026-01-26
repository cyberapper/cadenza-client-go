# MarketVenue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Venue** | [**Venue**](Venue.md) |  | 
**Status** | [**VenueStatus**](VenueStatus.md) |  | 

## Methods

### NewMarketVenue

`func NewMarketVenue(venue Venue, status VenueStatus, ) *MarketVenue`

NewMarketVenue instantiates a new MarketVenue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketVenueWithDefaults

`func NewMarketVenueWithDefaults() *MarketVenue`

NewMarketVenueWithDefaults instantiates a new MarketVenue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVenue

`func (o *MarketVenue) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *MarketVenue) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *MarketVenue) SetVenue(v Venue)`

SetVenue sets Venue field to given value.


### GetStatus

`func (o *MarketVenue) GetStatus() VenueStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketVenue) GetStatusOk() (*VenueStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketVenue) SetStatus(v VenueStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


