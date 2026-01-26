# WsPresenceResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Presence** | Pointer to [**map[string]WsClientInfo**](WsClientInfo.md) | Map of client ID to client info | [optional] 

## Methods

### NewWsPresenceResult

`func NewWsPresenceResult() *WsPresenceResult`

NewWsPresenceResult instantiates a new WsPresenceResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsPresenceResultWithDefaults

`func NewWsPresenceResultWithDefaults() *WsPresenceResult`

NewWsPresenceResultWithDefaults instantiates a new WsPresenceResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPresence

`func (o *WsPresenceResult) GetPresence() map[string]WsClientInfo`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *WsPresenceResult) GetPresenceOk() (*map[string]WsClientInfo, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *WsPresenceResult) SetPresence(v map[string]WsClientInfo)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *WsPresenceResult) HasPresence() bool`

HasPresence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


