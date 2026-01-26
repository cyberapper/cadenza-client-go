# RpcGetOrderBookResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**RpcOrderBook**](RpcOrderBook.md) |  | [optional] 
**Error** | Pointer to [**RpcError**](RpcError.md) |  | [optional] 

## Methods

### NewRpcGetOrderBookResult

`func NewRpcGetOrderBookResult() *RpcGetOrderBookResult`

NewRpcGetOrderBookResult instantiates a new RpcGetOrderBookResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcGetOrderBookResultWithDefaults

`func NewRpcGetOrderBookResultWithDefaults() *RpcGetOrderBookResult`

NewRpcGetOrderBookResultWithDefaults instantiates a new RpcGetOrderBookResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RpcGetOrderBookResult) GetData() RpcOrderBook`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RpcGetOrderBookResult) GetDataOk() (*RpcOrderBook, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RpcGetOrderBookResult) SetData(v RpcOrderBook)`

SetData sets Data field to given value.

### HasData

`func (o *RpcGetOrderBookResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *RpcGetOrderBookResult) GetError() RpcError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RpcGetOrderBookResult) GetErrorOk() (*RpcError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RpcGetOrderBookResult) SetError(v RpcError)`

SetError sets Error field to given value.

### HasError

`func (o *RpcGetOrderBookResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


