# WsUnsubscribeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | Channel name to unsubscribe from | 

## Methods

### NewWsUnsubscribeRequest

`func NewWsUnsubscribeRequest(channel string, ) *WsUnsubscribeRequest`

NewWsUnsubscribeRequest instantiates a new WsUnsubscribeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsUnsubscribeRequestWithDefaults

`func NewWsUnsubscribeRequestWithDefaults() *WsUnsubscribeRequest`

NewWsUnsubscribeRequestWithDefaults instantiates a new WsUnsubscribeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *WsUnsubscribeRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *WsUnsubscribeRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *WsUnsubscribeRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


