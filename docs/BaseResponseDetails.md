# BaseResponseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Error code identifier (e.g., INVALID_TOKEN, ACCESS_DENIED) | [optional] 
**Resource** | Pointer to **string** | Resource that was being accessed | [optional] 
**Action** | Pointer to **string** | Action that was being attempted | [optional] 
**Required** | Pointer to **[]string** | Required permissions or roles | [optional] 
**Provided** | Pointer to **[]string** | Provided permissions or roles | [optional] 
**TenantId** | Pointer to **string** | Tenant ID if relevant to the error | [optional] 
**RequestId** | Pointer to **string** | Request ID for tracking and debugging | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Additional metadata as key-value pairs | [optional] 

## Methods

### NewBaseResponseDetails

`func NewBaseResponseDetails() *BaseResponseDetails`

NewBaseResponseDetails instantiates a new BaseResponseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseResponseDetailsWithDefaults

`func NewBaseResponseDetailsWithDefaults() *BaseResponseDetails`

NewBaseResponseDetailsWithDefaults instantiates a new BaseResponseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *BaseResponseDetails) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BaseResponseDetails) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BaseResponseDetails) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *BaseResponseDetails) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetResource

`func (o *BaseResponseDetails) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *BaseResponseDetails) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *BaseResponseDetails) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *BaseResponseDetails) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetAction

`func (o *BaseResponseDetails) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BaseResponseDetails) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BaseResponseDetails) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *BaseResponseDetails) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetRequired

`func (o *BaseResponseDetails) GetRequired() []string`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *BaseResponseDetails) GetRequiredOk() (*[]string, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *BaseResponseDetails) SetRequired(v []string)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *BaseResponseDetails) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetProvided

`func (o *BaseResponseDetails) GetProvided() []string`

GetProvided returns the Provided field if non-nil, zero value otherwise.

### GetProvidedOk

`func (o *BaseResponseDetails) GetProvidedOk() (*[]string, bool)`

GetProvidedOk returns a tuple with the Provided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvided

`func (o *BaseResponseDetails) SetProvided(v []string)`

SetProvided sets Provided field to given value.

### HasProvided

`func (o *BaseResponseDetails) HasProvided() bool`

HasProvided returns a boolean if a field has been set.

### GetTenantId

`func (o *BaseResponseDetails) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BaseResponseDetails) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BaseResponseDetails) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BaseResponseDetails) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetRequestId

`func (o *BaseResponseDetails) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *BaseResponseDetails) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *BaseResponseDetails) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *BaseResponseDetails) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetMetadata

`func (o *BaseResponseDetails) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BaseResponseDetails) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BaseResponseDetails) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BaseResponseDetails) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


