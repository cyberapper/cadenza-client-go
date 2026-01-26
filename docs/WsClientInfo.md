# WsClientInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **string** | User ID | [optional] 
**Client** | Pointer to **string** | Client connection ID | [optional] 
**ConnInfo** | Pointer to **map[string]interface{}** | Connection info set during authentication | [optional] 
**ChanInfo** | Pointer to **map[string]interface{}** | Channel info set during subscription | [optional] 

## Methods

### NewWsClientInfo

`func NewWsClientInfo() *WsClientInfo`

NewWsClientInfo instantiates a new WsClientInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsClientInfoWithDefaults

`func NewWsClientInfoWithDefaults() *WsClientInfo`

NewWsClientInfoWithDefaults instantiates a new WsClientInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *WsClientInfo) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *WsClientInfo) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *WsClientInfo) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *WsClientInfo) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetClient

`func (o *WsClientInfo) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *WsClientInfo) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *WsClientInfo) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *WsClientInfo) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetConnInfo

`func (o *WsClientInfo) GetConnInfo() map[string]interface{}`

GetConnInfo returns the ConnInfo field if non-nil, zero value otherwise.

### GetConnInfoOk

`func (o *WsClientInfo) GetConnInfoOk() (*map[string]interface{}, bool)`

GetConnInfoOk returns a tuple with the ConnInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnInfo

`func (o *WsClientInfo) SetConnInfo(v map[string]interface{})`

SetConnInfo sets ConnInfo field to given value.

### HasConnInfo

`func (o *WsClientInfo) HasConnInfo() bool`

HasConnInfo returns a boolean if a field has been set.

### GetChanInfo

`func (o *WsClientInfo) GetChanInfo() map[string]interface{}`

GetChanInfo returns the ChanInfo field if non-nil, zero value otherwise.

### GetChanInfoOk

`func (o *WsClientInfo) GetChanInfoOk() (*map[string]interface{}, bool)`

GetChanInfoOk returns a tuple with the ChanInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanInfo

`func (o *WsClientInfo) SetChanInfo(v map[string]interface{})`

SetChanInfo sets ChanInfo field to given value.

### HasChanInfo

`func (o *WsClientInfo) HasChanInfo() bool`

HasChanInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


