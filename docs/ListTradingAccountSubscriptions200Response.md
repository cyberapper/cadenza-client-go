# ListTradingAccountSubscriptions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Indicates if the operation was successful | 
**Errno** | **int32** | Error code (0 for success, negative for errors) | 
**Error** | **NullableString** | Error message (null for successful operations) | 
**Details** | Pointer to [**NullableBaseResponseDetails**](BaseResponseDetails.md) |  | [optional] 
**Data** | Pointer to [**[]Subscription**](Subscription.md) |  | [optional] 

## Methods

### NewListTradingAccountSubscriptions200Response

`func NewListTradingAccountSubscriptions200Response(success bool, errno int32, error_ NullableString, ) *ListTradingAccountSubscriptions200Response`

NewListTradingAccountSubscriptions200Response instantiates a new ListTradingAccountSubscriptions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTradingAccountSubscriptions200ResponseWithDefaults

`func NewListTradingAccountSubscriptions200ResponseWithDefaults() *ListTradingAccountSubscriptions200Response`

NewListTradingAccountSubscriptions200ResponseWithDefaults instantiates a new ListTradingAccountSubscriptions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListTradingAccountSubscriptions200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListTradingAccountSubscriptions200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListTradingAccountSubscriptions200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetErrno

`func (o *ListTradingAccountSubscriptions200Response) GetErrno() int32`

GetErrno returns the Errno field if non-nil, zero value otherwise.

### GetErrnoOk

`func (o *ListTradingAccountSubscriptions200Response) GetErrnoOk() (*int32, bool)`

GetErrnoOk returns a tuple with the Errno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrno

`func (o *ListTradingAccountSubscriptions200Response) SetErrno(v int32)`

SetErrno sets Errno field to given value.


### GetError

`func (o *ListTradingAccountSubscriptions200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListTradingAccountSubscriptions200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListTradingAccountSubscriptions200Response) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *ListTradingAccountSubscriptions200Response) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ListTradingAccountSubscriptions200Response) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetDetails

`func (o *ListTradingAccountSubscriptions200Response) GetDetails() BaseResponseDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ListTradingAccountSubscriptions200Response) GetDetailsOk() (*BaseResponseDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ListTradingAccountSubscriptions200Response) SetDetails(v BaseResponseDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ListTradingAccountSubscriptions200Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *ListTradingAccountSubscriptions200Response) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *ListTradingAccountSubscriptions200Response) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetData

`func (o *ListTradingAccountSubscriptions200Response) GetData() []Subscription`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListTradingAccountSubscriptions200Response) GetDataOk() (*[]Subscription, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListTradingAccountSubscriptions200Response) SetData(v []Subscription)`

SetData sets Data field to given value.

### HasData

`func (o *ListTradingAccountSubscriptions200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


