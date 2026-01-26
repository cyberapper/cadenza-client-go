# WsSubRefreshRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | Channel to refresh | 
**Token** | **string** | New subscription token | 

## Methods

### NewWsSubRefreshRequest

`func NewWsSubRefreshRequest(channel string, token string, ) *WsSubRefreshRequest`

NewWsSubRefreshRequest instantiates a new WsSubRefreshRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsSubRefreshRequestWithDefaults

`func NewWsSubRefreshRequestWithDefaults() *WsSubRefreshRequest`

NewWsSubRefreshRequestWithDefaults instantiates a new WsSubRefreshRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *WsSubRefreshRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *WsSubRefreshRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *WsSubRefreshRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetToken

`func (o *WsSubRefreshRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *WsSubRefreshRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *WsSubRefreshRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


