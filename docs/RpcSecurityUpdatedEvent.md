# RpcSecurityUpdatedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RpcSecurity**](RpcSecurity.md) |  | [optional] 

## Methods

### NewRpcSecurityUpdatedEvent

`func NewRpcSecurityUpdatedEvent() *RpcSecurityUpdatedEvent`

NewRpcSecurityUpdatedEvent instantiates a new RpcSecurityUpdatedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcSecurityUpdatedEventWithDefaults

`func NewRpcSecurityUpdatedEventWithDefaults() *RpcSecurityUpdatedEvent`

NewRpcSecurityUpdatedEventWithDefaults instantiates a new RpcSecurityUpdatedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RpcSecurityUpdatedEvent) GetData() []RpcSecurity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RpcSecurityUpdatedEvent) GetDataOk() (*[]RpcSecurity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RpcSecurityUpdatedEvent) SetData(v []RpcSecurity)`

SetData sets Data field to given value.

### HasData

`func (o *RpcSecurityUpdatedEvent) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


