# WsDisconnect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | Disconnect reason code | [optional] 
**Reason** | Pointer to **string** | Disconnect reason message | [optional] 
**Reconnect** | Pointer to **bool** | Whether client should attempt to reconnect | [optional] 

## Methods

### NewWsDisconnect

`func NewWsDisconnect() *WsDisconnect`

NewWsDisconnect instantiates a new WsDisconnect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsDisconnectWithDefaults

`func NewWsDisconnectWithDefaults() *WsDisconnect`

NewWsDisconnectWithDefaults instantiates a new WsDisconnect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *WsDisconnect) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *WsDisconnect) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *WsDisconnect) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *WsDisconnect) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetReason

`func (o *WsDisconnect) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WsDisconnect) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WsDisconnect) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *WsDisconnect) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReconnect

`func (o *WsDisconnect) GetReconnect() bool`

GetReconnect returns the Reconnect field if non-nil, zero value otherwise.

### GetReconnectOk

`func (o *WsDisconnect) GetReconnectOk() (*bool, bool)`

GetReconnectOk returns a tuple with the Reconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconnect

`func (o *WsDisconnect) SetReconnect(v bool)`

SetReconnect sets Reconnect field to given value.

### HasReconnect

`func (o *WsDisconnect) HasReconnect() bool`

HasReconnect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


