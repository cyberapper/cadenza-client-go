# ListTradingAccounts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Indicates if the operation was successful | 
**Errno** | **int32** | Error code (0 for success, negative for errors) | 
**Error** | Pointer to **NullableString** | Error message (null for successful operations) | [optional] 
**Details** | Pointer to [**NullableBaseResponseDetails**](BaseResponseDetails.md) |  | [optional] 
**Data** | Pointer to [**[]TradingAccount**](TradingAccount.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListTradingAccounts200Response

`func NewListTradingAccounts200Response(success bool, errno int32, ) *ListTradingAccounts200Response`

NewListTradingAccounts200Response instantiates a new ListTradingAccounts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTradingAccounts200ResponseWithDefaults

`func NewListTradingAccounts200ResponseWithDefaults() *ListTradingAccounts200Response`

NewListTradingAccounts200ResponseWithDefaults instantiates a new ListTradingAccounts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListTradingAccounts200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListTradingAccounts200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListTradingAccounts200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetErrno

`func (o *ListTradingAccounts200Response) GetErrno() int32`

GetErrno returns the Errno field if non-nil, zero value otherwise.

### GetErrnoOk

`func (o *ListTradingAccounts200Response) GetErrnoOk() (*int32, bool)`

GetErrnoOk returns a tuple with the Errno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrno

`func (o *ListTradingAccounts200Response) SetErrno(v int32)`

SetErrno sets Errno field to given value.


### GetError

`func (o *ListTradingAccounts200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListTradingAccounts200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListTradingAccounts200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ListTradingAccounts200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ListTradingAccounts200Response) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ListTradingAccounts200Response) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetDetails

`func (o *ListTradingAccounts200Response) GetDetails() BaseResponseDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ListTradingAccounts200Response) GetDetailsOk() (*BaseResponseDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ListTradingAccounts200Response) SetDetails(v BaseResponseDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ListTradingAccounts200Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *ListTradingAccounts200Response) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *ListTradingAccounts200Response) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetData

`func (o *ListTradingAccounts200Response) GetData() []TradingAccount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListTradingAccounts200Response) GetDataOk() (*[]TradingAccount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListTradingAccounts200Response) SetData(v []TradingAccount)`

SetData sets Data field to given value.

### HasData

`func (o *ListTradingAccounts200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListTradingAccounts200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListTradingAccounts200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListTradingAccounts200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListTradingAccounts200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


