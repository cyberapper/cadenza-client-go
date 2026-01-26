# ListTradingAccountPortfolios200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Indicates if the operation was successful | 
**Errno** | **int32** | Error code (0 for success, negative for errors) | 
**Error** | **NullableString** | Error message (null for successful operations) | 
**Details** | Pointer to [**NullableBaseResponseDetails**](BaseResponseDetails.md) |  | [optional] 
**Data** | Pointer to [**[]Portfolio**](Portfolio.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListTradingAccountPortfolios200Response

`func NewListTradingAccountPortfolios200Response(success bool, errno int32, error_ NullableString, ) *ListTradingAccountPortfolios200Response`

NewListTradingAccountPortfolios200Response instantiates a new ListTradingAccountPortfolios200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTradingAccountPortfolios200ResponseWithDefaults

`func NewListTradingAccountPortfolios200ResponseWithDefaults() *ListTradingAccountPortfolios200Response`

NewListTradingAccountPortfolios200ResponseWithDefaults instantiates a new ListTradingAccountPortfolios200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListTradingAccountPortfolios200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListTradingAccountPortfolios200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListTradingAccountPortfolios200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetErrno

`func (o *ListTradingAccountPortfolios200Response) GetErrno() int32`

GetErrno returns the Errno field if non-nil, zero value otherwise.

### GetErrnoOk

`func (o *ListTradingAccountPortfolios200Response) GetErrnoOk() (*int32, bool)`

GetErrnoOk returns a tuple with the Errno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrno

`func (o *ListTradingAccountPortfolios200Response) SetErrno(v int32)`

SetErrno sets Errno field to given value.


### GetError

`func (o *ListTradingAccountPortfolios200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListTradingAccountPortfolios200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListTradingAccountPortfolios200Response) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *ListTradingAccountPortfolios200Response) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ListTradingAccountPortfolios200Response) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetDetails

`func (o *ListTradingAccountPortfolios200Response) GetDetails() BaseResponseDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ListTradingAccountPortfolios200Response) GetDetailsOk() (*BaseResponseDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ListTradingAccountPortfolios200Response) SetDetails(v BaseResponseDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ListTradingAccountPortfolios200Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *ListTradingAccountPortfolios200Response) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *ListTradingAccountPortfolios200Response) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetData

`func (o *ListTradingAccountPortfolios200Response) GetData() []Portfolio`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListTradingAccountPortfolios200Response) GetDataOk() (*[]Portfolio, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListTradingAccountPortfolios200Response) SetData(v []Portfolio)`

SetData sets Data field to given value.

### HasData

`func (o *ListTradingAccountPortfolios200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListTradingAccountPortfolios200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListTradingAccountPortfolios200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListTradingAccountPortfolios200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListTradingAccountPortfolios200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


