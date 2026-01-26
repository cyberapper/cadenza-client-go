# WsPublishRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | Channel to publish to | 
**Data** | **map[string]interface{}** | Message data to publish | 

## Methods

### NewWsPublishRequest

`func NewWsPublishRequest(channel string, data map[string]interface{}, ) *WsPublishRequest`

NewWsPublishRequest instantiates a new WsPublishRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsPublishRequestWithDefaults

`func NewWsPublishRequestWithDefaults() *WsPublishRequest`

NewWsPublishRequestWithDefaults instantiates a new WsPublishRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *WsPublishRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *WsPublishRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *WsPublishRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetData

`func (o *WsPublishRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WsPublishRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WsPublishRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


