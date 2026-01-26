# DeleteMarketSecurity200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Indicates if the operation was successful | 
**Errno** | **int32** | Error code (0 for success, negative for errors) | 
**Error** | **NullableString** | Error message (null for successful operations) | 
**Details** | Pointer to [**NullableBaseResponseDetails**](BaseResponseDetails.md) |  | [optional] 
**Data** | Pointer to [**FinancialSecurity**](FinancialSecurity.md) |  | [optional] 

## Methods

### NewDeleteMarketSecurity200Response

`func NewDeleteMarketSecurity200Response(success bool, errno int32, error_ NullableString, ) *DeleteMarketSecurity200Response`

NewDeleteMarketSecurity200Response instantiates a new DeleteMarketSecurity200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteMarketSecurity200ResponseWithDefaults

`func NewDeleteMarketSecurity200ResponseWithDefaults() *DeleteMarketSecurity200Response`

NewDeleteMarketSecurity200ResponseWithDefaults instantiates a new DeleteMarketSecurity200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *DeleteMarketSecurity200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DeleteMarketSecurity200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DeleteMarketSecurity200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetErrno

`func (o *DeleteMarketSecurity200Response) GetErrno() int32`

GetErrno returns the Errno field if non-nil, zero value otherwise.

### GetErrnoOk

`func (o *DeleteMarketSecurity200Response) GetErrnoOk() (*int32, bool)`

GetErrnoOk returns a tuple with the Errno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrno

`func (o *DeleteMarketSecurity200Response) SetErrno(v int32)`

SetErrno sets Errno field to given value.


### GetError

`func (o *DeleteMarketSecurity200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DeleteMarketSecurity200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DeleteMarketSecurity200Response) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *DeleteMarketSecurity200Response) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *DeleteMarketSecurity200Response) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetDetails

`func (o *DeleteMarketSecurity200Response) GetDetails() BaseResponseDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DeleteMarketSecurity200Response) GetDetailsOk() (*BaseResponseDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DeleteMarketSecurity200Response) SetDetails(v BaseResponseDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *DeleteMarketSecurity200Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *DeleteMarketSecurity200Response) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *DeleteMarketSecurity200Response) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetData

`func (o *DeleteMarketSecurity200Response) GetData() FinancialSecurity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteMarketSecurity200Response) GetDataOk() (*FinancialSecurity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteMarketSecurity200Response) SetData(v FinancialSecurity)`

SetData sets Data field to given value.

### HasData

`func (o *DeleteMarketSecurity200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


