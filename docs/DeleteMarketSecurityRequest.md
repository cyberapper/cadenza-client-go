# DeleteMarketSecurityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityId** | **string** | Security ID in the format of venue:symbol | 

## Methods

### NewDeleteMarketSecurityRequest

`func NewDeleteMarketSecurityRequest(securityId string, ) *DeleteMarketSecurityRequest`

NewDeleteMarketSecurityRequest instantiates a new DeleteMarketSecurityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteMarketSecurityRequestWithDefaults

`func NewDeleteMarketSecurityRequestWithDefaults() *DeleteMarketSecurityRequest`

NewDeleteMarketSecurityRequestWithDefaults instantiates a new DeleteMarketSecurityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityId

`func (o *DeleteMarketSecurityRequest) GetSecurityId() string`

GetSecurityId returns the SecurityId field if non-nil, zero value otherwise.

### GetSecurityIdOk

`func (o *DeleteMarketSecurityRequest) GetSecurityIdOk() (*string, bool)`

GetSecurityIdOk returns a tuple with the SecurityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityId

`func (o *DeleteMarketSecurityRequest) SetSecurityId(v string)`

SetSecurityId sets SecurityId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


