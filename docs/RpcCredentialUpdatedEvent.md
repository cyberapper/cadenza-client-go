# RpcCredentialUpdatedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**RpcTradingAccountCredential**](RpcTradingAccountCredential.md) |  | [optional] 

## Methods

### NewRpcCredentialUpdatedEvent

`func NewRpcCredentialUpdatedEvent() *RpcCredentialUpdatedEvent`

NewRpcCredentialUpdatedEvent instantiates a new RpcCredentialUpdatedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcCredentialUpdatedEventWithDefaults

`func NewRpcCredentialUpdatedEventWithDefaults() *RpcCredentialUpdatedEvent`

NewRpcCredentialUpdatedEventWithDefaults instantiates a new RpcCredentialUpdatedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RpcCredentialUpdatedEvent) GetData() RpcTradingAccountCredential`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RpcCredentialUpdatedEvent) GetDataOk() (*RpcTradingAccountCredential, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RpcCredentialUpdatedEvent) SetData(v RpcTradingAccountCredential)`

SetData sets Data field to given value.

### HasData

`func (o *RpcCredentialUpdatedEvent) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


