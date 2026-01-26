# WsUnsubscribePush

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | Unsubscribe reason code | [optional] 
**Reason** | Pointer to **string** | Unsubscribe reason message | [optional] 

## Methods

### NewWsUnsubscribePush

`func NewWsUnsubscribePush() *WsUnsubscribePush`

NewWsUnsubscribePush instantiates a new WsUnsubscribePush object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsUnsubscribePushWithDefaults

`func NewWsUnsubscribePushWithDefaults() *WsUnsubscribePush`

NewWsUnsubscribePushWithDefaults instantiates a new WsUnsubscribePush object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *WsUnsubscribePush) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *WsUnsubscribePush) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *WsUnsubscribePush) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *WsUnsubscribePush) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetReason

`func (o *WsUnsubscribePush) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WsUnsubscribePush) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WsUnsubscribePush) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *WsUnsubscribePush) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


