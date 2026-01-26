# Health200ResponseChecks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Database** | Pointer to [**HealthCheckComponent**](HealthCheckComponent.md) |  | [optional] 
**Temporal** | Pointer to [**HealthCheckComponent**](HealthCheckComponent.md) |  | [optional] 
**Redis** | Pointer to [**HealthCheckComponent**](HealthCheckComponent.md) |  | [optional] 

## Methods

### NewHealth200ResponseChecks

`func NewHealth200ResponseChecks() *Health200ResponseChecks`

NewHealth200ResponseChecks instantiates a new Health200ResponseChecks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealth200ResponseChecksWithDefaults

`func NewHealth200ResponseChecksWithDefaults() *Health200ResponseChecks`

NewHealth200ResponseChecksWithDefaults instantiates a new Health200ResponseChecks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabase

`func (o *Health200ResponseChecks) GetDatabase() HealthCheckComponent`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *Health200ResponseChecks) GetDatabaseOk() (*HealthCheckComponent, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *Health200ResponseChecks) SetDatabase(v HealthCheckComponent)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *Health200ResponseChecks) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetTemporal

`func (o *Health200ResponseChecks) GetTemporal() HealthCheckComponent`

GetTemporal returns the Temporal field if non-nil, zero value otherwise.

### GetTemporalOk

`func (o *Health200ResponseChecks) GetTemporalOk() (*HealthCheckComponent, bool)`

GetTemporalOk returns a tuple with the Temporal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporal

`func (o *Health200ResponseChecks) SetTemporal(v HealthCheckComponent)`

SetTemporal sets Temporal field to given value.

### HasTemporal

`func (o *Health200ResponseChecks) HasTemporal() bool`

HasTemporal returns a boolean if a field has been set.

### GetRedis

`func (o *Health200ResponseChecks) GetRedis() HealthCheckComponent`

GetRedis returns the Redis field if non-nil, zero value otherwise.

### GetRedisOk

`func (o *Health200ResponseChecks) GetRedisOk() (*HealthCheckComponent, bool)`

GetRedisOk returns a tuple with the Redis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedis

`func (o *Health200ResponseChecks) SetRedis(v HealthCheckComponent)`

SetRedis sets Redis field to given value.

### HasRedis

`func (o *Health200ResponseChecks) HasRedis() bool`

HasRedis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


