# WsSubscribeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | Channel name to subscribe to | 
**Token** | Pointer to **string** | Subscription token (for private channels) | [optional] 
**Recover** | Pointer to **bool** | Whether to recover missed messages | [optional] 
**Epoch** | Pointer to **string** | Stream epoch for recovery | [optional] 
**Offset** | Pointer to **int64** | Stream offset for recovery | [optional] 
**Data** | Pointer to **map[string]interface{}** | Custom subscription data | [optional] 
**Positioned** | Pointer to **bool** | Whether to receive position info in publications | [optional] 
**Recoverable** | Pointer to **bool** | Whether the subscription should be recoverable | [optional] 
**JoinLeave** | Pointer to **bool** | Whether to receive join/leave messages | [optional] 
**Delta** | Pointer to **string** | Delta compression mode | [optional] 

## Methods

### NewWsSubscribeRequest

`func NewWsSubscribeRequest(channel string, ) *WsSubscribeRequest`

NewWsSubscribeRequest instantiates a new WsSubscribeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsSubscribeRequestWithDefaults

`func NewWsSubscribeRequestWithDefaults() *WsSubscribeRequest`

NewWsSubscribeRequestWithDefaults instantiates a new WsSubscribeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *WsSubscribeRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *WsSubscribeRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *WsSubscribeRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetToken

`func (o *WsSubscribeRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *WsSubscribeRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *WsSubscribeRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *WsSubscribeRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetRecover

`func (o *WsSubscribeRequest) GetRecover() bool`

GetRecover returns the Recover field if non-nil, zero value otherwise.

### GetRecoverOk

`func (o *WsSubscribeRequest) GetRecoverOk() (*bool, bool)`

GetRecoverOk returns a tuple with the Recover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecover

`func (o *WsSubscribeRequest) SetRecover(v bool)`

SetRecover sets Recover field to given value.

### HasRecover

`func (o *WsSubscribeRequest) HasRecover() bool`

HasRecover returns a boolean if a field has been set.

### GetEpoch

`func (o *WsSubscribeRequest) GetEpoch() string`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *WsSubscribeRequest) GetEpochOk() (*string, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *WsSubscribeRequest) SetEpoch(v string)`

SetEpoch sets Epoch field to given value.

### HasEpoch

`func (o *WsSubscribeRequest) HasEpoch() bool`

HasEpoch returns a boolean if a field has been set.

### GetOffset

`func (o *WsSubscribeRequest) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *WsSubscribeRequest) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *WsSubscribeRequest) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *WsSubscribeRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetData

`func (o *WsSubscribeRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WsSubscribeRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WsSubscribeRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *WsSubscribeRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPositioned

`func (o *WsSubscribeRequest) GetPositioned() bool`

GetPositioned returns the Positioned field if non-nil, zero value otherwise.

### GetPositionedOk

`func (o *WsSubscribeRequest) GetPositionedOk() (*bool, bool)`

GetPositionedOk returns a tuple with the Positioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositioned

`func (o *WsSubscribeRequest) SetPositioned(v bool)`

SetPositioned sets Positioned field to given value.

### HasPositioned

`func (o *WsSubscribeRequest) HasPositioned() bool`

HasPositioned returns a boolean if a field has been set.

### GetRecoverable

`func (o *WsSubscribeRequest) GetRecoverable() bool`

GetRecoverable returns the Recoverable field if non-nil, zero value otherwise.

### GetRecoverableOk

`func (o *WsSubscribeRequest) GetRecoverableOk() (*bool, bool)`

GetRecoverableOk returns a tuple with the Recoverable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverable

`func (o *WsSubscribeRequest) SetRecoverable(v bool)`

SetRecoverable sets Recoverable field to given value.

### HasRecoverable

`func (o *WsSubscribeRequest) HasRecoverable() bool`

HasRecoverable returns a boolean if a field has been set.

### GetJoinLeave

`func (o *WsSubscribeRequest) GetJoinLeave() bool`

GetJoinLeave returns the JoinLeave field if non-nil, zero value otherwise.

### GetJoinLeaveOk

`func (o *WsSubscribeRequest) GetJoinLeaveOk() (*bool, bool)`

GetJoinLeaveOk returns a tuple with the JoinLeave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinLeave

`func (o *WsSubscribeRequest) SetJoinLeave(v bool)`

SetJoinLeave sets JoinLeave field to given value.

### HasJoinLeave

`func (o *WsSubscribeRequest) HasJoinLeave() bool`

HasJoinLeave returns a boolean if a field has been set.

### GetDelta

`func (o *WsSubscribeRequest) GetDelta() string`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *WsSubscribeRequest) GetDeltaOk() (*string, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *WsSubscribeRequest) SetDelta(v string)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *WsSubscribeRequest) HasDelta() bool`

HasDelta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


