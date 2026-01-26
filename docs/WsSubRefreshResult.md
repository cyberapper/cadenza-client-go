# WsSubRefreshResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | Pointer to **bool** | Whether subscription expires | [optional] 
**Ttl** | Pointer to **int32** | New TTL in seconds | [optional] 

## Methods

### NewWsSubRefreshResult

`func NewWsSubRefreshResult() *WsSubRefreshResult`

NewWsSubRefreshResult instantiates a new WsSubRefreshResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsSubRefreshResultWithDefaults

`func NewWsSubRefreshResultWithDefaults() *WsSubRefreshResult`

NewWsSubRefreshResultWithDefaults instantiates a new WsSubRefreshResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpires

`func (o *WsSubRefreshResult) GetExpires() bool`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *WsSubRefreshResult) GetExpiresOk() (*bool, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *WsSubRefreshResult) SetExpires(v bool)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *WsSubRefreshResult) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetTtl

`func (o *WsSubRefreshResult) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *WsSubRefreshResult) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *WsSubRefreshResult) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *WsSubRefreshResult) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


