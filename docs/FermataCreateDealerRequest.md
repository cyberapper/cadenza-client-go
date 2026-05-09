# FermataCreateDealerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable dealer name | 
**BaseCurrencies** | **[]string** | Base currencies the dealer settles in | 
**RiskThreshold** | Pointer to **string** | Positive decimal value as string | [optional] 
**Config** | Pointer to **map[string]interface{}** | Additional dealer configuration | [optional] 

## Methods

### NewFermataCreateDealerRequest

`func NewFermataCreateDealerRequest(name string, baseCurrencies []string, ) *FermataCreateDealerRequest`

NewFermataCreateDealerRequest instantiates a new FermataCreateDealerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFermataCreateDealerRequestWithDefaults

`func NewFermataCreateDealerRequestWithDefaults() *FermataCreateDealerRequest`

NewFermataCreateDealerRequestWithDefaults instantiates a new FermataCreateDealerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FermataCreateDealerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FermataCreateDealerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FermataCreateDealerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetBaseCurrencies

`func (o *FermataCreateDealerRequest) GetBaseCurrencies() []string`

GetBaseCurrencies returns the BaseCurrencies field if non-nil, zero value otherwise.

### GetBaseCurrenciesOk

`func (o *FermataCreateDealerRequest) GetBaseCurrenciesOk() (*[]string, bool)`

GetBaseCurrenciesOk returns a tuple with the BaseCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrencies

`func (o *FermataCreateDealerRequest) SetBaseCurrencies(v []string)`

SetBaseCurrencies sets BaseCurrencies field to given value.


### GetRiskThreshold

`func (o *FermataCreateDealerRequest) GetRiskThreshold() string`

GetRiskThreshold returns the RiskThreshold field if non-nil, zero value otherwise.

### GetRiskThresholdOk

`func (o *FermataCreateDealerRequest) GetRiskThresholdOk() (*string, bool)`

GetRiskThresholdOk returns a tuple with the RiskThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskThreshold

`func (o *FermataCreateDealerRequest) SetRiskThreshold(v string)`

SetRiskThreshold sets RiskThreshold field to given value.

### HasRiskThreshold

`func (o *FermataCreateDealerRequest) HasRiskThreshold() bool`

HasRiskThreshold returns a boolean if a field has been set.

### GetConfig

`func (o *FermataCreateDealerRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *FermataCreateDealerRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *FermataCreateDealerRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *FermataCreateDealerRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


