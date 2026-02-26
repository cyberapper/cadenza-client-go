# Root400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Indicates if the operation was successful | 
**Errno** | **int32** | Error code (0 for success, negative for errors) | 
**Error** | Pointer to **NullableString** | Error message (null for successful operations) | [optional] 
**Details** | Pointer to [**NullableBaseResponseDetails**](BaseResponseDetails.md) |  | [optional] 
**Data** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRoot400Response

`func NewRoot400Response(success bool, errno int32, ) *Root400Response`

NewRoot400Response instantiates a new Root400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoot400ResponseWithDefaults

`func NewRoot400ResponseWithDefaults() *Root400Response`

NewRoot400ResponseWithDefaults instantiates a new Root400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *Root400Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Root400Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Root400Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetErrno

`func (o *Root400Response) GetErrno() int32`

GetErrno returns the Errno field if non-nil, zero value otherwise.

### GetErrnoOk

`func (o *Root400Response) GetErrnoOk() (*int32, bool)`

GetErrnoOk returns a tuple with the Errno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrno

`func (o *Root400Response) SetErrno(v int32)`

SetErrno sets Errno field to given value.


### GetError

`func (o *Root400Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Root400Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Root400Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Root400Response) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *Root400Response) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *Root400Response) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetDetails

`func (o *Root400Response) GetDetails() BaseResponseDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Root400Response) GetDetailsOk() (*BaseResponseDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Root400Response) SetDetails(v BaseResponseDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Root400Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *Root400Response) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *Root400Response) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetData

`func (o *Root400Response) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Root400Response) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Root400Response) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *Root400Response) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *Root400Response) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *Root400Response) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


