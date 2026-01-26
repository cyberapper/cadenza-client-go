# RpcSubmitTradeOrderResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**RpcTradeOrder**](RpcTradeOrder.md) |  | [optional] 
**Error** | Pointer to [**RpcError**](RpcError.md) |  | [optional] 

## Methods

### NewRpcSubmitTradeOrderResult

`func NewRpcSubmitTradeOrderResult() *RpcSubmitTradeOrderResult`

NewRpcSubmitTradeOrderResult instantiates a new RpcSubmitTradeOrderResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcSubmitTradeOrderResultWithDefaults

`func NewRpcSubmitTradeOrderResultWithDefaults() *RpcSubmitTradeOrderResult`

NewRpcSubmitTradeOrderResultWithDefaults instantiates a new RpcSubmitTradeOrderResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RpcSubmitTradeOrderResult) GetData() RpcTradeOrder`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RpcSubmitTradeOrderResult) GetDataOk() (*RpcTradeOrder, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RpcSubmitTradeOrderResult) SetData(v RpcTradeOrder)`

SetData sets Data field to given value.

### HasData

`func (o *RpcSubmitTradeOrderResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *RpcSubmitTradeOrderResult) GetError() RpcError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RpcSubmitTradeOrderResult) GetErrorOk() (*RpcError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RpcSubmitTradeOrderResult) SetError(v RpcError)`

SetError sets Error field to given value.

### HasError

`func (o *RpcSubmitTradeOrderResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


