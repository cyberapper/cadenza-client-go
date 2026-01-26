# WsSubscribeResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | Pointer to **bool** | Whether the subscription expires | [optional] 
**Ttl** | Pointer to **int32** | Time to live in seconds | [optional] 
**Recoverable** | Pointer to **bool** | Whether missed messages can be recovered | [optional] 
**Epoch** | Pointer to **string** | Stream epoch | [optional] 
**Publications** | Pointer to [**[]WsPublication**](WsPublication.md) | Recovered publications (if recovery was requested) | [optional] 
**Recovered** | Pointer to **bool** | Whether recovery was successful | [optional] 
**Offset** | Pointer to **int64** | Current stream offset | [optional] 
**Positioned** | Pointer to **bool** | Whether position info is enabled | [optional] 
**Data** | Pointer to **map[string]interface{}** | Custom data from server | [optional] 
**Delta** | Pointer to **bool** | Whether delta compression is enabled | [optional] 

## Methods

### NewWsSubscribeResult

`func NewWsSubscribeResult() *WsSubscribeResult`

NewWsSubscribeResult instantiates a new WsSubscribeResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsSubscribeResultWithDefaults

`func NewWsSubscribeResultWithDefaults() *WsSubscribeResult`

NewWsSubscribeResultWithDefaults instantiates a new WsSubscribeResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpires

`func (o *WsSubscribeResult) GetExpires() bool`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *WsSubscribeResult) GetExpiresOk() (*bool, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *WsSubscribeResult) SetExpires(v bool)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *WsSubscribeResult) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetTtl

`func (o *WsSubscribeResult) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *WsSubscribeResult) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *WsSubscribeResult) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *WsSubscribeResult) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetRecoverable

`func (o *WsSubscribeResult) GetRecoverable() bool`

GetRecoverable returns the Recoverable field if non-nil, zero value otherwise.

### GetRecoverableOk

`func (o *WsSubscribeResult) GetRecoverableOk() (*bool, bool)`

GetRecoverableOk returns a tuple with the Recoverable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverable

`func (o *WsSubscribeResult) SetRecoverable(v bool)`

SetRecoverable sets Recoverable field to given value.

### HasRecoverable

`func (o *WsSubscribeResult) HasRecoverable() bool`

HasRecoverable returns a boolean if a field has been set.

### GetEpoch

`func (o *WsSubscribeResult) GetEpoch() string`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *WsSubscribeResult) GetEpochOk() (*string, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *WsSubscribeResult) SetEpoch(v string)`

SetEpoch sets Epoch field to given value.

### HasEpoch

`func (o *WsSubscribeResult) HasEpoch() bool`

HasEpoch returns a boolean if a field has been set.

### GetPublications

`func (o *WsSubscribeResult) GetPublications() []WsPublication`

GetPublications returns the Publications field if non-nil, zero value otherwise.

### GetPublicationsOk

`func (o *WsSubscribeResult) GetPublicationsOk() (*[]WsPublication, bool)`

GetPublicationsOk returns a tuple with the Publications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublications

`func (o *WsSubscribeResult) SetPublications(v []WsPublication)`

SetPublications sets Publications field to given value.

### HasPublications

`func (o *WsSubscribeResult) HasPublications() bool`

HasPublications returns a boolean if a field has been set.

### GetRecovered

`func (o *WsSubscribeResult) GetRecovered() bool`

GetRecovered returns the Recovered field if non-nil, zero value otherwise.

### GetRecoveredOk

`func (o *WsSubscribeResult) GetRecoveredOk() (*bool, bool)`

GetRecoveredOk returns a tuple with the Recovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecovered

`func (o *WsSubscribeResult) SetRecovered(v bool)`

SetRecovered sets Recovered field to given value.

### HasRecovered

`func (o *WsSubscribeResult) HasRecovered() bool`

HasRecovered returns a boolean if a field has been set.

### GetOffset

`func (o *WsSubscribeResult) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *WsSubscribeResult) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *WsSubscribeResult) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *WsSubscribeResult) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetPositioned

`func (o *WsSubscribeResult) GetPositioned() bool`

GetPositioned returns the Positioned field if non-nil, zero value otherwise.

### GetPositionedOk

`func (o *WsSubscribeResult) GetPositionedOk() (*bool, bool)`

GetPositionedOk returns a tuple with the Positioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositioned

`func (o *WsSubscribeResult) SetPositioned(v bool)`

SetPositioned sets Positioned field to given value.

### HasPositioned

`func (o *WsSubscribeResult) HasPositioned() bool`

HasPositioned returns a boolean if a field has been set.

### GetData

`func (o *WsSubscribeResult) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WsSubscribeResult) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WsSubscribeResult) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *WsSubscribeResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDelta

`func (o *WsSubscribeResult) GetDelta() bool`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *WsSubscribeResult) GetDeltaOk() (*bool, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *WsSubscribeResult) SetDelta(v bool)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *WsSubscribeResult) HasDelta() bool`

HasDelta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


