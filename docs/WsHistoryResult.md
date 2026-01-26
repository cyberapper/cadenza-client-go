# WsHistoryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Publications** | Pointer to [**[]WsPublication**](WsPublication.md) |  | [optional] 
**Epoch** | Pointer to **string** | Stream epoch | [optional] 
**Offset** | Pointer to **int64** | Current stream offset | [optional] 

## Methods

### NewWsHistoryResult

`func NewWsHistoryResult() *WsHistoryResult`

NewWsHistoryResult instantiates a new WsHistoryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsHistoryResultWithDefaults

`func NewWsHistoryResultWithDefaults() *WsHistoryResult`

NewWsHistoryResultWithDefaults instantiates a new WsHistoryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublications

`func (o *WsHistoryResult) GetPublications() []WsPublication`

GetPublications returns the Publications field if non-nil, zero value otherwise.

### GetPublicationsOk

`func (o *WsHistoryResult) GetPublicationsOk() (*[]WsPublication, bool)`

GetPublicationsOk returns a tuple with the Publications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublications

`func (o *WsHistoryResult) SetPublications(v []WsPublication)`

SetPublications sets Publications field to given value.

### HasPublications

`func (o *WsHistoryResult) HasPublications() bool`

HasPublications returns a boolean if a field has been set.

### GetEpoch

`func (o *WsHistoryResult) GetEpoch() string`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *WsHistoryResult) GetEpochOk() (*string, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *WsHistoryResult) SetEpoch(v string)`

SetEpoch sets Epoch field to given value.

### HasEpoch

`func (o *WsHistoryResult) HasEpoch() bool`

HasEpoch returns a boolean if a field has been set.

### GetOffset

`func (o *WsHistoryResult) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *WsHistoryResult) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *WsHistoryResult) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *WsHistoryResult) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


