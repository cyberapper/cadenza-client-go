# RpcHealthCheckResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**RpcHealthCheck**](RpcHealthCheck.md) |  | [optional] 
**Error** | Pointer to [**RpcError**](RpcError.md) |  | [optional] 

## Methods

### NewRpcHealthCheckResult

`func NewRpcHealthCheckResult() *RpcHealthCheckResult`

NewRpcHealthCheckResult instantiates a new RpcHealthCheckResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcHealthCheckResultWithDefaults

`func NewRpcHealthCheckResultWithDefaults() *RpcHealthCheckResult`

NewRpcHealthCheckResultWithDefaults instantiates a new RpcHealthCheckResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RpcHealthCheckResult) GetData() RpcHealthCheck`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RpcHealthCheckResult) GetDataOk() (*RpcHealthCheck, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RpcHealthCheckResult) SetData(v RpcHealthCheck)`

SetData sets Data field to given value.

### HasData

`func (o *RpcHealthCheckResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *RpcHealthCheckResult) GetError() RpcError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RpcHealthCheckResult) GetErrorOk() (*RpcError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RpcHealthCheckResult) SetError(v RpcError)`

SetError sets Error field to given value.

### HasError

`func (o *RpcHealthCheckResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


