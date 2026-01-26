# WsPresenceStatsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumClients** | Pointer to **int32** | Number of connected clients | [optional] 
**NumUsers** | Pointer to **int32** | Number of unique users | [optional] 

## Methods

### NewWsPresenceStatsResult

`func NewWsPresenceStatsResult() *WsPresenceStatsResult`

NewWsPresenceStatsResult instantiates a new WsPresenceStatsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsPresenceStatsResultWithDefaults

`func NewWsPresenceStatsResultWithDefaults() *WsPresenceStatsResult`

NewWsPresenceStatsResultWithDefaults instantiates a new WsPresenceStatsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumClients

`func (o *WsPresenceStatsResult) GetNumClients() int32`

GetNumClients returns the NumClients field if non-nil, zero value otherwise.

### GetNumClientsOk

`func (o *WsPresenceStatsResult) GetNumClientsOk() (*int32, bool)`

GetNumClientsOk returns a tuple with the NumClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumClients

`func (o *WsPresenceStatsResult) SetNumClients(v int32)`

SetNumClients sets NumClients field to given value.

### HasNumClients

`func (o *WsPresenceStatsResult) HasNumClients() bool`

HasNumClients returns a boolean if a field has been set.

### GetNumUsers

`func (o *WsPresenceStatsResult) GetNumUsers() int32`

GetNumUsers returns the NumUsers field if non-nil, zero value otherwise.

### GetNumUsersOk

`func (o *WsPresenceStatsResult) GetNumUsersOk() (*int32, bool)`

GetNumUsersOk returns a tuple with the NumUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUsers

`func (o *WsPresenceStatsResult) SetNumUsers(v int32)`

SetNumUsers sets NumUsers field to given value.

### HasNumUsers

`func (o *WsPresenceStatsResult) HasNumUsers() bool`

HasNumUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


