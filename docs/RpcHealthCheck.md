# RpcHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**HealthStatus**](HealthStatus.md) |  | 
**Timestamp** | **time.Time** | Timestamp in ISO 8601 format (RFC3339). This is the native format used by Go&#39;s time.Time. | 
**Version** | **string** | API version | 
**Checks** | Pointer to [**Health200ResponseChecks**](Health200ResponseChecks.md) |  | [optional] 

## Methods

### NewRpcHealthCheck

`func NewRpcHealthCheck(status HealthStatus, timestamp time.Time, version string, ) *RpcHealthCheck`

NewRpcHealthCheck instantiates a new RpcHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcHealthCheckWithDefaults

`func NewRpcHealthCheckWithDefaults() *RpcHealthCheck`

NewRpcHealthCheckWithDefaults instantiates a new RpcHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RpcHealthCheck) GetStatus() HealthStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RpcHealthCheck) GetStatusOk() (*HealthStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RpcHealthCheck) SetStatus(v HealthStatus)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *RpcHealthCheck) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RpcHealthCheck) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RpcHealthCheck) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetVersion

`func (o *RpcHealthCheck) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RpcHealthCheck) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RpcHealthCheck) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetChecks

`func (o *RpcHealthCheck) GetChecks() Health200ResponseChecks`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *RpcHealthCheck) GetChecksOk() (*Health200ResponseChecks, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *RpcHealthCheck) SetChecks(v Health200ResponseChecks)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *RpcHealthCheck) HasChecks() bool`

HasChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


