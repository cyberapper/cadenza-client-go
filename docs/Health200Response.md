# Health200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**HealthStatus**](HealthStatus.md) |  | 
**Timestamp** | **int64** | Health check timestamp in milliseconds | 
**IsoDateTime** | Pointer to **time.Time** | Health check timestamp in ISO 8601 format | [optional] 
**Version** | **string** | API version | 
**Checks** | Pointer to [**Health200ResponseChecks**](Health200ResponseChecks.md) |  | [optional] 

## Methods

### NewHealth200Response

`func NewHealth200Response(status HealthStatus, timestamp int64, version string, ) *Health200Response`

NewHealth200Response instantiates a new Health200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealth200ResponseWithDefaults

`func NewHealth200ResponseWithDefaults() *Health200Response`

NewHealth200ResponseWithDefaults instantiates a new Health200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Health200Response) GetStatus() HealthStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Health200Response) GetStatusOk() (*HealthStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Health200Response) SetStatus(v HealthStatus)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *Health200Response) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Health200Response) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Health200Response) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetIsoDateTime

`func (o *Health200Response) GetIsoDateTime() time.Time`

GetIsoDateTime returns the IsoDateTime field if non-nil, zero value otherwise.

### GetIsoDateTimeOk

`func (o *Health200Response) GetIsoDateTimeOk() (*time.Time, bool)`

GetIsoDateTimeOk returns a tuple with the IsoDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoDateTime

`func (o *Health200Response) SetIsoDateTime(v time.Time)`

SetIsoDateTime sets IsoDateTime field to given value.

### HasIsoDateTime

`func (o *Health200Response) HasIsoDateTime() bool`

HasIsoDateTime returns a boolean if a field has been set.

### GetVersion

`func (o *Health200Response) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Health200Response) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Health200Response) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetChecks

`func (o *Health200Response) GetChecks() Health200ResponseChecks`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *Health200Response) GetChecksOk() (*Health200ResponseChecks, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *Health200Response) SetChecks(v Health200ResponseChecks)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *Health200Response) HasChecks() bool`

HasChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


