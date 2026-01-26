# WsConnectResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Client** | Pointer to **string** | Unique client ID assigned by server | [optional] 
**Version** | Pointer to **string** | Server version | [optional] 
**Expires** | Pointer to **bool** | Whether the connection expires | [optional] 
**Ttl** | Pointer to **int32** | Time to live in seconds (0 &#x3D; no expiry) | [optional] 
**Data** | Pointer to **map[string]interface{}** | Custom data from server | [optional] 
**Subs** | Pointer to [**map[string]WsSubscribeResult**](WsSubscribeResult.md) | Results of initial subscriptions | [optional] 
**Ping** | Pointer to **int32** | Server ping interval in seconds | [optional] 
**Pong** | Pointer to **bool** | Whether client should send pong responses | [optional] 
**Session** | Pointer to **string** | Session ID | [optional] 
**Node** | Pointer to **string** | Server node ID | [optional] 
**Time** | Pointer to **int64** | Server time in milliseconds | [optional] 

## Methods

### NewWsConnectResult

`func NewWsConnectResult() *WsConnectResult`

NewWsConnectResult instantiates a new WsConnectResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsConnectResultWithDefaults

`func NewWsConnectResultWithDefaults() *WsConnectResult`

NewWsConnectResultWithDefaults instantiates a new WsConnectResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClient

`func (o *WsConnectResult) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *WsConnectResult) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *WsConnectResult) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *WsConnectResult) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetVersion

`func (o *WsConnectResult) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WsConnectResult) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WsConnectResult) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WsConnectResult) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetExpires

`func (o *WsConnectResult) GetExpires() bool`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *WsConnectResult) GetExpiresOk() (*bool, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *WsConnectResult) SetExpires(v bool)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *WsConnectResult) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetTtl

`func (o *WsConnectResult) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *WsConnectResult) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *WsConnectResult) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *WsConnectResult) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetData

`func (o *WsConnectResult) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WsConnectResult) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WsConnectResult) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *WsConnectResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSubs

`func (o *WsConnectResult) GetSubs() map[string]WsSubscribeResult`

GetSubs returns the Subs field if non-nil, zero value otherwise.

### GetSubsOk

`func (o *WsConnectResult) GetSubsOk() (*map[string]WsSubscribeResult, bool)`

GetSubsOk returns a tuple with the Subs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubs

`func (o *WsConnectResult) SetSubs(v map[string]WsSubscribeResult)`

SetSubs sets Subs field to given value.

### HasSubs

`func (o *WsConnectResult) HasSubs() bool`

HasSubs returns a boolean if a field has been set.

### GetPing

`func (o *WsConnectResult) GetPing() int32`

GetPing returns the Ping field if non-nil, zero value otherwise.

### GetPingOk

`func (o *WsConnectResult) GetPingOk() (*int32, bool)`

GetPingOk returns a tuple with the Ping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPing

`func (o *WsConnectResult) SetPing(v int32)`

SetPing sets Ping field to given value.

### HasPing

`func (o *WsConnectResult) HasPing() bool`

HasPing returns a boolean if a field has been set.

### GetPong

`func (o *WsConnectResult) GetPong() bool`

GetPong returns the Pong field if non-nil, zero value otherwise.

### GetPongOk

`func (o *WsConnectResult) GetPongOk() (*bool, bool)`

GetPongOk returns a tuple with the Pong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPong

`func (o *WsConnectResult) SetPong(v bool)`

SetPong sets Pong field to given value.

### HasPong

`func (o *WsConnectResult) HasPong() bool`

HasPong returns a boolean if a field has been set.

### GetSession

`func (o *WsConnectResult) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *WsConnectResult) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *WsConnectResult) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *WsConnectResult) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetNode

`func (o *WsConnectResult) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *WsConnectResult) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *WsConnectResult) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *WsConnectResult) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetTime

`func (o *WsConnectResult) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *WsConnectResult) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *WsConnectResult) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *WsConnectResult) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


