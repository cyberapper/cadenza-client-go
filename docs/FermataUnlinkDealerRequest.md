# FermataUnlinkDealerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DealerAccountId** | **string** | UUID string | 
**TradingAccountId** | **string** | UUID string | 

## Methods

### NewFermataUnlinkDealerRequest

`func NewFermataUnlinkDealerRequest(dealerAccountId string, tradingAccountId string, ) *FermataUnlinkDealerRequest`

NewFermataUnlinkDealerRequest instantiates a new FermataUnlinkDealerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFermataUnlinkDealerRequestWithDefaults

`func NewFermataUnlinkDealerRequestWithDefaults() *FermataUnlinkDealerRequest`

NewFermataUnlinkDealerRequestWithDefaults instantiates a new FermataUnlinkDealerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDealerAccountId

`func (o *FermataUnlinkDealerRequest) GetDealerAccountId() string`

GetDealerAccountId returns the DealerAccountId field if non-nil, zero value otherwise.

### GetDealerAccountIdOk

`func (o *FermataUnlinkDealerRequest) GetDealerAccountIdOk() (*string, bool)`

GetDealerAccountIdOk returns a tuple with the DealerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealerAccountId

`func (o *FermataUnlinkDealerRequest) SetDealerAccountId(v string)`

SetDealerAccountId sets DealerAccountId field to given value.


### GetTradingAccountId

`func (o *FermataUnlinkDealerRequest) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *FermataUnlinkDealerRequest) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *FermataUnlinkDealerRequest) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


