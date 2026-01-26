# RpcSecurityQuantity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to **string** | Asset symbol | [optional] 
**Quantity** | Pointer to **string** | Quantity (decimal string) | [optional] 

## Methods

### NewRpcSecurityQuantity

`func NewRpcSecurityQuantity() *RpcSecurityQuantity`

NewRpcSecurityQuantity instantiates a new RpcSecurityQuantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcSecurityQuantityWithDefaults

`func NewRpcSecurityQuantityWithDefaults() *RpcSecurityQuantity`

NewRpcSecurityQuantityWithDefaults instantiates a new RpcSecurityQuantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *RpcSecurityQuantity) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *RpcSecurityQuantity) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *RpcSecurityQuantity) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *RpcSecurityQuantity) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetQuantity

`func (o *RpcSecurityQuantity) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *RpcSecurityQuantity) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *RpcSecurityQuantity) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *RpcSecurityQuantity) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


