# FermataTransferWalletRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceWalletId** | **string** | UUID string | 
**DestWalletId** | **string** | UUID string | 
**Asset** | **string** | Asset to transfer | 
**Amount** | **string** | Positive decimal value as string | 

## Methods

### NewFermataTransferWalletRequest

`func NewFermataTransferWalletRequest(sourceWalletId string, destWalletId string, asset string, amount string, ) *FermataTransferWalletRequest`

NewFermataTransferWalletRequest instantiates a new FermataTransferWalletRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFermataTransferWalletRequestWithDefaults

`func NewFermataTransferWalletRequestWithDefaults() *FermataTransferWalletRequest`

NewFermataTransferWalletRequestWithDefaults instantiates a new FermataTransferWalletRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceWalletId

`func (o *FermataTransferWalletRequest) GetSourceWalletId() string`

GetSourceWalletId returns the SourceWalletId field if non-nil, zero value otherwise.

### GetSourceWalletIdOk

`func (o *FermataTransferWalletRequest) GetSourceWalletIdOk() (*string, bool)`

GetSourceWalletIdOk returns a tuple with the SourceWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceWalletId

`func (o *FermataTransferWalletRequest) SetSourceWalletId(v string)`

SetSourceWalletId sets SourceWalletId field to given value.


### GetDestWalletId

`func (o *FermataTransferWalletRequest) GetDestWalletId() string`

GetDestWalletId returns the DestWalletId field if non-nil, zero value otherwise.

### GetDestWalletIdOk

`func (o *FermataTransferWalletRequest) GetDestWalletIdOk() (*string, bool)`

GetDestWalletIdOk returns a tuple with the DestWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestWalletId

`func (o *FermataTransferWalletRequest) SetDestWalletId(v string)`

SetDestWalletId sets DestWalletId field to given value.


### GetAsset

`func (o *FermataTransferWalletRequest) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *FermataTransferWalletRequest) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *FermataTransferWalletRequest) SetAsset(v string)`

SetAsset sets Asset field to given value.


### GetAmount

`func (o *FermataTransferWalletRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FermataTransferWalletRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FermataTransferWalletRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


