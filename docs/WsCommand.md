# WsCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique command ID for request-response correlation | 
**Connect** | Pointer to [**WsConnectRequest**](WsConnectRequest.md) |  | [optional] 
**Subscribe** | Pointer to [**WsSubscribeRequest**](WsSubscribeRequest.md) |  | [optional] 
**Unsubscribe** | Pointer to [**WsUnsubscribeRequest**](WsUnsubscribeRequest.md) |  | [optional] 
**Publish** | Pointer to [**WsPublishRequest**](WsPublishRequest.md) |  | [optional] 
**Presence** | Pointer to [**WsPresenceRequest**](WsPresenceRequest.md) |  | [optional] 
**PresenceStats** | Pointer to [**WsPresenceStatsRequest**](WsPresenceStatsRequest.md) |  | [optional] 
**History** | Pointer to [**WsHistoryRequest**](WsHistoryRequest.md) |  | [optional] 
**Ping** | Pointer to **map[string]interface{}** | Ping request to keep connection alive | [optional] 
**Send** | Pointer to [**WsSendRequest**](WsSendRequest.md) |  | [optional] 
**Rpc** | Pointer to [**WsRPCRequest**](WsRPCRequest.md) |  | [optional] 
**Refresh** | Pointer to [**WsRefreshRequest**](WsRefreshRequest.md) |  | [optional] 
**SubRefresh** | Pointer to [**WsSubRefreshRequest**](WsSubRefreshRequest.md) |  | [optional] 

## Methods

### NewWsCommand

`func NewWsCommand(id int32, ) *WsCommand`

NewWsCommand instantiates a new WsCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsCommandWithDefaults

`func NewWsCommandWithDefaults() *WsCommand`

NewWsCommandWithDefaults instantiates a new WsCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WsCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WsCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WsCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetConnect

`func (o *WsCommand) GetConnect() WsConnectRequest`

GetConnect returns the Connect field if non-nil, zero value otherwise.

### GetConnectOk

`func (o *WsCommand) GetConnectOk() (*WsConnectRequest, bool)`

GetConnectOk returns a tuple with the Connect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnect

`func (o *WsCommand) SetConnect(v WsConnectRequest)`

SetConnect sets Connect field to given value.

### HasConnect

`func (o *WsCommand) HasConnect() bool`

HasConnect returns a boolean if a field has been set.

### GetSubscribe

`func (o *WsCommand) GetSubscribe() WsSubscribeRequest`

GetSubscribe returns the Subscribe field if non-nil, zero value otherwise.

### GetSubscribeOk

`func (o *WsCommand) GetSubscribeOk() (*WsSubscribeRequest, bool)`

GetSubscribeOk returns a tuple with the Subscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribe

`func (o *WsCommand) SetSubscribe(v WsSubscribeRequest)`

SetSubscribe sets Subscribe field to given value.

### HasSubscribe

`func (o *WsCommand) HasSubscribe() bool`

HasSubscribe returns a boolean if a field has been set.

### GetUnsubscribe

`func (o *WsCommand) GetUnsubscribe() WsUnsubscribeRequest`

GetUnsubscribe returns the Unsubscribe field if non-nil, zero value otherwise.

### GetUnsubscribeOk

`func (o *WsCommand) GetUnsubscribeOk() (*WsUnsubscribeRequest, bool)`

GetUnsubscribeOk returns a tuple with the Unsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribe

`func (o *WsCommand) SetUnsubscribe(v WsUnsubscribeRequest)`

SetUnsubscribe sets Unsubscribe field to given value.

### HasUnsubscribe

`func (o *WsCommand) HasUnsubscribe() bool`

HasUnsubscribe returns a boolean if a field has been set.

### GetPublish

`func (o *WsCommand) GetPublish() WsPublishRequest`

GetPublish returns the Publish field if non-nil, zero value otherwise.

### GetPublishOk

`func (o *WsCommand) GetPublishOk() (*WsPublishRequest, bool)`

GetPublishOk returns a tuple with the Publish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublish

`func (o *WsCommand) SetPublish(v WsPublishRequest)`

SetPublish sets Publish field to given value.

### HasPublish

`func (o *WsCommand) HasPublish() bool`

HasPublish returns a boolean if a field has been set.

### GetPresence

`func (o *WsCommand) GetPresence() WsPresenceRequest`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *WsCommand) GetPresenceOk() (*WsPresenceRequest, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *WsCommand) SetPresence(v WsPresenceRequest)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *WsCommand) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetPresenceStats

`func (o *WsCommand) GetPresenceStats() WsPresenceStatsRequest`

GetPresenceStats returns the PresenceStats field if non-nil, zero value otherwise.

### GetPresenceStatsOk

`func (o *WsCommand) GetPresenceStatsOk() (*WsPresenceStatsRequest, bool)`

GetPresenceStatsOk returns a tuple with the PresenceStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceStats

`func (o *WsCommand) SetPresenceStats(v WsPresenceStatsRequest)`

SetPresenceStats sets PresenceStats field to given value.

### HasPresenceStats

`func (o *WsCommand) HasPresenceStats() bool`

HasPresenceStats returns a boolean if a field has been set.

### GetHistory

`func (o *WsCommand) GetHistory() WsHistoryRequest`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *WsCommand) GetHistoryOk() (*WsHistoryRequest, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *WsCommand) SetHistory(v WsHistoryRequest)`

SetHistory sets History field to given value.

### HasHistory

`func (o *WsCommand) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetPing

`func (o *WsCommand) GetPing() map[string]interface{}`

GetPing returns the Ping field if non-nil, zero value otherwise.

### GetPingOk

`func (o *WsCommand) GetPingOk() (*map[string]interface{}, bool)`

GetPingOk returns a tuple with the Ping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPing

`func (o *WsCommand) SetPing(v map[string]interface{})`

SetPing sets Ping field to given value.

### HasPing

`func (o *WsCommand) HasPing() bool`

HasPing returns a boolean if a field has been set.

### GetSend

`func (o *WsCommand) GetSend() WsSendRequest`

GetSend returns the Send field if non-nil, zero value otherwise.

### GetSendOk

`func (o *WsCommand) GetSendOk() (*WsSendRequest, bool)`

GetSendOk returns a tuple with the Send field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSend

`func (o *WsCommand) SetSend(v WsSendRequest)`

SetSend sets Send field to given value.

### HasSend

`func (o *WsCommand) HasSend() bool`

HasSend returns a boolean if a field has been set.

### GetRpc

`func (o *WsCommand) GetRpc() WsRPCRequest`

GetRpc returns the Rpc field if non-nil, zero value otherwise.

### GetRpcOk

`func (o *WsCommand) GetRpcOk() (*WsRPCRequest, bool)`

GetRpcOk returns a tuple with the Rpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpc

`func (o *WsCommand) SetRpc(v WsRPCRequest)`

SetRpc sets Rpc field to given value.

### HasRpc

`func (o *WsCommand) HasRpc() bool`

HasRpc returns a boolean if a field has been set.

### GetRefresh

`func (o *WsCommand) GetRefresh() WsRefreshRequest`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *WsCommand) GetRefreshOk() (*WsRefreshRequest, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *WsCommand) SetRefresh(v WsRefreshRequest)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *WsCommand) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetSubRefresh

`func (o *WsCommand) GetSubRefresh() WsSubRefreshRequest`

GetSubRefresh returns the SubRefresh field if non-nil, zero value otherwise.

### GetSubRefreshOk

`func (o *WsCommand) GetSubRefreshOk() (*WsSubRefreshRequest, bool)`

GetSubRefreshOk returns a tuple with the SubRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRefresh

`func (o *WsCommand) SetSubRefresh(v WsSubRefreshRequest)`

SetSubRefresh sets SubRefresh field to given value.

### HasSubRefresh

`func (o *WsCommand) HasSubRefresh() bool`

HasSubRefresh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


