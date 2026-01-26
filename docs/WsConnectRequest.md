# WsConnectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | JWT authentication token | 
**Data** | Pointer to **map[string]interface{}** | Custom data to send with connect | [optional] 
**Subs** | Pointer to [**map[string]WsSubscribeRequest**](WsSubscribeRequest.md) | Initial subscriptions to establish on connect | [optional] 
**Name** | Pointer to **string** | Client name for identification | [optional] 
**Version** | Pointer to **string** | Client version | [optional] 

## Methods

### NewWsConnectRequest

`func NewWsConnectRequest(token string, ) *WsConnectRequest`

NewWsConnectRequest instantiates a new WsConnectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsConnectRequestWithDefaults

`func NewWsConnectRequestWithDefaults() *WsConnectRequest`

NewWsConnectRequestWithDefaults instantiates a new WsConnectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *WsConnectRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *WsConnectRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *WsConnectRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetData

`func (o *WsConnectRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WsConnectRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WsConnectRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *WsConnectRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSubs

`func (o *WsConnectRequest) GetSubs() map[string]WsSubscribeRequest`

GetSubs returns the Subs field if non-nil, zero value otherwise.

### GetSubsOk

`func (o *WsConnectRequest) GetSubsOk() (*map[string]WsSubscribeRequest, bool)`

GetSubsOk returns a tuple with the Subs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubs

`func (o *WsConnectRequest) SetSubs(v map[string]WsSubscribeRequest)`

SetSubs sets Subs field to given value.

### HasSubs

`func (o *WsConnectRequest) HasSubs() bool`

HasSubs returns a boolean if a field has been set.

### GetName

`func (o *WsConnectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WsConnectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WsConnectRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WsConnectRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *WsConnectRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WsConnectRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WsConnectRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WsConnectRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


