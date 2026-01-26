# WsPush

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel the push is for | [optional] 
**Pub** | Pointer to [**WsPublication**](WsPublication.md) |  | [optional] 
**Join** | Pointer to [**WsJoin**](WsJoin.md) |  | [optional] 
**Leave** | Pointer to [**WsLeave**](WsLeave.md) |  | [optional] 
**Unsubscribe** | Pointer to [**WsUnsubscribePush**](WsUnsubscribePush.md) |  | [optional] 
**Message** | Pointer to [**WsMessage**](WsMessage.md) |  | [optional] 
**Subscribe** | Pointer to [**WsSubscribePush**](WsSubscribePush.md) |  | [optional] 
**Connect** | Pointer to [**WsConnectPush**](WsConnectPush.md) |  | [optional] 
**Disconnect** | Pointer to [**WsDisconnect**](WsDisconnect.md) |  | [optional] 
**Refresh** | Pointer to [**WsRefreshPush**](WsRefreshPush.md) |  | [optional] 

## Methods

### NewWsPush

`func NewWsPush() *WsPush`

NewWsPush instantiates a new WsPush object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsPushWithDefaults

`func NewWsPushWithDefaults() *WsPush`

NewWsPushWithDefaults instantiates a new WsPush object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *WsPush) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *WsPush) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *WsPush) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *WsPush) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetPub

`func (o *WsPush) GetPub() WsPublication`

GetPub returns the Pub field if non-nil, zero value otherwise.

### GetPubOk

`func (o *WsPush) GetPubOk() (*WsPublication, bool)`

GetPubOk returns a tuple with the Pub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPub

`func (o *WsPush) SetPub(v WsPublication)`

SetPub sets Pub field to given value.

### HasPub

`func (o *WsPush) HasPub() bool`

HasPub returns a boolean if a field has been set.

### GetJoin

`func (o *WsPush) GetJoin() WsJoin`

GetJoin returns the Join field if non-nil, zero value otherwise.

### GetJoinOk

`func (o *WsPush) GetJoinOk() (*WsJoin, bool)`

GetJoinOk returns a tuple with the Join field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoin

`func (o *WsPush) SetJoin(v WsJoin)`

SetJoin sets Join field to given value.

### HasJoin

`func (o *WsPush) HasJoin() bool`

HasJoin returns a boolean if a field has been set.

### GetLeave

`func (o *WsPush) GetLeave() WsLeave`

GetLeave returns the Leave field if non-nil, zero value otherwise.

### GetLeaveOk

`func (o *WsPush) GetLeaveOk() (*WsLeave, bool)`

GetLeaveOk returns a tuple with the Leave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeave

`func (o *WsPush) SetLeave(v WsLeave)`

SetLeave sets Leave field to given value.

### HasLeave

`func (o *WsPush) HasLeave() bool`

HasLeave returns a boolean if a field has been set.

### GetUnsubscribe

`func (o *WsPush) GetUnsubscribe() WsUnsubscribePush`

GetUnsubscribe returns the Unsubscribe field if non-nil, zero value otherwise.

### GetUnsubscribeOk

`func (o *WsPush) GetUnsubscribeOk() (*WsUnsubscribePush, bool)`

GetUnsubscribeOk returns a tuple with the Unsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribe

`func (o *WsPush) SetUnsubscribe(v WsUnsubscribePush)`

SetUnsubscribe sets Unsubscribe field to given value.

### HasUnsubscribe

`func (o *WsPush) HasUnsubscribe() bool`

HasUnsubscribe returns a boolean if a field has been set.

### GetMessage

`func (o *WsPush) GetMessage() WsMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WsPush) GetMessageOk() (*WsMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WsPush) SetMessage(v WsMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WsPush) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSubscribe

`func (o *WsPush) GetSubscribe() WsSubscribePush`

GetSubscribe returns the Subscribe field if non-nil, zero value otherwise.

### GetSubscribeOk

`func (o *WsPush) GetSubscribeOk() (*WsSubscribePush, bool)`

GetSubscribeOk returns a tuple with the Subscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribe

`func (o *WsPush) SetSubscribe(v WsSubscribePush)`

SetSubscribe sets Subscribe field to given value.

### HasSubscribe

`func (o *WsPush) HasSubscribe() bool`

HasSubscribe returns a boolean if a field has been set.

### GetConnect

`func (o *WsPush) GetConnect() WsConnectPush`

GetConnect returns the Connect field if non-nil, zero value otherwise.

### GetConnectOk

`func (o *WsPush) GetConnectOk() (*WsConnectPush, bool)`

GetConnectOk returns a tuple with the Connect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnect

`func (o *WsPush) SetConnect(v WsConnectPush)`

SetConnect sets Connect field to given value.

### HasConnect

`func (o *WsPush) HasConnect() bool`

HasConnect returns a boolean if a field has been set.

### GetDisconnect

`func (o *WsPush) GetDisconnect() WsDisconnect`

GetDisconnect returns the Disconnect field if non-nil, zero value otherwise.

### GetDisconnectOk

`func (o *WsPush) GetDisconnectOk() (*WsDisconnect, bool)`

GetDisconnectOk returns a tuple with the Disconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnect

`func (o *WsPush) SetDisconnect(v WsDisconnect)`

SetDisconnect sets Disconnect field to given value.

### HasDisconnect

`func (o *WsPush) HasDisconnect() bool`

HasDisconnect returns a boolean if a field has been set.

### GetRefresh

`func (o *WsPush) GetRefresh() WsRefreshPush`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *WsPush) GetRefreshOk() (*WsRefreshPush, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *WsPush) SetRefresh(v WsRefreshPush)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *WsPush) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


