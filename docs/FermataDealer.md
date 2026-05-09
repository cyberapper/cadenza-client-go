# FermataDealer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DealerAccountId** | **string** | UUID string | 
**Name** | **string** | Human-readable dealer name | 
**Status** | [**DealerStatus**](DealerStatus.md) |  | 
**BaseCurrencies** | **[]string** | Base currencies the dealer settles in | 
**RiskThreshold** | Pointer to **string** | Positive decimal value as string | [optional] 
**LinkedAccountIds** | Pointer to **[]string** | Trading account IDs of exchange accounts linked as liquidity providers | [optional] 
**Config** | Pointer to **map[string]interface{}** | Additional dealer configuration (spreads, fees, etc.) | [optional] 
**CreatedAt** | **int64** | Unix timestamp in milliseconds | 
**CreatedAtDateTime** | Pointer to **time.Time** | Creation timestamp in ISO 8601 format | [optional] 
**UpdatedAt** | **int64** | Unix timestamp in milliseconds | 
**UpdatedAtDateTime** | Pointer to **time.Time** | Last update timestamp in ISO 8601 format | [optional] 

## Methods

### NewFermataDealer

`func NewFermataDealer(dealerAccountId string, name string, status DealerStatus, baseCurrencies []string, createdAt int64, updatedAt int64, ) *FermataDealer`

NewFermataDealer instantiates a new FermataDealer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFermataDealerWithDefaults

`func NewFermataDealerWithDefaults() *FermataDealer`

NewFermataDealerWithDefaults instantiates a new FermataDealer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDealerAccountId

`func (o *FermataDealer) GetDealerAccountId() string`

GetDealerAccountId returns the DealerAccountId field if non-nil, zero value otherwise.

### GetDealerAccountIdOk

`func (o *FermataDealer) GetDealerAccountIdOk() (*string, bool)`

GetDealerAccountIdOk returns a tuple with the DealerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealerAccountId

`func (o *FermataDealer) SetDealerAccountId(v string)`

SetDealerAccountId sets DealerAccountId field to given value.


### GetName

`func (o *FermataDealer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FermataDealer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FermataDealer) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *FermataDealer) GetStatus() DealerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FermataDealer) GetStatusOk() (*DealerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FermataDealer) SetStatus(v DealerStatus)`

SetStatus sets Status field to given value.


### GetBaseCurrencies

`func (o *FermataDealer) GetBaseCurrencies() []string`

GetBaseCurrencies returns the BaseCurrencies field if non-nil, zero value otherwise.

### GetBaseCurrenciesOk

`func (o *FermataDealer) GetBaseCurrenciesOk() (*[]string, bool)`

GetBaseCurrenciesOk returns a tuple with the BaseCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrencies

`func (o *FermataDealer) SetBaseCurrencies(v []string)`

SetBaseCurrencies sets BaseCurrencies field to given value.


### GetRiskThreshold

`func (o *FermataDealer) GetRiskThreshold() string`

GetRiskThreshold returns the RiskThreshold field if non-nil, zero value otherwise.

### GetRiskThresholdOk

`func (o *FermataDealer) GetRiskThresholdOk() (*string, bool)`

GetRiskThresholdOk returns a tuple with the RiskThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskThreshold

`func (o *FermataDealer) SetRiskThreshold(v string)`

SetRiskThreshold sets RiskThreshold field to given value.

### HasRiskThreshold

`func (o *FermataDealer) HasRiskThreshold() bool`

HasRiskThreshold returns a boolean if a field has been set.

### GetLinkedAccountIds

`func (o *FermataDealer) GetLinkedAccountIds() []string`

GetLinkedAccountIds returns the LinkedAccountIds field if non-nil, zero value otherwise.

### GetLinkedAccountIdsOk

`func (o *FermataDealer) GetLinkedAccountIdsOk() (*[]string, bool)`

GetLinkedAccountIdsOk returns a tuple with the LinkedAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountIds

`func (o *FermataDealer) SetLinkedAccountIds(v []string)`

SetLinkedAccountIds sets LinkedAccountIds field to given value.

### HasLinkedAccountIds

`func (o *FermataDealer) HasLinkedAccountIds() bool`

HasLinkedAccountIds returns a boolean if a field has been set.

### GetConfig

`func (o *FermataDealer) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *FermataDealer) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *FermataDealer) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *FermataDealer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FermataDealer) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FermataDealer) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FermataDealer) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedAtDateTime

`func (o *FermataDealer) GetCreatedAtDateTime() time.Time`

GetCreatedAtDateTime returns the CreatedAtDateTime field if non-nil, zero value otherwise.

### GetCreatedAtDateTimeOk

`func (o *FermataDealer) GetCreatedAtDateTimeOk() (*time.Time, bool)`

GetCreatedAtDateTimeOk returns a tuple with the CreatedAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtDateTime

`func (o *FermataDealer) SetCreatedAtDateTime(v time.Time)`

SetCreatedAtDateTime sets CreatedAtDateTime field to given value.

### HasCreatedAtDateTime

`func (o *FermataDealer) HasCreatedAtDateTime() bool`

HasCreatedAtDateTime returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FermataDealer) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FermataDealer) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FermataDealer) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpdatedAtDateTime

`func (o *FermataDealer) GetUpdatedAtDateTime() time.Time`

GetUpdatedAtDateTime returns the UpdatedAtDateTime field if non-nil, zero value otherwise.

### GetUpdatedAtDateTimeOk

`func (o *FermataDealer) GetUpdatedAtDateTimeOk() (*time.Time, bool)`

GetUpdatedAtDateTimeOk returns a tuple with the UpdatedAtDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtDateTime

`func (o *FermataDealer) SetUpdatedAtDateTime(v time.Time)`

SetUpdatedAtDateTime sets UpdatedAtDateTime field to given value.

### HasUpdatedAtDateTime

`func (o *FermataDealer) HasUpdatedAtDateTime() bool`

HasUpdatedAtDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


