# WsRefreshResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Client** | Pointer to **string** | Client ID | [optional] 
**Version** | Pointer to **string** | Server version | [optional] 
**Expires** | Pointer to **bool** | Whether connection expires | [optional] 
**Ttl** | Pointer to **int32** | New TTL in seconds | [optional] 

## Methods

### NewWsRefreshResult

`func NewWsRefreshResult() *WsRefreshResult`

NewWsRefreshResult instantiates a new WsRefreshResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsRefreshResultWithDefaults

`func NewWsRefreshResultWithDefaults() *WsRefreshResult`

NewWsRefreshResultWithDefaults instantiates a new WsRefreshResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClient

`func (o *WsRefreshResult) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *WsRefreshResult) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *WsRefreshResult) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *WsRefreshResult) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetVersion

`func (o *WsRefreshResult) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WsRefreshResult) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WsRefreshResult) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WsRefreshResult) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetExpires

`func (o *WsRefreshResult) GetExpires() bool`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *WsRefreshResult) GetExpiresOk() (*bool, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *WsRefreshResult) SetExpires(v bool)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *WsRefreshResult) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetTtl

`func (o *WsRefreshResult) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *WsRefreshResult) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *WsRefreshResult) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *WsRefreshResult) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


