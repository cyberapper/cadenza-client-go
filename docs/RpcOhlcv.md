# RpcOhlcv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**T** | Pointer to **time.Time** | Open time | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**O** | Pointer to **string** | Open price (decimal) | [optional] 
**H** | Pointer to **string** | High price (decimal) | [optional] 
**L** | Pointer to **string** | Low price (decimal) | [optional] 
**C** | Pointer to **string** | Close price (decimal) | [optional] 
**V** | Pointer to **string** | Volume (decimal) | [optional] 

## Methods

### NewRpcOhlcv

`func NewRpcOhlcv() *RpcOhlcv`

NewRpcOhlcv instantiates a new RpcOhlcv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcOhlcvWithDefaults

`func NewRpcOhlcvWithDefaults() *RpcOhlcv`

NewRpcOhlcvWithDefaults instantiates a new RpcOhlcv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetT

`func (o *RpcOhlcv) GetT() time.Time`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *RpcOhlcv) GetTOk() (*time.Time, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *RpcOhlcv) SetT(v time.Time)`

SetT sets T field to given value.

### HasT

`func (o *RpcOhlcv) HasT() bool`

HasT returns a boolean if a field has been set.

### GetInterval

`func (o *RpcOhlcv) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *RpcOhlcv) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *RpcOhlcv) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *RpcOhlcv) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetO

`func (o *RpcOhlcv) GetO() string`

GetO returns the O field if non-nil, zero value otherwise.

### GetOOk

`func (o *RpcOhlcv) GetOOk() (*string, bool)`

GetOOk returns a tuple with the O field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO

`func (o *RpcOhlcv) SetO(v string)`

SetO sets O field to given value.

### HasO

`func (o *RpcOhlcv) HasO() bool`

HasO returns a boolean if a field has been set.

### GetH

`func (o *RpcOhlcv) GetH() string`

GetH returns the H field if non-nil, zero value otherwise.

### GetHOk

`func (o *RpcOhlcv) GetHOk() (*string, bool)`

GetHOk returns a tuple with the H field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH

`func (o *RpcOhlcv) SetH(v string)`

SetH sets H field to given value.

### HasH

`func (o *RpcOhlcv) HasH() bool`

HasH returns a boolean if a field has been set.

### GetL

`func (o *RpcOhlcv) GetL() string`

GetL returns the L field if non-nil, zero value otherwise.

### GetLOk

`func (o *RpcOhlcv) GetLOk() (*string, bool)`

GetLOk returns a tuple with the L field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL

`func (o *RpcOhlcv) SetL(v string)`

SetL sets L field to given value.

### HasL

`func (o *RpcOhlcv) HasL() bool`

HasL returns a boolean if a field has been set.

### GetC

`func (o *RpcOhlcv) GetC() string`

GetC returns the C field if non-nil, zero value otherwise.

### GetCOk

`func (o *RpcOhlcv) GetCOk() (*string, bool)`

GetCOk returns a tuple with the C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC

`func (o *RpcOhlcv) SetC(v string)`

SetC sets C field to given value.

### HasC

`func (o *RpcOhlcv) HasC() bool`

HasC returns a boolean if a field has been set.

### GetV

`func (o *RpcOhlcv) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *RpcOhlcv) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *RpcOhlcv) SetV(v string)`

SetV sets V field to given value.

### HasV

`func (o *RpcOhlcv) HasV() bool`

HasV returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


