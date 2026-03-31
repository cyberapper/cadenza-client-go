# Quote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteId** | **string** | UUID string | 
**QuoteRequestId** | Pointer to **string** | Client-provided reference for idempotency and tracking | [optional] 
**ExternalQuoteId** | Pointer to **string** | External quote identifier from upstream venue | [optional] 
**Venue** | Pointer to [**Venue**](Venue.md) |  | [optional] 
**InstrumentId** | Pointer to **string** | Instrument ID in format {VENUE}:{BASE}/{QUOTE} | [optional] 
**DealerAccountId** | **string** | UUID string | 
**TraderAccountId** | Pointer to **string** | UUID string | [optional] 
**BaseAsset** | **string** | Base asset being traded | 
**QuoteAsset** | **string** | Quote asset (payment currency) | 
**OrderSide** | [**OrderSide**](OrderSide.md) |  | 
**QuantityType** | Pointer to [**QuantityType**](QuantityType.md) |  | [optional] 
**Quantity** | **string** | Positive decimal value as string | 
**QuoteQuantity** | **string** | Positive decimal value as string | 
**Price** | **string** | Positive decimal value as string | 
**Status** | [**QuoteStatus**](QuoteStatus.md) |  | 
**RejectReason** | Pointer to **string** | Reason for quote rejection, if status is REJECTED | [optional] 
**CreatedAt** | **int64** | Unix timestamp in milliseconds | 
**CreatedAtDateTime** | Pointer to **time.Time** | Quote creation timestamp in ISO 8601 format | [optional] 
**UpdatedAt** | Pointer to **int64** | Unix timestamp in milliseconds | [optional] 
**UpdatedAtDateTime** | Pointer to **time.Time** | Last update timestamp in ISO 8601 format | [optional] 
**ExpireAt** | **int64** | Unix timestamp in milliseconds | 
**ExpireAtDateTime** | Pointer to **time.Time** | Quote expiration timestamp in ISO 8601 format | [optional] 
**PricingProfileId** | Pointer to **string** | UUID string | [optional] 

## Methods

### NewQuote

`func NewQuote(quoteId string, dealerAccountId string, baseAsset string, quoteAsset string, orderSide OrderSide, quantity string, quoteQuantity string, price string, status QuoteStatus, createdAt int64, expireAt int64, ) *Quote`

NewQuote instantiates a new Quote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteWithDefaults

`func NewQuoteWithDefaults() *Quote`

NewQuoteWithDefaults instantiates a new Quote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteId

`func (o *Quote) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *Quote) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *Quote) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetQuoteRequestId

`func (o *Quote) GetQuoteRequestId() string`

GetQuoteRequestId returns the QuoteRequestId field if non-nil, zero value otherwise.

### GetQuoteRequestIdOk

`func (o *Quote) GetQuoteRequestIdOk() (*string, bool)`

GetQuoteRequestIdOk returns a tuple with the QuoteRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteRequestId

`func (o *Quote) SetQuoteRequestId(v string)`

SetQuoteRequestId sets QuoteRequestId field to given value.

### HasQuoteRequestId

`func (o *Quote) HasQuoteRequestId() bool`

HasQuoteRequestId returns a boolean if a field has been set.

### GetExternalQuoteId

`func (o *Quote) GetExternalQuoteId() string`

GetExternalQuoteId returns the ExternalQuoteId field if non-nil, zero value otherwise.

### GetExternalQuoteIdOk

`func (o *Quote) GetExternalQuoteIdOk() (*string, bool)`

GetExternalQuoteIdOk returns a tuple with the ExternalQuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalQuoteId

`func (o *Quote) SetExternalQuoteId(v string)`

SetExternalQuoteId sets ExternalQuoteId field to given value.

### HasExternalQuoteId

`func (o *Quote) HasExternalQuoteId() bool`

HasExternalQuoteId returns a boolean if a field has been set.

### GetVenue

`func (o *Quote) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *Quote) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *Quote) SetVenue(v Venue)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *Quote) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetInstrumentId

`func (o *Quote) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *Quote) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *Quote) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *Quote) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetDealerAccountId

`func (o *Quote) GetDealerAccountId() string`

GetDealerAccountId returns the DealerAccountId field if non-nil, zero value otherwise.

### GetDealerAccountIdOk

`func (o *Quote) GetDealerAccountIdOk() (*string, bool)`

GetDealerAccountIdOk returns a tuple with the DealerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealerAccountId

`func (o *Quote) SetDealerAccountId(v string)`

SetDealerAccountId sets DealerAccountId field to given value.


### GetTraderAccountId

`func (o *Quote) GetTraderAccountId() string`

GetTraderAccountId returns the TraderAccountId field if non-nil, zero value otherwise.

### GetTraderAccountIdOk

`func (o *Quote) GetTraderAccountIdOk() (*string, bool)`

GetTraderAccountIdOk returns a tuple with the TraderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraderAccountId

`func (o *Quote) SetTraderAccountId(v string)`

SetTraderAccountId sets TraderAccountId field to given value.

### HasTraderAccountId

`func (o *Quote) HasTraderAccountId() bool`

HasTraderAccountId returns a boolean if a field has been set.

### GetBaseAsset

`func (o *Quote) GetBaseAsset() string`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *Quote) GetBaseAssetOk() (*string, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *Quote) SetBaseAsset(v string)`

SetBaseAsset sets BaseAsset field to given value.


### GetQuoteAsset

`func (o *Quote) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *Quote) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *Quote) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.


### GetOrderSide

`func (o *Quote) GetOrderSide() OrderSide`

GetOrderSide returns the OrderSide field if non-nil, zero value otherwise.

### GetOrderSideOk

`func (o *Quote) GetOrderSideOk() (*OrderSide, bool)`

GetOrderSideOk returns a tuple with the OrderSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSide

`func (o *Quote) SetOrderSide(v OrderSide)`

SetOrderSide sets OrderSide field to given value.


### GetQuantityType

`func (o *Quote) GetQuantityType() QuantityType`

GetQuantityType returns the QuantityType field if non-nil, zero value otherwise.

### GetQuantityTypeOk

`func (o *Quote) GetQuantityTypeOk() (*QuantityType, bool)`

GetQuantityTypeOk returns a tuple with the QuantityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityType

`func (o *Quote) SetQuantityType(v QuantityType)`

SetQuantityType sets QuantityType field to given value.

### HasQuantityType

`func (o *Quote) HasQuantityType() bool`

HasQuantityType returns a boolean if a field has been set.

### GetQuantity

`func (o *Quote) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Quote) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Quote) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetQuoteQuantity

`func (o *Quote) GetQuoteQuantity() string`

GetQuoteQuantity returns the QuoteQuantity field if non-nil, zero value otherwise.

### GetQuoteQuantityOk

`func (o *Quote) GetQuoteQuantityOk() (*string, bool)`

GetQuoteQuantityOk returns a tuple with the QuoteQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteQuantity

`func (o *Quote) SetQuoteQuantity(v string)`

SetQuoteQuantity sets QuoteQuantity field to given value.


### GetPrice

`func (o *Quote) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Quote) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Quote) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetStatus

`func (o *Quote) GetStatus() QuoteStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Quote) GetStatusOk() (*QuoteStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Quote) SetStatus(v QuoteStatus)`

SetStatus sets Status field to given value.


### GetRejectReason

`func (o *Quote) GetRejectReason() string`

GetRejectReason returns the RejectReason field if non-nil, zero value otherwise.

### GetRejectReasonOk

`func (o *Quote) GetRejectReasonOk() (*string, bool)`

GetRejectReasonOk returns a tuple with the RejectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectReason

`func (o *Quote) SetRejectReason(v string)`

SetRejectReason sets RejectReason field to given value.

### HasRejectReason

`func (o *Quote) HasRejectReason() bool`

HasRejectReason returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Quote) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Quote) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Quote) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedAtDateTime

`func (o *Quote) GetCreatedAtDateTime() time.Time`

GetCreatedAtDateTime returns the CreatedAtDateTime field if non-nil, zero value otherwise.

### GetCreatedAtDateTimeOk

`func (o *Quote) GetCreatedAtDateTimeOk() (*time.Time, bool)`

GetCreatedAtDateTimeOk returns a tuple with the CreatedAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtDateTime

`func (o *Quote) SetCreatedAtDateTime(v time.Time)`

SetCreatedAtDateTime sets CreatedAtDateTime field to given value.

### HasCreatedAtDateTime

`func (o *Quote) HasCreatedAtDateTime() bool`

HasCreatedAtDateTime returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Quote) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Quote) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Quote) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Quote) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedAtDateTime

`func (o *Quote) GetUpdatedAtDateTime() time.Time`

GetUpdatedAtDateTime returns the UpdatedAtDateTime field if non-nil, zero value otherwise.

### GetUpdatedAtDateTimeOk

`func (o *Quote) GetUpdatedAtDateTimeOk() (*time.Time, bool)`

GetUpdatedAtDateTimeOk returns a tuple with the UpdatedAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtDateTime

`func (o *Quote) SetUpdatedAtDateTime(v time.Time)`

SetUpdatedAtDateTime sets UpdatedAtDateTime field to given value.

### HasUpdatedAtDateTime

`func (o *Quote) HasUpdatedAtDateTime() bool`

HasUpdatedAtDateTime returns a boolean if a field has been set.

### GetExpireAt

`func (o *Quote) GetExpireAt() int64`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *Quote) GetExpireAtOk() (*int64, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *Quote) SetExpireAt(v int64)`

SetExpireAt sets ExpireAt field to given value.


### GetExpireAtDateTime

`func (o *Quote) GetExpireAtDateTime() time.Time`

GetExpireAtDateTime returns the ExpireAtDateTime field if non-nil, zero value otherwise.

### GetExpireAtDateTimeOk

`func (o *Quote) GetExpireAtDateTimeOk() (*time.Time, bool)`

GetExpireAtDateTimeOk returns a tuple with the ExpireAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAtDateTime

`func (o *Quote) SetExpireAtDateTime(v time.Time)`

SetExpireAtDateTime sets ExpireAtDateTime field to given value.

### HasExpireAtDateTime

`func (o *Quote) HasExpireAtDateTime() bool`

HasExpireAtDateTime returns a boolean if a field has been set.

### GetPricingProfileId

`func (o *Quote) GetPricingProfileId() string`

GetPricingProfileId returns the PricingProfileId field if non-nil, zero value otherwise.

### GetPricingProfileIdOk

`func (o *Quote) GetPricingProfileIdOk() (*string, bool)`

GetPricingProfileIdOk returns a tuple with the PricingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingProfileId

`func (o *Quote) SetPricingProfileId(v string)`

SetPricingProfileId sets PricingProfileId field to given value.

### HasPricingProfileId

`func (o *Quote) HasPricingProfileId() bool`

HasPricingProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


