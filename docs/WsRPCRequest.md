# WsRPCRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | [**RpcMethod**](RpcMethod.md) |  | 
**Data** | Pointer to [**WsRPCRequestData**](WsRPCRequestData.md) |  | [optional] 

## Methods

### NewWsRPCRequest

`func NewWsRPCRequest(method RpcMethod, ) *WsRPCRequest`

NewWsRPCRequest instantiates a new WsRPCRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsRPCRequestWithDefaults

`func NewWsRPCRequestWithDefaults() *WsRPCRequest`

NewWsRPCRequestWithDefaults instantiates a new WsRPCRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *WsRPCRequest) GetMethod() RpcMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *WsRPCRequest) GetMethodOk() (*RpcMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *WsRPCRequest) SetMethod(v RpcMethod)`

SetMethod sets Method field to given value.


### GetData

`func (o *WsRPCRequest) GetData() WsRPCRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WsRPCRequest) GetDataOk() (*WsRPCRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WsRPCRequest) SetData(v WsRPCRequestData)`

SetData sets Data field to given value.

### HasData

`func (o *WsRPCRequest) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


