# Root404Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Indicates if the operation was successful | 
**Errno** | **int32** | Error code (0 for success, negative for errors) | 
**Error** | Pointer to **NullableString** | Error message (null for successful operations) | [optional] 
**Details** | Pointer to [**NullableBaseResponseDetails**](BaseResponseDetails.md) |  | [optional] 
**Data** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRoot404Response

`func NewRoot404Response(success bool, errno int32, ) *Root404Response`

NewRoot404Response instantiates a new Root404Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoot404ResponseWithDefaults

`func NewRoot404ResponseWithDefaults() *Root404Response`

NewRoot404ResponseWithDefaults instantiates a new Root404Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *Root404Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Root404Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Root404Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetErrno

`func (o *Root404Response) GetErrno() int32`

GetErrno returns the Errno field if non-nil, zero value otherwise.

### GetErrnoOk

`func (o *Root404Response) GetErrnoOk() (*int32, bool)`

GetErrnoOk returns a tuple with the Errno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrno

`func (o *Root404Response) SetErrno(v int32)`

SetErrno sets Errno field to given value.


### GetError

`func (o *Root404Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Root404Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Root404Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Root404Response) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *Root404Response) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *Root404Response) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetDetails

`func (o *Root404Response) GetDetails() BaseResponseDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Root404Response) GetDetailsOk() (*BaseResponseDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Root404Response) SetDetails(v BaseResponseDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Root404Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *Root404Response) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *Root404Response) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetData

`func (o *Root404Response) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Root404Response) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Root404Response) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *Root404Response) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *Root404Response) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *Root404Response) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


