# FermataListDealers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | Indicates if the operation was successful | [optional] 
**Errno** | **int32** | Error code (0 for success, non-zero indicates error). Format: AABBB where AA is the module code and BBB is the error code | 
**Error** | Pointer to **NullableString** | Error message (null for successful operations) | [optional] 
**Details** | Pointer to [**NullableBaseResponseDetails**](BaseResponseDetails.md) |  | [optional] 
**Data** | Pointer to [**[]FermataDealer**](FermataDealer.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewFermataListDealers200Response

`func NewFermataListDealers200Response(errno int32, ) *FermataListDealers200Response`

NewFermataListDealers200Response instantiates a new FermataListDealers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFermataListDealers200ResponseWithDefaults

`func NewFermataListDealers200ResponseWithDefaults() *FermataListDealers200Response`

NewFermataListDealers200ResponseWithDefaults instantiates a new FermataListDealers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *FermataListDealers200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *FermataListDealers200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *FermataListDealers200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *FermataListDealers200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetErrno

`func (o *FermataListDealers200Response) GetErrno() int32`

GetErrno returns the Errno field if non-nil, zero value otherwise.

### GetErrnoOk

`func (o *FermataListDealers200Response) GetErrnoOk() (*int32, bool)`

GetErrnoOk returns a tuple with the Errno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrno

`func (o *FermataListDealers200Response) SetErrno(v int32)`

SetErrno sets Errno field to given value.


### GetError

`func (o *FermataListDealers200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FermataListDealers200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FermataListDealers200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *FermataListDealers200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *FermataListDealers200Response) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *FermataListDealers200Response) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetDetails

`func (o *FermataListDealers200Response) GetDetails() BaseResponseDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *FermataListDealers200Response) GetDetailsOk() (*BaseResponseDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *FermataListDealers200Response) SetDetails(v BaseResponseDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *FermataListDealers200Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *FermataListDealers200Response) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *FermataListDealers200Response) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetData

`func (o *FermataListDealers200Response) GetData() []FermataDealer`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FermataListDealers200Response) GetDataOk() (*[]FermataDealer, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FermataListDealers200Response) SetData(v []FermataDealer)`

SetData sets Data field to given value.

### HasData

`func (o *FermataListDealers200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *FermataListDealers200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *FermataListDealers200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *FermataListDealers200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *FermataListDealers200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


