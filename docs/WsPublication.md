# WsPublication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** | Publication payload | [optional] 
**Info** | Pointer to [**WsClientInfo**](WsClientInfo.md) |  | [optional] 
**Offset** | Pointer to **int64** | Stream offset for recovery | [optional] 
**Tags** | Pointer to **map[string]string** | Publication tags | [optional] 
**Delta** | Pointer to **bool** | Whether this is a delta update | [optional] 
**Time** | Pointer to **int64** | Publication time in milliseconds | [optional] 
**Channel** | Pointer to **string** | Channel name (for wildcard subscriptions) | [optional] 

## Methods

### NewWsPublication

`func NewWsPublication() *WsPublication`

NewWsPublication instantiates a new WsPublication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsPublicationWithDefaults

`func NewWsPublicationWithDefaults() *WsPublication`

NewWsPublicationWithDefaults instantiates a new WsPublication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WsPublication) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WsPublication) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WsPublication) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *WsPublication) HasData() bool`

HasData returns a boolean if a field has been set.

### GetInfo

`func (o *WsPublication) GetInfo() WsClientInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WsPublication) GetInfoOk() (*WsClientInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WsPublication) SetInfo(v WsClientInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WsPublication) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetOffset

`func (o *WsPublication) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *WsPublication) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *WsPublication) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *WsPublication) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTags

`func (o *WsPublication) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WsPublication) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WsPublication) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WsPublication) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDelta

`func (o *WsPublication) GetDelta() bool`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *WsPublication) GetDeltaOk() (*bool, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *WsPublication) SetDelta(v bool)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *WsPublication) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### GetTime

`func (o *WsPublication) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *WsPublication) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *WsPublication) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *WsPublication) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetChannel

`func (o *WsPublication) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *WsPublication) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *WsPublication) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *WsPublication) HasChannel() bool`

HasChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


