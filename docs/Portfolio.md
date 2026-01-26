# Portfolio

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradingAccountId** | **string** | UUID string | 
**Venue** | [**Venue**](Venue.md) |  | 
**Positions** | [**[]PositionEntry**](PositionEntry.md) |  | 
**Balances** | [**[]BalanceEntry**](BalanceEntry.md) |  | 
**Summary** | [**PortfolioSummary**](PortfolioSummary.md) |  | 
**UpdatedAt** | **int64** | Unix timestamp in milliseconds | 
**UpdatedAtDateTime** | Pointer to **time.Time** | Last update timestamp in ISO 8601 format | [optional] 

## Methods

### NewPortfolio

`func NewPortfolio(tradingAccountId string, venue Venue, positions []PositionEntry, balances []BalanceEntry, summary PortfolioSummary, updatedAt int64, ) *Portfolio`

NewPortfolio instantiates a new Portfolio object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortfolioWithDefaults

`func NewPortfolioWithDefaults() *Portfolio`

NewPortfolioWithDefaults instantiates a new Portfolio object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradingAccountId

`func (o *Portfolio) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *Portfolio) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *Portfolio) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.


### GetVenue

`func (o *Portfolio) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *Portfolio) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *Portfolio) SetVenue(v Venue)`

SetVenue sets Venue field to given value.


### GetPositions

`func (o *Portfolio) GetPositions() []PositionEntry`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *Portfolio) GetPositionsOk() (*[]PositionEntry, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *Portfolio) SetPositions(v []PositionEntry)`

SetPositions sets Positions field to given value.


### GetBalances

`func (o *Portfolio) GetBalances() []BalanceEntry`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *Portfolio) GetBalancesOk() (*[]BalanceEntry, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *Portfolio) SetBalances(v []BalanceEntry)`

SetBalances sets Balances field to given value.


### GetSummary

`func (o *Portfolio) GetSummary() PortfolioSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Portfolio) GetSummaryOk() (*PortfolioSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Portfolio) SetSummary(v PortfolioSummary)`

SetSummary sets Summary field to given value.


### GetUpdatedAt

`func (o *Portfolio) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Portfolio) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Portfolio) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpdatedAtDateTime

`func (o *Portfolio) GetUpdatedAtDateTime() time.Time`

GetUpdatedAtDateTime returns the UpdatedAtDateTime field if non-nil, zero value otherwise.

### GetUpdatedAtDateTimeOk

`func (o *Portfolio) GetUpdatedAtDateTimeOk() (*time.Time, bool)`

GetUpdatedAtDateTimeOk returns a tuple with the UpdatedAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtDateTime

`func (o *Portfolio) SetUpdatedAtDateTime(v time.Time)`

SetUpdatedAtDateTime sets UpdatedAtDateTime field to given value.

### HasUpdatedAtDateTime

`func (o *Portfolio) HasUpdatedAtDateTime() bool`

HasUpdatedAtDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


