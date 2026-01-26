# HealthCheckComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**HealthStatus**](HealthStatus.md) |  | 

## Methods

### NewHealthCheckComponent

`func NewHealthCheckComponent(status HealthStatus, ) *HealthCheckComponent`

NewHealthCheckComponent instantiates a new HealthCheckComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckComponentWithDefaults

`func NewHealthCheckComponentWithDefaults() *HealthCheckComponent`

NewHealthCheckComponentWithDefaults instantiates a new HealthCheckComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HealthCheckComponent) GetStatus() HealthStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthCheckComponent) GetStatusOk() (*HealthStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthCheckComponent) SetStatus(v HealthStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


