# RpcListPortfoliosResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RpcPortfolio**](RpcPortfolio.md) |  | [optional] 
**Pagination** | Pointer to [**RpcPagination**](RpcPagination.md) |  | [optional] 
**Error** | Pointer to [**RpcError**](RpcError.md) |  | [optional] 

## Methods

### NewRpcListPortfoliosResult

`func NewRpcListPortfoliosResult() *RpcListPortfoliosResult`

NewRpcListPortfoliosResult instantiates a new RpcListPortfoliosResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcListPortfoliosResultWithDefaults

`func NewRpcListPortfoliosResultWithDefaults() *RpcListPortfoliosResult`

NewRpcListPortfoliosResultWithDefaults instantiates a new RpcListPortfoliosResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RpcListPortfoliosResult) GetData() []RpcPortfolio`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RpcListPortfoliosResult) GetDataOk() (*[]RpcPortfolio, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RpcListPortfoliosResult) SetData(v []RpcPortfolio)`

SetData sets Data field to given value.

### HasData

`func (o *RpcListPortfoliosResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *RpcListPortfoliosResult) GetPagination() RpcPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RpcListPortfoliosResult) GetPaginationOk() (*RpcPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RpcListPortfoliosResult) SetPagination(v RpcPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RpcListPortfoliosResult) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetError

`func (o *RpcListPortfoliosResult) GetError() RpcError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RpcListPortfoliosResult) GetErrorOk() (*RpcError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RpcListPortfoliosResult) SetError(v RpcError)`

SetError sets Error field to given value.

### HasError

`func (o *RpcListPortfoliosResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


