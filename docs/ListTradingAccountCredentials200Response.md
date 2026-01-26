# ListTradingAccountCredentials200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Indicates if the operation was successful | 
**Errno** | **int32** | Error code (0 for success, negative for errors) | 
**Error** | **NullableString** | Error message (null for successful operations) | 
**Details** | Pointer to [**NullableBaseResponseDetails**](BaseResponseDetails.md) |  | [optional] 
**Data** | Pointer to [**[]TradingAccountCredential**](TradingAccountCredential.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListTradingAccountCredentials200Response

`func NewListTradingAccountCredentials200Response(success bool, errno int32, error_ NullableString, ) *ListTradingAccountCredentials200Response`

NewListTradingAccountCredentials200Response instantiates a new ListTradingAccountCredentials200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTradingAccountCredentials200ResponseWithDefaults

`func NewListTradingAccountCredentials200ResponseWithDefaults() *ListTradingAccountCredentials200Response`

NewListTradingAccountCredentials200ResponseWithDefaults instantiates a new ListTradingAccountCredentials200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListTradingAccountCredentials200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListTradingAccountCredentials200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListTradingAccountCredentials200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetErrno

`func (o *ListTradingAccountCredentials200Response) GetErrno() int32`

GetErrno returns the Errno field if non-nil, zero value otherwise.

### GetErrnoOk

`func (o *ListTradingAccountCredentials200Response) GetErrnoOk() (*int32, bool)`

GetErrnoOk returns a tuple with the Errno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrno

`func (o *ListTradingAccountCredentials200Response) SetErrno(v int32)`

SetErrno sets Errno field to given value.


### GetError

`func (o *ListTradingAccountCredentials200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListTradingAccountCredentials200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListTradingAccountCredentials200Response) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *ListTradingAccountCredentials200Response) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ListTradingAccountCredentials200Response) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetDetails

`func (o *ListTradingAccountCredentials200Response) GetDetails() BaseResponseDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ListTradingAccountCredentials200Response) GetDetailsOk() (*BaseResponseDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ListTradingAccountCredentials200Response) SetDetails(v BaseResponseDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ListTradingAccountCredentials200Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *ListTradingAccountCredentials200Response) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *ListTradingAccountCredentials200Response) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetData

`func (o *ListTradingAccountCredentials200Response) GetData() []TradingAccountCredential`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListTradingAccountCredentials200Response) GetDataOk() (*[]TradingAccountCredential, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListTradingAccountCredentials200Response) SetData(v []TradingAccountCredential)`

SetData sets Data field to given value.

### HasData

`func (o *ListTradingAccountCredentials200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListTradingAccountCredentials200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListTradingAccountCredentials200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListTradingAccountCredentials200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListTradingAccountCredentials200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


