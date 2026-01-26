# RpcOrderBookLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **string** | Price level | [optional] 
**Quantity** | Pointer to **string** | Quantity at this price | [optional] 

## Methods

### NewRpcOrderBookLevel

`func NewRpcOrderBookLevel() *RpcOrderBookLevel`

NewRpcOrderBookLevel instantiates a new RpcOrderBookLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcOrderBookLevelWithDefaults

`func NewRpcOrderBookLevelWithDefaults() *RpcOrderBookLevel`

NewRpcOrderBookLevelWithDefaults instantiates a new RpcOrderBookLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *RpcOrderBookLevel) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *RpcOrderBookLevel) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *RpcOrderBookLevel) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *RpcOrderBookLevel) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *RpcOrderBookLevel) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *RpcOrderBookLevel) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *RpcOrderBookLevel) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *RpcOrderBookLevel) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


