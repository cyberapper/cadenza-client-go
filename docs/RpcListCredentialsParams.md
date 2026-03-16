# RpcListCredentialsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialIds** | Pointer to **[]string** |  | [optional] 
**Venue** | Pointer to [**Venue**](Venue.md) |  | [optional] 
**CredentialType** | Pointer to [**CredentialType**](CredentialType.md) |  | [optional] 
**Status** | Pointer to [**CredentialStatus**](CredentialStatus.md) |  | [optional] 
**Pagination** | Pointer to [**RpcPagination**](RpcPagination.md) |  | [optional] 

## Methods

### NewRpcListCredentialsParams

`func NewRpcListCredentialsParams() *RpcListCredentialsParams`

NewRpcListCredentialsParams instantiates a new RpcListCredentialsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcListCredentialsParamsWithDefaults

`func NewRpcListCredentialsParamsWithDefaults() *RpcListCredentialsParams`

NewRpcListCredentialsParamsWithDefaults instantiates a new RpcListCredentialsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialIds

`func (o *RpcListCredentialsParams) GetCredentialIds() []string`

GetCredentialIds returns the CredentialIds field if non-nil, zero value otherwise.

### GetCredentialIdsOk

`func (o *RpcListCredentialsParams) GetCredentialIdsOk() (*[]string, bool)`

GetCredentialIdsOk returns a tuple with the CredentialIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialIds

`func (o *RpcListCredentialsParams) SetCredentialIds(v []string)`

SetCredentialIds sets CredentialIds field to given value.

### HasCredentialIds

`func (o *RpcListCredentialsParams) HasCredentialIds() bool`

HasCredentialIds returns a boolean if a field has been set.

### GetVenue

`func (o *RpcListCredentialsParams) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *RpcListCredentialsParams) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *RpcListCredentialsParams) SetVenue(v Venue)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *RpcListCredentialsParams) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetCredentialType

`func (o *RpcListCredentialsParams) GetCredentialType() CredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *RpcListCredentialsParams) GetCredentialTypeOk() (*CredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *RpcListCredentialsParams) SetCredentialType(v CredentialType)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *RpcListCredentialsParams) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### GetStatus

`func (o *RpcListCredentialsParams) GetStatus() CredentialStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RpcListCredentialsParams) GetStatusOk() (*CredentialStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RpcListCredentialsParams) SetStatus(v CredentialStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RpcListCredentialsParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPagination

`func (o *RpcListCredentialsParams) GetPagination() RpcPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RpcListCredentialsParams) GetPaginationOk() (*RpcPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RpcListCredentialsParams) SetPagination(v RpcPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RpcListCredentialsParams) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


