# WsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Command ID matching the request (0 for pushes) | [optional] 
**Error** | Pointer to [**WsError**](WsError.md) |  | [optional] 
**Push** | Pointer to [**WsPush**](WsPush.md) |  | [optional] 
**Connect** | Pointer to [**WsConnectResult**](WsConnectResult.md) |  | [optional] 
**Subscribe** | Pointer to [**WsSubscribeResult**](WsSubscribeResult.md) |  | [optional] 
**Unsubscribe** | Pointer to **map[string]interface{}** | Unsubscribe result (empty on success) | [optional] 
**Publish** | Pointer to **map[string]interface{}** | Publish result (empty on success) | [optional] 
**Presence** | Pointer to [**WsPresenceResult**](WsPresenceResult.md) |  | [optional] 
**PresenceStats** | Pointer to [**WsPresenceStatsResult**](WsPresenceStatsResult.md) |  | [optional] 
**History** | Pointer to [**WsHistoryResult**](WsHistoryResult.md) |  | [optional] 
**Ping** | Pointer to **map[string]interface{}** | Pong response | [optional] 
**Rpc** | Pointer to [**WsRPCResult**](WsRPCResult.md) |  | [optional] 
**Refresh** | Pointer to [**WsRefreshResult**](WsRefreshResult.md) |  | [optional] 
**SubRefresh** | Pointer to [**WsSubRefreshResult**](WsSubRefreshResult.md) |  | [optional] 

## Methods

### NewWsReply

`func NewWsReply() *WsReply`

NewWsReply instantiates a new WsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsReplyWithDefaults

`func NewWsReplyWithDefaults() *WsReply`

NewWsReplyWithDefaults instantiates a new WsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WsReply) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WsReply) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WsReply) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WsReply) HasId() bool`

HasId returns a boolean if a field has been set.

### GetError

`func (o *WsReply) GetError() WsError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WsReply) GetErrorOk() (*WsError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WsReply) SetError(v WsError)`

SetError sets Error field to given value.

### HasError

`func (o *WsReply) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPush

`func (o *WsReply) GetPush() WsPush`

GetPush returns the Push field if non-nil, zero value otherwise.

### GetPushOk

`func (o *WsReply) GetPushOk() (*WsPush, bool)`

GetPushOk returns a tuple with the Push field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPush

`func (o *WsReply) SetPush(v WsPush)`

SetPush sets Push field to given value.

### HasPush

`func (o *WsReply) HasPush() bool`

HasPush returns a boolean if a field has been set.

### GetConnect

`func (o *WsReply) GetConnect() WsConnectResult`

GetConnect returns the Connect field if non-nil, zero value otherwise.

### GetConnectOk

`func (o *WsReply) GetConnectOk() (*WsConnectResult, bool)`

GetConnectOk returns a tuple with the Connect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnect

`func (o *WsReply) SetConnect(v WsConnectResult)`

SetConnect sets Connect field to given value.

### HasConnect

`func (o *WsReply) HasConnect() bool`

HasConnect returns a boolean if a field has been set.

### GetSubscribe

`func (o *WsReply) GetSubscribe() WsSubscribeResult`

GetSubscribe returns the Subscribe field if non-nil, zero value otherwise.

### GetSubscribeOk

`func (o *WsReply) GetSubscribeOk() (*WsSubscribeResult, bool)`

GetSubscribeOk returns a tuple with the Subscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribe

`func (o *WsReply) SetSubscribe(v WsSubscribeResult)`

SetSubscribe sets Subscribe field to given value.

### HasSubscribe

`func (o *WsReply) HasSubscribe() bool`

HasSubscribe returns a boolean if a field has been set.

### GetUnsubscribe

`func (o *WsReply) GetUnsubscribe() map[string]interface{}`

GetUnsubscribe returns the Unsubscribe field if non-nil, zero value otherwise.

### GetUnsubscribeOk

`func (o *WsReply) GetUnsubscribeOk() (*map[string]interface{}, bool)`

GetUnsubscribeOk returns a tuple with the Unsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribe

`func (o *WsReply) SetUnsubscribe(v map[string]interface{})`

SetUnsubscribe sets Unsubscribe field to given value.

### HasUnsubscribe

`func (o *WsReply) HasUnsubscribe() bool`

HasUnsubscribe returns a boolean if a field has been set.

### GetPublish

`func (o *WsReply) GetPublish() map[string]interface{}`

GetPublish returns the Publish field if non-nil, zero value otherwise.

### GetPublishOk

`func (o *WsReply) GetPublishOk() (*map[string]interface{}, bool)`

GetPublishOk returns a tuple with the Publish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublish

`func (o *WsReply) SetPublish(v map[string]interface{})`

SetPublish sets Publish field to given value.

### HasPublish

`func (o *WsReply) HasPublish() bool`

HasPublish returns a boolean if a field has been set.

### GetPresence

`func (o *WsReply) GetPresence() WsPresenceResult`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *WsReply) GetPresenceOk() (*WsPresenceResult, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *WsReply) SetPresence(v WsPresenceResult)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *WsReply) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetPresenceStats

`func (o *WsReply) GetPresenceStats() WsPresenceStatsResult`

GetPresenceStats returns the PresenceStats field if non-nil, zero value otherwise.

### GetPresenceStatsOk

`func (o *WsReply) GetPresenceStatsOk() (*WsPresenceStatsResult, bool)`

GetPresenceStatsOk returns a tuple with the PresenceStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceStats

`func (o *WsReply) SetPresenceStats(v WsPresenceStatsResult)`

SetPresenceStats sets PresenceStats field to given value.

### HasPresenceStats

`func (o *WsReply) HasPresenceStats() bool`

HasPresenceStats returns a boolean if a field has been set.

### GetHistory

`func (o *WsReply) GetHistory() WsHistoryResult`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *WsReply) GetHistoryOk() (*WsHistoryResult, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *WsReply) SetHistory(v WsHistoryResult)`

SetHistory sets History field to given value.

### HasHistory

`func (o *WsReply) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetPing

`func (o *WsReply) GetPing() map[string]interface{}`

GetPing returns the Ping field if non-nil, zero value otherwise.

### GetPingOk

`func (o *WsReply) GetPingOk() (*map[string]interface{}, bool)`

GetPingOk returns a tuple with the Ping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPing

`func (o *WsReply) SetPing(v map[string]interface{})`

SetPing sets Ping field to given value.

### HasPing

`func (o *WsReply) HasPing() bool`

HasPing returns a boolean if a field has been set.

### GetRpc

`func (o *WsReply) GetRpc() WsRPCResult`

GetRpc returns the Rpc field if non-nil, zero value otherwise.

### GetRpcOk

`func (o *WsReply) GetRpcOk() (*WsRPCResult, bool)`

GetRpcOk returns a tuple with the Rpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpc

`func (o *WsReply) SetRpc(v WsRPCResult)`

SetRpc sets Rpc field to given value.

### HasRpc

`func (o *WsReply) HasRpc() bool`

HasRpc returns a boolean if a field has been set.

### GetRefresh

`func (o *WsReply) GetRefresh() WsRefreshResult`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *WsReply) GetRefreshOk() (*WsRefreshResult, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *WsReply) SetRefresh(v WsRefreshResult)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *WsReply) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetSubRefresh

`func (o *WsReply) GetSubRefresh() WsSubRefreshResult`

GetSubRefresh returns the SubRefresh field if non-nil, zero value otherwise.

### GetSubRefreshOk

`func (o *WsReply) GetSubRefreshOk() (*WsSubRefreshResult, bool)`

GetSubRefreshOk returns a tuple with the SubRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRefresh

`func (o *WsReply) SetSubRefresh(v WsSubRefreshResult)`

SetSubRefresh sets SubRefresh field to given value.

### HasSubRefresh

`func (o *WsReply) HasSubRefresh() bool`

HasSubRefresh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


