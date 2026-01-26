# WsSubscribePush

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recoverable** | Pointer to **bool** |  | [optional] 
**Epoch** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Positioned** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWsSubscribePush

`func NewWsSubscribePush() *WsSubscribePush`

NewWsSubscribePush instantiates a new WsSubscribePush object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsSubscribePushWithDefaults

`func NewWsSubscribePushWithDefaults() *WsSubscribePush`

NewWsSubscribePushWithDefaults instantiates a new WsSubscribePush object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoverable

`func (o *WsSubscribePush) GetRecoverable() bool`

GetRecoverable returns the Recoverable field if non-nil, zero value otherwise.

### GetRecoverableOk

`func (o *WsSubscribePush) GetRecoverableOk() (*bool, bool)`

GetRecoverableOk returns a tuple with the Recoverable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverable

`func (o *WsSubscribePush) SetRecoverable(v bool)`

SetRecoverable sets Recoverable field to given value.

### HasRecoverable

`func (o *WsSubscribePush) HasRecoverable() bool`

HasRecoverable returns a boolean if a field has been set.

### GetEpoch

`func (o *WsSubscribePush) GetEpoch() string`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *WsSubscribePush) GetEpochOk() (*string, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *WsSubscribePush) SetEpoch(v string)`

SetEpoch sets Epoch field to given value.

### HasEpoch

`func (o *WsSubscribePush) HasEpoch() bool`

HasEpoch returns a boolean if a field has been set.

### GetOffset

`func (o *WsSubscribePush) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *WsSubscribePush) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *WsSubscribePush) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *WsSubscribePush) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetPositioned

`func (o *WsSubscribePush) GetPositioned() bool`

GetPositioned returns the Positioned field if non-nil, zero value otherwise.

### GetPositionedOk

`func (o *WsSubscribePush) GetPositionedOk() (*bool, bool)`

GetPositionedOk returns a tuple with the Positioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositioned

`func (o *WsSubscribePush) SetPositioned(v bool)`

SetPositioned sets Positioned field to given value.

### HasPositioned

`func (o *WsSubscribePush) HasPositioned() bool`

HasPositioned returns a boolean if a field has been set.

### GetData

`func (o *WsSubscribePush) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WsSubscribePush) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WsSubscribePush) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *WsSubscribePush) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


