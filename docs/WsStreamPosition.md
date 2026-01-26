# WsStreamPosition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **int64** | Stream offset | [optional] 
**Epoch** | Pointer to **string** | Stream epoch | [optional] 

## Methods

### NewWsStreamPosition

`func NewWsStreamPosition() *WsStreamPosition`

NewWsStreamPosition instantiates a new WsStreamPosition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsStreamPositionWithDefaults

`func NewWsStreamPositionWithDefaults() *WsStreamPosition`

NewWsStreamPositionWithDefaults instantiates a new WsStreamPosition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *WsStreamPosition) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *WsStreamPosition) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *WsStreamPosition) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *WsStreamPosition) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetEpoch

`func (o *WsStreamPosition) GetEpoch() string`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *WsStreamPosition) GetEpochOk() (*string, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *WsStreamPosition) SetEpoch(v string)`

SetEpoch sets Epoch field to given value.

### HasEpoch

`func (o *WsStreamPosition) HasEpoch() bool`

HasEpoch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


