# QuoteRfqRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DealerAccountId** | **string** | UUID string | 
**TraderAccountId** | Pointer to **string** | UUID string | [optional] 
**InstrumentId** | Pointer to **string** | Instrument ID in format {VENUE}:{BASE}/{QUOTE} | [optional] 
**BaseAsset** | Pointer to **string** | Base asset to trade. Used with quoteAsset for symbol-based venues. | [optional] 
**QuoteAsset** | Pointer to **string** | Quote asset (payment currency). Used with baseAsset for symbol-based venues. | [optional] 
**OrderSide** | [**OrderSide**](OrderSide.md) |  | 
**Quantity** | Pointer to **string** | Positive decimal value as string | [optional] 
**QuoteQuantity** | Pointer to **string** | Positive decimal value as string | [optional] 
**QuoteRequestId** | Pointer to **string** | Client-provided reference for idempotency and tracking | [optional] 

## Methods

### NewQuoteRfqRequest

`func NewQuoteRfqRequest(dealerAccountId string, orderSide OrderSide, ) *QuoteRfqRequest`

NewQuoteRfqRequest instantiates a new QuoteRfqRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteRfqRequestWithDefaults

`func NewQuoteRfqRequestWithDefaults() *QuoteRfqRequest`

NewQuoteRfqRequestWithDefaults instantiates a new QuoteRfqRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDealerAccountId

`func (o *QuoteRfqRequest) GetDealerAccountId() string`

GetDealerAccountId returns the DealerAccountId field if non-nil, zero value otherwise.

### GetDealerAccountIdOk

`func (o *QuoteRfqRequest) GetDealerAccountIdOk() (*string, bool)`

GetDealerAccountIdOk returns a tuple with the DealerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealerAccountId

`func (o *QuoteRfqRequest) SetDealerAccountId(v string)`

SetDealerAccountId sets DealerAccountId field to given value.


### GetTraderAccountId

`func (o *QuoteRfqRequest) GetTraderAccountId() string`

GetTraderAccountId returns the TraderAccountId field if non-nil, zero value otherwise.

### GetTraderAccountIdOk

`func (o *QuoteRfqRequest) GetTraderAccountIdOk() (*string, bool)`

GetTraderAccountIdOk returns a tuple with the TraderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraderAccountId

`func (o *QuoteRfqRequest) SetTraderAccountId(v string)`

SetTraderAccountId sets TraderAccountId field to given value.

### HasTraderAccountId

`func (o *QuoteRfqRequest) HasTraderAccountId() bool`

HasTraderAccountId returns a boolean if a field has been set.

### GetInstrumentId

`func (o *QuoteRfqRequest) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *QuoteRfqRequest) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *QuoteRfqRequest) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *QuoteRfqRequest) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetBaseAsset

`func (o *QuoteRfqRequest) GetBaseAsset() string`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *QuoteRfqRequest) GetBaseAssetOk() (*string, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *QuoteRfqRequest) SetBaseAsset(v string)`

SetBaseAsset sets BaseAsset field to given value.

### HasBaseAsset

`func (o *QuoteRfqRequest) HasBaseAsset() bool`

HasBaseAsset returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *QuoteRfqRequest) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *QuoteRfqRequest) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *QuoteRfqRequest) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *QuoteRfqRequest) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetOrderSide

`func (o *QuoteRfqRequest) GetOrderSide() OrderSide`

GetOrderSide returns the OrderSide field if non-nil, zero value otherwise.

### GetOrderSideOk

`func (o *QuoteRfqRequest) GetOrderSideOk() (*OrderSide, bool)`

GetOrderSideOk returns a tuple with the OrderSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSide

`func (o *QuoteRfqRequest) SetOrderSide(v OrderSide)`

SetOrderSide sets OrderSide field to given value.


### GetQuantity

`func (o *QuoteRfqRequest) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuoteRfqRequest) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuoteRfqRequest) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QuoteRfqRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetQuoteQuantity

`func (o *QuoteRfqRequest) GetQuoteQuantity() string`

GetQuoteQuantity returns the QuoteQuantity field if non-nil, zero value otherwise.

### GetQuoteQuantityOk

`func (o *QuoteRfqRequest) GetQuoteQuantityOk() (*string, bool)`

GetQuoteQuantityOk returns a tuple with the QuoteQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteQuantity

`func (o *QuoteRfqRequest) SetQuoteQuantity(v string)`

SetQuoteQuantity sets QuoteQuantity field to given value.

### HasQuoteQuantity

`func (o *QuoteRfqRequest) HasQuoteQuantity() bool`

HasQuoteQuantity returns a boolean if a field has been set.

### GetQuoteRequestId

`func (o *QuoteRfqRequest) GetQuoteRequestId() string`

GetQuoteRequestId returns the QuoteRequestId field if non-nil, zero value otherwise.

### GetQuoteRequestIdOk

`func (o *QuoteRfqRequest) GetQuoteRequestIdOk() (*string, bool)`

GetQuoteRequestIdOk returns a tuple with the QuoteRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteRequestId

`func (o *QuoteRfqRequest) SetQuoteRequestId(v string)`

SetQuoteRequestId sets QuoteRequestId field to given value.

### HasQuoteRequestId

`func (o *QuoteRfqRequest) HasQuoteRequestId() bool`

HasQuoteRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


