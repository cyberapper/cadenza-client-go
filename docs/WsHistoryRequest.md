# WsHistoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | Channel name | 
**Limit** | Pointer to **int32** | Maximum number of publications to return | [optional] 
**Since** | Pointer to [**WsStreamPosition**](WsStreamPosition.md) |  | [optional] 
**Reverse** | Pointer to **bool** | Whether to return publications in reverse order | [optional] 

## Methods

### NewWsHistoryRequest

`func NewWsHistoryRequest(channel string, ) *WsHistoryRequest`

NewWsHistoryRequest instantiates a new WsHistoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsHistoryRequestWithDefaults

`func NewWsHistoryRequestWithDefaults() *WsHistoryRequest`

NewWsHistoryRequestWithDefaults instantiates a new WsHistoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *WsHistoryRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *WsHistoryRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *WsHistoryRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetLimit

`func (o *WsHistoryRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *WsHistoryRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *WsHistoryRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *WsHistoryRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetSince

`func (o *WsHistoryRequest) GetSince() WsStreamPosition`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *WsHistoryRequest) GetSinceOk() (*WsStreamPosition, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *WsHistoryRequest) SetSince(v WsStreamPosition)`

SetSince sets Since field to given value.

### HasSince

`func (o *WsHistoryRequest) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetReverse

`func (o *WsHistoryRequest) GetReverse() bool`

GetReverse returns the Reverse field if non-nil, zero value otherwise.

### GetReverseOk

`func (o *WsHistoryRequest) GetReverseOk() (*bool, bool)`

GetReverseOk returns a tuple with the Reverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverse

`func (o *WsHistoryRequest) SetReverse(v bool)`

SetReverse sets Reverse field to given value.

### HasReverse

`func (o *WsHistoryRequest) HasReverse() bool`

HasReverse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


